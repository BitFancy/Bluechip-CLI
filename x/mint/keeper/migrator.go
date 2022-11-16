package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/smartdev0328/bluechip/x/mint/migrations/v2"
)

// Migrator is a struct for handling in-place state migrations.
type Migrator struct {
	keeper Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{
		keeper: k,
	}
}

// Migrate migrates the x/mint module state from the consensus version
// 1 to version 2
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.Migrate(ctx, ctx.KVStore(m.keeper.storeKey), m.keeper.cdc)
}
