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

// swagger:route GET /auth/event_attendees event_attendee AttendeesIndex
// Get attendees list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route POST /auth/event_attendees event_attendee AttendeesCreate
// Create a new attendee.
//
// parameters:
//	+ name: attendee
//    in: body
//    type: Attendee
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/event_attendees/id_event/{id_event} event_attendee AttendeesShowByIdEvent
// Show an attendee by event id.
//
// parameters:
//	+ name: id_event
//    in: path
//    type: integer
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/event_attendees/id_adherent/{id_adherent} event_attendee AttendeesShowByIdAdherent
// Show an attendee by adherent id.
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

// swagger:route PUT /auth/event_attendees/{id} event_attendee AttendeesUpdate
// Update an attendee.
//
// parameters:
//	+ name: id
//    in: path
//    type: integer
//    required: true
//	+ name: attendee
//    in: body
//    type: Attendee
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route DELETE /auth/event_attendees/{id} event_attendee AttendeesDelete
// Delete an attendee.
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
