package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"block-csv2json/internal/validate"
)

type Block struct {
	BlockNumber int    `json:"block_number"`
	Timestamp   string `json:"timestamp"`
	TxCount     int    `json:"tx_count"`
	TxHexSample string `json:"tx_hex_sample"`
}

func main() {
	filePath := flag.String("file", "", "path to CSV file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide --file flag")
		os.Exit(1)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	var blocks []Block

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		blockNumber, err := validate.Int(row[0], "block_number")
		if err != nil {
			fmt.Println(err)
			continue
		}

		txCount, err := validate.Int(row[2], "tx_count")
		if err != nil {
			fmt.Println(err)
			continue
		}

		block := Block{
			BlockNumber: blockNumber,
			Timestamp:   row[1],
			TxCount:     txCount,
			TxHexSample: row[3],
		}

		blocks = append(blocks, block)
	}

	output, err := json.MarshalIndent(blocks, "", "  ")
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

