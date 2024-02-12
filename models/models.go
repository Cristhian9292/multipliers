package models

type RequestBody struct {
	Number int `json:"number" example:"65"`
}

type Response struct {
	Message string `json:"message" example:"Saved!"`
}

type ResponseGet struct {
	Message string `json:"message" example:"Type 1"`
}

type ResponseCollection struct {
	Message []int `json:"message" example:"1,2,3"`
}
