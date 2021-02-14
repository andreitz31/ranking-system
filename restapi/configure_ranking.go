// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"ranking/handlers"
	"ranking/restapi/operations"
	"ranking/restapi/operations/top"
)

//go:generate swagger generate server --target ..\..\ranking --name Ranking --spec ..\api\swagger.yaml --principal interface{}

func configureFlags(api *operations.RankingAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RankingAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TopTopListHandler = handlers.NewTopPlayersListHandler()
	api.TopAddPlayerHandler = handlers.NewAddPlayerScoreHandler()
	api.TopTopListRelativeHandler = handlers.NewTopPlayersRelativeListHandler()

	if api.TopTopListHandler == nil {
		api.TopTopListHandler = top.TopListHandlerFunc(func(params top.TopListParams) middleware.Responder {
			return middleware.NotImplemented("operation top.TopList has not yet been implemented")
		})
	}
	if api.TopTopListRelativeHandler == nil {
		api.TopTopListRelativeHandler = top.TopListRelativeHandlerFunc(func(params top.TopListRelativeParams) middleware.Responder {
			return middleware.NotImplemented("operation top.TopListRelative has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
