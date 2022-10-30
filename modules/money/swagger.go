package money

// swagger:route GET /auth/moneys money MoneysIndex
// Get money operations list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/moneys/label/{label} money MoneysShowByLabel
// Show a money operation.
//
// parameters:
//	+ name: label
//    in: path
//    type: string
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route POST /auth/moneys money MoneysCreate
// Create a new money operation.
//
// parameters:
//	+ name: money
//    in: body
//    type: Money
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess
