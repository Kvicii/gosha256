package main

import (
	"fmt"
	"io"
	"os"

	"github.com/jwatson0/go/gosha256/sha2"
)

func main() {
	// read from stdin until EOF
	var f *os.File = os.Stdin

	// create buffer
	var buf []byte = make([]byte, sha2.Sha256BlocksizeBytes*2048) // large reading buffer
	var n int
	var err error

	// TODO set up sha processor with pointer to buf, start, and end indexes
	// use channel to send message with new end index when new data available
	// carefully handle buffer wraparound

	// read and loop until io.EOF
	n, err = f.Read(buf)
	for n > 0 { // while we have read any bytes

		// print any errors
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "Read error: %v", err)
		}
		if err == io.EOF {
			break
		}
		// FIXME for now, bail after one pass
		fmt.Fprintf(os.Stderr, "Warning: Data left on input queue. Read %d bytes but no EOF", n)
		// n, err = f.Read(buf[n: ??? ])
		break
	}

}
