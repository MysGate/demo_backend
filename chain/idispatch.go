package chain

import "github.com/MysGate/demo_backend/model"

type IDispatcher interface {
	PayForDest(*model.Order) error
	GenerateZkProof(*model.Order) error
	VerifyZkProof(*model.Order) error
	OrderSucceed(*model.Order) error
}
