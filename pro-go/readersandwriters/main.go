package main

import (
	"io"
	// "strings"
)

/*
Understanding readers
The io package provides abstract ways to read and write data without being
tied to where the data is coming from or going to.
- The Read(byteslice) it read data into the specified []byte. It returns the
number of bytes that were read expressed as an int and an error.
*/

func processData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("read %v bytes: %v", count, string(b[:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

/*
Understanding Writers
It defines a method write(byteslice) that writes data from the specified
byte slice. The method returns the number of bytes that were written and an error.
The error will be non-nil if the number of the bytes written are less than
the length of the slice.
*/

func writeData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	// r := strings.NewReader("kayak")

	// This function call demonstrates how the reader works
	// processData(r)

	// This demonstrates how the writer works
	// var builder strings.Builder
	// writeData(r, &builder)
	// Printfln("String builder contents: %s", builder.String())

	// This now demostrates how the pipe function from the io package works
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)

	ConsumeData(pipeReader)
}
