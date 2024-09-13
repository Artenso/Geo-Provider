package controller

import (
	"context"

	converter "github.com/Artenso/Geo-Provider/internal/converter/grpc_geo_provider"
	"github.com/Artenso/Geo-Provider/internal/service"
	desc "github.com/Artenso/Geo-Provider/pkg/grpc_geo_provider"
)

type Controller struct {
	desc.UnimplementedGeoProviderServer

	service service.GeoProvider
}

func NewController(service service.GeoProvider) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) AddressSearch(ctx context.Context, req *desc.AddressSearchRequest) (*desc.AddressSearchResponse, error) {
	addresses, err := c.service.AddressSearch(req.Input)
	if err != nil {
		return nil, err
	}

	return converter.ToAddressSearchResponse(addresses), nil
}

func (c *Controller) GeoCode(ctx context.Context, req *desc.GeoCodeRequest) (*desc.GeoCodeResponse, error) {
	addresses, err := c.service.GeoCode(req.Lat, req.Lng)
	if err != nil {
		return nil, err
	}

	return converter.ToGeoCodeResponse(addresses), nil
}
