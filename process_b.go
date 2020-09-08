package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Message struct {
	To      string
	From    string
	Date    string
	Title   string
	Content string
}

func main() {

	//Checks if port number is provided on the CLI
	//If not, a request is made
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	//Starts a server
	address := ":" + arguments[1]
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	//Acknowledges connection to process_a
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//Reads message from process_a
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))

		//formats the message (netData) for email struct
		input := strings.Split(string(netData), ";")
		currentTime := time.Now()
		message := Message{input[0], input[1], currentTime.String(), input[2], input[3]}

		//returns formatted data to process_a
		c.Write([]byte(fmt.Sprintf("\nTitle: %s\nTo: %s\nFrom: %s\nDate: %s\nContent: %s\t", message.Title, message.To, message.From, message.Date, message.Content)))
	}
}
