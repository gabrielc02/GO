package report

type ReportedAsset struct {
	ID            string `json:"id"`
	Hostname      string `json:"hostname"`
	Status        string `json:"status"`
	CPUUsage      int    `json:"cpu_usage"`
	MemoryUsage   int    `json:"memory_usage"`
	Region        string `json:"region"`
	SeverityLevel string `json:"severity_level"`
}
