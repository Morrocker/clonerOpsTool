package confeditor

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

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

type backup struct {
	URL   string
	Magic string
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

	// spew.Dump(data)
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
					// ln := len(sortedStores) + 1
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
	// for _, s := range c.Stores {
	// 	fmt.Printf("Store: %s#%d,%d\n", s.Data.dns, s.Data.StoreNum, s.Data.PointNum)
	// }
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

			fmt.Printf("Point %d duplicated in config %v.\n Check manually to fix inconsistency or possible error", lastPoint, s)
			err := errors.New("point duplicated in config. Check manually to fix inconsistency or possible error")
			return 0, err
		}
	}
	return lastPoint, nil
}

// GenerateSlave adsfa asdf asdf asdf
func (c *StorageConfig) GenerateSlave() {
}

// ExtendStore lkjfda lkasjd asdf lkj
func (c *StorageConfig) ExtendStore(svName string, stNum, toPoint int, master bool) error {
	lastPoint, err := c.GetLastPoint(svName, stNum)
	if err != nil {
		return err
	} else if lastPoint == 0 {
		err := errors.New("No last point found, check if parameters are correct")
		return err
	}
	err = c.AddStore(svName, stNum, lastPoint, toPoint, master)
	if err != nil {
		return err
	}
	return nil

}

// AddStore lkjsdf lkajslkdf asd klj
func (c *StorageConfig) AddStore(svName string, stNum, fromPoint, toPoint int, master bool) error {

	fmt.Printf("Creating new store %s#%d from %d to %d in with master %v\n", svName, stNum, fromPoint, toPoint, master)

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

	if fromPoint >= toPoint {
		err := errors.New("Start point (" + strconv.Itoa(point) + ") is larger or equal to finish point (" + strconv.Itoa(toPoint) + "). Cancelling operation")
		return err
	}

	for point < toPoint {
		point++
		lastPort++
		port := strconv.Itoa(lastPort)
		point := strconv.Itoa(point)
		var newStore = Store{
			Capacity: 960000000000,
			URL:      "https://" + svName + ".cloner.cl:" + port,
			Magic:    svName + "_s" + store + "_" + point,
			CertFile: "/etc/letsencrypt/live/" + svName + ".cloner.cl/fullchain.pem",
			KeyFile:  "/etc/letsencrypt/live/" + svName + ".cloner.cl/privkey.pem",
			Insecure: false,
			Open:     true,
			Run:      master,
		}

		newStore.Options.BasePath = "/storage" + store + "/point" + point
		newStore.Options.Backend = "disk"

		newStores = append(newStores, newStore)

	}
	c.Stores = append(c.Stores, newStores...)
	c.GetStoresData()
	c.SortStores()
	issues := c.Check()
	if issues {
		err := errors.New("Errors found in resulting stores. Doing rollback")
		c.Stores = bkpStores
		return err
	}
	// spew.Dump(newStores)
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
