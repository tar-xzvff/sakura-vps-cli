package object

type PublicScript struct {
	ID                     string        `json:"id"`
	Name                   string        `json:"name"`
	Description            string        `json:"description"`
	SupportedDistributions []int         `json:"supported_distributions"`
	ScriptBody             string `json:"script_body"`
	Parameters             []struct {
		Type    string `json:"type"`
		Label   string `json:"label"`
		Name    string `json:"name"`
		Options struct {
			Ex       string `json:"ex"`
			Required bool   `json:"required"`
		} `json:"options"`
	} `json:"parameters"`
	Ordering               int           `json:"ordering"`
	CanUseReinstall        bool          `json:"can_use_reinstall"`
	CanUseContract         bool          `json:"can_use_contract"`
	ExtraData              struct {
	} `json:"extra_data"`
}

type CustomScript struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ScriptBody  string `json:"script_body"`
}