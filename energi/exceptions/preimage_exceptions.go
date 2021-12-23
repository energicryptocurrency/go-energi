package exceptions

import (
	"math/big"
	"sync/atomic"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/params"
)

/*
The package main purpose is to make it possible for nodes syncing from genesis block circumvent
blocks which were "mined" with incorrect preimage database and as they became part of canonical chain in order
to follow the same blocks we need to make calculations with incorrect preimage key-value(s) listed below
to produce the same block hash and, accordingly, continue syncing to the last block
*/
var (
  // current block head height
  blockNum uint64

  // mainnet/testnet network id
  networkID uint64
)

const (
  // block number when we stop using exception and get back to preimate database
  exceptionEndBlockNum = 693000
)


/*
preimage key-value pair that is damaged - storing the
block number where this damage was first detected
*/
type PreimagePoint struct {

  // shakey for the address
  shaKey string

  // damaged value that corresponds to shaKey in canonical chain
  correspondingDamagedValue string

  // block number damage was first detected
  blockNumStart uint64

}

var preimagePoint = PreimagePoint{
  "0x975d1872ac9ae8ba471c229105acabcc5154be9f349304071c442ffa7405326d",
  "0x0000000000000000000000000000000000000000000000000000000000000012",
  325054,
}


// set current block height
func SetCurrentHead(currentBlockNum *big.Int) {
  atomic.StoreUint64(&blockNum, currentBlockNum.Uint64())
}

// sets current network id
func SetNetworkID(currentNetworkID *big.Int) {
  if currentNetworkID != nil {
    networkID = currentNetworkID.Uint64()
  }
}

/*
GetPreimage function returns damaged key corresponding to shaKey value from canonical chains
checking if the current blockchain head passed the damage time
*/
func GetPreimage(shaKey []byte) []byte {
  if networkID == params.EnergiMainnetChainConfig.ChainID.Uint64() && common.BytesToHash(shaKey).String() == preimagePoint.shaKey && atomic.LoadUint64(&blockNum) >= preimagePoint.blockNumStart && atomic.LoadUint64(&blockNum) <= exceptionEndBlockNum {
    return common.HexToHash(preimagePoint.correspondingDamagedValue).Bytes();
  }
  return nil
}
