package model

type Log struct {
	Service string `json:"service"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
