package lib

import (
	"encoding/json"
	"gopkg.in/resty.v1"
	"../model/api"
	"../model/object"
	"strconv"
)

type API struct {
	ApiUrl string
	Token  string
}

func NewAPI(url string, token string) *API {
	a := new(API)
	a.ApiUrl = url
	a.Token = token
	return a
}

func (t *API)GetServers() []object.Server {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "servers")
	var Response api.ResponseServers
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.Servers
	}
}

func (t *API)GetServer(serverId int) object.Server {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "servers/" + strconv.Itoa(serverId))
	var Response api.ResponseServer
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.Server
	}
}

func (t *API)GetSwitches() []object.Switch {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "switches")
	var Response api.ResponseSwitches
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.Switches
	}
}

func (t *API)GetSwitch(switchId int) object.Switch {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "switches/" + strconv.Itoa(switchId))
	var Response api.ResponseSwitch
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.Switch
	}
}

func (t *API)GetPublicScripts() []object.PublicScript {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "startupscripts/public")
	var Response api.ResponsePublicScripts
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.PublicScripts
	}
}

func (t *API)GetPublicScript(scriptId string) object.PublicScript {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "startupscripts/public/" + scriptId)
	var Response api.ResponsePublicScript
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.PublicScript
	}
}

func (t *API) GetCustomScripts() []object.CustomScript {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "startupscripts/custom")
	var Response api.ResponseCustomScripts
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.CustomScripts
	}
}

func (t *API) GetCustomScript(scriptId string) object.CustomScript {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "startupscripts/custom/" + scriptId)
	var Response api.ResponseCustomScript
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.CustomScript
	}
}

func (t *API) GetMonitorings() []object.Monitoring {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "monitorings")
	var Response api.ResponseMonitorings
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.Monitorings
	}
}

func (t *API) GetMonitoring(monitoringId int) object.Monitoring {
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + t.Token).
		Get(t.ApiUrl + "monitorings/" + strconv.Itoa(monitoringId))
	var Response api.ResponseMonitoring
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		return Response.Monitoring
	}
}