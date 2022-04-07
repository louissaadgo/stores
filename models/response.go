package models

type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
