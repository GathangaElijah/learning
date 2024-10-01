package main

import (
	"fmt"
	"io"
)

func GenerateData(writer io.Writer) {
	data := []byte("Kayak, Lifejacket")
	fmt.Println("This is the length of the data: ", len(data))
	writeSize := 4

	for i := 0; i < len(data); i += writeSize {
		end := i + writeSize
		if end > len(data) {
			end = len(data)
		}
		count, err := writer.Write(data[i:end])
		Printfln("Wrote %v byte(s): %v", count, string(data[i:end]))
		if err != nil {
			Printfln("Error: %v", err.Error())
		}
	}
	if closer, ok := writer.(io.Closer); ok{
		closer.Close()
	}
}

func ConsumeData(reader io.Reader) {
	data := make([]byte, 0, 10)
	slice := make([]byte, 2)

	for {
		count, err := reader.Read(slice)
		if (count > 0) {
			Printfln("Read data: %v", string(slice[:count]))
			data = append(data, slice[:count]...)
		}
		if err == io.EOF{
			break
		}
	}
	Printfln("Read data: %v", string(data))
}
