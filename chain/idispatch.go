package chain

import "github.com/MysGate/demo_backend/module"

type IDispatcher interface {
	PayForDest(*module.Order) error
	GenerateZkProof(*module.Order) error
	VerifyZkProof(*module.Order) error
	OrderSucceed(*module.Order) error
}
