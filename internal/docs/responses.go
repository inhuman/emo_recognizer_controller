package docs

type commonErrorResponse struct {
	Code    int           `json:"code"`
	Error   string        `json:"error"`
	Details []interface{} `json:"details"`
}

// Success (200)
// swagger:response OK
type okResponse struct{}

// Bad data (400)
// swagger:response badDataResponse
type badDataResponse struct {
	// in:body
	Body commonErrorResponse
}

// Not authorized (401)
// swagger:response unauthorizedResponse
type unauthorizedResponse struct {
	// in:body
	Body commonErrorResponse
}

// Not found (404)
// swagger:response notFoundResponse
type notFoundResponse struct {
	// in:body
	Body commonErrorResponse
}

// Too many requests (429)
// swagger:response tooManyRequestsResponse
type tooManyRequestsResponse struct {
	// in:body
	Body commonErrorResponse
}

// Internal server error (500)
// swagger:response internalErrorResponse
type internalErrorResponse struct {
	// in:body
	Body commonErrorResponse
}
