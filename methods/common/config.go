package common

// StorageConfig provides the structure to contain a blockmaster storage configuration
type StorageConfig struct {
	Stores  []store
	Backups []backup
	Master  master `json:"Master"`
}

type store struct {
	Capacity int64
	Options  option
	URL      string
	Magic    string
	CertFile string
	KeyFile  string
	Insecure bool
	Open     bool
	Run      bool
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
	Backend  string
	BasePath string
}
