package mobwithdrawalfee

import (
	"mavianceutililtyservice/internal/domain/mobwithdrawalfee"
)

type repository struct {
}

func NewRepository() mobwithdrawalfee.Repository {
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

func (r *repository) GetListFeesForMobileMoneyProvider() {
	// Domain interface implementation
}

func (r *repository) GetCalculateWithdrawalFees() {
	// Domain interface implementation

}

func (r *repository) ListCalculateWithdrawalFees() {
	// Domain interface implementation

}
