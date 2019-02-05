package did

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/ethereum"
	id "github.com/centrifuge/go-centrifuge/identity"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DID stores the identity address of the user
type DID common.Address

func (d DID) toAddress() common.Address {
	return common.Address(d)
}

// NewDID returns a DID based on a common.Address
func NewDID(address common.Address) DID {
	return DID(address)
}

// NewDIDFromString returns a DID based on a hex string
func NewDIDFromString(address string) DID {
	return DID(common.HexToAddress(address))
}

// Identity interface contains the methods to interact with the identity contract
type Identity interface {
	// AddKey adds a key to identity contract
	AddKey(key Key) (chan *ethereum.WatchTransaction, error)

	// GetKey return a key from the identity contract
	GetKey(key [32]byte) (*KeyResponse, error)

	// RawExecute calls the execute method on the identity contract
	RawExecute(to common.Address, data []byte) (chan *ethereum.WatchTransaction, error)

	// Execute creates the abi encoding an calls the execute method on the identity contract
	Execute(to common.Address, contractAbi, methodName string, args ...interface{}) (chan *ethereum.WatchTransaction, error)
}

type contract interface {

	// calls
	GetKey(opts *bind.CallOpts, _key [32]byte) (struct {
		Key       [32]byte
		Purposes  []*big.Int
		RevokedAt *big.Int
	}, error)

	// transactions
	AddKey(opts *bind.TransactOpts, _key [32]byte, _purpose *big.Int, _keyType *big.Int) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error)
}

type identity struct {
	config id.Config
	client ethereum.Client
	did    *DID
}

func (i identity) prepareTransaction(did DID) (contract, *bind.TransactOpts, error) {
	opts, err := i.client.GetTxOpts(i.config.GetEthereumDefaultAccountName())
	if err != nil {
		log.Infof("Failed to get txOpts from Ethereum client: %v", err)
		return nil, nil, err
	}

	contract, err := i.bindContract(did)
	if err != nil {
		return nil, nil, err
	}

	return contract, opts, nil

}

func (i identity) prepareCall(did DID) (contract, *bind.CallOpts, context.CancelFunc, error) {
	opts, cancelFunc := i.client.GetGethCallOpts(false)

	contract, err := i.bindContract(did)
	if err != nil {
		return nil, nil, nil, err
	}

	return contract, opts, cancelFunc, nil

}

func (i identity) bindContract(did DID) (contract, error) {
	contract, err := NewIdentityContract(did.toAddress(), i.client.GetEthClient())
	if err != nil {
		return nil, errors.New("Could not bind identity contract: %v", err)
	}

	return contract, nil

}

// NewIdentity creates a instance of an identity
func NewIdentity(config id.Config, client ethereum.Client, did *DID) Identity {
	// TODO use DID stored in config file
	return identity{config: config, client: client, did: did}
}

// TODO: will be replaced with statusTask
func waitForTransaction(client ethereum.Client, txHash common.Hash, txStatus chan *ethereum.WatchTransaction) {
	time.Sleep(3000 * time.Millisecond)
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		txStatus <- &ethereum.WatchTransaction{Error: err}
	}
	txStatus <- &ethereum.WatchTransaction{Status: receipt.Status}

}

func logTxHash(tx *types.Transaction) {
	log.Infof("Ethereum transaction created. Hash [%x] and Nonce [%v] and Check [%v]", tx.Hash(), tx.Nonce(), tx.CheckNonce())
	log.Infof("Transfer pending: 0x%x\n", tx.Hash())
}

func (i identity) AddKey(key Key) (chan *ethereum.WatchTransaction, error) {
	contract, opts, err := i.prepareTransaction(*i.did)
	if err != nil {
		return nil, err
	}

	tx, err := i.client.SubmitTransactionWithRetries(contract.AddKey, opts, key.GetKey(), key.GetPurpose(), key.GetType())
	if err != nil {
		log.Infof("could not addKey to identity contract: %v[txHash: %s] : %v", tx.Hash(), err)
		return nil, errors.New("could not addKey to identity contract: %v", err)
	}
	logTxHash(tx)

	txStatus := make(chan *ethereum.WatchTransaction)

	// TODO will be replaced with transaction Status task
	go waitForTransaction(i.client, tx.Hash(), txStatus)

	return txStatus, nil

}

func (i identity) GetKey(key [32]byte) (*KeyResponse, error) {
	contract, opts, _, err := i.prepareCall(*i.did)
	if err != nil {
		return nil, err
	}

	result, err := contract.GetKey(opts, key)

	if err != nil {
		return nil, errors.New("Could not call identity contract: %v", err)
	}

	return &KeyResponse{result.Key, result.Purposes, result.RevokedAt}, nil

}

func (i identity) RawExecute(to common.Address, data []byte) (chan *ethereum.WatchTransaction, error) {
	contract, opts, err := i.prepareTransaction(*i.did)
	if err != nil {
		return nil, err
	}

	// default: no ether should be send
	value := big.NewInt(0)

	tx, err := i.client.SubmitTransactionWithRetries(contract.Execute, opts, to, value, data)
	if err != nil {
		log.Infof("could not call execute method on identity contract: %v[txHash: %s] toAddress: %s : %v", tx.Hash(), to.String(), err)
		return nil, errors.New("could not execute to identity contract: %v", err)
	}
	logTxHash(tx)

	txStatus := make(chan *ethereum.WatchTransaction)
	// TODO will be replaced with transaction Status task
	go waitForTransaction(i.client, tx.Hash(), txStatus)

	return txStatus, nil

}

func (i identity) Execute(to common.Address, contractAbi, methodName string, args ...interface{}) (chan *ethereum.WatchTransaction, error) {
	abi, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	// Pack encodes the parameters and additionally checks if the method and arguments are defined correctly
	data, err := abi.Pack(methodName, args...)
	if err != nil {
		return nil, err
	}
	return i.RawExecute(to, data)
}
