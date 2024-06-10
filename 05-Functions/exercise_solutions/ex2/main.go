package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func fileLen(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	data := make([]byte, 2048)
	total := 0
	for {
		count, err := file.Read(data)
		total += count
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	return total, nil
}
