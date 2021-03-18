// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/t1cg/VL/api/backend/restapi/operations"
	"github.com/t1cg/VL/api/backend/restapi/operations/login"
	"github.com/t1cg/VL/api/backend/restapi/operations/record"
	"github.com/t1cg/VL/api/backend/restapi/operations/records"
	"github.com/t1cg/VL/api/backend/restapi/operations/register"
)

//go:generate swagger generate server --target ../../backend --name Medicare --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.MedicareAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MedicareAPI) http.Handler {
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

	//if api.RecordGetRecordHandler == nil {
		api.RecordGetRecordHandler = record.GetRecordHandlerFunc(func(params record.GetRecordParams) middleware.Responder {
			return record.GetFunc(params.ID)
		})
	//}
	//if api.RecordsGetRecordsHandler == nil {
		api.RecordsGetRecordsHandler = records.GetRecordsHandlerFunc(func(params records.GetRecordsParams) middleware.Responder {
			return records.GetFunc(params.City, params.State, params.Specialty, params.DrugName)
		})
	//}
	//if api.LoginPostLoginHandler == nil {
		api.LoginPostLoginHandler = login.PostLoginHandlerFunc(func(params login.PostLoginParams) middleware.Responder {
			return login.PostFunc(params.Login)
		})
	//}
	//if api.RegisterPostRegisterHandler == nil {
		api.RegisterPostRegisterHandler = register.PostRegisterHandlerFunc(func(params register.PostRegisterParams) middleware.Responder {
			return register.PostFunc(params.Register)
		})
	//}

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
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler = c.Handler(handler)
	
	return handler
}
