package confeditor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// StorageConfig provides the structure to contain a blockmaster storage configuration
type StorageConfig struct {
	Stores  []Store     `json:"Stores"`
	Backups []Backup    `json:"Backups"`
	Master  interface{} `json:"Master"`
}

// Store asfdasfd
type Store struct {
	Capacity int    `json:"Capacity"`
	Options  option `json:"Options"`
	URL      string `json:"URL"`
	Magic    string `json:"Magic"`
	CertFile string `json:"CertFile"`
	KeyFile  string `json:"KeyFile"`
	Insecure bool   `json:"Insecure"`
	Open     bool   `json:"Open"`
	Run      bool   `json:"Run"`
	Data     IDData `json:"-"`
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

// Load asdf aasdf asdf
func (c *StorageConfig) Load(st string) error {
	f, err := ioutil.ReadFile(st)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(f), c)
	if err != nil {
		return err
	}
	// spew.Dump(data)

	c.GetStoresData()
	c.SortStores()
	return nil
}

// Write asdf adf asdf
func (c *StorageConfig) Write(name string) error {
	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(name, file, 0644)
	if err != nil {
		return err

	}
	fmt.Printf("JSON data writen to file %s\n", name)
	return nil
}

// GetStoresData asdfsa asdf asdf
func (c *StorageConfig) GetStoresData() error {
	for i := range c.Stores {
		if err := c.Stores[i].getData(); err != nil {
			return err
		}
	}
	return nil
}

// getData asdlfk asldkfa lkdfja sk
func (s *Store) getData() error {
	var data IDData

	// fmt.Println("Trying to get store data")
	re := regexp.MustCompile(`\d[\d]*`)
	pNumbers := re.FindAllString(s.Options.BasePath, -1)

	if len(pNumbers) <= 1 {
		err := errors.New("Couldn't detect store and point numbers from basepath:" + s.Options.BasePath)
		return err
	}
	storeInt, err := strconv.Atoi(pNumbers[0])
	if err != nil {
		return err
	}
	pointInt, err := strconv.Atoi(pNumbers[1])
	if err != nil {
		return err
	}
	data.StoreNum = storeInt
	data.PointNum = pointInt

	fullURL, err := url.Parse(s.URL)
	data.dns = fullURL.Hostname()
	port, err := strconv.Atoi(fullURL.Port())
	if err != nil {
		return err
	}
	data.port = port

	s.Data = data
	return nil
}

// GetLastPort asdfas sadfasfas asdfasfd
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

			fmt.Printf("Port %d duplicated in config. Check manually to fix inconsistency or possible error", lastPort)
			err := errors.New("Port duplicated in config. Check manually to fix inconsistency or possible error")
			return 0, err
		}
	}
	return lastPort, nil
}

// SortStores asdfasdf asdfasdf aasdf a
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

// GetLastPoint asdfklj alsdfka lfdk alskdfja
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

			fmt.Printf("Point %d duplicated in config\n", lastPoint)
			spew.Dump(s)
			err := errors.New("point duplicated in config. Check manually to fix inconsistency or possible error")
			return 0, err
		}
	}
	return lastPoint, nil
}

// GenerateSlave adsfa asdf asdf asdf
func (c *StorageConfig) GenerateSlave(svName string) StorageConfig {
	var Bkps = make([]Backup, 0)
	var SlaveConf = StorageConfig{
		Backups: Bkps,
		Master:  nil,
	}
	// var SlaveStores []Stores
	for _, store := range c.Stores {
		if !strings.Contains(store.Data.dns, svName) {
			continue
		}
		store.Run = true
		store.CertFile = "/etc/letsencrypt/live/" + svName + ".cloner.cl/fullchain.pem"
		store.KeyFile = "/etc/letsencrypt/live/" + svName + ".cloner.cl/privkey.pem"
		SlaveConf.Stores = append(SlaveConf.Stores, store)
	}
	return SlaveConf
}

// ExtendStore lkjfda lkasjd asdf lkj
func (c *StorageConfig) ExtendStore(svName string, stNum, toPoint int, master bool) error {
	fmt.Printf("Extending store %d\n", stNum)
	lastPoint, err := c.GetLastPoint(svName, stNum)
	if err != nil {
		return err
	} else if lastPoint == 0 {
		err := errors.New("No last point found, check if parameters are correct")
		return err
	}
	err = c.AddStore(svName, stNum, lastPoint+1, toPoint, master)
	if err != nil {
		return err
	}
	return nil
}

// AddStore lkjsdf lkajslkdf asd klj
func (c *StorageConfig) AddStore(svName string, stNum, fromPoint, toPoint int, master bool) error {

	fmt.Printf("Creating store %d on %s from point %d to %d with master set to %v\n", stNum, svName, fromPoint, toPoint, master)

	bkpStores := c.Stores
	var newStores []Store
	point := fromPoint
	store := strconv.Itoa(stNum)
	lastPort, err := c.GetLastPort(svName)
	if err != nil {
		return err
	} else if lastPort == 0 {
		err := errors.New("No last port found, check if parameters are correct")
		return err
	}

	if fromPoint > toPoint {
		err := errors.New("Start point (" + strconv.Itoa(point) + ") is larger than finish point (" + strconv.Itoa(toPoint) + "). Cancelling operation")
		return err
	}

	for point <= toPoint {
		lastPort++
		port := strconv.Itoa(lastPort)
		sPoint := strconv.Itoa(point)
		var newStore = Store{
			Capacity: 960000000000,
			URL:      "https://" + svName + ".cloner.cl:" + port,
			Magic:    svName + "_s" + store + "_" + sPoint,
			CertFile: "/etc/letsencrypt/live/" + svName + ".cloner.cl/fullchain.pem",
			KeyFile:  "/etc/letsencrypt/live/" + svName + ".cloner.cl/privkey.pem",
			Insecure: false,
			Open:     true,
			Run:      master,
		}

		newStore.Options.BasePath = "/storage" + store + "/point" + sPoint
		newStore.Options.Backend = "disk"

		newStores = append(newStores, newStore)
		point++

	}
	c.Stores = append(c.Stores, newStores...)
	c.GetStoresData()
	c.SortStores()
	// spew.Dump(c.Stores)
	issues := c.Check()
	if issues {
		err := errors.New("Errors found in resulting stores. Doing rollback")
		c.Stores = bkpStores
		return err
	}
	return nil
}

// Check aslkdfj aasfkdja sjfas dk
func (c *StorageConfig) Check() bool {
	var lstStore, curStore, lstPoint, curPoint int
	var portList []int
	var lstDNS, curDNS string
	var err bool
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
					fmt.Printf("Port %d duplicated in %s\n", port, curDNS)
					err = true
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
			fmt.Printf("Point duplicated in %s#%d.%d\n", curDNS, curStore, curPoint)
			err = true
			update()
			continue
		} else if lstPoint != curPoint-1 {
			fmt.Printf("Point missing between %s#%d.%d and %s#%d.%d\n", lstDNS, lstStore, lstPoint, curDNS, curStore, curPoint)
			err = true
			update()
			continue
		}
		update()
	}
	return err
}

// ModifyStore akjdsf lskdfa jfdlka jfdl
func (s *Store) ModifyStore(c changeStore) error {
	fmt.Printf("Modifying store %d, point %d according to instruction\n", s.Data.StoreNum, s.Data.PointNum)
	r := reflect.ValueOf(&c).Elem()
	for x := 0; x < r.NumField(); x++ {
		key := strings.ToLower(r.Type().Field(x).Name)
		val := r.Field(x).Interface()
		if val == nil {
			continue
		}
		switch v := val.(type) {
		case string:
			switch key {
			case "backend":
				s.Options.Backend = v
			case "basepath":
				s.Options.BasePath = v
			case "url":
				s.URL = v
			case "magic":
				s.Magic = v
			case "certfile":
				s.CertFile = v
			case "keyfile":
				s.KeyFile = v
			default:
				fmt.Printf("given key ( %v / type %v ) does not match any know storage key", key, v)
			}
		case int:
			switch key {
			case "capacity":
				s.Capacity = v
			default:
				fmt.Printf("given key ( %v  / type %v ) does not match any know storage key", key, v)
			}
		case float64:
			var value int = int(v)
			switch key {
			case "capacity":
				s.Capacity = value
			default:
				fmt.Printf("given key ( %v  / type %v ) does not match any know storage key", key, v)
			}

		case bool:
			switch key {
			case "insecure":
				s.Insecure = v
			case "open":
				s.Open = v
			case "run":
				s.Run = v
			default:
				fmt.Printf("given key ( %v  / type %v ) does not match any know storage key", key, v)
			}
		default:
			fmt.Println("Given key does not match any known storage key. Cancelling Operation")
			err := errors.New("Given type doesn't exist the storage config")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetStore asdf adaf sadlkf aldkf
func (c *StorageConfig) GetStore(svName string, stNum, ptNum int) (*Store, error) {
	// spew.Dump(c)
	fmt.Printf("Searching for store %d point %d on server %s\n", stNum, ptNum, svName)
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
	err := errors.New("Didn't find store")

	return nil, err
}
