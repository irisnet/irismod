package v160_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v160 "github.com/irisnet/irismod/modules/nft/migrations/v160"
	"github.com/irisnet/irismod/modules/nft/types"
	"github.com/irisnet/irismod/simapp"
)

type legacyKeeper struct {
	ctx      sdk.Context
	cdc      codec.Codec
	storeKey storetypes.StoreKey
}

func TestMigrateStore(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	cdc := app.AppCodec()
	storeKey := app.GetKey(types.ModuleName)
	nftKeeper := app.NFTKeeper

	legacy := legacyKeeper{
		ctx:      ctx,
		cdc:      cdc,
		storeKey: storeKey,
	}

	type testcase = struct {
		col    types.Collection
		before func() error
		after  func(col types.Collection) error
	}

	var cases = []testcase{
		{
			col: types.Collection{
				Denom: types.Denom{
					Id:               "denom1",
					Name:             "Name",
					Schema:           "Schema",
					Creator:          sdk.AccAddress([]byte("Creator1")).String(),
					Symbol:           "Symbol",
					MintRestricted:   false,
					UpdateRestricted: false,
				},
				NFTs: []types.BaseNFT{
					{
						Id:    "1",
						Name:  "token1",
						Owner: sdk.AccAddress([]byte("owner")).String(),
						URI:   "https://example.com/token/1",
						Data:  "metadata1",
					},
					{
						Id:    "2",
						Name:  "token2",
						Owner: sdk.AccAddress([]byte("owner")).String(),
						URI:   "https://example.com/token/2",
						Data:  "metadata2",
					},
				},
			},
			before: func() error {
				return v160.MigrateStore(ctx, storeKey, cdc, nftKeeper.IssueDenom, nftKeeper.MintNFT)
			},
			after: func(col types.Collection) error {
				expCol, err := nftKeeper.GetCollection(ctx, col.Denom.Id)
				require.NoError(t, err)
				require.EqualValues(t, col, expCol)
				require.EqualValues(t, len(col.NFTs), nftKeeper.GetTotalSupply(ctx, col.Denom.Id))
				return nil
			},
		},
		{
			col: types.Collection{
				Denom: types.Denom{
					Id:               "denom2",
					Name:             "Name",
					Schema:           "Schema",
					Creator:          sdk.AccAddress([]byte("Creator2")).String(),
					Symbol:           "Symbol",
					MintRestricted:   false,
					UpdateRestricted: false,
				},
				NFTs: []types.BaseNFT{
					{
						Id:    "1",
						Name:  "token1",
						Owner: sdk.AccAddress([]byte("owner")).String(),
						URI:   "https://example.com/token/1",
						Data:  "metadata1",
					},
					{
						Id:    "2",
						Name:  "token2",
						Owner: sdk.AccAddress([]byte("owner")).String(),
						URI:   "https://example.com/token/2",
						Data:  "metadata2",
					},
				},
			},
			before: func() error {
				return v160.MigrateStore(ctx, storeKey, cdc, nftKeeper.IssueDenom, nftKeeper.MintNFT)
			},
			after: func(col types.Collection) error {
				expCol, err := nftKeeper.GetCollection(ctx, col.Denom.Id)
				require.NoError(t, err)
				require.EqualValues(t, col, expCol)
				require.EqualValues(t, len(col.NFTs), nftKeeper.GetTotalSupply(ctx, col.Denom.Id))
				return nil
			},
		},
	}

	for _, c := range cases {
		err := legacy.importCollection(c.col)
		require.NoError(t, err)

		err = c.before()
		require.NoError(t, err)

		err = c.after(c.col)
		require.NoError(t, err)
	}
}

func (legacy legacyKeeper) importCollection(col types.Collection) error {
	err := legacy.setDenom(col.Denom)
	if err != nil {
		return err
	}
	for _, nft := range col.NFTs {
		owner, err := sdk.AccAddressFromBech32(nft.Owner)
		if err != nil {
			return err
		}

		err = legacy.mintNFT(col.Denom.Id, nft.Id, nft.Name, nft.URI, nft.Data, owner)
		if err != nil {
			return err
		}
	}
	return nil
}

func (legacy legacyKeeper) mintNFT(denomID,
	tokenID,
	tokenNm,
	tokenURI,
	tokenData string,
	owner sdk.AccAddress,
) error {
	legacy.setNFT(denomID,
		types.NewBaseNFT(
			tokenID,
			tokenNm,
			owner,
			tokenURI,
			tokenData,
		),
	)
	legacy.setOwner(denomID, tokenID, owner)
	legacy.increaseSupply(denomID)
	return nil
}

func (legacy legacyKeeper) setDenom(denom types.Denom) error {
	store := legacy.ctx.KVStore(legacy.storeKey)
	bz := legacy.cdc.MustMarshal(&denom)
	store.Set(v160.KeyDenomID(denom.Id), bz)
	store.Set(v160.KeyDenomName(denom.Name), []byte(denom.Id))
	return nil
}

func (legacy legacyKeeper) setNFT(
	denomID string,
	nft types.BaseNFT,
) {
	store := legacy.ctx.KVStore(legacy.storeKey)

	bz := legacy.cdc.MustMarshal(&nft)
	store.Set(v160.KeyNFT(denomID, nft.GetID()), bz)
}

func (legacy legacyKeeper) setOwner(
	denomID, tokenID string,
	owner sdk.AccAddress) {
	store := legacy.ctx.KVStore(legacy.storeKey)
	bz := types.MustMarshalTokenID(legacy.cdc, tokenID)
	store.Set(v160.KeyOwner(owner, denomID, tokenID), bz)
}

func (legacy legacyKeeper) increaseSupply(denomID string) {
	supply := legacy.getTotalSupply(denomID)
	supply++

	store := legacy.ctx.KVStore(legacy.storeKey)
	bz := types.MustMarshalSupply(legacy.cdc, supply)
	store.Set(v160.KeyCollection(denomID), bz)
}

// GetTotalSupply returns the number of NFTs by the specified denom ID
func (legacy legacyKeeper) getTotalSupply(denomID string) uint64 {
	store := legacy.ctx.KVStore(legacy.storeKey)
	bz := store.Get(v160.KeyCollection(denomID))
	if len(bz) == 0 {
		return 0
	}
	return types.MustUnMarshalSupply(legacy.cdc, bz)
}
