package client

import (
	"context"

	"github.com/Artenso/Geo-Provider/internal/model"
)

type Client interface {
	AddressSearch(ctx context.Context, input string) ([]*model.Address, error)
	GeoCode(ctx context.Context, lat, lng string) ([]*model.Address, error)
}
