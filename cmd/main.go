package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	in, err := os.ReadFile("./speedtest.csv")

	if err != nil {
		return
	}

	r := csv.NewReader(bytes.NewReader((in)))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
