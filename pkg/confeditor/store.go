package confeditor

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/morrocker/errors"
)

// getData parses a single Store object to obtain and set its IDData
func (s *Store) getData() error {
	op := "store.getData()"
	var data IDData

	// fmt.Println("Trying to get store data")
	re := regexp.MustCompile(`\d[\d]*`)
	pNumbers := re.FindAllString(s.Options.BasePath, -1)

	if len(pNumbers) <= 1 {
		return errors.New(op, fmt.Sprintf("Couldn't detect store and point numbers from basepath:%s", s.Options.BasePath))
	}
	storeInt, err := strconv.Atoi(pNumbers[0])
	if err != nil {
		return errors.Extend(op, err)
	}
	pointInt, err := strconv.Atoi(pNumbers[1])
	if err != nil {
		return errors.Extend(op, err)
	}
	data.StoreNum = storeInt
	data.PointNum = pointInt

	fullURL, err := url.Parse(s.URL)
	if err != nil {
		return errors.Extend(op, err)
	}
	data.dns = fullURL.Hostname()
	port, err := strconv.Atoi(fullURL.Port())
	if err != nil {
		return errors.Extend(op, err)
	}
	data.port = port

	s.Data = data
	return nil
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
			return errors.New("store.ModifyStore()", "Given type doesn't exist the storage config")
		}
	}
	return nil
}
