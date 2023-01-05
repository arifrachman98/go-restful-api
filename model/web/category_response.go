package web

type CategoryResponse struct {
	Id   int
	Name string
}

type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}
