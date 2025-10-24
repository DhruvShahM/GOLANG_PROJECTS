package main

import (
	"fmt"
	"log"
	"os"

	"qr-code-generator/internal/db"
	"qr-code-generator/internal/factory"
	"qr-code-generator/internal/generator"
	"qr-code-generator/internal/scanner"
)

func main() {
	db.InitDB()

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/main.go")
	}

	command := os.Args[1]

	switch command {
	case "generate":
		gen := factory.QRFactory("standard")
		file, err := gen.Generate(os.Args[2], nil)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Standard QR generated at:", file)

	case "custom":
		gen := factory.QRFactory("custom")
		opts := map[string]string{}
		if len(os.Args) > 3 {
			opts["color"] = os.Args[3]
		}
		file, err := gen.Generate(os.Args[2], opts)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Custom QR generated at:", file)

	case "batch":
		gen := &generator.BatchQRGenerator{}
		opts := map[string]string{}
		if len(os.Args) > 3 {
			opts["format"] = os.Args[3]
		} else {
			opts["format"] = "csv"
		}
		_, err := gen.Generate(os.Args[2], opts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Batch QR generated")

	case "scan":
		content, err := scanner.Scan(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Scanned QR Content:", content)

	// case "cache":
	// 	if len(os.Args) < 3 {
	// 		log.Fatal("usage : go run cmd/main.go cache ")
	// 	}
	// 	data := os.Args[2]
	// 	record, found = cache.GetFromCache(data)

	// 	if found {
	// 		fmt.Println("Cache HIT:", record)
	// 	} else {
	// 		fmt.Println("Cache MISS for:", data)
	// 	}

	default:
		fmt.Println("Unknown command")
	}

}
