package methods

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	cm "github.com/clonerOpsTool/pkg/common"
	st "github.com/clonerOpsTool/pkg/structs"
)

// StartMaster starts an Iperf listener on master server
func StartMaster(cfg st.Config, port, scantime, location, name string) (func(), st.Server, error) {
	var rmtCmd string
	var cmd *exec.Cmd
	s, err := GetServer(cfg, name, location)
	if err != nil {
		return func() {}, s, err
	}

	pFlag := "-p " + s.Port
	addr := s.User + "@" + s.VpnIP
	fmt.Printf("Starting Iperf Master in %s (%s)\n", strings.Title(name), s.LocalIP)

	isHost, err := cm.IsHost(s)
	if err != nil {
		return func() {}, s, err
	}

	if isHost {
		rmtCmd = "iperf -s -p " + port
		cmd = exec.Command(rmtCmd)
	} else {
		rmtCmd := ""
		if s.Os == "macos" {
			rmtCmd = "/usr/local/bin/iperf -s -p " + port

		} else {
			rmtCmd = "iperf -s -p " + port
		}
		cmd = exec.Command("ssh", pFlag, addr, rmtCmd)
	}

	err = cmd.Start()
	if err != nil {
		return func() {}, s, err
	}

	time.Sleep(2 * time.Second)
	cmd2 := exec.Command("ssh", pFlag, addr, "pgrep iperf")
	temp, err := cmd2.Output()
	if err != nil {
		return func() {}, s, err
	}

	pid := string(temp)
	fmt.Printf("Master started. PID: %s \n", pid)
	return func() {
		cmd := exec.Command("ssh", pFlag, addr, "sudo kill -9 "+pid)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("\nError while stopping master: %v", err)
		}
		fmt.Printf("Master stopped\n\n")
	}, s, nil
}

// ScanServers receives a config variable and a "row" and scans through the servers given that they are not the master and are in the set location
func ScanServers(mst st.Server, port, scantime, location string, cfg st.Config) []string {
	row := []string{mst.Name}
	for _, server := range cfg.Servers {
		if mst.Name == server.Name {
			row = append(row, "-")
			continue
		}
		if server.Location != location {
			continue
		}

		output, err := RunScan(server, mst, port, scantime)
		if err != nil {
			fmt.Printf("Error while running iperf server: %s", err)
		}
		row = append(row, output)
	}
	return row

}

// RunScan takes a client server, master server, iperf port and scantime and runs a bidirectional net test
func RunScan(s, m st.Server, p string, t string) (string, error) {
	fmt.Printf("Running Iperf client on %s (%s) to %s (%s)\n", strings.Title(s.Name), s.LocalIP, strings.Title(m.Name), m.LocalIP)
	var cmd *exec.Cmd
	var ret, rmtCmd string
	pFlag := "-p " + s.Port
	addr := s.User + "@" + s.VpnIP

	isHost, err := cm.IsHost(s)
	if err != nil {
		return "", err
	}
	if isHost {
		rmtCmd = "iperf -c " + m.LocalIP + " -p " + p
		cmd = exec.Command(rmtCmd)
	} else {
		if s.Os == "macos" {
			rmtCmd = "/usr/local/bin/iperf -c " + m.LocalIP + " -p " + p

		} else {
			rmtCmd = "iperf -c " + m.LocalIP + " -p " + p

		}
		cmd = exec.Command("ssh", pFlag, addr, rmtCmd)
	}
	// cmd.Stdout = os.Stdout
	out, err := cmd.StdoutPipe()
	if err != nil {
		return ret, err
	}
	err = cmd.Start()
	if err != nil {
		return ret, err
	}

	scanner := bufio.NewScanner(out)
	ret = readStuff(scanner)
	cmd.Wait() // }

	fmt.Printf("Output: %s\n", ret)
	return ret, nil
}

// CreateHeader Takes the servers config and location and returns a header to append
func CreateHeader(cfg st.Config, location string) []string {
	var header = []string{""}
	for _, server := range cfg.Servers {
		if !inLocation(location, server) {
			continue
		}
		header = append(header, server.Name)
	}
	return header

}

// GetServer adfasdf asfda
func GetServer(cfg st.Config, name, location string) (st.Server, error) {
	var s st.Server
	for _, server := range cfg.Servers {
		if !inLocation(location, server) {
			continue
		}
		if strings.ToLower(name) == strings.ToLower(server.Name) {
			s = server
			return s, nil
		}
	}
	err := errors.New("Master server not found in configuration list")
	return s, err
}

func readStuff(scanner *bufio.Scanner) string {
	var ret string
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
	return ret
}

func inLocation(l string, s st.Server) bool {
	if strings.ToLower(l) == strings.ToLower(s.Location) {
		return true
	}
	return false
}
