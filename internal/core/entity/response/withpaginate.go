package response

type ResponseWithPagination struct {
	Items  interface{} `json:"items"`
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
}
