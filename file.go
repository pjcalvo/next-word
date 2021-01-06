package main

import (
	//"bufio"
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

func getWord(level string) (word string, err error) {

	file := filepath.Join("words", level+".csv")
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
		return "", err
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()

	random := rand.Intn(len(records))
	return records[random][0], nil
}
