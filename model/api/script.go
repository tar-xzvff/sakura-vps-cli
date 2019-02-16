package api

import "sakura-vps-cli/model/object"

type ResponsePublicScripts struct {
	PublicScripts []object.PublicScript `json:"public_scripts"`
}

type ResponsePublicScript struct {
	PublicScript object.PublicScript `json:"public_scripts"`
}

type ResponseCustomScripts struct {
	CustomScripts []object.CustomScript `json:"custom_scripts"`
}

type ResponseCustomScript struct {
	CustomScript object.CustomScript `json:"custom_scripts"`
}