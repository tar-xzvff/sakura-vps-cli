package object

import (
	"log"
	"io/ioutil"
	"strconv"
	"gopkg.in/resty.v1"
	"encoding/json"
	"os/user"
)

type Server struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceCd   string `json:"service_cd"`
	Status      int    `json:"status"`
	CPUCores    int    `json:"cpu_cores"`
	MemoryBytes int    `json:"memory_bytes"`
	Zone        int    `json:"zone"`
	Power       string `json:"power"`
	Interfaces  []int  `json:"interfaces"`
	Storage     []struct {
	Port      int `json:"port"`
	Type      int `json:"type"`
	SizeBytes int `json:"size_bytes"`
	} `json:"storage"`
	RaidStatus interface{}   `json:"raid_status"`
	Options    []interface{} `json:"options"`
	HostType   string        `json:"host_type"`
	Ipv4       []struct {
	Address     string   `json:"address"`
	Netmask     string   `json:"netmask"`
	Ptr         string   `json:"ptr"`
	Hostname    string   `json:"hostname"`
	Gateway     string   `json:"gateway"`
	NameServers []string `json:"nameservers"`
	} `json:"ipv4"`
	Ipv6 []struct {
	Address     string   `json:"address"`
	Ptr         string   `json:"ptr"`
	PrefixLen   int      `json:"prefixlen"`
	Hostname    string   `json:"hostname"`
	Gateway     string   `json:"gateway"`
	NameServers []string `json:"nameservers"`
	} `json:"ipv6"`
	Limitation int    `json:"limitation"`
	Version    string `json:"version"`
	DNSRecords int    `json:"dns_records"`
}

type Power struct {
	ID     int    `json:"id"`
	Server int    `json:"server"`
	Status string `json:"status"`
}

type ResponsePowers struct {
	Power Power	`json:"powers"`
}

func LoadToken() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}

	dat, err := ioutil.ReadFile(usr.HomeDir + "/.sakura_vps_token")
	if err != nil {
		log.Fatal( err )
	}

	return string(dat)
}

func (t *Server) GetPower() string {
	var url = "https://secure.sakura.ad.jp/vps/api/v6.5/powers/" + strconv.Itoa(t.ID)
	resp, _ := resty.R().
		SetHeader("Cookie", "SIDv2=" + LoadToken()).
		Get(url)

	var Response ResponsePowers
	er := json.Unmarshal(resp.Body(), &Response)
	if er != nil {
		panic(er)
	} else {
		if resp.StatusCode() == 200 {
			return Response.Power.Status
		}
		return ""
	}
}

func (t *Server) PowerOn() bool {
	var url = "https://secure.sakura.ad.jp/vps/api/v6.5/powers/" + strconv.Itoa(t.ID) + "/start"
	resp, _ := resty.R().
		SetHeaders(map[string]string{
		"X-Requested-With": "XMLHttpRequest",
		"Host": "secure.sakura.ad.jp",
		"Origin": "https://secure.sakura.ad.jp",
		"Content-Type": "application/json",
		"Cookie": "SIDv2=" + LoadToken(),
	}).Post(url)

	if resp.StatusCode() == 202 {
		t.GetPower()
		return true
	} else {
		return false
	}
}

func (t *Server) PowerOff() bool {
	var url = "https://secure.sakura.ad.jp/vps/api/v6.5/powers/" + strconv.Itoa(t.ID) + "/destroy"
	resp, _ := resty.R().
		SetHeaders(map[string]string{
		"X-Requested-With": "XMLHttpRequest",
		"Host": "secure.sakura.ad.jp",
		"Origin": "https://secure.sakura.ad.jp",
		"Content-Type": "application/json",
		"Cookie": "SIDv2=",
	}).Post(url)

	if resp.StatusCode() == 202 {
		t.GetPower()
		return true
	} else {
		return false
	}
}

func (t *Server) Shutdown() bool {
	var url = "https://secure.sakura.ad.jp/vps/api/v6.5/powers/" + strconv.Itoa(t.ID) + "/shutdown"
	resp, _ := resty.R().
		SetHeaders(map[string]string{
		"X-Requested-With": "XMLHttpRequest",
		"Host": "secure.sakura.ad.jp",
		"Origin": "https://secure.sakura.ad.jp",
		"Content-Type": "application/json",
		"Cookie": "SIDv2=" + LoadToken(),
	}).Post(url)

	if resp.StatusCode() == 202 {
		t.GetPower()
		return true
	} else {
		return false
	}
}

func (t *Server) Reset() bool {
	var url = "https://secure.sakura.ad.jp/vps/api/v6.5/powers/" + strconv.Itoa(t.ID) + "/reset"
	resp, _ := resty.R().
		SetHeaders(map[string]string{
		"X-Requested-With": "XMLHttpRequest",
		"Host": "secure.sakura.ad.jp",
		"Origin": "https://secure.sakura.ad.jp",
		"Content-Type": "application/json",
		"Cookie": "SIDv2=" + LoadToken(),
	}).Post(url)

	if resp.StatusCode() == 202 {
		t.GetPower()
		return true
	} else {
		return false
	}
}