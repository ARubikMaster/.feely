package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		if os.Args[1] == "convert" {
			path := os.Args[2]
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				wd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}

				_, err = os.Stat(wd + path)

				if err == nil {
					path = wd + path
				}
			}

			raw_data, width, height, err := decoder(path)
			if err != nil {
				log.Fatal(err)
			}

			//fmt.Println(raw_data)

			file_content := strconv.Itoa(width) + ";" + strconv.Itoa(height) + ";" + raw_data

			err = os.WriteFile(os.Args[3]+".feely", []byte(file_content), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	//raw_data, err := decoder()
}
