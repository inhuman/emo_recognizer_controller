// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations"
	"github.com/inhuman/emo_recognizer_controller/pkg/gen/restapi/operations/job"
)

func configureFlags(api *operations.EmotionsRecognizerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EmotionsRecognizerAPI) http.Handler {
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

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// job.CreateJobMaxParseMemory = 32 << 20

	if api.JobCreateJobHandler == nil {
		api.JobCreateJobHandler = job.CreateJobHandlerFunc(func(params job.CreateJobParams) middleware.Responder {
			return middleware.NotImplemented("operation job.CreateJob has not yet been implemented")
		})
	}
	if api.JobGetJobHandler == nil {
		api.JobGetJobHandler = job.GetJobHandlerFunc(func(params job.GetJobParams) middleware.Responder {
			return middleware.NotImplemented("operation job.GetJob has not yet been implemented")
		})
	}
	if api.JobGetJobsHandler == nil {
		api.JobGetJobsHandler = job.GetJobsHandlerFunc(func(params job.GetJobsParams) middleware.Responder {
			return middleware.NotImplemented("operation job.GetJobs has not yet been implemented")
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
	c := cors.AllowAll()

	return c.Handler(handler)
}
