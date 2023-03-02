package report

// swagger:route GET /auth/reports report ReportsIndex
// Get reports list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/reports/{id} report ReportsShowById
// Show an report by the id.
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

// swagger:route POST /auth/reports report ReportsCreate
// Create a new report.
//
// parameters:
//	+ name: report
//    in: body
//    type: Report
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route PUT /auth/reports/{id} report ReportsUpdate
// Update a report.
//
// parameters:
//	+ name: id
//    in: path
//    type: integer
//    required: true
//	+ name: report
//    in: body
//    type: Report
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route DELETE /auth/reports/{id} report ReportsDelete
// Delete a report.
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
