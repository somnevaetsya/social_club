package models

//easyjson:json
type Created struct {
	CreatedInfo bool `json:"created"`
}

//easyjson:json
type CustomError struct {
	CustomErr string `json:"error"`
}
