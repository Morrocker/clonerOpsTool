package methods

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

// SetMaster takes a server name and returns a server with the same name existing in the configuration file
func SetMaster(n, l string, c Config) (Server, error) {
	var s Server
	for _, server := range c.Servers {
		if !inLocation(l, server) {
			continue
		}
		if strings.Contains(strings.ToLower(n), strings.ToLower(server.Name)) {
			return server, nil
		}
	}
	err := errors.New("Master server not found in configuration list")
	return s, err
}

// StartMaster starts an Iperf listener on master server
func StartMaster(s Server, p string, t string) string {

	pFlag := "-p " + s.Port
	addr := s.User + "@" + s.VpnIP
	rmtCmd := "iperf -s -p " + p
	cmd := exec.Command("ssh", pFlag, addr, rmtCmd)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("ssh", pFlag, addr, "pgrep iperf")
	temp, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	pid := string(temp)
	return pid
}

func inLocation(l string, s Server) bool {
	if strings.Contains(strings.ToLower(l), strings.ToLower(s.Location)) {
		return true
	}
	return false
}
