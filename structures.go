package methods

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// ------------------------------------------------------------------------------------------------------------------------------------

// StorageConfig provides the structure to contain a blockmaster storage configuration
type StorageConfig struct {
	Stores  []Store
	Backups []BlkBackup
	Master  BlkMaster `json:"Master"`
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

type stData struct {
	storeNum int
	pointNum int
	dns      string
	port     int
}

// SortStores asdfsadfa asldkfjalsdkfjal afsldkj as
func (s *StorageConfig) SortStores() {
	var storesData []stData
	for _, store := range s.Stores {
		stData, err := store.GetData()
		if err != nil {
			// return storesData, err
		}
		storesData = append(storesData, stData)
	}
	spew.Dump(storesData)

}

// GetData asdfa asdf asdf a
func (s *Store) GetData() (stData, error) {
	basepath := s.Options.BasePath
	URL := s.URL
	var data stData

	re := regexp.MustCompile(`\d[\d]*`)
	numbers := re.FindAllString(basepath, -1)

	if len(numbers) <= 1 {
		err := errors.New("Couldn't detect store and point numbers from basepath:" + basepath)
		return data, err
	}
	storeInt, err := strconv.Atoi(numbers[0])
	if err != nil {
		return data, err
	}
	pointInt, err := strconv.Atoi(numbers[1])
	if err != nil {
		return data, err
	}
	data.storeNum = storeInt
	data.pointNum = pointInt

	fullURL, err := url.Parse(URL)
	data.dns = fullURL.Hostname()
	port, err := strconv.Atoi(fullURL.Port())
	if err != nil {
		return data, err
	}
	data.port = port

	return data, nil
}

// BlkMaster asdfas asdf asf
type BlkMaster struct {
	DSN              string `json:"DSN"`
	URL              string `json:"URL"`
	Magic            string `json:"Magic"`
	CertFile         string `json:"CertFile"`
	KeyFile          string `json:"KeyFile"`
	BackupCacheLimit int    `json:"BackupCacheLimit"`
	BackupQueueLimit int    `json:"BackupQueueLimit"`
	Insecure         bool   `json:"Insecure"`
}

// Backup asdf asas dfas fas df
type BlkBackup struct {
	URL   string
	Magic string
}

type option struct {
	Backend  string `json:"backend"`
	BasePath string `json:"basePath"`
}
