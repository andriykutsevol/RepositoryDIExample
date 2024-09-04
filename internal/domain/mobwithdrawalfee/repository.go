package mobwithdrawalfee

type Repository interface {

	//---------------------------------
	// Management API
	//---------------------------------

	// API to add new entries
	Add() // TODO: What is an "entry" in this case?
	// API to update existing entries
	Update() // TODO: What is an "entry" in this case?

	//---------------------------------
	// Functional API
	//---------------------------------

	// TODO: Filters, pagination

	// Provide a set of fees for a selected mobile money provider
	GetListFeesForMobileMoneyProvider()

	// Calculate and return the withdrawal fees for given amount
	GetCalculateWithdrawalFees()

	// Calculate and return the withdrawal fees for a list of amounts.
	ListCalculateWithdrawalFees()

	//---------------------------------
	// For requests with filters we use "Retrieve" prefix.

}
