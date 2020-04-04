package confeditor

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
