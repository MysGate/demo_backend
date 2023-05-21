package router

import (
	"strings"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/model"
)

var m *Manager

type Manager struct {
	Cfg *conf.MysGateConfig
}

func GetManager(c *conf.MysGateConfig) *Manager {
	if m == nil {
		m = &Manager{
			Cfg: c,
		}
	}
	return m
}

func (m *Manager) SelectPorters(req *model.RouterReq) (porters []*model.Porter) {
	if strings.ToLower(m.Cfg.Router.Type) == "fixed" {
		for _, addr := range m.Cfg.Router.Porters {
			p := &model.Porter{
				Address: addr,
			}
			porters = append(porters, p)
		}
	}
	return
}
