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

// swagger:route GET /auth/moneys/stats/monthly money MonthlyStatsIndex
// Get stats list.
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route GET /auth/moneys/stats/monthly/archive_date/{archive_date} money MonthlyStatsShowByDate
// Show stats for the archive_date.
//
// parameters:
//	+ name: archive_date
//    in: path
//    type: string
//    required: true
//
// responses:
//	400: CommonError
//	200: CommonSuccess

// swagger:route POST /auth/moneys/stats/monthly/create money MonthlyStatCreate
// Generate stats of the past month.
//
// responses:
//	400: CommonError
//	200: CommonSuccess
