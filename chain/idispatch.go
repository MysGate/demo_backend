package chain

import "github.com/MysGate/demo_backend/module"

type IDispatcher interface {
	DispatchCrossChainOrder(*module.Order) error
}
