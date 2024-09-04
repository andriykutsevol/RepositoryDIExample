package handler

import (
	"mavianceutililtyservice/internal/app/application"
)

type AdminRegion interface {
	Add()
	Update()
	SomeAggregateFunction()
}

type adminRegionHandler struct {
	adminRegionApp application.AdminRegion
}

func NewAdminRegion(adminRegionApp application.AdminRegion) AdminRegion {
	return &adminRegionHandler{
		adminRegionApp: adminRegionApp}
}

func (h *adminRegionHandler) Add() {

}

func (h *adminRegionHandler) Update() {

}

func (h *adminRegionHandler) SomeAggregateFunction() {

}
