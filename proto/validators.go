package proto

import (
	"encoding/hex"
	"fmt"

	bbn "github.com/babylonchain/babylon/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func (v *StoreValidator) GetBabylonPK() *secp256k1.PubKey {
	return &secp256k1.PubKey{
		Key: v.BabylonPk,
	}
}

func NewBabylonPkFromHex(hexStr string) (*secp256k1.PubKey, error) {
	pkBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}

	return &secp256k1.PubKey{Key: pkBytes}, nil
}

func (v *StoreValidator) GetBabylonPkHexString() string {
	return hex.EncodeToString(v.BabylonPk)
}

func (v *StoreValidator) MustGetBTCPK() *btcec.PublicKey {
	btcPubKey, err := schnorr.ParsePubKey(v.BtcPk)
	if err != nil {
		panic(fmt.Errorf("failed to parse BTC PK: %w", err))
	}
	return btcPubKey
}

func (v *StoreValidator) MustGetBIP340BTCPK() *bbn.BIP340PubKey {
	btcPK := v.MustGetBTCPK()
	return bbn.NewBIP340PubKeyFromBTCPK(btcPK)
}

func NewValidatorInfo(v *StoreValidator) *ValidatorInfo {
	return &ValidatorInfo{
		BabylonPkHex:        v.GetBabylonPkHexString(),
		BtcPkHex:            v.MustGetBIP340BTCPK().MarshalHex(),
		LastVotedHeight:     v.LastVotedHeight,
		LastCommittedHeight: v.LastCommittedHeight,
		Status:              v.Status,
	}
}