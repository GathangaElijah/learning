package main

import "io"

type CustomReader struct{
	reader io.Reader
	readCount int
}

func NewCustomReader(reader io.Reader) *CustomReader{
	return &CustomReader{reader, 0}
}

func (cr *CustomReader) Read(slice []byte) (count int, err error){
	count, err = cr.reader.Read(slice)
	cr.readCount++
	Printfln("Custom Reader: %v bytes", count)
	if err == io.EOF{
		Printfln("Total Reads: %v", cr.readCount)
	}
	return
}