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

// Instructions receives the JSON info that details instructions about how to modify the storage_config.json file
type Instructions struct {
	Instructions []Instruction
}

// Instruction receives the JSON info that details instructions about how to modify the storage_config.json file
type Instruction struct {
	FromStore  int
	ToStore    int
	FromPort   int
	ToPort     int
	TargetURLS []string
	Store      changeStore
}

type changeStore struct {
	Capacity string `json:"Capacity"`
	Backend  string `json:"backend"`
	BasePath string `json:"basePath"`
	URL      string `json:"URL"`
	Magic    string `json:"Magic"`
	CertFile string `json:"CertFile"`
	KeyFile  string `json:"KeyFile"`
	Insecure string `json:"Insecure"`
	Open     string `json:"Open"`
	Run      string `json:"Run"`
}

// StorageConfig provides the structure to contain a blockmaster storage configuration
type StorageConfig struct {
	Stores  []Store
	Backups []backup
	Master  master `json:"Master"`
}

// Store asfdasfd
type Store struct {
	Capacity int64  `json:"Capacity"`
	Options  option `json:"Options"`
	URL      string `json:"URL"`
	Magic    string `json:"Magic"`
	CertFile string `json:"CertFile"`
	KeyFile  string `json:"KeyFile"`
	Insecure bool   `json:"Insecure"`
	Open     bool   `json:"Open"`
	Run      bool   `json:"Run"`
}

// Master asdfas asdf asf
type master struct {
	DSN              string `json:"DSN"`
	URL              string `json:"URL"`
	Magic            string `json:"Magic"`
	CertFile         string `json:"CertFile"`
	KeyFile          string `json:"KeyFile"`
	BackupCacheLimit int    `json:"BackupCacheLimit"`
	BackupQueueLimit int    `json:"BackupQueueLimit"`
	Insecure         bool   `json:"Insecure"`
}

type backup struct {
	URL   string
	Magic string
}

type option struct {
	Backend  string `json:"backend"`
	BasePath string `json:"basePath"`
}
