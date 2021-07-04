package confeditor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/morrocker/errors"
	"github.com/morrocker/log"
)

// StorageConfig provides the structure to contain a blockmaster storage configuration
type StorageConfig struct {
	Stores  []Store  `json:"Stores"`
	Backups []Backup `json:"Backups"`
	Master  Master   `json:"Master"`
}

// Store asfdasfd
type Store struct {
	Capacity int    `json:"Capacity"`
	Options  option `json:"Options"`
	URL      string `json:"URL"`
	Magic    string `json:"Magic"`
	CertFile string `json:"CertFile"`
	KeyFile  string `json:"KeyFile"`
	Legacy   bool   `json:"Legacy"`
	Insecure bool   `json:"Insecure"`
	Open     bool   `json:"Open"`
	Run      bool   `json:"Run"`
	Data     IDData `json:"-"`
}

// Master asdfas asdf asf
type Master struct {
	DSN              string `json:"DSN"`
	URL              string `json:"URL"`
	Magic            string `json:"Magic"`
	CertFile         string `json:"CertFile"`
	KeyFile          string `json:"KeyFile"`
	BackupCacheLimit int    `json:"BackupCacheLimit"`
	BackupQueueLimit int    `json:"BackupQueueLimit"`
	Insecure         bool   `json:"Insecure"`
}

// Backup asdfas
type Backup struct {
	URL   string `json:"URL"`
	Magic string `json:"Magic"`
}

type option struct {
	Backend  string `json:"backend"`
	BasePath string `json:"basePath"`
}

// IDData asdfa laksdfjl aslaks jdf
type IDData struct {
	StoreNum int
	PointNum int
	dns      string
	port     int
}

// Load loads the storage configguration JSON file into the StorageConfig object
func (c *StorageConfig) Load(st string) error {
	op := "storageconfig.Load()"
	f, err := ioutil.ReadFile(st)
	if err != nil {
		return errors.Extend(op, err)
	}

	err = json.Unmarshal([]byte(f), c)
	if err != nil {
		return errors.Extend(op, err)
	}
	c.GetStoresData()
	c.SortStores()
	return nil
}

// Write StorageConfig into a JSON file with the given name
func (c *StorageConfig) Write(name string) error {
	op := "storageconfig.Write()"
	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return errors.Extend(op, err)
	}
	err = ioutil.WriteFile(name, file, 0644)
	if err != nil {
		return errors.Extend(op, err)

	}
	log.Info("JSON data writen to file %s", name)
	return nil
}

// GetStoresData travels all stores in the StorageConfig and set the IDData for each one
func (c *StorageConfig) GetStoresData() error {
	for i := range c.Stores {
		if err := c.Stores[i].getData(); err != nil {
			return errors.Extend("storageconfig.GetStoreData()", err)
		}
	}
	return nil
}

// GetLastPort searchs the storage config for the highest existing port asociated to the given server name
func (c *StorageConfig) GetLastPort(svName string) (int, error) {
	lastPort := 0
	for _, s := range c.Stores {
		data := s.Data
		if !strings.Contains(data.dns, svName) {
			continue
		}
		if lastPort < data.port {
			lastPort = data.port
			continue
		} else if lastPort == data.port {

			return 0, errors.New("storageconfig.GetLastPort()", fmt.Sprintf("Port %d duplicated in config. Check manually to fix inconsistency or possible error", lastPort))
		}
	}
	return lastPort, nil
}

// SortStores sorts all stores inside the storage config. First by servername, then by storage and each storage by point.
func (c *StorageConfig) SortStores() {
	var dnsList []string
	var FinalStores []Store
	for _, s := range c.Stores {
		dns := s.Data.dns
		for _, d := range dnsList {
			if d == dns {
				goto NextDNS
			}
		}
		dnsList = append(dnsList, dns)
	NextDNS:
	}
	for _, dns := range dnsList {
		var sortedStores []Store
		for _, s := range c.Stores {
			if s.Data.dns != dns {
				continue
			}
			if len(sortedStores) == 0 {
				sortedStores = append(sortedStores, s)
			} else {
				for i, sorted := range sortedStores {
					var temp []Store
					if s.Data.StoreNum < sorted.Data.StoreNum {
						temp = append(temp, sortedStores[0:i]...)
						temp = append(temp, s)
						temp = append(temp, sortedStores[i:]...)
						sortedStores = temp
						break
					} else if s.Data.StoreNum == sorted.Data.StoreNum && s.Data.PointNum < sorted.Data.PointNum {
						temp = append(temp, sortedStores[0:i]...)
						temp = append(temp, s)
						temp = append(temp, sortedStores[i:]...)
						sortedStores = temp
						break
					} else if i == len(sortedStores)-1 {
						sortedStores = append(sortedStores, s)
						break
					}
				}
			}
		}
		FinalStores = append(FinalStores, sortedStores...)
	}
	c.Stores = FinalStores
}

// GetLastPoint searches the last existing point for the given server name and store
func (c *StorageConfig) GetLastPoint(svName string, storeNum int) (int, error) {
	lastPoint := 0
	for _, s := range c.Stores {
		if !strings.Contains(s.Data.dns, svName) || s.Data.StoreNum != storeNum {
			continue
		}
		if lastPoint < s.Data.PointNum {
			lastPoint = s.Data.PointNum
			continue
		} else if lastPoint == s.Data.PointNum {
			return 0, errors.New("storageconfig.GetLastPoint()", fmt.Sprintf("point %d duplicated in config. Check manually to fix inconsistency or possible error", lastPoint))
		}
	}
	return lastPoint, nil
}

// GenerateSlave generates a storage config corresponding to the slave server managed by the master config
func (c *StorageConfig) GenerateSlave(svName string) StorageConfig {
	var Bkps = make([]Backup, 0)
	var SlaveConf = StorageConfig{
		Backups: Bkps,
		Master:  Master{},
	}
	// var SlaveStores []Stores
	for _, store := range c.Stores {
		if !strings.Contains(store.Data.dns, svName) {
			continue
		}
		store.Run = true
		store.CertFile = fmt.Sprintf("/etc/letsencrypt/live/%s.cloner.cl/fullchain.pem", svName)
		store.KeyFile = fmt.Sprintf("/etc/letsencrypt/live/%s.cloner.cl/privkey.pem", svName)
		SlaveConf.Stores = append(SlaveConf.Stores, store)
	}
	return SlaveConf
}

// ExtendStore creates new store points up to the given number. It will only create points from the last point onwards
func (c *StorageConfig) ExtendStore(svName string, stNum, toPoint int, master bool) error {
	op := "storageconfig.ExtendStore()"
	log.Task("Extending store %d", stNum)
	lastPoint, err := c.GetLastPoint(svName, stNum)
	if err != nil {
		return errors.Extend(op, err)
	} else if lastPoint == 0 {
		return errors.New(op, "No last point found, check if parameters are correct")
	}
	err = c.AddStore(svName, stNum, lastPoint+1, toPoint, master)
	if err != nil {
		return errors.Extend(op, err)
	}
	return nil
}

// AddStore creates new store points
func (c *StorageConfig) AddStore(svName string, stNum, fromPoint, toPoint int, master bool) error {
	op := "storageconfig.AddStore()"
	log.Task("Creating store %d on %s from point %d to %d with master set to %v", stNum, svName, fromPoint, toPoint, master)

	bkpStores := c.Stores
	var newStores []Store
	point := fromPoint
	store := strconv.Itoa(stNum)
	lastPort, err := c.GetLastPort(svName)
	if err != nil {
		return errors.Extend(op, err)
	} else if lastPort == 0 {
		return errors.New(op, "No last port found, check if parameters are correct")
	}

	if fromPoint > toPoint {
		return errors.New(op, fmt.Sprintf("Start point (%d) is larger than finish point (%d). Cancelling operation", point, toPoint))
	}

	for point <= toPoint {
		lastPort++
		port := strconv.Itoa(lastPort)
		sPoint := strconv.Itoa(point)
		var newStore = Store{
			Capacity: 960000000000,
			URL:      fmt.Sprintf("https://%s.cloner.cl:%s", svName, port),
			Magic:    fmt.Sprintf("%s_s%s_%s", svName, store, sPoint),
			CertFile: fmt.Sprintf("/etc/letsencrypt/live/%s.cloner.cl/fullchain.pem", svName),
			KeyFile:  fmt.Sprintf("/etc/letsencrypt/live/%s.cloner.cl/privkey.pem", svName),
			Legacy:   false,
			Insecure: false,
			Open:     true,
			Run:      master,
		}

		newStore.Options.BasePath = fmt.Sprintf("/storage%s/point%s", store, sPoint)
		newStore.Options.Backend = "block_bank"

		newStores = append(newStores, newStore)
		point++

	}
	c.Stores = append(c.Stores, newStores...)
	c.GetStoresData()
	c.SortStores()
	if err := c.Check(); err != nil {
		c.Stores = bkpStores
		log.Info("Errors found in resulting stores. Doing rollback")
		return errors.Extend(op, err)
	}
	return nil
}

// RemoveStore removes the target stores from the StorageConfig object
func (c *StorageConfig) RemoveStore(svName string, st, from, to int) {
	log.Task("Removing store %d from point %d to %d from server %s", st, from, to, svName)
	svDns := fmt.Sprintf("%s.cloner.cl", svName)
	newStores := []Store{}
	for _, store := range c.Stores {
		if svDns != store.Data.dns || store.Data.StoreNum != st {
			newStores = append(newStores, store)
		} else {
			if store.Data.PointNum >= from && store.Data.PointNum <= to {
				continue
			} else {
				newStores = append(newStores, store)
			}
		}
	}
	c.Stores = newStores
}

// RenewStore replaces an existing point for a new one. Maintaining directory path, but assigning new port.
func (c *StorageConfig) RenewStore(svName string, st, from, to int, master bool) error {
	c.RemoveStore(svName, st, from, to)
	if err := c.AddStore(svName, st, from, to, master); err != nil {
		return errors.Extend("storageconfig.RenewStore()", err)
	}
	return nil
}

// Check travels the storage config stores checking for errors or inconsistencies
func (c *StorageConfig) Check() error {
	op := "storageconfig.Check()"
	var lstStore, curStore, lstPoint, curPoint int
	var portList []int
	var err error
	var lstDNS, curDNS string
	update := func() {
		lstStore = curStore
		lstPoint = curPoint
		lstDNS = curDNS
	}
	for i, s := range c.Stores {
		curStore = s.Data.StoreNum
		curPoint = s.Data.PointNum
		curDNS = s.Data.dns
		if i == 0 {
			portList = append(portList, s.Data.port)
			continue
		}
		if lstDNS == curDNS {
			for _, port := range portList {
				if port == s.Data.port {
					log.Error("Port %d duplicated in %s", port, curDNS)
					err = errors.New(op, "Errors found in the storage config file")
					goto skipPort
				}
			}
			portList = append(portList, s.Data.port)
		}
	skipPort:
		if lstDNS != curDNS {
			portList = []int{s.Data.port}
			update()
			continue
		} else if lstStore != curStore {
			update()
			continue
		} else if lstPoint == curPoint {
			log.Error("Point duplicated in %s#%d.%d", curDNS, curStore, curPoint)
			err = errors.New(op, "Errors found in the storage config file")
			update()
			continue
		} else if lstPoint != curPoint-1 {
			log.Error("Point missing between %s#%d.%d and %s#%d.%d", lstDNS, lstStore, lstPoint, curDNS, curStore, curPoint)
			err = errors.New(op, "Errors found in the storage config file")
			update()
			continue
		}
		update()
	}
	return err
}

// GetStore searchs the config for the given store and returns the object
func (c *StorageConfig) GetStore(svName string, stNum, ptNum int) (*Store, error) {
	log.Task("Searching for store %d point %d on server %s", stNum, ptNum, svName)
	for i, store := range c.Stores {
		if !strings.Contains(store.Data.dns, svName) {
			continue
		}
		if store.Data.StoreNum != stNum {
			continue
		}
		if store.Data.PointNum != ptNum {
			continue
		}
		return &c.Stores[i], nil
	}
	return nil, errors.New("storageconfig.GetStore()", "Didn't find store")
}
