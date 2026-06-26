package types

// QueryDenomAdminRequest is the request type for the Query/DenomAdmin RPC method.
type QueryDenomAdminRequest struct {
	Denom string
}

// QueryDenomAdminResponse is the response type for the Query/DenomAdmin RPC method.
type QueryDenomAdminResponse struct {
	Admin string
}

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct{}

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	Params Params
}
