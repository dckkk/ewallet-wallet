package cmd

import (
	"ewallet-wallet/external"
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	r.GET("/health", d.HealthcheckAPI.HealthcheckHandlerHTTP)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletAPI.Create)
	walletV1.PUT("/balance/credit", d.MiddlewareValidateToken, d.WalletAPI.CreditBalance)
	walletV1.PUT("/balance/debit", d.MiddlewareValidateToken, d.WalletAPI.DebitBalance)
	walletV1.GET("/balance", d.MiddlewareValidateToken, d.WalletAPI.GetBalance)
	walletV1.GET("/history", d.MiddlewareValidateToken, d.WalletAPI.GetWalletHistory)

	exWalletV1 := walletV1.Group("/ex")
	exWalletV1.Use(d.MiddlewareSignatureValidation)
	exWalletV1.POST("/link", d.WalletAPI.CreateWalletLink)
	exWalletV1.PUT("/link/:wallet_id/confirmation", d.WalletAPI.WalletLinkConfirmation)
	exWalletV1.DELETE("/:wallet_id/unlink", d.WalletAPI.WalletUnlink)
	exWalletV1.GET("/:wallet_id/balance", d.WalletAPI.ExGetBalance)
	exWalletV1.POST("/transaction", d.WalletAPI.ExternalTransaction)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckAPI
	WalletAPI      interfaces.IWalletAPI
	External       interfaces.IExternal
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	walletRepo := &repository.WalletRepo{
		DB: helpers.DB,
	}

	walletSvc := &services.WalletService{
		WalletRepo: walletRepo,
	}
	walletAPI := &api.WalletAPI{
		WalletService: walletSvc,
	}

	external := &external.External{}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		WalletAPI:      walletAPI,
		External:       external,
	}
}
