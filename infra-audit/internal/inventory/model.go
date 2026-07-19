package inventory

type Asset struct {
	ID          string `json:"id"`
	Hostname    string `json:"hostname"`
	Status      string `json:"status"`
	CPUUsage    int    `json:"cpu_usage"`
	MemoryUsage int    `json:"memory_usage"`
	Region      string `json:"region"`
}
