package ssh

import (
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

// ClientConf provides structure to create an ssh Client
type ClientConf struct {
	User, Key, IP, Port string
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	return ssh.PublicKeys(signer)
}

func runCommand(cmd string, conn *ssh.Client) {
	sess, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	sessStdOut, err := sess.StdoutPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stdout, sessStdOut)
	sessStderr, err := sess.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)
	err = sess.Run(cmd) // eg., /usr/bin/whoami
	if err != nil {
		panic(err)
	}
}

func newClient(c ClientConf) (*ssh.Client, error) {

	config := &ssh.ClientConfig{
		User: c.User,
		Auth: []ssh.AuthMethod{
			publicKey(c.Key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", c.IP+":"+c.Port, config)
	return conn, err
}

// SendSSHcomm sends an ssh command by providing a client connection config and a single command
func SendSSHcomm(c ClientConf, cmd string) {
	conn, err := newClient(c)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	runCommand(cmd, conn)
}
