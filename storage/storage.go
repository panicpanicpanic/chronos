// Package storage includes functions and methods that allow Cycle data to be stored
package storage

import (
	"encoding/csv"
	"os"
)

// Insert saves Cycle data in storage
// This storage can be a CSV file or
// an array of SQL databases (in the future)
func Insert(data []string) error {
	var err error

	file, err := os.OpenFile("cycle_data.csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	rows := [][]string{data}
	csvWriter.WriteAll(rows)
	csvWriter.Flush()

	return nil
}
