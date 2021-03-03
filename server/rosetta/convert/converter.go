package convert

import (
	"encoding/json"
	"fmt"
	"reflect"

	rosettatypes "github.com/coinbase/rosetta-sdk-go/types"
	"github.com/gogo/protobuf/proto"
	crgerrs "github.com/tendermint/cosmos-rosetta-gateway/errors"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/rosetta/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// converter is a utility that can be used to convert
// back and forth from rosetta to sdk and tendermint types
type Converter interface {
	FromRosetta() FromRosettaConverter
	ToRosetta() ToRosettaConverter
}

// ToRosettaConverter is an interface that exposes
// all the functions used to convert sdk and
// tendermint types to rosetta known types
type ToRosettaConverter interface {
	// CoinsToAmounts converts sdk.Coins to amounts
	CoinsToAmounts(ownedCoins []sdk.Coin, availableCoins sdk.Coins) []*rosettatypes.Amount
	// MsgToOps converts an sdk.Msg to a rosetta operation
	MsgToOps(status string, msg sdk.Msg) ([]*rosettatypes.Operation, error)
	// MsgToMeta converts an sdk.Msg to rosetta metadata
	MsgToMeta(msg sdk.Msg) (meta map[string]interface{}, err error)
	// Tx converts a tendermint transaction and tx result if provided to a rosetta tx
	Tx(rawTx tmtypes.Tx, txResult *abci.ResponseDeliverTx) (*rosettatypes.Transaction, error)
	// EventsToBalanceOps converts events to balance operations
	EventsToBalanceOps(status string, events []abci.Event) []*rosettatypes.Operation
}

// FromRosettaConverter is an interface that exposes
// all the functions used to convert rosetta types
// to tendermint and sdk types
type FromRosettaConverter interface {
	OpsToUnsignedTx(ops []*rosettatypes.Operation) (tx authsigning.Tx, err error)
	MetaToMsg(meta map[string]interface{}, msg sdk.Msg) (err error)
}

type converter struct {
	newTxBuilder func() sdkclient.TxBuilder
	txDecode     sdk.TxDecoder
	ir           codectypes.InterfaceRegistry
	cdc          *codec.ProtoCodec
}

func NewConverter(cdc *codec.ProtoCodec, ir codectypes.InterfaceRegistry, cfg sdkclient.TxConfig) Converter {
	return converter{
		newTxBuilder: func() sdkclient.TxBuilder {
			return cfg.NewTxBuilder()
		},
		ir:       ir,
		cdc:      cdc,
		txDecode: cfg.TxDecoder(),
	}
}

func (c converter) FromRosetta() FromRosettaConverter {
	return c
}

func (c converter) ToRosetta() ToRosettaConverter {
	return c
}

// OpsToUnsignedTx returns all the sdk.Msgs given the operations
func (c converter) OpsToUnsignedTx(ops []*rosettatypes.Operation) (tx authsigning.Tx, err error) {
	builder := c.newTxBuilder()

	var msgs []sdk.Msg

	for i := 0; i < len(ops); i++ {
		op := ops[i]

		protoMessage, err := c.ir.Resolve("/" + op.Type)
		if err != nil {
			return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, "operation not found: "+op.Type)
		}

		msg, ok := protoMessage.(sdk.Msg)
		if !ok {
			return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, "operation is not a valid supported sdk.Msg: "+op.Type)
		}

		err = c.MetaToMsg(op.Metadata, msg)
		if err != nil {
			return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
		}

		// verify message correctness
		if err = msg.ValidateBasic(); err != nil {
			return nil, crgerrs.WrapError(
				crgerrs.ErrBadArgument,
				fmt.Sprintf("validation of operation at index %d failed: %s", op.OperationIdentifier.Index, err),
			)
		}
		signers := msg.GetSigners()
		// check if there are enough signers
		if len(signers) == 0 {
			return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, fmt.Sprintf("operation at index %d got no signers", op.OperationIdentifier.Index))
		}
		// append the msg
		msgs = append(msgs, msg)
		// if there's only one signer then simply continue
		if len(signers) == 1 {
			continue
		}
		// after we have got the msg, we need to verify if the message has multiple signers
		// if it has got multiple signers, then we need to fetch all the related operations
		// which involve the other signers of the msg, we expect to find them in order
		// so if the msg is named "v1.test.Send" and it expects 3 signers, the next 3 operations
		// must be with the same name "v1.test.Send" and contain the other signers
		// then we can just skip their processing
		for j := 0; j < len(signers)-1; j++ {
			skipOp := ops[i+j] // get the next index
			// verify that the operation is equal to the new one
			if skipOp.Type != op.Type {
				return nil, crgerrs.WrapError(
					crgerrs.ErrBadArgument,
					fmt.Sprintf("operation at index %d should have had type %s got: %s", i+j, op.Type, skipOp.Type),
				)
			}

			if !reflect.DeepEqual(op.Metadata, skipOp.Metadata) {
				return nil, crgerrs.WrapError(
					crgerrs.ErrBadArgument,
					fmt.Sprintf("operation at index %d should have had metadata equal to %#v, got: %#v", i+j, op.Metadata, skipOp.Metadata))
			}

			i++ // increase so we skip it
		}
	}

	if err := builder.SetMsgs(msgs...); err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, err.Error())
	}

	return builder.GetTx(), nil

}

// MetaToMsg unmarshals the rosetta metadata to the given sdk.Msg
func (c converter) MetaToMsg(meta map[string]interface{}, msg sdk.Msg) error {
	metaBytes, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	return c.cdc.UnmarshalJSON(metaBytes, msg)
}

func (c converter) MsgToMeta(msg sdk.Msg) (meta map[string]interface{}, err error) {
	b, err := c.cdc.MarshalJSON(msg)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	err = json.Unmarshal(b, &meta)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	return
}

// MsgToOps will create an operation for each msg signer
// with the message proto name as type, and the raw fields
// as metadata
func (c converter) MsgToOps(status string, msg sdk.Msg) ([]*rosettatypes.Operation, error) {
	opName := proto.MessageName(msg)
	// in case proto does not recognize the message name
	// then we should try to cast it to service msg, to
	// check if it was wrapped or not, in case the cast
	// from sdk.ServiceMsg to sdk.Msg fails, then a
	// codec error is returned
	if opName == "" {
		unwrappedMsg, ok := msg.(sdk.ServiceMsg)
		if !ok {
			return nil, crgerrs.WrapError(crgerrs.ErrCodec, fmt.Sprintf("unrecognized message type: %T", msg))
		}

		msg, ok = unwrappedMsg.Request.(sdk.Msg)
		if !ok {
			return nil, crgerrs.WrapError(
				crgerrs.ErrCodec,
				fmt.Sprintf("unable to cast %T to sdk.Msg, method: %s", unwrappedMsg.Request, unwrappedMsg.MethodName),
			)
		}

		opName = proto.MessageName(msg)
		if opName == "" {
			return nil, crgerrs.WrapError(crgerrs.ErrCodec, fmt.Sprintf("unrecognized message type: %T", msg))
		}
	}

	meta, err := c.MsgToMeta(msg)
	if err != nil {
		return nil, err
	}

	ops := make([]*rosettatypes.Operation, len(msg.GetSigners()))
	for i, signer := range msg.GetSigners() {
		op := &rosettatypes.Operation{
			Type:     opName,
			Status:   status,
			Account:  &rosettatypes.AccountIdentifier{Address: signer.String()},
			Metadata: meta,
		}

		ops[i] = op
	}

	return ops, nil
}

// Tx converts a tendermint raw transaction and its result (if provided) to a rosetta transaction
func (c converter) Tx(rawTx tmtypes.Tx, txResult *abci.ResponseDeliverTx) (*rosettatypes.Transaction, error) {
	// decode tx
	tx, err := c.txDecode(rawTx)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}
	// get initial status, as per sdk design, if one msg fails
	// the whole TX will be considered failing, so we can't have
	// 1 msg being success and 1 msg being reverted
	status := types.StatusTxSuccess
	switch txResult {
	// if nil, we're probably checking an unconfirmed tx
	case nil:
		status = types.StatusTxUnconfirmed
	// set the status
	default:
		if txResult.Code != abci.CodeTypeOK {
			status = types.StatusTxReverted
		}
	}
	// get operations from msgs
	msgs := tx.GetMsgs()
	var rawTxOps []*rosettatypes.Operation
	for _, msg := range msgs {
		ops, err := c.MsgToOps(status, msg)
		if err != nil {
			return nil, err
		}
		rawTxOps = append(rawTxOps, ops...)
	}

	// now get balance events from response deliver tx
	var balanceOps []*rosettatypes.Operation
	// tx result might be nil, in case we're querying an unconfirmed tx from the mempool
	if txResult != nil {
		balanceOps = c.EventsToBalanceOps(status, txResult.Events)
	}

	// now normalize indexes
	totalOps := AddOperationIndexes(rawTxOps, balanceOps)

	return &rosettatypes.Transaction{
		TransactionIdentifier: &rosettatypes.TransactionIdentifier{Hash: fmt.Sprintf("%x", rawTx.Hash())},
		Operations:            totalOps,
	}, nil
}

func (c converter) EventsToBalanceOps(status string, events []abci.Event) []*rosettatypes.Operation {
	var ops []*rosettatypes.Operation

	for _, e := range events {
		balanceOps, ok := sdkEventToBalanceOperations(status, e)
		if !ok {
			continue
		}
		ops = append(ops, balanceOps...)
	}

	return ops
}

// sdkEventToBalanceOperations converts an event to a rosetta balance operation
// it will panic if the event is malformed because it might mean the sdk spec
// has changed and rosetta needs to reflect those changes too.
// The balance operations are multiple, one for each denom.
func sdkEventToBalanceOperations(status string, event abci.Event) (operations []*rosettatypes.Operation, isBalanceEvent bool) {

	var (
		accountIdentifier string
		coinChange        sdk.Coins
		isSub             bool
	)

	switch event.Type {
	default:
		return nil, false
	case banktypes.EventTypeCoinSpent:
		spender, err := sdk.AccAddressFromBech32((string)(event.Attributes[0].Value))
		if err != nil {
			panic(err)
		}
		coins, err := sdk.ParseCoinsNormalized((string)(event.Attributes[1].Value))
		if err != nil {
			panic(err)
		}

		isSub = true
		coinChange = coins
		accountIdentifier = spender.String()

	case banktypes.EventTypeCoinReceived:
		receiver, err := sdk.AccAddressFromBech32((string)(event.Attributes[0].Value))
		if err != nil {
			panic(err)
		}
		coins, err := sdk.ParseCoinsNormalized((string)(event.Attributes[1].Value))
		if err != nil {
			panic(err)
		}

		isSub = false
		coinChange = coins
		accountIdentifier = receiver.String()

	// rosetta does not have the concept of burning coins, so we need to mock
	// the burn as a send to an address that cannot be resolved to anything
	case banktypes.EventTypeCoinBurn:
		coins, err := sdk.ParseCoinsNormalized((string)(event.Attributes[1].Value))
		if err != nil {
			panic(err)
		}

		coinChange = coins
		accountIdentifier = types.BurnerAddressIdentifier
	}

	operations = make([]*rosettatypes.Operation, len(coinChange))

	for i, coin := range coinChange {

		value := coin.Amount.String()
		// in case the event is a subtract balance one the rewrite value with
		// the negative coin identifier
		if isSub {
			value = "-" + value
		}

		op := &rosettatypes.Operation{
			Type:    event.Type,
			Status:  status,
			Account: &rosettatypes.AccountIdentifier{Address: accountIdentifier},
			Amount: &rosettatypes.Amount{
				Value: value,
				Currency: &rosettatypes.Currency{
					Symbol:   coin.Denom,
					Decimals: 0,
				},
			},
		}

		operations[i] = op
	}
	return operations, true
}

// sdkCoinsToRosettaAmounts converts []sdk.Coin to rosetta amounts
func (c converter) CoinsToAmounts(ownedCoins []sdk.Coin, availableCoins sdk.Coins) []*rosettatypes.Amount {
	amounts := make([]*rosettatypes.Amount, len(availableCoins))
	ownedCoinsMap := make(map[string]sdk.Int, len(availableCoins))

	for _, ownedCoin := range ownedCoins {
		ownedCoinsMap[ownedCoin.Denom] = ownedCoin.Amount
	}

	for i, coin := range availableCoins {
		value, owned := ownedCoinsMap[coin.Denom]
		if !owned {
			amounts[i] = &rosettatypes.Amount{
				Value: sdk.NewInt(0).String(),
				Currency: &rosettatypes.Currency{
					Symbol: coin.Denom,
				},
			}
			continue
		}
		amounts[i] = &rosettatypes.Amount{
			Value: value.String(),
			Currency: &rosettatypes.Currency{
				Symbol: coin.Denom,
			},
		}
	}

	return amounts
}

// AddOperationIndexes adds the indexes to operations adhering to specific rules:
// operations related to messages will be always before than the balance ones
func AddOperationIndexes(msgOps []*rosettatypes.Operation, balanceOps []*rosettatypes.Operation) (finalOps []*rosettatypes.Operation) {
	lenMsgOps := len(msgOps)
	lenBalanceOps := len(balanceOps)
	finalOps = make([]*rosettatypes.Operation, 0, lenMsgOps+lenBalanceOps)

	var currentIndex int64
	// add indexes to msg ops
	for _, op := range msgOps {
		op.OperationIdentifier = &rosettatypes.OperationIdentifier{
			Index: currentIndex,
		}

		finalOps = append(finalOps, op)
		currentIndex++
	}

	// add indexes to balance ops
	for _, op := range balanceOps {
		op.OperationIdentifier = &rosettatypes.OperationIdentifier{
			Index: currentIndex,
		}

		finalOps = append(finalOps, op)
		currentIndex++
	}

	return finalOps
}
