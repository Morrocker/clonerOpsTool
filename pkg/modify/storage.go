package modify

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
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
			params, err := getStoreParams(store.Options.BasePath, store.URL)
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
					// here v has type S
				default:
					// fmt.Println("couldnt identify the value type")
					// no match; here v has the same type as i
				}
				if value == "" {
					fmt.Println("value is empty, skipping")
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
			fmt.Println("")
		}
		spew.Dump(stores)
	}
	return stores, nil
}

func getStoreParams(basepath, URL string) (storedata, error) {
	var data storedata
	// fmt.Println(basepath)

	numbers := getNumbers(basepath)
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

	// base, point := path.Split(basepath)
	// pathArray := strings.Split(basepath, "/")
	// store := pathArray[1]
	// point := pathArray[2]
	// store := path.Base(base)

	// storeSplit := strings.Split(store, "e")
	// storeInt, err := strconv.Atoi(storeSplit[1])
	if err != nil {
		return data, err
	}
	// data.storeNum = storeInt

	// pointSplit := strings.Split(point, "t")
	// pointInt, err := strconv.Atoi(pointSplit[1])
	// if err != nil {
	// 	return data, err
	// }
	// data.pointNum = pointInt
	// if err != nil {
	// 	return data, err
	// }

	fullURL, err := url.Parse(URL)
	data.dns = fullURL.Hostname()
	port, err := strconv.Atoi(fullURL.Port())
	if err != nil {
		return data, err
	}
	data.port = port
	// httpSplit := strings.Split(url, "/")
	// urlEnd := httpSplit[2]
	// urlSplit := strings.Split(urlEnd, ":")

	// data.dns = urlSplit[0]
	// port, err := strconv.Atoi(urlSplit[1])
	// if err != nil {
	// 	return data, err
	// }
	// fmt.Printf("DNS: %v, Port: %v\n", data.dns, data.port)

	return data, nil
}

func getNumbers(str string) []string {
	re := regexp.MustCompile(`\d[\d]*`)
	submatchall := re.FindAllString(str, -1)
	return submatchall
}
