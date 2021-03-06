// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eventlog

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*depositMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (d DepositEvent) MarshalJSON() ([]byte, error) {
	type DepositEvent struct {
		Depositor    common.Address `json:"depositor"     gencodec:"required"`
		DepositIndex hexutil.Uint64 `json:"depositIndex"  gencodec:"required"`
		Denomination hexutil.Uint64 `json:"denomination"  gencodec:"required"`
		TokenID      hexutil.Uint64 `json:"tokenID"       gencodec:"required"`
	}
	var enc DepositEvent
	enc.Depositor = d.Depositor
	enc.DepositIndex = hexutil.Uint64(d.DepositIndex)
	enc.Denomination = hexutil.Uint64(d.Denomination)
	enc.TokenID = hexutil.Uint64(d.TokenID)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (d *DepositEvent) UnmarshalJSON(input []byte) error {
	type DepositEvent struct {
		Depositor    *common.Address `json:"depositor"     gencodec:"required"`
		DepositIndex *hexutil.Uint64 `json:"depositIndex"  gencodec:"required"`
		Denomination *hexutil.Uint64 `json:"denomination"  gencodec:"required"`
		TokenID      *hexutil.Uint64 `json:"tokenID"       gencodec:"required"`
	}
	var dec DepositEvent
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Depositor == nil {
		return errors.New("missing required field 'depositor' for DepositEvent")
	}
	d.Depositor = *dec.Depositor
	if dec.DepositIndex == nil {
		return errors.New("missing required field 'depositIndex' for DepositEvent")
	}
	d.DepositIndex = uint64(*dec.DepositIndex)
	if dec.Denomination == nil {
		return errors.New("missing required field 'denomination' for DepositEvent")
	}
	d.Denomination = uint64(*dec.Denomination)
	if dec.TokenID == nil {
		return errors.New("missing required field 'tokenID' for DepositEvent")
	}
	d.TokenID = uint64(*dec.TokenID)
	return nil
}
