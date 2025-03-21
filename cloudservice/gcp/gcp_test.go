// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.

package gcp

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGCP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudService GCP Suite")
}

var _ = Describe("GCP Parsing", func() {
	Context("Parse GCP IP Range file", func() {
		It("should parse the ip range file", func() {
			ipRangeFile := "{\"syncToken\":\"1000000000\",\"createDate\":\"2022-09-01-22-03-06\",\"prefixes\":[{\"ipv4Prefix\": \"34.80.0.0/15\"},{\"ipv4Prefix\": \"5.80.0.0/15\"},{\"ipv4Prefix\": \"150.81.0.0/15\"},{\"ipv4Prefix\": \"127.80.0.0/15\"}]}"
			gcpManager := New()

			info, err := gcpManager.ConvertToGenericIPRanges([]byte(ipRangeFile))

			By("Ensuring that no error was thrown")
			Expect(err).ToNot(HaveOccurred())

			By("Ensuring that the right version string was parsed")
			Expect(info.Version).To(Equal("1000000000"))

			By("Ensuring that we have the right info")
			Expect(info.IPRanges[GoogleCloudService]).To(HaveLen(4))
		})

		It("should remove 8.8.8.8 of the ip range file", func() {
			ipRangeFile := "{\"syncToken\":\"1000000000\",\"createDate\":\"2022-09-01-22-03-06\",\"prefixes\":[{\"ipv4Prefix\": \"8.8.8.0/15\"},{\"ipv4Prefix\": \"5.80.0.0/15\"},{\"ipv4Prefix\": \"150.81.0.0/15\"},{\"ipv4Prefix\": \"127.80.0.0/15\"}]}"
			gcpManager := New()

			info, err := gcpManager.ConvertToGenericIPRanges([]byte(ipRangeFile))

			By("Ensuring that no error was thrown")
			Expect(err).ToNot(HaveOccurred())

			By("Ensuring that we have the right info")
			Expect(info.IPRanges[GoogleCloudService]).To(HaveLen(3))
		})
	})

	Context("Verify GCP New version of the file", func() {
		ipRangeFile := "{\"syncToken\":\"1000000000\",\"createDate\":\"2022-09-01-22-03-06\",\"prefixes\":[{\"ipv4Prefix\": \"34.80.0.0/15\"},{\"ipv4Prefix\": \"5.80.0.0/15\"},{\"ipv4Prefix\": \"150.81.0.0/15\"},{\"ipv4Prefix\": \"127.80.0.0/15\"}]}"
		gcpManager := New()

		isNewVersion, err := gcpManager.IsNewVersion([]byte(ipRangeFile), "20")

		It("Should indicate is a new version", func() {
			By("Ensuring that no error was thrown")
			Expect(err).ToNot(HaveOccurred())

			By("Ensuring that the version is new")
			Expect(isNewVersion).To(BeTrue())
		})
	})

	Context("Verify GCP handle of errors", func() {
		It("Should throw an error on empty ip ranges file", func() {
			ipRangeFile := ""
			gcpManager := New()

			_, errConvert := gcpManager.ConvertToGenericIPRanges([]byte(ipRangeFile))
			_, errIsNewVersion := gcpManager.IsNewVersion([]byte(ipRangeFile), "20")

			By("Ensuring that an error was thrown on ConvertToGenericIPRanges")
			Expect(errConvert).To(HaveOccurred())

			By("Ensuring that an error was thrown on IsNewVersion")
			Expect(errIsNewVersion).To(HaveOccurred())
		})

		It("Should throw an error on empty ip ranges file", func() {
			gcpManager := New()

			_, errConvert := gcpManager.ConvertToGenericIPRanges(make([]byte, 0))
			_, errIsNewVersion := gcpManager.IsNewVersion(make([]byte, 0), "20")

			By("Ensuring that an error was thrown on ConvertToGenericIPRanges")
			Expect(errConvert).To(HaveOccurred())

			By("Ensuring that an error was thrown on IsNewVersion")
			Expect(errIsNewVersion).To(HaveOccurred())
		})

		It("Should throw an error on nil ip ranges file", func() {
			gcpManager := New()

			_, errConvert := gcpManager.ConvertToGenericIPRanges(nil)
			_, errIsNewVersion := gcpManager.IsNewVersion(nil, "20")

			By("Ensuring that an error was thrown on ConvertToGenericIPRanges")
			Expect(errConvert).To(HaveOccurred())

			By("Ensuring that an error was thrown on IsNewVersion")
			Expect(errIsNewVersion).To(HaveOccurred())
		})
	})
})
