// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

//go:build js
// +build js

package arbvid

import espressoTypes "github.com/EspressoSystems/espresso-sequencer-go/types"

func verifyNamespace(namespace uint64, proof espressoTypes.Bytes, block_comm espressoTypes.NmtRoot, txs []espressoTypes.Bytes, srs espressoTypes.Bytes)
