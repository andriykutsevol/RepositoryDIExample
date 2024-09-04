package administrativeregion

import (
	"mavianceutililtyservice/internal/domain/administrativeregion"
)

type repository struct {
}

func NewRepository() administrativeregion.Repository {
	return &repository{}
}

//---------------------------------
// Management API
//---------------------------------

func (r *repository) Add() {
	// Domain interface implementation

}

func (r *repository) Update() {
	// Domain interface implementation

}

//---------------------------------
// Functional API
//---------------------------------

func (r *repository) ListRetionsInCoutry() {
	// Domain interface implementation
}

func (r *repository) ListCitiesInRegion() {
	// Domain interface implementation

}

func (r *repository) ListRetionsAndCitiesInCoutry() {
	// Domain interface implementation

}
