package object

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
