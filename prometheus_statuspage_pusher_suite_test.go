package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPrometheusStatuspagePusher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrometheusStatuspagePusher Suite")
}
