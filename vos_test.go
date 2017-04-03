package vos

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/virtual-go/fs"
	"os"
)

var _ = Context("VOS", func() {
	Context("New", func() {
		It("should return VOS", func() {
			/* arrange/act/assert */
			Expect(New(new(fs.Fake))).
				Should(Not(BeNil()))
		})
	})
	Context("FindProcess", func() {
		It("should return expected process", func() {
			/* arrange */
			providedPID := os.Getpid()
			expectedProcess, _ := os.FindProcess(providedPID)
			objectUnderTest := _VOS{}

			/* act */
			actualProcess, actualErr := objectUnderTest.FindProcess(providedPID)

			/* assert */
			Expect(actualProcess).To(Equal(expectedProcess))
			Expect(actualErr).To(BeNil())
		})
	})
	Context("Getpid", func() {
		It("should return expected PID", func() {
			/* arrange */
			expectedPID := os.Getpid()
			objectUnderTest := _VOS{}

			/* act */
			actualPID := objectUnderTest.Getpid()

			/* assert */
			Expect(actualPID).To(Equal(expectedPID))
		})
	})
	Context("Getenv proceeding Setenv", func() {
		It("should return value set by Setenv", func() {
			/* arrange */
			providedName := "dummyName"
			providedValue := "dummyValue"
			expectedValue := providedValue
			objectUnderTest := _VOS{}

			objectUnderTest.Setenv(providedName, providedValue)

			/* act */
			actualValue := objectUnderTest.Getenv(providedName)

			/* assert */
			Expect(actualValue).To(Equal(expectedValue))
		})
	})
	Context("Getwd", func() {
		It("should return expected process", func() {
			/* arrange */
			expectedWd, _ := os.Getwd()
			objectUnderTest := _VOS{}

			/* act */
			actualWd, actualErr := objectUnderTest.Getwd()

			/* assert */
			Expect(actualWd).To(Equal(expectedWd))
			Expect(actualErr).To(BeNil())
		})
	})
})
