package planning

// swagger:route GET /auth/plannings planning PlanningsIndex
// Get plannings list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/plannings/{id} planning PlanningsShow
// Show an planning.
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

// swagger:route POST /auth/plannings planning PlanningsCreate
// Create a new planning.
//
// parameters:
//	+ name: planning
//    in: body
//    type: Planning
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route PUT /auth/plannings/{id} planning PlanningsUpdate
// Update an planning.
//
// parameters:
//	+ name: id
//    in: path
//    type: integer
//    required: true
//	+ name: planning
//    in: body
//    type: Planning
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route DELETE /auth/plannings/{id} planning PlanningsDelete
// Delete an planning.
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

// swagger:route GET /auth/planning_attendees planning_attendee AttendeesIndex
// Get attendees list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/planning_attendees/id_planning/{id_planning} planning_attendee AttendeesShowByIdPlanning
// Show an attendee by planning id.
//
// parameters:
//	+ name: id_planning
//    in: path
//    type: integer
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/planning_attendees/id_adherent/{id_adherent} planning_attendee AttendeesShowByIdAdherent
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

// swagger:route POST /auth/planning_attendees planning_attendee AttendeesCreate
// Create a new attendee.
//
// parameters:
//	+ name: attendee
//    in: body
//    type: Planning_Attendee
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route PUT /auth/planning_attendees/{id} planning_attendee AttendeesUpdate
// Update an attendee.
//
// parameters:
//	+ name: id
//    in: path
//    type: integer
//    required: true
//	+ name: attendee
//    in: body
//    type: Planning_Attendee
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route DELETE /auth/planning_attendees/{id} planning_attendee AttendeesDelete
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
