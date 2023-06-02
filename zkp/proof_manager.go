package zkp

import (
	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/model"
)

var pm *ProofManager

type ProofManager struct {
	Cfg *conf.MysGateConfig
}

func GetProofManager(cfg *conf.MysGateConfig) *ProofManager {
	if pm == nil {
		pm = &ProofManager{
			Cfg: cfg,
		}
	}
	return pm
}

func (m *ProofManager) GetZKProof() *model.ZkProof {
	// url := m.Cfg.ZkpUrl
	// hc := util.GetHTTPClient()
	// body, err := util.HTTPGet("POST", url, hc)
	// if err != nil {
	// 	util.Logger().Error(fmt.Sprintf("GetZKProof err:+v", err))
	// 	return nil
	// }
	zkp := &model.ZkProof{Proof: "Success"}
	// err = json.Unmarshal(body, zkp)
	// if err != nil {
	// 	util.Logger().Error(fmt.Sprintf("GetZKProof err:+v", err))
	// 	return nil
	// }

	return zkp
}
