package zkp

import (
	"encoding/json"
	"fmt"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/model"
	"github.com/MysGate/demo_backend/util"
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

func (m *ProofManager) GetZkRawProof(req *model.ProofReq) (*model.RawZkProof, string) {
	url := m.Cfg.ZkVerify.ProofUrl
	content, err := json.Marshal(req)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("GetZkRawProof HTTPReq err:%+v", err))
		return nil, ""
	}

	headers := make(map[string]string)
	headers["Content-Type"] = " application/json"
	hc := util.GetHTTPClient()
	body, err := util.HTTPReq("POST", url, hc, content, headers)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("GetZkRawProof HTTPReq err:%+v", err))
		return nil, ""
	}

	zkp := &model.RawZkProof{}
	err = json.Unmarshal(body, zkp)
	if err != nil {
		util.Logger().Error(fmt.Sprintf("GetZkRawProof Unmarshal err:%+v", err))
		return nil, ""
	}

	return zkp, string(body)
}
