package models

type Node struct {
	Id     uint
	Weight uint
}

//easyjson:json
type UserInfo struct {
	Id       uint   `json:"user_id"`
	Messages string `json:"messages"`
}

//easyjson:json
type Info struct {
	Graph    []UserInfo `json:"graph"`
	MinValue uint       `json:"min_value"`
	AvgValue float32    `json:"avg_value"`
	MaxValue uint       `json:"max_value"`
	IsEmpty  bool       `json:"is_empty"`
}
