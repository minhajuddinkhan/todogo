package utils

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

//ReadCsv ReadCsv
func ReadCsv(dir, filename string) [][]string {

	b, _ := ioutil.ReadFile(dir + filename)
	r := csv.NewReader(strings.NewReader(string(b)))

	var records [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}

	return records
}
