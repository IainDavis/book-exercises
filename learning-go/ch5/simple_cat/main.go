package main

import (
	"io"
	"log"
	"os"
)

/*
Let's take a look at how to use `defer` to release resources.
You'll do this by writing a simple version of `cat`, the Unix
utility for printing the contents of a file. You can't pass in
arguments on The Go Playground, but you can find the code for this
example in the sample_code/simple_cat directory in the Chapter 5
repository
*/
func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// defer closure of the file handle until after this method returns
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}
		break
	}
}
