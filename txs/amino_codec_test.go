package txs

import (
	"testing"

	"fmt"

	acm "github.com/hyperledger/burrow/account"
	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/txs/payload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAminoEncodeTxDecodeTx(t *testing.T) {
	codec := NewAminoCodec()
	inputAddress := crypto.Address{1, 2, 3, 4, 5}
	outputAddress := crypto.Address{5, 4, 3, 2, 1}
	amount := uint64(2)
	sequence := uint64(3)
	tx := &payload.SendTx{
		Inputs: []*payload.TxInput{{
			Address:  inputAddress,
			Amount:   amount,
			Sequence: sequence,
		}},
		Outputs: []*payload.TxOutput{{
			Address: outputAddress,
			Amount:  amount,
		}},
	}
	txEnv := Enclose(chainID, tx)
	txBytes, err := codec.EncodeTx(txEnv)
	if err != nil {
		t.Fatal(err)
	}
	txEnvOut, err := codec.DecodeTx(txBytes)
	assert.NoError(t, err, "DecodeTx error")
	assert.Equal(t, txEnv, txEnvOut)
}

func TestAminoEncodeTxDecodeTx_CallTx(t *testing.T) {
	codec := NewAminoCodec()
	inputAccount := acm.GeneratePrivateAccountFromSecret("fooo")
	amount := uint64(2)
	sequence := uint64(3)
	tx := &payload.CallTx{
		Input: &payload.TxInput{
			Address:  inputAccount.Address(),
			Amount:   amount,
			Sequence: sequence,
		},
		GasLimit: 233,
		Fee:      2,
		Address:  nil,
		Data:     []byte("code"),
	}
	txEnv := Enclose(chainID, tx)
	require.NoError(t, txEnv.Sign(inputAccount))
	txBytes, err := codec.EncodeTx(txEnv)
	if err != nil {
		t.Fatal(err)
	}
	txEnvOut, err := codec.DecodeTx(txBytes)
	assert.NoError(t, err, "DecodeTx error")
	assert.Equal(t, txEnv, txEnvOut)
}

func TestAminoTxEnvelope(t *testing.T) {
	codec := NewAminoCodec()
	privAccFrom := acm.GeneratePrivateAccountFromSecret("foo")
	privAccTo := acm.GeneratePrivateAccountFromSecret("bar")
	toAddress := privAccTo.Address()
	txEnv := Enclose("testChain", payload.NewCallTxWithSequence(privAccFrom.PublicKey(), &toAddress,
		[]byte{3, 4, 5, 5}, 343, 2323, 12, 3))
	err := txEnv.Sign(privAccFrom)
	require.NoError(t, err)

	bs, err := codec.EncodeTx(txEnv)
	require.NoError(t, err)
	txEnvOut, err := codec.DecodeTx(bs)
	require.NoError(t, err)
	fmt.Println(txEnvOut)
	assert.Equal(t, txEnv, txEnvOut)
}
