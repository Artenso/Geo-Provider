package model

import "github.com/Artenso/Geo-Provider/internal/model"

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type ResponseAddress struct {
	Addresses []*model.Address `json:"addresses"`
}
