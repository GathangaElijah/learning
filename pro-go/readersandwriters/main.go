package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
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
	/*
		r := strings.NewReader("kayak")

		This function call demonstrates how the reader works
		processData(r)
	*/

	/*
		This demonstrates how the writer works
		var builder strings.Builder
		writeData(r, &builder)
		Printfln("String builder contents: %s", builder.String())
	*/

	/*
		This now demostrates how the pipe function from the io package works
		pipeReader, pipeWriter := io.Pipe()
		go GenerateData(pipeWriter)

		ConsumeData(pipeReader)
	*/

	/*
		Concatenating multiple readers
		The MultiReader function concentrates the input from multiple readers
		so they can be processed in the sequence.
		- The reader returned by the Multireader function responds to the
		read method with the content from one of the underlying Reader values.
		- When one reader returns and EOF the content is read from the second reader
		until it returns an EOF and continues until the last reader returns an EOF.
		r1 := strings.NewReader("Kayak")
		r2 := strings.NewReader("LifeJacket")
		r3 := strings.NewReader("Canoe")

		concatReader := io.MultiReader(r1, r2, r3)

		ConsumeData(concatReader)
	*/

	/*
		Combining Multiple Writers
		- The multiple writer function combines multiple writers so that data is sent
		to all of them.

		var w1 strings.Builder
		var w2 strings.Builder
		var w3 strings.Builder

		combinedWriter := io.MultiWriter(&w1, &w2, &w3)

		GenerateData(combinedWriter)

		Printfln("Writer #1: %v", w1.String())
		Printfln("Writer #2: %v", w2.String())
		Printfln("Writer #3: %v", w3.String())
	*/

	/*
		Echoing Reads to a Writer
		- TeeReader function returns a Reader that echoes the data that it receives
		to a writer.
	*/

	/*
		Buffering Data
		Bufio pakage has the fuctions that implement the buffered reads.
		NewReader(r)- returns a buffered reader with the default buffer size.
		NewReaderSize(r, size)- This function returns a buffered Reader with the
		specified buffer size.
		- This functions reduce the number of operations made to the underlying
		data source.
		text := "It was a boat. A small boat."
		var reader io.Reader = NewCustomReader(strings.NewReader(text))
		var writer strings.Builder
		slice := make([]byte, 5)

		reader = bufio.NewReader(reader)

		for {
			count, err := reader.Read(slice)
			if count > 0{
				writer.Write(slice[:count])
			}
			if err != nil{
				break
			}
		}
		Printfln("Read data: %v", writer.String())
	*/

	/*
	Decoding data
	*/
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	fmt.Println(vals...)
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
}
