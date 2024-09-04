package app

import (
	"mavianceutililtyservice/internal/app/application"
	adminRegionsInfra "mavianceutililtyservice/internal/infrastructure/inmemory/administrativeregion"
	"mavianceutililtyservice/internal/presentation/http/handler"
	"mavianceutililtyservice/internal/presentation/http/router"
)

func Run() {

	//=========================================================================
	// Dependency injection for Infrastructure layer (the database for example). No depencencies yet.
	// Infrastructure layer packages implementing corresponding domain's interfaces.
	// The infrastructure layer imports domain only, no other layers.

	adminRegionRepo := adminRegionsInfra.NewRepository()

	//=========================================================================
	//Dependency injections for Application layer

	// Here we usually inject infrastructure layer. For example the dataabse. The mock object?
	// On the other hand, the applications define it's own interface, to contract with the presentation layer (http).
	// But it can also implement domain interface, but we'll lose application abstraction layer in this case.
	// The apllication layer imports domain only, no other layers.

	// We're allowed to inject any application that implements corresponding interface. mock object?

	administrativeregionApp := application.NewAdminRegion(adminRegionRepo)
	_ = administrativeregionApp

	//=========================================================================
	//Dependency injections for Presentation layer
	// We're allowed to inject any application that implements corresponding interface. mock object?

	administrativeregionHandler := handler.NewAdminRegion(administrativeregionApp)
	_ = administrativeregionHandler

	// The same for the mobileMoneyWithdrawalFeesHandler, etc...
	// mobileMoneyWithdrawalFeesHandler := handler.WithdrawalFees(mobileMoneyWithdrawalFeesApp)

	// !!! HENCE. The HTTP router is as loosely coupled as possible. !!!

	// Now just inject
	routerRouter := router.NewRouter(
		//infrarepos.AuthRepository,
		administrativeregionHandler,
		//mobileMoneyWithdrawalFeesHandler,
	)
	_ = routerRouter

}
