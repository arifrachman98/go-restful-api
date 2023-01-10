package web

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
