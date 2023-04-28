package conf

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v3"
)

var c *MysGateConfig

type Log struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

type Service struct {
	ServicePort string `yaml:"service_port"`
}

type Chain struct {
	RpcUrl          string `yaml:"rpc_url"`
	PrivateKey      string `yaml:"private_key"`
	Key             *ecdsa.PrivateKey
	ChainID         *big.Int
	Value           *big.Int
	GasLimit        uint64
	GasSuggest      *big.Int
	ContractAddress string `yaml:"contract_address"`
	Addr            common.Address
	Client          *ethclient.Client
}

type CrossChain struct {
	Name      string `yaml:"name"`
	SrcChain  *Chain `yaml:"src_chain"`
	DestChain *Chain `yaml:"dest_chain"`
}

type MysGateConfig struct {
	ChainMap map[string]*CrossChain
	Chains   []*CrossChain `yaml:"cross_chains"`
	Service  *Service      `yaml:"service_port"`
	Logger   *Log          `yaml:"log"`
	Debug    bool          `yaml:"debug"`
}

func GetConfig() *MysGateConfig {
	if c == nil {
		util.Logger().Fatal("conf is empty")
	}
	return c
}

func initChain(s *Chain) {
	privateKey, err := crypto.HexToECDSA(s.PrivateKey)
	if err != nil {
		util.Logger().Fatal("Dial err")
	}

	s.Key = privateKey
	s.Addr = common.HexToAddress(s.ContractAddress)
}

func (c *MysGateConfig) initConfig() {
	c.ChainMap = make(map[string]*CrossChain)

	for _, s := range c.Chains {
		initChain(s.SrcChain)
		initChain(s.DestChain)

		c.ChainMap[s.Name] = s
	}
}

func (c *MysGateConfig) findCrossChain(k string) *CrossChain {
	if v, ok := c.ChainMap[k]; ok {
		return v
	}

	return nil
}

func (c *MysGateConfig) initEthClient(cc *Chain) error {
	conn, err := ethclient.Dial(cc.RpcUrl)
	if err != nil {
		errMsg := fmt.Sprintf("conn err is: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	id, err := conn.ChainID(context.Background())
	if err != nil {
		errMsg := fmt.Sprintf("get chain id err: %+v", err)
		util.Logger().Error(errMsg)
		return err
	}

	cc.ChainID = id

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(30000000)             // in units
	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		errMsg := fmt.Sprintf("get suggest gas price err:", err)
		util.Logger().Error(errMsg)
		return err
	}

	cc.Value = value
	cc.GasLimit = gasLimit
	cc.GasSuggest = gasPrice
	cc.Client = conn

	return nil
}

func (c *MysGateConfig) getEthClient(k string) (*ethclient.Client, *ethclient.Client) {
	cc := c.findCrossChain(k)
	if cc == nil {
		util.Logger().Error("Dial err")
		return nil, nil
	}

	err := c.initEthClient(cc.SrcChain)
	if err != nil {
		errMsg := fmt.Sprintf("getEthClient src err:", err)
		util.Logger().Error(errMsg)
		return nil, nil
	}

	err = c.initEthClient(cc.DestChain)
	if err != nil {
		errMsg := fmt.Sprintf("getEthClient dest err:", err)
		util.Logger().Error(errMsg)
		return nil, nil
	}

	return cc.SrcChain.Client, cc.SrcChain.Client
}

func (c *MysGateConfig) CloseClient() {
	for _, s := range c.Chains {
		s.SrcChain.Client.Close()
		s.DestChain.Client.Close()
	}
}

func ParseYaml(configFile string) error {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		os.Exit(0)
	}

	c = &MysGateConfig{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		os.Exit(0)
	}

	c.initConfig()
	return nil
}
