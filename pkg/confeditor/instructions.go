package confeditor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// Instructions receives the JSON info that details instructions about how to modify the storage_config.json file
type Instructions struct {
	Instructions []Instruction
}

// Instruction receives the JSON info that details instructions about how to modify the storage_config.json file
type Instruction struct {
	Type          string      `json:"Type"`
	FromStore     int         `json:"fromStore"`
	ToStore       int         `json:"toStore"`
	FromPoint     int         `json:"fromPoint"`
	ToPoint       int         `json:"toPoint"`
	URL           string      `json:"URL"`
	Master        bool        `json:"master"`
	ChangedParams changeStore `json:"ChangedParams"`
}

type changeStore struct {
	Capacity interface{} `json:"Capacity"`
	Backend  interface{} `json:"backend"`
	BasePath interface{} `json:"basePath"`
	URL      interface{} `json:"URL"`
	Magic    interface{} `json:"Magic"`
	CertFile interface{} `json:"CertFile"`
	KeyFile  interface{} `json:"KeyFile"`
	Insecure interface{} `json:"Insecure"`
	Open     interface{} `json:"Open"`
	Run      interface{} `json:"Run"`
}

// Run asdlkfj asdflkja sldfkj a
func (i *Instructions) Run(s *StorageConfig) error {
	fmt.Println("Starting run")
	for x, ins := range i.Instructions {
		fmt.Printf("\nRunning instruction %d, type %s\n", x+1, ins.Type)
		if e := ins.run(s); e != nil {
			return e
		}
		// fmt.Println("Resulting Stores")
	}
	// spew.Dump(s)
	return nil
}

func (i *Instruction) run(c *StorageConfig) error {
	t := strings.ToLower(i.Type)
	if i.FromStore == 0 {
		i.FromStore = 1
	}
	if i.FromPoint == 0 {
		i.FromPoint = 1
	}
	switch t {
	case "add":
		for s := i.FromStore; s <= i.ToStore; s++ {
			if e := c.AddStore(i.URL, s, i.FromPoint, i.ToPoint, i.Master); e != nil {
				return e
			}
		}
	case "extend":
		for s := i.FromStore; s <= i.ToStore; s++ {
			if e := c.ExtendStore(i.URL, s, i.ToPoint, i.Master); e != nil {
				return e
			}
		}
	case "change":
		for s := i.FromStore; s <= i.ToStore; s++ {
			for p := i.FromPoint; p <= i.ToPoint; p++ {
				st, e := c.GetStore(i.URL, s, p)
				if e != nil {
					fmt.Println(e)
					continue
				}
				if e := st.ModifyStore(i.ChangedParams); e != nil {
					return e
				}
			}
		}
		c.SortStores()
		if e := c.Check(); e {
			err := errors.New("Resulting modifications create issues. Exiting")
			return err
		}
	default:
		err := errors.New("Instruction type not found")
		return err
	}
	return nil
}

// Load asdfa dfa asdfk asldfkj a
func (i *Instructions) Load(name string) error {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &i)
	if err != nil {
		return err
	}
	// spew.Dump(data)
	return nil

}
