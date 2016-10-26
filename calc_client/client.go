package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Input file is missing.")
		os.Exit(1)
	}
	resp, err := http.Get("http://localhost:8081/" + os.Args[1] + "?A=" + os.Args[2] + "&B=" + os.Args[3])
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		_, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
