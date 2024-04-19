package common

type successResponse struct {
	Data   interface{} `json:"data"` //interface{}: means any data type
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func SuccessResponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging, Filter: filter}
}

func SingleSuccessResponse(data interface{}) *successResponse {
	return SuccessResponse(data, nil, nil)
}
