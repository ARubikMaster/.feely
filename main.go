package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		if os.Args[1] == "convert" {
			raw_data, width, height, err := decoder(os.Args[2])
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
