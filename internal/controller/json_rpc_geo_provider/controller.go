package controller

import (
	model "github.com/Artenso/Geo-Provider/internal/model/json_rpc_geo_provider"
	"github.com/Artenso/Geo-Provider/internal/service"
)

type Controller struct {
	service service.GeoProvider
}

func NewController(service service.GeoProvider) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) AddressSearch(req *model.RequestAddressSearch, resp *model.ResponseAddress) error {
	addresses, err := c.service.AddressSearch(req.Query)
	if err != nil {
		return err
	}

	resp.Addresses = addresses

	return nil
}

func (c *Controller) GeoCode(req *model.RequestAddressGeocode, resp *model.ResponseAddress) error {
	addresses, err := c.service.GeoCode(req.Lat, req.Lng)
	if err != nil {
		return err
	}

	resp.Addresses = addresses

	return nil
}
