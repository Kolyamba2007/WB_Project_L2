package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	myTelnet()
}

func myTelnet() error {
	conn, err := net.DialTimeout("tcp", ":5555", 10*time.Second)

	if err != nil {
		fmt.Println(err)
		return err
	}

	go func() {
		for {
			msg, _ := bufio.NewReader(conn).ReadString('\n')
			if msg == "" {
				conn.Close()
				os.Exit(0)
			}
			fmt.Println(msg)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')
		if strings.Contains(line, string([]byte{4})) {
			conn.Close()
			return nil
		}
		conn.Write([]byte(line))
	}
}
