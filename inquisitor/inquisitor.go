package inquisitor

import (
	"encoding/csv"
	"io"
	"fmt"
	"net/http"
	"strings"
	"errors"
)

type Inquisitor interface {
	ReadCsv() ([][]string, error)
	ConnectionArrayFromCsv(csvContent [][]string) []string
}

type HttpInquisitor struct {
	username 	string
	password 	string
	ip			string
}

func NewHttpInquisitor(username string, password string, ip string) Inquisitor {
	return &HttpInquisitor{
		username : username,
		password : password,
		ip : ip,
	}
}

func (iq *HttpInquisitor) ReadCsv() ([][]string, error) {

	resp, err := http.Get(iq.connectionString())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	parsed := parseCsv(resp.Body)

	return parsed, nil
}

func (iq *HttpInquisitor) ConnectionArrayFromCsv(csvContent [][]string) []string {
	// Find index of 'scur' in first line (we know first line is header)
	header := csvContent[0]
	scurIdx, err := findIndex(header, "scur")
	if err != nil {
		panic(err)
	}

	// Find index of 'svname' in first line
	var svnameIdx int
	svnameIdx, err = findIndex(header, "svname")
	if err != nil {
		panic(err)
	}

	// Find the lines that contain 'mysql-x' at the index of svname
	var lineIdx []int
	for idx,element := range csvContent {
		if strings.HasPrefix(element[svnameIdx],"mysql-") {
			lineIdx = append(lineIdx, idx)
		}
	}

	// For these lines, store the value at 'scur'
	conns := make([]string, len(lineIdx))
	for idx,element := range lineIdx {
		line := csvContent[element]
		conns[idx] = line[scurIdx]
	}
	return conns
}


//private
func findIndex(array []string, key string) (int, error) {
	for idx,element := range array {
		if element == key {
			return idx, nil
		}
	}
	return -1, errors.New(key + "not found in array")
}

func parseCsv(r io.Reader) [][]string {

	result, err := csv.NewReader(r).ReadAll()
	if err != nil {
		panic(err)
	}
	return result
}

func (iq *HttpInquisitor) connectionString() string {
	return fmt.Sprintf("http://%s:%s@%s:1936/;csv", iq.username, iq.password, iq.ip)
}
