/* Copyright (c) 2021 Forte Labs
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of
 * Forte Labs and their suppliers, if any.  The intellectual and technical
 * concepts contained herein are proprietary to Forte Labs and their suppliers
 * and may be covered by U.S. and Foreign Patents, patents in process, and are
 * protected by trade secret or copyright law. Dissemination of this information
 * or reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Forte Labs.
 */

package aleo

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/rpc"
)

type Connection struct {
	endpoint string
	http     bool
	log      log15.Logger
	client   *rpc.Client
	stop     chan int
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(endpoint string, http bool, log log15.Logger) *Connection {
	return &Connection{
		endpoint: endpoint,
		http:     http,
		log:      log,
		stop:     make(chan int),
	}
}

// Connect starts the http or WS connection to the Aleo Custodian, this is established via RPC but does not make the same calls
func (c *Connection) Connect() error {
	c.log.Info("Connecting to the aleo custodian...", "url", c.endpoint)
	var rpcClient *rpc.Client
	var err error

	if c.http {
		rpcClient, err = rpc.DialHTTP(c.endpoint)
	} else {
		rpcClient, err = rpc.DialContext(context.Background(), c.endpoint)
	}

	if err != nil {
		return err
	}
	c.client = rpcClient

	return nil
}

// LatestBlock queries the Aleo Custodian for it's advertised "latest block", from our custodian's RPC replication endpoint
func (c *Connection) LatestBlock() (*big.Int, error) {
	var latestBlock *big.Int
	err := c.client.CallContext(context.Background(), &latestBlock, "latest_block")
	return latestBlock, err
}

// DepositEvents queries the Aleo Custodian for the deposit events at the latest block
func (c *Connection) DepositEvents(latestBlock *big.Int) ([]DepositLog, error) {
	var results []DepositLog
	arg := map[string]interface{}{
		"latest_block": latestBlock,
	}
	err := c.client.CallContext(context.Background(), &results, "deposit_events", arg)
	return results, err
}

// Arc721DepositRecord gets a requested ARC721 Deposit record from the custodian
func (c *Connection) Arc721DepositRecord(destId msg.ChainId, nonce msg.Nonce) (ARC721DepositRecord, error) {
	var record ARC721DepositRecord
	arg := map[string]interface{}{
		"destination_chain_id": uint64(destId),
		"nonce":                uint64(nonce),
	}
	err := c.client.CallContext(context.Background(), &record, "deposit_record", arg)
	return record, err
}

// ProposalEvents queries the Aleo Custodian for the proposal events at the latest block
func (c *Connection) ProposalEvents(latestBlock *big.Int) ([]ProposalRecord, error) {
	var results []ProposalRecord
	arg := map[string]interface{}{
		"latest_block": latestBlock,
	}
	err := c.client.CallContext(context.Background(), &results, "proposal_events", arg)
	return results, err
}

// GetProposal queries the custodian for the status of a proposal
func (c *Connection) GetProposal(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) (ProposalRecord, error) {
	var prop ProposalRecord
	arg := map[string]interface{}{
		"source_chain_id": uint8(srcId),
		"nonce": uint64(nonce),
		"data_hash": dataHash,
	}
	err := c.client.CallContext(context.Background(), &prop, "get_proposal_record", arg)
	return prop, err
}

// HasVotedOnProposal queries the custodian to check if this relayer has voted on proposal
func (c *Connection) HasVotedOnProposal(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte, from [20]byte) (bool, error){
	var hasVoted bool
	arg := map[string]interface{}{
		"source_chain_id": uint8(srcId),
		"nonce": uint64(nonce),
		"data_hash": dataHash,
		"from": from,
	}
	err := c.client.CallContext(context.Background(), &hasVoted, "get_has_voted_on_proposal", arg)
	return hasVoted, err
}

// VoteProposal vote a proposal on the custodian
func (c *Connection) VoteProposal(srcId msg.ChainId, depositNonce msg.Nonce, resourceId [32]byte, dataHash [32]byte) (string, error){
	var propId string
	arg := map[string]interface{}{
		"source_chain_id": uint8(srcId),
		"nonce": uint64(depositNonce),
		"resource_id": resourceId,
		"data_hash": dataHash,
	}

	err := c.client.CallContext(context.Background(), &propId, "vote_proposal", arg)
	return propId, err
}

// ExecuteProposal execute a proposal on the custodian
func (c *Connection) ExecuteProposal(srcId msg.ChainId, depositNonce msg.Nonce, data []byte, dataHash [32]byte, resourceId [32]byte) (string, error) {

	var propId string
	arg := map[string]interface{}{
		"source_chain_id": uint8(srcId),
		"nonce": uint64(depositNonce),
		"resource_id": resourceId,
		"data_hash": dataHash,
		"data": data,
	}

	err := c.client.CallContext(context.Background(), &propId, "execute_proposal", arg)
	return propId, err
}

func (c *Connection) WaitForBlock(targetBlock *big.Int, delay *big.Int) error {
	for {
		select {
		case <-c.stop:
			return errors.New("connection terminated")
		default:
			currBlock, err := c.LatestBlock()
			if err != nil {
				return err
			}

			if delay != nil {
				currBlock.Sub(currBlock, delay)
			}

			// Equal or greater than target
			if currBlock.Cmp(targetBlock) >= 0 {
				return nil
			}
			c.log.Trace("Block not ready, waiting", "target", targetBlock, "current", currBlock, "delay", delay)
			time.Sleep(CustodianRetryInterval)
			continue
		}
	}
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	if c.client != nil {
		c.client.Close()
	}
	close(c.stop)
}
