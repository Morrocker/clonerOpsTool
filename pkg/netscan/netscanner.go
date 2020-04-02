package netscan

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

// ServerList provides the structure to store the configuration file information
type ServerList struct {
	Servers []Server
}

type nsParams struct {
	site     string
	port     string
	scantime string
	pid      string
}

// Server provides the structure to store any server to be used
type Server struct {
	Name       string
	User       string
	Category   string
	HostIperf  bool
	Site       string
	SSHPort    string
	LocalIP    string
	VpnIP      string
	ExternalIP string
	Os         string
}

// ScanMaster provides the structure to set a master server to do a netscan
type ScanMaster struct {
	Server Server
	Params nsParams
}

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

// SetIMaster starts an Iperf listener on master server
func (sm *ScanMaster) SetIMaster(port, scantime, location, name string, servers []Server) error {
	var s Server

	sm.Params = nsParams{
		site:     location,
		scantime: scantime,
		port:     port,
	}

	err := errors.New("Server not found in configuration list")
	for _, s = range servers {
		if !inLocation(location, s) {
			continue
		}
		if strings.ToLower(name) == strings.ToLower(s.Name) {
			err = nil
			goto End
		}
	}
	return err
End:
	// fmt.Println("Master server set")
	sm.Server = s
	return nil
}

// StartMaster asfdas asdfasdf asdfasd
func (sm *ScanMaster) StartMaster() error {
	var rmtCmd string
	var cmd *exec.Cmd
	pFlag := "-p " + sm.Server.SSHPort
	addr := sm.Server.User + "@" + sm.Server.VpnIP
	// fmt.Println(pFlag)
	// fmt.Println(addr)

	fmt.Printf("Starting Iperf Master in %s (%s)\n", strings.Title(sm.Server.Name), sm.Server.LocalIP)

	isHost, err := isHost(sm.Server)
	if err != nil {
		return err
	}

	iPort := sm.Params.port
	if isHost {
		rmtCmd = "iperf -s -p " + iPort
		fmt.Println(rmtCmd)
		cmd = exec.Command(rmtCmd)
	} else {
		if sm.Server.Os == "macos" {
			rmtCmd = "/usr/local/bin/iperf -s -p " + iPort

		} else {
			rmtCmd = "iperf -s -p " + iPort
		}
		// fmt.Println(rmtCmd)
		cmd = exec.Command("ssh", pFlag, addr, rmtCmd)
	}

	err = cmd.Start()
	if err != nil {
		return err
	}
	// fmt.Println("Command started")

	time.Sleep(2 * time.Second)
	cmd2 := exec.Command("ssh", pFlag, addr, "pgrep iperf")
	rmtPid, err := cmd2.Output()
	if err != nil {
		return err
	}

	pid := string(rmtPid)
	sm.Params.pid = pid
	fmt.Printf("Master started. PID: %s \n", pid)
	return nil
}

// StopMaster asfdasdf asdfasdf asdf asdf
func (sm *ScanMaster) StopMaster() error {
	pFlag := "-p " + sm.Server.SSHPort
	addr := sm.Server.User + "@" + sm.Server.VpnIP

	cmd := exec.Command("ssh", pFlag, addr, "sudo kill -9 "+sm.Params.pid)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\nError while stopping master: %v", err)
		return err
	}
	fmt.Printf("Master stopped\n\n")
	return nil
}

// RunTest receives a config variable and a "row" and scans through the servers given that they are not the master and are in the set location
func (sm *ScanMaster) RunTest(servers []Server) ([]string, error) {
	row := []string{sm.Server.Name}
	for _, server := range servers {
		if sm.Server.Name == server.Name {
			row = append(row, "-")
			continue
		}
		if server.Site != sm.Params.site {
			continue
		}

		output, err := sm.runScan(server)
		if err != nil {
			fmt.Printf("Error while running iperf server: %s", err)
			return row, err
		}
		row = append(row, output)
	}
	return row, nil
}

// RunScan takes a client server, master server, iperf port and scantime and runs a bidirectional net test
func (sm *ScanMaster) runScan(s Server) (string, error) {
	fmt.Printf("Running Iperf client on %s (%s) to %s (%s)\n", strings.Title(s.Name), s.LocalIP, strings.Title(sm.Server.Name), sm.Server.LocalIP)
	var cmd *exec.Cmd
	var ret, rmtCmd string
	pFlag := "-p " + s.SSHPort
	addr := s.User + "@" + s.VpnIP

	isHost, err := isHost(s)
	if err != nil {
		return "", err
	}
	if isHost {
		rmtCmd = "iperf -c " + sm.Server.LocalIP + " -p " + sm.Params.port
		cmd = exec.Command(rmtCmd)
	} else {
		if s.Os == "macos" {
			rmtCmd = "/usr/local/bin/iperf -c " + sm.Server.LocalIP + " -p " + sm.Params.port

		} else {
			rmtCmd = "iperf -c " + sm.Server.LocalIP + " -p " + sm.Params.port

		}
		cmd = exec.Command("ssh", pFlag, addr, rmtCmd)
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		return ret, err
	}
	err = cmd.Start()
	if err != nil {
		return ret, err
	}

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "Bytes") || strings.Contains(text, "bits") {
			words := strings.Fields(text)
			ret = words[6] + " " + words[7]
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	cmd.Wait() // }

	fmt.Printf("Output: %s\n", ret)
	return ret, nil
}

// CreateHeader Takes the servers config and location and returns a header to append
func CreateHeader(cfg ServerList, location string) []string {
	var header = []string{""}
	for _, server := range cfg.Servers {
		if !inLocation(location, server) {
			continue
		}
		header = append(header, server.Name)
	}
	return header

}

func inLocation(l string, s Server) bool {
	if strings.ToLower(l) == strings.ToLower(s.Site) {
		return true
	}
	return false
}

// isHost checks if the device running the application is part of the scanned
func isHost(s Server) (bool, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return false, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return false, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.String() == s.LocalIP {
				// fmt.Printf("LocalIP match: %s vs %s.\n", ip, s.LocalIP)
				return true, nil
			} else if ip.String() == s.VpnIP {
				// fmt.Printf("VpnIP match: %s vs %s.\n", ip, s.VpnIP)
				return true, nil
			}
		}
	}
	return false, nil
}
