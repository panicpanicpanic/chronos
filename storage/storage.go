// Package storage includes functions and methods that allow Cycle data to be stored
package storage

import (
	"encoding/csv"
	"os"
)

// Insert saves Cycle data in storage
// This storage can be a CSV file or
// an array of SQL databases (in the future)
func Insert(f *os.File, data []string) error {
	var err error

	csvWriter := csv.NewWriter(f)
	rows := [][]string{data}
	for _, row := range rows {
		err = csvWriter.Write(row)
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()
	f.Close()

	return nil
}

// WriteCSV creates a CSV file if one does
// not already exist. Returns an error if something fails
func WriteCSV() (*os.File, error) {
	var err error

	file, err := os.OpenFile("cycle_data.csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return file, nil
}
