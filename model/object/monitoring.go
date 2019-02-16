package object

type 	Monitoring struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Settings    struct {
		SimpleMonitor struct {
			NotifyEmail struct {
				HTML    bool `json:"html"`
				Enabled bool `json:"enabled"`
			} `json:"notify_email"`
			Enabled     bool `json:"enabled"`
			NotifySlack struct {
				Enabled             bool   `json:"enabled"`
				IncomingWebhooksURL string `json:"incoming_webhooks_url"`
			} `json:"notify_slack"`
			DelayLoop   int `json:"delay_loop"`
			HealthCheck struct {
				Protocol string `json:"protocol"`
			} `json:"health_check"`
		} `json:"simple_monitor"`
	} `json:"settings"`
	Server          int    `json:"server"`
	CloudResourceID string `json:"cloud_resource_id"`
}