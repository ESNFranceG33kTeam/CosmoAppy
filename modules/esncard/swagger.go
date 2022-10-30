package esncard

// swagger:route GET /auth/esncards esncard ESNcardsIndex
// Get esncards list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/esncards/id_adherent/{id_adherent} esncard ESNcardsShowByIdAdherent
// Show an esncard by the id_adherent.
//
// parameters:
//	+ name: id_adherent
//    in: path
//    type: integer
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/esncards/esncard/{esncard} esncard ESNcardsShowByESNcard
// Show an esncard by the esncard code.
//
// parameters:
//	+ name: esncard
//    in: path
//    type: string
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route POST /auth/esncards esncard ESNcardsCreate
// Create a new esncard.
//
// parameters:
//	+ name: esncard
//    in: body
//    type: ESNcard
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route DELETE /auth/esncards/{id} esncard ESNcardsDelete
// Delete an esncard.
//
// parameters:
//	+ name: id
//    in: path
//    type: integer
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess
