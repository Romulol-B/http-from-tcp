package main

import (
	"bytes"
	"net"
	"errors"
	"fmt"
	"io"
	"log"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)
		currentLineContents := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			data = data[:n]

			if i := bytes.IndexByte(data, '\n'); i != -1 {
				currentLineContents += string(data[:i])
				data = data[i+1:]
				lines <- currentLineContents
				currentLineContents = ""
			}
			currentLineContents += string(data)

		}
		if len(currentLineContents) != 0 {
			lines <- currentLineContents
		}
	}()

	return lines
}
func main() {
	port := ":42069"
	ear,err := net.Listen("tcp",port)

	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer ear.Close()
	for {
		connection , err := ear.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
		fmt.Println("Connection accepted")


		/*
		path := "message.txt"
		file, err := os.Open(path)
		if err != nil {
			log.Fatal("error", "error", err)
		}
		*/
		lines := getLinesChannel(connection)
		for line := range lines {
			fmt.Printf("%s\n",line)
		}
	fmt.Println("The connection is being closed ")
}
}
