package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	//Checks if host:port is provided on the CLI
	//If not, a request is made
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	//Connect to the server in process_b
	address := arguments[1]
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//allows the user to input an email via the command line
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		//writes the message to process_b
		fmt.Fprintf(conn, text+"\n")

		//reads the message from the process_b and prints it out
		//TODO eliminate necessity of tab delimiter to signify the end of the message
		message, _ := bufio.NewReader(conn).ReadString('\t')

		fmt.Print("->: " + message)

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}

	}
}
