package docs

// swagger:route GET /adherents adherent AdherentsIndex
// Get adherents list.
//
// security:
// - apiKey: []
//
// responses:
//	401: CommonError
//	200: CommonSuccess

// swagger:route POST /adherents adherent AdherentsCreate
// Create a new adherent.
//
// parameters:
// + name: adherent
//   in: body
//   type: Adherent
//   required: true
//
// security:
// - apiKey: []
//
// responses:
//	401: CommonError
//	200: CommonSuccess

// swagger:route  GET /adherents/{id} adherent AdherentsShow
// Show an adherent.
//
// consumes:
//   - application/x-www-form-urlencoded
//
// parameters:
// + name: id
//   in: path
//   type: integer
//   required: true
//
// security:
// - apiKey: []
//
// responses:
//	401: CommonError
//	200: CommonSuccess

// swagger:route  PUT /adherents/{id} adherent AdherentsUpdate
// Update an adherent.
//
// consumes:
//   - application/x-www-form-urlencoded
//
// parameters:
// + name: id
//   in: path
//   type: integer
//   required: true
// + name: adherent
//   in: body
//   type: Adherent
//   required: true
//
// security:
// - apiKey: []
//
// responses:
//	401: CommonError
//	200: CommonSuccess

// swagger:route  DELETE /adherents/{id} adherent AdherentsDelete
// Delete an adherent.
//
// consumes:
//   - application/x-www-form-urlencoded
//
// parameters:
// + name: id
//   in: path
//   type: integer
//   required: true
//
// security:
// - apiKey: []
//
// responses:
//	401: CommonError
//	200: CommonSuccess
