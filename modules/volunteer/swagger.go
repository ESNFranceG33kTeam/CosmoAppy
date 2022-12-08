package volunteer

// swagger:route GET /auth/volunteers volunteer VolunteersIndex
// Get volunteers list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/volunteers/{id} volunteer VolunteersShowById
// Show an volunteer by the id.
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

// swagger:route POST /auth/volunteers volunteer VolunteersCreate
// Create a new volunteer.
//
// parameters:
//	+ name: volunteer
//    in: body
//    type: Volunteer
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route PUT /auth/volunteers/{id} volunteer VolunteersUpdate
// Update a volunteer.
//
// parameters:
//	+ name: id
//    in: path
//    type: integer
//    required: true
//	+ name: volunteer
//    in: body
//    type: Volunteer
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route DELETE /auth/volunteers/{id} volunteer VolunteersDelete
// Delete a volunteer.
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
