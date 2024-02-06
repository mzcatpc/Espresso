// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package arbvid

import (
	espressoTypes "github.com/EspressoSystems/espresso-sequencer-go/types"
)

func VerifyNamespace(namespace uint64, proof espressoTypes.Bytes, block_comm espressoTypes.NmtRoot, txs []espressoTypes.Bytes, srs espressoTypes.Bytes) error {
	verifyNamespace(namespace, proof, block_comm, txs, srs)
}
