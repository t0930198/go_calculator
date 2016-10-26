package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var op, ip string = "", ""
	var op1, op2 string = "", ""
	if len(os.Args) < 2 {
		input(&op, &op1, &op2, &ip)
	} else {
		op = os.Args[1]
		op1 = os.Args[2]
		op2 = os.Args[3]
		ip = "localhost"
	}
	fmt.Println("http://" + ip + ":3000/" + op + "?A=" + op1 + "&B=" + op2)
	resp, err := http.Get("http://" + ip + ":3000/" + op + "?A=" + op1 + "&B=" + op2)
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
func input(op, op1, op2, ip *string) {
	var err error
	fmt.Println("Input operator.")
	_, err = fmt.Scanf("%s\n", op)
	fmt.Println("Input operand1.")
	_, err = fmt.Scanf("%s\n", op1)
	fmt.Println("Input operand2.")
	_, err = fmt.Scanf("%s\n", op2)
	fmt.Println("Input IP address.")
	_, err = fmt.Scanf("%s\n", ip)
	if err != nil {
		fmt.Println("input err")
	}
}
