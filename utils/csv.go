package utils

import (
	"encoding/csv"
	"os"
)

//WriteCsv WriteCsv
func WriteCsv(dir string, filename string, records [][]string) error {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}
	f, err := os.Create(dir + filename)
	if err != nil {
		return err
	}
	defer f.Close()
	csvWriter := csv.NewWriter(f)
	err = csvWriter.WriteAll(records)
	if err != nil {
		return err
	}
	csvWriter.Flush()
	return nil

}
