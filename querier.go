// Code generated by sqlc. DO NOT EDIT.

package main

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAsset(ctx context.Context, arg CreateAssetParams) error
	CreateAssetHost(ctx context.Context, arg CreateAssetHostParams) error
	CreateHost(ctx context.Context, arg CreateHostParams) error
	DeleteAsset(ctx context.Context, id string) error
	DeleteAssetHost(ctx context.Context, id string) error
	DeleteHost(ctx context.Context, id string) error
	GetAllAssetHosts(ctx context.Context) ([]AssetHost, error)
	GetAllAssets(ctx context.Context) ([]Asset, error)
	GetAllHosts(ctx context.Context) ([]Host, error)
	GetAsset(ctx context.Context, id string) (Asset, error)
	GetHost(ctx context.Context, id string) (Host, error)
	GetHostsForAsset(ctx context.Context, assetID sql.NullString) ([]GetHostsForAssetRow, error)
	UpdateHostAddress(ctx context.Context, arg UpdateHostAddressParams) error
}

var _ Querier = (*Queries)(nil)