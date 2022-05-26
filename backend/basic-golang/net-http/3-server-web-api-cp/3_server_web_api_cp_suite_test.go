package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test3ServerWebApiCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "3ServerWebApiCp Suite")
}
