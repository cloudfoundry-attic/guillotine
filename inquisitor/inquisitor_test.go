package inquisitor_test

import (
	"encoding/base64"
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"os"
	"encoding/csv"

	"github.com/cloudfoundry-incubator/guillotine/inquisitor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
func init() {

	http.HandleFunc("/;csv", func (w http.ResponseWriter, req *http.Request) {
			basicAuthHeader := strings.Split(req.Header.Get("Authorization"), " ")[1]
			decodedAuthHeader, _ := base64.StdEncoding.DecodeString(basicAuthHeader)
			Expect(string(decodedAuthHeader)).To(Equal("admin:password"))

			csvStuff, _ := ioutil.ReadFile("../fixtures/simple_csv.csv")
			fmt.Fprintf(w, "%s", csvStuff)
		})
	go http.ListenAndServe(":1936", nil)
}

var _ = Describe("Inquisitor", func() {
	Describe("ReadCsv", func(){
		It("Returns a nested array with the parsed csv content", func(){
			iq := inquisitor.NewHttpInquisitor("admin", "password", "localhost")
			expected_result := [][]string{
				[]string{ "foo", "bar",	},
				[]string{ "1", "2", },
			}

			result, _ := iq.ReadCsv()
			Expect(result).To(Equal(expected_result))
		})

		Context("When http.Get returns an error", func() {
			It("Returns the error", func(){
				iq := inquisitor.NewHttpInquisitor("this", "will", "fail")
				_, err := iq.ReadCsv()
				Expect(err).To(HaveOccurred())
			})
		})

	})

	Describe("ConnectionArrayFromCsv", func(){
		It("Returns an array of connections to each node", func(){
				iq := inquisitor.NewHttpInquisitor("admin", "password", "localhost")
				csvFile, _ := os.Open("../fixtures/example_haproxy.csv")
				csvData, _ := csv.NewReader(csvFile).ReadAll()
				expected_result := []string{"1", "0", "0"}
				Expect(iq.ConnectionArrayFromCsv(csvData)).To(Equal(expected_result))
			})
	})
})
