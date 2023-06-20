package model

import (
	"math/big"

	"github.com/MysGate/demo_backend/util"
	"github.com/ethereum/go-ethereum/crypto"
)

type ProofReq struct {
	Addr   string `json:"addr"`
	Url    string `json:"url"`
	TxHash string `json:"txhash"`
	CmtIdx int    `json:"cmtIdx"`
}

type Proof struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

type ZkProof struct {
	Proof      *Proof
	PublicInfo [2]*big.Int
}

type RawProof struct {
	A []string   `json:"a"`
	B [][]string `json:"b"`
	C []string   `json:"c"`
}

type RawZkProof struct {
	Proof      *RawProof `json:"proof"`
	PublicInfo []string  `json:"publicInfo"`
}

func (zp *ZkProof) Keccak256EncodePackedZkProof() (proofHash [32]byte) {
	result := util.EncodePacked(
		util.EncodeBigInt(zp.Proof.A[0]),
		util.EncodeBigInt(zp.Proof.A[1]),
		util.EncodeBigInt(zp.Proof.B[0][0]),
		util.EncodeBigInt(zp.Proof.B[0][1]),
		util.EncodeBigInt(zp.Proof.B[1][0]),
		util.EncodeBigInt(zp.Proof.B[1][1]),
		util.EncodeBigInt(zp.Proof.C[0]),
		util.EncodeBigInt(zp.Proof.C[1]))
	hash := crypto.Keccak256Hash(result)
	copy(proofHash[:], hash.Bytes())
	return
}

func (rp *RawZkProof) RawProofToZkProof() (zProof *ZkProof) {
	zProof = &ZkProof{Proof: &Proof{}, PublicInfo: [2]*big.Int{}}
	A0, A1, err := util.ConvertBigIntFromString(rp.Proof.A[0], rp.Proof.A[1])
	if err != nil {
		util.Logger().Error(err.Error())
		return
	}
	zProof.Proof.A[0] = A0
	zProof.Proof.A[1] = A1

	B00, B01, err := util.ConvertBigIntFromString(rp.Proof.B[0][0], rp.Proof.B[0][1])
	if err != nil {
		util.Logger().Error(err.Error())
		return
	}
	zProof.Proof.B[0][0] = B00
	zProof.Proof.B[0][1] = B01

	B10, B11, err := util.ConvertBigIntFromString(rp.Proof.B[1][0], rp.Proof.B[1][1])
	if err != nil {
		util.Logger().Error(err.Error())
		return
	}
	zProof.Proof.B[1][0] = B10
	zProof.Proof.B[1][1] = B11

	C0, C1, err := util.ConvertBigIntFromString(rp.Proof.C[0], rp.Proof.C[1])
	if err != nil {
		util.Logger().Error(err.Error())
		return
	}
	zProof.Proof.C[0] = C0
	zProof.Proof.C[1] = C1

	pi0, pi1, err := util.ConvertBigIntFromString(rp.PublicInfo[0], rp.PublicInfo[1])
	if err != nil {
		util.Logger().Error(err.Error())
		return
	}
	zProof.PublicInfo[0] = pi0
	zProof.PublicInfo[1] = pi1

	return
}
