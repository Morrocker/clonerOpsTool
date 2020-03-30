package modify

import (
	"fmt"
	"strconv"
	"strings"

	st "github.com/clonerOpsTool/pkg/structs"
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
	// key := strings.ToLower(k)

	for _, it := range its.Instructions {
		fmt.Printf("Params: Port:%d/%d, Store: %d/%d, Target:%v\n", it.FromPort, it.ToPort, it.FromStore, it.ToStore, it.TargetURLS)

		for _, store := range stores {
			params, err := getStoreParams(store.Options.BasePath, store.URL)
			if err != nil {
				return Stores, err
			}

			if it.FromPort != 0 && (params.port < it.FromPort || params.port > it.ToPort) {
				// fmt.Printf("Store port:%v, From port:%v, To port:%v\n", params.port, it.FromPort, it.ToPort)
				continue
			}
			if it.FromStore != 0 && (params.storeNum < it.FromStore || params.storeNum > it.ToStore) {
				// fmt.Printf("Store num:%v, From store:%v, To store:%v\n", params.storeNum, it.FromStore, it.ToStore)
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
				// fmt.Printf("URL: %s does not match any target:%v\n", params.dns, it.TargetURLS)
				continue
			}

			fmt.Printf("Changing: path=%s, url=%s\n", store.Options.BasePath, store.URL)

			// 	switch key {
			// 	case "capacity":
			// 		cap, err := strconv.ParseInt(value, 10, 64)
			// 		if err != nil {
			// 			return Stores, err
			// 		}
			// 		stores[i].Capacity = cap
			// 	case "backend":
			// 		stores[i].Options.Backend = value
			// 	case "basepath":
			// 		stores[i].Options.BasePath = value
			// 	case "url":
			// 		stores[i].URL = value
			// 	case "magic":
			// 		stores[i].Magic = value
			// 	case "certfile":
			// 		stores[i].CertFile = value
			// 	case "keyfile":
			// 		stores[i].KeyFile = value
			// 	case "insecure":
			// 		b, err := strconv.ParseBool(value)
			// 		if err != nil {
			// 			return Stores, err
			// 		}
			// 		stores[i].Insecure = b
			// 	case "open":
			// 		b, err := strconv.ParseBool(value)
			// 		if err != nil {
			// 			return Stores, err
			// 		}
			// 		stores[i].Open = b
			// 	case "run":
			// 		b, err := strconv.ParseBool(value)
			// 		if err != nil {
			// 			return Stores, err
			// 		}
			// 		stores[i].Run = b
			// 	default:
			// 		fmt.Println("given key does not match any know storage key")
			// 	}
		}
	}
	return Stores, nil
}

func getStoreParams(basepath, url string) (storedata, error) {
	var data storedata
	pathArray := strings.Split(basepath, "/")
	store := pathArray[1]
	point := pathArray[2]

	storeSplit := strings.Split(store, "e")
	storeInt, err := strconv.Atoi(storeSplit[1])
	if err != nil {
		return data, err
	}
	data.storeNum = storeInt

	pointSplit := strings.Split(point, "t")
	pointInt, err := strconv.Atoi(pointSplit[1])
	if err != nil {
		return data, err
	}
	data.pointNum = pointInt

	httpSplit := strings.Split(url, "/")
	urlEnd := httpSplit[2]
	urlSplit := strings.Split(urlEnd, ":")

	data.dns = urlSplit[0]
	port, err := strconv.Atoi(urlSplit[1])
	if err != nil {
		return data, err
	}
	data.port = port

	// fmt.Printf("DNS: %v, Port: %v\n", data.dns, data.port)

	return data, nil
}
