package main

import (
	"fmt"
	"log"
	"os"
	)
/* 
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
*/

/*1*/type FooReader struct{}

/*2*/func (fooReader *FooReader) Read(b []byte) (int, error){
	fmt.Print("in >")
	return os.Stdin.Read(b)/*3*/
}

/*4*/type FooWriter struct{}

/*5*/func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)/*6*/

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	//Create buffer to hold input/output
	/*7*/input := make([]byte, 4096)

	s,err := reader.Read(input)/*8*/
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	s,err = writer.Write(input)/*9*/
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
