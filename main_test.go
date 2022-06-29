package main

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	When("Startup", func() {
		Context("Environment Variables", func() {
			BeforeEach(func() {
				os.Setenv("PROM", "http://something.else")
				os.Setenv("APIKEY", "abcdefg")
				os.Setenv("PAGEID", "123456789")
				os.Setenv("CONFIG", "queries.yml")
				os.Setenv("INTERVAL", "10s")
				os.Setenv("LOGLEVEL", "debug")
			})

			It("returns the prom env var", func() {
				test := *getEnvOrFlag("prom", "http://localhost:9090", "")
				Expect(test).To(Equal("http://something.else"))
			})

			It("returns the apikey env var", func() {
				test := *getEnvOrFlag("apikey", "", "")
				Expect(test).To(Equal("abcdefg"))
			})

			It("returns the pageid env var", func() {
				test := *getEnvOrFlag("pageid", "", "")
				Expect(test).To(Equal("123456789"))
			})

			It("returns the config env var", func() {
				test := *getEnvOrFlag("config", "queries.yaml", "")
				Expect(test).To(Equal("queries.yml"))
			})

			It("returns the interval env var", func() {
				test := *getEnvOrFlag("interval", "300s", "")
				Expect(test).To(Equal("10s"))
			})

			It("returns the loglevel env var", func() {
				test := *getEnvOrFlag("loglevel", "info", "")
				Expect(test).To(Equal("debug"))
			})
		})

		Context("Flags", func() {
			BeforeEach(func() {
				os.Unsetenv("PROM")
				os.Unsetenv("APIKEY")
				os.Unsetenv("PAGEID")
				os.Unsetenv("CONFIG")
				os.Unsetenv("INTERVAL")
				os.Unsetenv("LOGLEVEL")
			})

			It("returns the default prom flag", func() {
				Expect(*prometheusURL).To(Equal("http://localhost:9090"))
			})

			It("returns the default apikey flag", func() {
				Expect(*statusPageAPIKey).To(Equal(""))
			})

			It("returns the default pageid flag", func() {
				Expect(*statusPageID).To(Equal(""))
			})

			It("returns the default config flag", func() {
				Expect(*queryConfigFile).To(Equal("queries.yaml"))
			})

			It("returns the default interval flag", func() {
				Expect(*metricInterval).To(Equal("300s"))
			})

			It("returns the default loglevel flag", func() {
				Expect(*logLevel).To(Equal("info"))
			})
		})
	})
})
