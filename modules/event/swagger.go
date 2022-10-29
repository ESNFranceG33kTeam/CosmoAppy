package event

// swagger:route GET /auth/events event EventsIndex
// Get events list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route POST /auth/events event EventsCreate
// Create a new event.
//
// parameters:
//	+ name: event
//    in: body
//    type: Event
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/events/{id} event EventsShow
// Show an event.
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

// swagger:route PUT /auth/events/{id} event EventsUpdate
// Update an event.
//
// parameters:
//	+ name: id
//    in: path
//    type: integer
//    required: true
//	+ name: event
//    in: body
//    type: Event
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route DELETE /auth/events/{id} event EventsDelete
// Delete an event.
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
