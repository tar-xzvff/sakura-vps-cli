package api

import "../object"

type ResponseServers struct {
	Servers []object.Server `json:"servers"`
	Meta struct {
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
		PerPage    int `json:"per_page"`
		Page       int `json:"page"`
	} `json:"meta"`
}

type ResponseServer struct {
	Server object.Server `json:"servers"`
}

