package router

import (
	"strings"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/module"
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

func (m *Manager) SelectPorters(req *module.RoterReq) (porters []*module.Porter) {
	if strings.ToLower(m.Cfg.Router.Type) == "fixed" {
		for _, addr := range m.Cfg.Router.Porters {
			p := &module.Porter{
				Address: addr,
			}
			porters = append(porters, p)
		}
	}
	return
}
