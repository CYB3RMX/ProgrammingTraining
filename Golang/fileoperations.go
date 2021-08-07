package main

import (
	"bufio"
	"fmt"
	"os"
)

func WriteToFiles(file_name string, buffer_write string) {
	/*
		This block creates file with writable mode if it doesnt exist
		os.O_WRONLY => Write only
		os.O_APPEND => Append mode
		os.O_CREATE => Create file if it doesnt exist
		0755 => File permissions
	*/
	fhandler_write, err := os.OpenFile(file_name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer fhandler_write.Close() // Close the file when everything ends
	fmt.Fprintln(fhandler_write, buffer_write)
}

func ReadFromFile(file_name string) {
	/*
		os.O_RDONLY => Read only
	*/
	var counter int
	counter = 0
	fhandler_read, err := os.OpenFile(file_name, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer fhandler_read.Close()
	reader := bufio.NewScanner(fhandler_read)
	for reader.Scan() {
		fmt.Printf("\nLine => %d\t%s\n", counter, reader.Text())
		counter++
	}
}

func MetaDataExtract(file_name string) {
	fhandler_meta, err := os.Stat(file_name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nFile Name: %s\n", fhandler_meta.Name())
	fmt.Printf("File Size: %d\n", fhandler_meta.Size())
	fmt.Printf("File Permissions: %s\n", fhandler_meta.Mode())
	fmt.Printf("File TimeDateStamp: %s\n\n", fhandler_meta.ModTime())
}

func main() {
	// Variables
	var fname_read string
	var fname_write string

	// File reader
	fmt.Printf("Enter file name to read: ")
	fmt.Scanln(&fname_read)
	ReadFromFile(fname_read)

	// File writer
	fmt.Printf("\nEnter file name to write: ")
	fmt.Scanln(&fname_write)
	WriteToFiles(fname_write, "is it appended?")

	// Meta data extractor
	MetaDataExtract(fname_read)
}
