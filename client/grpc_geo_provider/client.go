package client

import (
	"context"

	cli "github.com/Artenso/Geo-Provider/client"
	"github.com/Artenso/Geo-Provider/internal/model"
	desc "github.com/Artenso/Geo-Provider/pkg/grpc_geo_provider"
	"google.golang.org/grpc"
)

type client struct {
	client desc.GeoProviderClient
}

func NewGRPCclient(conn *grpc.ClientConn) cli.Client {
	return &client{
		client: desc.NewGeoProviderClient(conn),
	}
}

func (c *client) AddressSearch(ctx context.Context, input string) ([]*model.Address, error) {
	req := &desc.AddressSearchRequest{
		Input: input,
	}

	resp, err := c.client.AddressSearch(ctx, req)
	if err != nil {
		return nil, err
	}

	addresses := make([]*model.Address, 0, len(resp.Addresses))

	for _, descAddr := range resp.Addresses {
		addr := &model.Address{
			City:   descAddr.City,
			Street: descAddr.Street,
			House:  descAddr.House,
			Lat:    descAddr.Lat,
			Lon:    descAddr.Lon,
		}

		addresses = append(addresses, addr)
	}

	return addresses, nil
}

func (c *client) GeoCode(ctx context.Context, lat, lng string) ([]*model.Address, error) {
	req := &desc.GeoCodeRequest{
		Lat: lat,
		Lng: lng,
	}

	resp, err := c.client.GeoCode(ctx, req)
	if err != nil {
		return nil, err
	}

	addresses := make([]*model.Address, 0, len(resp.Addresses))

	for _, descAddr := range resp.Addresses {
		addr := &model.Address{
			City:   descAddr.City,
			Street: descAddr.Street,
			House:  descAddr.House,
			Lat:    descAddr.Lat,
			Lon:    descAddr.Lon,
		}

		addresses = append(addresses, addr)
	}

	return addresses, nil
}
