package api

import "sakura-vps-cli/model/object"

type ResponseMonitorings struct {
	Monitorings []object.Monitoring `json:"monitorings"`
	Meta struct {
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
		PerPage    int `json:"per_page"`
		Page       int `json:"page"`
	} `json:"meta"`
}

type ResponseMonitoring struct {
	Monitoring object.Monitoring `json:"monitorings"`
}