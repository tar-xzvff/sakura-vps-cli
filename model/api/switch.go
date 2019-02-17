package api

import "sakura-vps-cli/model/object"

type ResponseSwitches struct {
	Switches []object.Switch `json:"switches"`
	Meta struct {
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
		PerPage    int `json:"per_page"`
		Page       int `json:"page"`
	} `json:"meta"`
}

type ResponseSwitch struct {
	Switch object.Switch `json:"switches"`
}