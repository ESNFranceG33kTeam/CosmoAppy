// Package documentation CosmoAppy
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// Terms Of Service:
//
// not yet...
//
//	Schemes: http, https
//	BasePath: /
//	Version: latest
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
	// example: 200
	// required: true
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	// example: successed
	Message string `json:"message"`
}

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int64
	// example: 400
	// required: true
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	// example: error
	// required: true
	Message string `json:"message"`
}
