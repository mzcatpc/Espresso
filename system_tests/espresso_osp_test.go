package arbtest

import (
	"context"
	"math/big"
	"testing"

	"github.com/EspressoSystems/espresso-sequencer-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/offchainlabs/nitro/solgen/go/ospgen"
	"github.com/offchainlabs/nitro/solgen/go/test_helpersgen"
	"github.com/offchainlabs/nitro/validator"
	"github.com/offchainlabs/nitro/validator/server_arb"
	"github.com/offchainlabs/nitro/validator/server_common"
	"github.com/offchainlabs/nitro/validator/valnode"
)

func TestEspressoOsp(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	initialBalance := new(big.Int).Lsh(big.NewInt(1), 200)
	l1Info := NewL1TestInfo(t)
	l1Info.GenerateGenesisAccount("deployer", initialBalance)

	deployerTxOpts := l1Info.GetDefaultTransactOpts("deployer", ctx)

	chainConfig := params.ArbitrumDevTestChainConfig()
	l1Info, l1Backend, _, _ := createTestL1BlockChain(t, l1Info)
	hotshotAddr, tx, _, err := test_helpersgen.DeployMockHotShot(&deployerTxOpts, l1Backend)
	Require(t, err)
	_, err = EnsureTxSucceeded(ctx, l1Backend, tx)
	Require(t, err)

	rollup, _ := DeployOnTestL1(t, ctx, l1Info, l1Backend, chainConfig, hotshotAddr)

	ospEntryAddr := common.HexToAddress("0xffd0c2C95214aa9980D7419bd87c260C80Ce2546")

	locator, err := server_common.NewMachineLocator("")
	if err != nil {
		Fatal(t, err)
	}
	wasmModuleRoot := locator.LatestWasmModuleRoot()
	if (wasmModuleRoot == common.Hash{}) {
		Fatal(t, "latest machine not found")
	}
	fetcher := func() *server_arb.ArbitratorSpawnerConfig {
		return &valnode.DefaultValidationConfig.Arbitrator
	}
	arbSpawner, err := server_arb.NewArbitratorSpawner(locator, fetcher)
	Require(t, err)
	err = arbSpawner.Start(ctx)
	Require(t, err)

	// Read the validation input from json file
	input := validator.ValidationInput{
		StartState: validator.GoGlobalState{
			BlockHash: common.HexToHash("0xead0c2C95214aa2480D7409bd87c260C89Ce2548"),
		},
	}

	runPromise := arbSpawner.CreateExecutionRun(common.Hash{}, &input)
	goodRun, err := runPromise.Await(ctx)
	Require(t, err)
	input.HotShotCommitment = types.Commitment([32]byte{1: 1})

	runPromise = arbSpawner.CreateExecutionRun(common.Hash{}, &input)
	badRun, err := runPromise.Await(ctx)
	Require(t, err)

	step := uint64(64)
	beforeHash := common.Hash{}
	expectedAfterHash := common.Hash{}
	for {
		goodResult, err := goodRun.GetStepAt(step).Await(ctx)
		Require(t, err)
		badResult, err := badRun.GetStepAt(step).Await(ctx)
		Require(t, err)

		if goodResult != badResult {
			expectedAfterHash = goodResult.Hash
			break
		}

		beforeHash = goodResult.Hash
		step += 1
	}

	goodProof, err := goodRun.GetProofAt(uint64(step)).Await(ctx)
	Require(t, err)
	_, err = badRun.GetProofAt(uint64(step)).Await(ctx)
	Require(t, err)

	ospEntry, err := ospgen.NewOneStepProofEntry(ospEntryAddr, l1Backend)
	Require(t, err)
	afterhash, err := ospEntry.ProveOneStep(
		l1Info.GetDefaultCallOpts("deployer", ctx),
		ospgen.ExecutionContext{
			MaxInboxMessagesRead: big.NewInt(1),
			Bridge:               rollup.Bridge,
		}, big.NewInt(int64(step)),
		beforeHash,
		goodProof,
	)
	if expectedAfterHash != afterhash {
		t.Fatal("")
	}
}
