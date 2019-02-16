package object

type Switch struct {
	ID                  int           `json:"id"`
	Zone                int           `json:"zone"`
	Name                string        `json:"name"`
	Description         string        `json:"description"`
	Interfaces          []int         `json:"interfaces"`
	ServerInterfaces    []int         `json:"server_interfaces"`
	ApplianceInterfaces []interface{} `json:"appliance_interfaces"`
	InterfacesCount     int           `json:"interfaces_count"`
	IsHybridConnection  bool          `json:"is_hybrid_connection"`
	SwitchCode          string        `json:"switch_code"`
	HybridConnection    interface{}   `json:"hybrid_connection"`
	CanUseHybrid        bool          `json:"can_use_hybrid"`
}