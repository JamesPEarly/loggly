# loggly
A Go (Golang) package for communicating with Loggly using user credentials

## Setup
This package expects an environment variable called `LOGGLY_TOKEN` to be assigned a value corresponding to the Customer Token provided by Loggly. This can be found in the Source Setup under the Customer Tokens tab.

In order to include this package in your Go application, you must import it as follows
```
import loggly "github.com/jamespearly/loggly"
```

## Initialization
A connection to Loggly must be established before messages can be sent. Your program establishes its connection
by first instantiating a client. The constructor requires a string parameter used as a tag for all messages sent by your application:
```
client = loggly.New('MyApplication')
```

## Sending Messages
There are two functions provided for sending messages -- `Send` and `EchoSend`. Each function sends a valid
log message to Loggly, but `EchoSend` additionally logs the message to the local console.

Both functions require two string parameters. The first is the message `level`.
The following is a list of valid levels accepted by Loggly (in order of decreasing severity):
```
"error"
"warn"
"info"
"verbose"
"debug"
"silly"
```
The second string parameter is the `message text`. The following are valid example messages:
```
client.Send("debug", "This is a debug message");
client.EchoSend("silly", "This is a silly message that is logged to the console");
```

## Example Program
Here is an example Go program that opens a connection and sends various types of messages. It also shows
a variety of errors:
```
// This package is used to test the Loggly package
package main

import (
	"fmt"
	loggly "github.com/jamespearly/loggly"
)

func main() {

	var tag string
	tag = "My-Go-Demo"

	// Instantiate the client
	client := loggly.New(tag)

	// Valid EchoSend (message echoed to console and no error returned)
	err := client.EchoSend("info", "Good morning!")
	fmt.Println("err:", err)

	// Valid Send (no error returned)
	err = client.Send("error", "Good morning! No echo.")
	fmt.Println("err:", err)

	// Invalid EchoSend -- message level error
	err = client.EchoSend("blah", "blah")
	fmt.Println("err:", err)

}
```
