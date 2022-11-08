// Package documentation CosmoAppy.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// Terms Of Service:
//
// not yet...
//
//	Schemes: http, https
//	Host: 127.0.0.1:8080
//	BasePath: /
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- key:
//
//	SecurityDefinitions:
//	key:
//	     name: X-Session-Token
//	     type: apiKey
//	     in: header
//
// swagger:meta
package docs

// swagger:model CommonSuccess
type CommonSuccess struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}
