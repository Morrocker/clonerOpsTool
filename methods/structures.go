package methods

// Config provides the structure to store the configuration file information
type Config struct {
	Servers []Server
}

// Server provides the structure to store any server to be used
type Server struct {
	Name       string
	User       string
	Category   string
	HostIperf  bool
	Location   string
	Port       string
	LocalIP    string
	VpnIP      string
	ExternalIP string
	Os         string
}

// Master provides the structure to set a master server to do a netscan
type Master struct {
	Name     string
	Category string
	Port     string
	VpnIP    string
	LocalIP  string
}
