package conf

import (
	"crypto/ecdsa"
	"io/ioutil"
	"os"
	"strings"

	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

type MySQL struct {
	Uri     string `yaml:"uri"`
	ShowSQL bool   `yaml:"show_sql"`
}

type Chain struct {
	WssRpcUrl       string `yaml:"wss_rpc_url"`
	HttpRpcUrl      string `yaml:"http_rpc_url"`
	PrivateKey      string `yaml:"private_key"`
	Key             *ecdsa.PrivateKey
	ChainID         uint64 `yaml:"chain_id"`
	ContractAddress string `yaml:"contract_address"`
	ContractAddr    common.Address
	Name            string   `yaml:"name"`
	SuppirtCoins    []string `yaml:"support_coins"`
}

type CrossChain struct {
	SrcChainId  uint64 `yaml:"src_chain_id"`
	DestChainId uint64 `yaml:"dest_chain_id"`
}

type CrossChainFee struct {
	Name      string  `yaml:"name"`
	Fixed     float64 `yaml:"fixed"`
	FloatRate float64 `yaml:"float_rate"`
}

type CoinAmountLimit struct {
	Name      string  `yaml:"name"`
	MinAmount float64 `yaml:"min_amount"`
	MaxAmount float64 `yaml:"max_amount"`
}

type CrossChainCoin struct {
	Name     string `yaml:"name"`
	CoinType string `yaml:"type"`
}

type Router struct {
	Type    string   `yaml:"type"`
	Porters []string `yaml:"porters"`
}

type MysGateConfig struct {
	SupportChains     map[uint64]*Chain
	SupportCrossChain map[uint64][]uint64
	Coins             map[string]*CrossChainCoin
	Fee               map[string]*CrossChainFee
	Limit             map[string]*CoinAmountLimit
	Router            *Router            `yaml:"router"`
	SupportCoins      []string           `yaml:"support_coins"`
	CoinAmountLimits  []*CoinAmountLimit `yaml:"cross_chain_coin_limit"`
	CrossChainCoins   []*CrossChainCoin  `yaml:"cross_chain_coins"`
	Crosschainfee     []*CrossChainFee   `yaml:"cross_chain_fees"`
	Chains            []*Chain           `yaml:"chains"`
	Crosschains       []*CrossChain      `yaml:"cross_chains"`
	Service           *Service           `yaml:"service"`
	Logger            *Log               `yaml:"log"`
	Debug             bool               `yaml:"debug"`
	ZkpUrl            string             `yaml:"zkp_url"`
	MySql             *MySQL             `yaml:"mysql"`
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
	s.ContractAddr = common.HexToAddress(s.ContractAddress)
}

func (c *MysGateConfig) initConfig() {
	c.SupportChains = make(map[uint64]*Chain)
	c.SupportCrossChain = make(map[uint64][]uint64)
	c.Fee = make(map[string]*CrossChainFee)
	c.Coins = make(map[string]*CrossChainCoin)
	c.Limit = make(map[string]*CoinAmountLimit)

	for _, s := range c.Chains {
		initChain(s)
		c.SupportChains[s.ChainID] = s
	}

	for _, cc := range c.Crosschains {
		c.SupportCrossChain[cc.SrcChainId] = append(c.SupportCrossChain[cc.SrcChainId], cc.DestChainId)
	}

	for _, ccf := range c.Crosschainfee {
		c.Fee[strings.ToLower(ccf.Name)] = ccf
	}

	for _, ccc := range c.CrossChainCoins {
		c.Coins[strings.ToLower(ccc.Name)] = ccc
	}

	for _, cal := range c.CoinAmountLimits {
		c.Limit[strings.ToLower(cal.Name)] = cal
	}
}

func (c *MysGateConfig) FindCrossChain(cid uint64) *Chain {
	if v, ok := c.SupportChains[cid]; ok {
		return v
	}

	return nil
}

func (c *MysGateConfig) GetCrossChainFee(coin string) *CrossChainFee {
	lc := strings.ToLower(coin)
	token, ok := c.Coins[lc]
	if !ok {
		util.Logger().Error("GetCrossChainFee coin err")
		return nil
	}

	ccf, ok := c.Fee[token.CoinType]
	if !ok {
		util.Logger().Error("GetCrossChainFee coin type err")
		return nil
	}

	return ccf
}

func (c *MysGateConfig) GetCoinLimit(coin string) *CoinAmountLimit {
	lc := strings.ToLower(coin)
	token, ok := c.Coins[lc]
	if !ok {
		util.Logger().Error("GetCoinLimit coin err")
		return nil
	}

	cal, ok := c.Limit[token.CoinType]
	if !ok {
		util.Logger().Error("GetCoinLimit coin type err")
		return nil
	}

	return cal
}

// func (c *MysGateConfig) CloseClient() {
// 	for _, v := range c.SupportChains {
// 		if v.SrcClient != nil {
// 			v.SrcClient.Close()
// 		}

// 		if v.DestClient != nil {
// 			v.DestClient.Close()
// 		}
// 	}
// }

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
