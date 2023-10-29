package model

type ResponseDataCreate struct {
	Status bool `json:"status"`
}

type CreateResponse struct {
	Message string             `json:"message"`
	Data    ResponseDataCreate `json:"data"`
}

type DetailResponse struct {
	Message string `json:"message"`
	Data    *Task  `json:"data"`
}

type ListData struct {
	Data       []Task
	Pagination Pagination
}

type ListResponse struct {
	Message string
	Data    *Pagination
}

type UpdateResponse struct {
	Message string             `json:"message"`
	Data    ResponseDataCreate `json:"data"`
}

type DeleteResponse struct {
	Message string             `json:"message"`
	Data    ResponseDataCreate `json:"data"`
}
