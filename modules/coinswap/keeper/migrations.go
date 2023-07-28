package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/irisnet/irismod/modules/coinswap/migrations/v2"
	v3 "github.com/irisnet/irismod/modules/coinswap/migrations/v3"
	v4 "github.com/irisnet/irismod/modules/coinswap/migrations/v4"
	"github.com/irisnet/irismod/types/exported"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k              Keeper
	legacySubspace exported.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(k Keeper, legacySubspace exported.Subspace) Migrator {
	return Migrator{k: k, legacySubspace: legacySubspace}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.Migrate(ctx, m.k, m.k.bk, m.k.ak)
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.Migrate(ctx, m.k, m.legacySubspace)
}

func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v4.Migrate(ctx, m.k, m.legacySubspace)
}
