package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Upotreba: unesite ime fajla ili foldera")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Previse argumenata")
		return
	}

	fmt.Println(os.Args)

	fileName := os.Args[1]

	fileInfo, err := os.Stat(fileName)

	if err != nil {
		log.Println("Invalid name", err)
		return
	}

	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size())
	fmt.Println("IsDir:", fileInfo.IsDir())
	fmt.Println("ModTime:", fileInfo.ModTime())
	fmt.Println("Mode:", fileInfo.Mode())
}
