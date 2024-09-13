package client

import (
	"context"
	"io"
	"net/rpc"
	"net/rpc/jsonrpc"

	cli "github.com/Artenso/Geo-Provider/client"
	"github.com/Artenso/Geo-Provider/internal/model"
	jsonRPCmodel "github.com/Artenso/Geo-Provider/internal/model/json_rpc_geo_provider"
)

type client struct {
	client *rpc.Client
}

func NewJSONrpcClient(conn io.ReadWriteCloser) cli.Client {
	return &client{
		client: jsonrpc.NewClient(conn),
	}
}

func (c *client) AddressSearch(ctx context.Context, input string) ([]*model.Address, error) {
	req := &jsonRPCmodel.RequestAddressSearch{
		Query: input,
	}

	resp := &jsonRPCmodel.ResponseAddress{}

	if err := c.client.Call("Controller.AddressSearch", req, resp); err != nil {
		return nil, err
	}

	return resp.Addresses, nil
}

func (c *client) GeoCode(ctx context.Context, lat, lng string) ([]*model.Address, error) {
	req := &jsonRPCmodel.RequestAddressGeocode{
		Lat: lat,
		Lng: lng,
	}

	resp := &jsonRPCmodel.ResponseAddress{}

	if err := c.client.Call("Controller.GeoCode", req, resp); err != nil {
		return nil, err
	}

	return resp.Addresses, nil
}
