package docs

// swagger:route GET /auth/adherents adherent AdherentsIndex
// Get adherents list.
//
// responses:
//	401: CommonError
//	200: CommonSuccess

// swagger:route POST /auth/adherents adherent AdherentsCreate
// Create a new adherent.
//
// parameters:
// + name: adherent
//   in: body
//   type: Adherent
//   required: true
//
// responses:
//	401: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/adherents/{id} adherent AdherentsShow
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
// responses:
//	401: CommonError
//	200: CommonSuccess

// swagger:route PUT /auth/adherents/{id} adherent AdherentsUpdate
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
// responses:
//	401: CommonError
//	200: CommonSuccess

// swagger:route DELETE /auth/adherents/{id} adherent AdherentsDelete
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
// responses:
//	401: CommonError
//	200: CommonSuccess
