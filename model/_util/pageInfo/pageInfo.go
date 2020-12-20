package pageInfo

type Request struct {
	Current int `json:"current_page"`
	Size    int `json:"page_size"`
	Skip    int `json:"skip"`
}

type Response struct {
	TotalPage   int `json:"total_page"`
	TotalCount  int `json:"total_count"`
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
}
