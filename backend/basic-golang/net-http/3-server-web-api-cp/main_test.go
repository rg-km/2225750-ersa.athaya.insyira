package main_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	httpserver "github.com/ruang-guru/playground/backend/basic-golang/net-http/3-server-web-api-cp"
)

var _ = Describe("GET data by ID", func() {

	Describe("GET data by ID", func() {
		It("get data by ID", func() {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/tables", nil)
			if err != nil {
				log.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()
			httpserver.TablesHandler(rec, req)

			result := rec.Result()
			defer result.Body.Close()

			if result.StatusCode != http.StatusOK {
				log.Fatalf("expected status ok 200; got %v", result.StatusCode)
			}

			b, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalf("could not read response: %v", err)
			}

			Expect(string(b)).To(Equal("Table ID:  not found"))

		})
	})
})
