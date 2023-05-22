package chain

import "github.com/MysGate/demo_backend/model"

type IDispatcher interface {
	PayForDest(*model.Order) error
}
