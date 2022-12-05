package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/smartdev0328/bluechip/testutil/network"
	"github.com/smartdev0328/bluechip/testutil/nullify"
	"github.com/smartdev0328/bluechip/x/pageinflation/client/cli"
	"github.com/smartdev0328/bluechip/x/pageinflation/types"
)

func networkWithStartBlockObjects(t *testing.T) (*network.Network, types.StartBlock) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	startBlock := &types.StartBlock{}
	nullify.Fill(&startBlock)
	state.StartBlock = startBlock
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.StartBlock
}

func TestShowStartBlock(t *testing.T) {
	net, obj := networkWithStartBlockObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.StartBlock
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowStartBlock(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetStartBlockResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.StartBlock)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.StartBlock),
				)
			}
		})
	}
}
