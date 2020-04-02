package modify

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	st "github.com/clonerOpsTool/pkg/structs"
	"github.com/davecgh/go-spew/spew"
)

type storedata struct {
	storeNum int
	pointNum int
	dns      string
	port     int
}

// ExecInstr takes a Store and Instruction to modify it and executes it.
func ExecInstr(stores []st.Store, its st.Instructions) ([]st.Store, error) {
	var Stores []st.Store

	for _, it := range its.Instructions {
		// fmt.Printf("Executing instruction #%v. Params: Port:%d/%d, Store: %d/%d, Target:%v\n", z, it.FromPort, it.ToPort, it.FromStore, it.ToStore, it.TargetURLS)
		for i, store := range stores {
			// fmt.Printf("checking store: %s\n", store.URL)
			params, err := getStoreData(store)
			if err != nil {
				return Stores, err
			}
			if it.FromPort != 0 && params.port < it.FromPort {
				continue
			}
			if it.ToPort != 0 && params.port > it.ToPort {
				continue
			}
			if it.FromStore != 0 && params.storeNum < it.FromStore {
				continue
			}
			if it.ToStore != 0 && params.storeNum > it.ToStore {
				continue
			}

			found := false
			for _, url := range it.TargetURLS {
				if strings.Contains(params.dns, url) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
			// fmt.Printf("Changing: path=%s, url=%s\n", store.Options.BasePath, store.URL)
			r := reflect.ValueOf(&it.Store).Elem()
			for x := 0; x < r.NumField(); x++ {
				key := strings.ToLower(r.Type().Field(x).Name)
				val := r.Field(x).Interface()
				var value string

				switch v := val.(type) {
				case string:
					value = v
				default:
					err := errors.New("Only string types can be used on an instruction store conf")
					if err != nil {
						return Stores, err
					}
				}
				if value == "" {
					// fmt.Println("value is empty, skipping")
					continue
				}
				// fmt.Printf("store: %s. key: %s, value:%v\n", store.Options.BasePath, key, value)
				switch key {
				case "capacity":
					cap, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return Stores, err
					}
					stores[i].Capacity = cap
				case "backend":
					stores[i].Options.Backend = value
				case "basepath":
					stores[i].Options.BasePath = value
				case "url":
					stores[i].URL = value
				case "magic":
					stores[i].Magic = value
				case "certfile":
					stores[i].CertFile = value
				case "keyfile":
					stores[i].KeyFile = value
				case "insecure":
					b, err := strconv.ParseBool(value)
					if err != nil {
						return Stores, err
					}
					stores[i].Insecure = b
				case "open":
					b, err := strconv.ParseBool(value)
					if err != nil {
						return Stores, err
					}
					stores[i].Open = b
				case "run":
					b, err := strconv.ParseBool(value)
					if err != nil {
						return Stores, err
					}
					stores[i].Run = b
				default:
					fmt.Println("given key does not match any know storage key")
				}
			}
			// fmt.Println("")
		}
		// spew.Dump(stores)
	}
	return stores, nil
}

// func getStoreData(store st.Store) (storedata, error) {
// 	basepath := store.Options.BasePath
// 	URL := store.URL
// 	var data storedata
// 	// fmt.Println(basepath)

// 	numbers := getNumbers(basepath)
// 	if len(numbers) <= 1 {
// 		err := errors.New("Couldn't detect store and point numbers from basepath:" + basepath)
// 		return data, err
// 	}
// 	storeInt, err := strconv.Atoi(numbers[0])
// 	if err != nil {
// 		return data, err
// 	}
// 	pointInt, err := strconv.Atoi(numbers[1])
// 	if err != nil {
// 		return data, err
// 	}
// 	data.storeNum = storeInt
// 	data.pointNum = pointInt

// 	fullURL, err := url.Parse(URL)
// 	data.dns = fullURL.Hostname()
// 	port, err := strconv.Atoi(fullURL.Port())
// 	if err != nil {
// 		return data, err
// 	}
// 	data.port = port

// 	return data, nil
// }

// func getNumbers(str string) []string {
// 	re := regexp.MustCompile(`\d[\d]*`)
// 	submatchall := re.FindAllString(str, -1)
// 	return submatchall
// }

// GetLastPort asdfas sadfasfas asdfasfd
func GetLastPort(name string, stores []st.Store) (int, error) {
	lastPort := 0
	for _, store := range stores {
		if !strings.Contains(store.URL, name) {
			continue
		}
		stData, err := getStoreData(store)
		if err != nil {
			return 0, err
		}
		if lastPort < stData.port {
			lastPort = stData.port
			continue
		} else if lastPort == stData.port {

			fmt.Printf("Port %d duplicated in config. Check manually to fix inconsistency or possible error", lastPort)
			err := errors.New("Port duplicated in config. Check manually to fix inconsistency or possible error")
			return 0, err
		}
	}
	return lastPort, nil
}

// GetStoreCluster asfdas asdfadfa sdf a
func GetStoreCluster(svName string, stNum int, stores []st.Store) ([]st.Store, error) {
	var stFound []st.Store

	for _, store := range stores {
		if !strings.Contains(store.URL, svName) {
			continue
		}
		stData, err := getStoreData(store)
		if err != nil {
			return stFound, err
		}
		if stNum == stData.storeNum {
			stFound = append(stFound, store)
		}
	}
	return stFound, nil
}

// GetLastPoint asdfklj alsdfka lfdk alskdfja
func GetLastPoint(stores []st.Store) (int, error) {
	var lastPoint int
	for _, store := range stores {
		stData, err := getStoreData(store)
		if err != nil {
			return 0, err
		}
		if lastPoint < stData.pointNum {
			lastPoint = stData.pointNum
			continue
		} else if lastPoint == stData.pointNum {

			fmt.Printf("Point %d duplicated in config %v.\n Check manually to fix inconsistency or possible error", lastPoint, store)
			err := errors.New("point duplicated in config. Check manually to fix inconsistency or possible error")
			return 0, err
		}
	}
	return lastPoint, nil
}

// ExtendStore asfdas asfdasdf sadf asdf
func ExtendStore(svName string, stNum, toPoint int, master bool, stores []st.Store) ([]st.Store, error) {
	newStores := stores
	lastPort, err := GetLastPort(svName, stores)
	if err != nil {
		return newStores, err
	} else if lastPort == 0 {
		err := errors.New("No last port found, check if parameters are correct")
		return newStores, err
	}
	cluster, err := GetStoreCluster(svName, stNum, stores)
	if err != nil {
		return newStores, err
	} else if cluster == nil {
		err := errors.New("No store cluster found found, check if parameters are correct")
		return newStores, err
	}
	lastPoint, err := GetLastPoint(cluster)
	if err != nil {
		return newStores, err
	}

	for lastPoint < toPoint {
		lastPoint++
		lastPort++
		port := strconv.Itoa(lastPort)
		store := strconv.Itoa(stNum)
		point := strconv.Itoa(lastPoint)
		var newStore = st.Store{
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
	return newStores, nil
}

// SortStores asdfa asf asfdlkas fdlkjsdf
func SortStores(stores []st.Store) ([]st.Store, error) {
	var storesData []storedata
	for _, store := range stores {
		stData, err := getStoreData(store)
		if err != nil {
			return stores, err
		}
		storesData = append(storesData, stData)
	}
	spew.Dump(storesData)
	return stores, nil

}

func getDistinctDNS() {

}
