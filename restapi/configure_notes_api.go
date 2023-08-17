// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/almalii/swagger-contracts/restapi/operations"
	"github.com/almalii/swagger-contracts/restapi/operations/auth"
	"github.com/almalii/swagger-contracts/restapi/operations/notes"
	"github.com/almalii/swagger-contracts/restapi/operations/users"
)

//go:generate swagger generate server --target ../../swagger-contracts --name NotesAPI --spec ../docs/swagger.json --principal interface{}

func configureFlags(api *operations.NotesAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.NotesAPIAPI) http.Handler {
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

	// Applies when the "Authorization" header is set
	if api.JWTAuthAuth == nil {
		api.JWTAuthAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (JWTAuth) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.NotesDeleteNotesIDHandler == nil {
		api.NotesDeleteNotesIDHandler = notes.DeleteNotesIDHandlerFunc(func(params notes.DeleteNotesIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation notes.DeleteNotesID has not yet been implemented")
		})
	}
	if api.UsersDeleteUsersHandler == nil {
		api.UsersDeleteUsersHandler = users.DeleteUsersHandlerFunc(func(params users.DeleteUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteUsers has not yet been implemented")
		})
	}
	if api.NotesGetNotesHandler == nil {
		api.NotesGetNotesHandler = notes.GetNotesHandlerFunc(func(params notes.GetNotesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation notes.GetNotes has not yet been implemented")
		})
	}
	if api.NotesGetNotesIDHandler == nil {
		api.NotesGetNotesIDHandler = notes.GetNotesIDHandlerFunc(func(params notes.GetNotesIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation notes.GetNotesID has not yet been implemented")
		})
	}
	if api.UsersGetUsersHandler == nil {
		api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
		})
	}
	if api.NotesPatchNotesIDHandler == nil {
		api.NotesPatchNotesIDHandler = notes.PatchNotesIDHandlerFunc(func(params notes.PatchNotesIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation notes.PatchNotesID has not yet been implemented")
		})
	}
	if api.UsersPatchUsersHandler == nil {
		api.UsersPatchUsersHandler = users.PatchUsersHandlerFunc(func(params users.PatchUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.PatchUsers has not yet been implemented")
		})
	}
	if api.AuthPostAuthLoginHandler == nil {
		api.AuthPostAuthLoginHandler = auth.PostAuthLoginHandlerFunc(func(params auth.PostAuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthLogin has not yet been implemented")
		})
	}
	if api.AuthPostAuthRegisterHandler == nil {
		api.AuthPostAuthRegisterHandler = auth.PostAuthRegisterHandlerFunc(func(params auth.PostAuthRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthRegister has not yet been implemented")
		})
	}
	if api.NotesPostNotesHandler == nil {
		api.NotesPostNotesHandler = notes.PostNotesHandlerFunc(func(params notes.PostNotesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation notes.PostNotes has not yet been implemented")
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
