package application

import (
	domain_adminregion "mavianceutililtyservice/internal/domain/administrativeregion"
)

type AdminRegion interface {
	Add()
	Update()
	// TODO: this implements filters
	Query()
}

type adminRegionApp struct {
	adminregion domain_adminregion.Repository
}

func NewAdminRegion(adminregion domain_adminregion.Repository) AdminRegion {
	return &adminRegionApp{
		adminregion: adminregion,
	}
}

func (a *adminRegionApp) Add() {

}

func (a *adminRegionApp) Update() {

}

func (a *adminRegionApp) Query() {

}
