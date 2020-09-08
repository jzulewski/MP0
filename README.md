**MP0**


**To Run**

`go run process_b go 1234`

Open A New Terminal

`go run process_a.go 127.0.0.1:1234`

You should be prompted to enter the email field

`>>`

Input should be formatted in the following format with semicolons in-between 

To Email; From Email; Subject Line; Content



`>> zulewskj@bc.edu; zulewski@bc.edu; Greetings!; How have you been? Hope all is well!`

should produce the following output

````
 ->:
 Title:  Greetings!    
 To: zulewskj@bc.edu
 From:  zulewski@bc.edu
 Date: 2020-09-07 19:18:21.599589 -0400 EDT m=+71.453165540
 Content:  How have you been? Hope all is well!
    >>
`````

**Structure and Design**

The message uses the struct
````
type Message struct {
	To      string
	From    string
	Date    string
	Title   string
	Content string
}
````
Date made into a string with the use of
```
currentTime := time.Now()
time.String()
```

Process A is the TCP Client which:

1. Accepts the user input for the email, 

2. Sends the message to the TCP server (Process B), 

3. Prints the formatted message from the TCP server

Process B is the TCP Server which:

1. Receives the message from Process A

2. Prints the received message

3. Writes the formatted message to Process A

**Citations**

To create the TCP Server and Client
https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/

Comments in process_a.go and process_b.go provide unique explanations for how the TCP client and server are created


