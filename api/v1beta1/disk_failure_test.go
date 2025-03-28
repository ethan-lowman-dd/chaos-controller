// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.

package v1beta1_test

import (
	"fmt"
	. "github.com/DataDog/chaos-controller/api/v1beta1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/rand"
)

var _ = Describe("DiskFailureSpec", func() {
	When("Call the 'Validate' method", func() {
		DescribeTable("success cases",
			func(diskFailureSpec DiskFailureSpec) {
				// Action && Assert
				Expect(diskFailureSpec.Validate()).Should(Succeed())
			},
			Entry("with a valid path not exceeding 62 characters",
				DiskFailureSpec{
					Paths: []string{randStringRunes(rand.IntnRange(1, 62)), randStringRunes(rand.IntnRange(1, 62))},
				},
			),
			Entry("with a valid path containing spaces",
				DiskFailureSpec{
					Paths: []string{"   " + randStringRunes(rand.IntnRange(61, 62)) + "   ", randStringRunes(rand.IntnRange(1, 62))},
				},
			),
		)

		pathGreaterThan62Characters := randStringRunes(rand.IntnRange(63, 10000))

		DescribeTable("error cases",
			func(df DiskFailureSpec, expectedError string) {
				// Action
				err := df.Validate()

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(Equal(expectedError))
			},
			Entry("with a path exceeding 62 characters",
				DiskFailureSpec{
					Paths: []string{randStringRunes(rand.IntnRange(1, 62)), pathGreaterThan62Characters},
				},
				fmt.Sprintf("the path of the disk failure disruption must not exceed 62 characters, found %d", len(pathGreaterThan62Characters)),
			),
			Entry("with an empty path",
				DiskFailureSpec{
					Paths: []string{""},
				},
				"the path of the disk failure disruption must not be empty",
			),
			Entry("with a blank path",
				DiskFailureSpec{
					Paths: []string{randStringRunes(rand.IntnRange(1, 62)), "   "},
				},
				"the path of the disk failure disruption must not be empty",
			),
		)
	})

	When("Call the 'GenerateArgs' method", func() {
		DescribeTable("success cases",
			func(diskFailureSpec DiskFailureSpec, expectedArgs []string) {
				// Arrange
				expectedArgs = append([]string{"disk-failure"}, expectedArgs...)

				//	Action
				args := diskFailureSpec.GenerateArgs()

				// Assert
				Expect(args).Should(Equal(expectedArgs))
			},
			Entry("with a '/' path",
				DiskFailureSpec{
					Paths: []string{"/"},
				},
				[]string{"--path", "/"},
			),
			Entry("with a '/sub/path/'",
				DiskFailureSpec{
					Paths: []string{"/sub/path/"},
				},
				[]string{"--path", "/sub/path/"},
			),
			Entry("with multiple paths",
				DiskFailureSpec{
					Paths: []string{"/path-1", "/path-2"},
				},
				[]string{"--path", "/path-1", "--path", "/path-2"},
			),
			Entry("with a path containing spaces",
				DiskFailureSpec{
					Paths: []string{"  /  "},
				},
				[]string{"--path", "/"},
			),
			Entry("with an EACCES exit code",
				DiskFailureSpec{
					Paths:         []string{"/"},
					OpenatSyscall: &OpenatSyscallSpec{ExitCode: "EACCES"},
				},
				[]string{"--path", "/", "--exit-code", "EACCES"},
			),
			Entry("with an empty exit code",
				DiskFailureSpec{
					Paths:         []string{"/"},
					OpenatSyscall: &OpenatSyscallSpec{},
				},
				[]string{"--path", "/"},
			),
		)
	})

	Describe("OpenatSyscallSpec", func() {
		DescribeTable("Call the 'GetExitCode' method",
			func(exitCode string, expectedExitCode int) {
				openatSyscallSpec := OpenatSyscallSpec{
					ExitCode: exitCode,
				}
				Expect(openatSyscallSpec.GetExitCodeInt()).Should(Equal(expectedExitCode))
			},
			Entry("EACCES exit code",
				"EACCES",
				13,
			),
			Entry("EDQUOT exit code",
				"EDQUOT",
				122,
			),
			Entry("EEXIST exit code",
				"EEXIST",
				17,
			),
			Entry("EFAULT exit code",
				"EFAULT",
				14,
			),
			Entry("EFBIG exit code",
				"EFBIG",
				27,
			),
			Entry("EINTR exit code",
				"EINTR",
				4,
			),
			Entry("EISDIR exit code",
				"EISDIR",
				21,
			),
			Entry("ELOOP exit code",
				"ELOOP",
				40,
			),
			Entry("EMFILE exit code",
				"EMFILE",
				24,
			),
			Entry("ENAMETOOLONG exit code",
				"ENAMETOOLONG",
				36,
			),
			Entry("ENFILE exit code",
				"ENFILE",
				23,
			),
			Entry("ENODEV exit code",
				"ENODEV",
				19,
			),
			Entry("ENOENT exit code",
				"ENOENT",
				2,
			),
			Entry("ENOMEM exit code",
				"ENOMEM",
				12,
			),
			Entry("ENOSPC exit code",
				"ENOSPC",
				28,
			),
			Entry("ENOTDIR exit code",
				"ENOTDIR",
				20,
			),
			Entry("ENXIO exit code",
				"ENXIO",
				6,
			),
			Entry("EOVERFLOW exit code",
				"EOVERFLOW",
				75,
			),
			Entry("EPERM exit code",
				"EPERM",
				1,
			),
			Entry("EROFS exit code",
				"EROFS",
				30,
			),
			Entry("ETXTBSY exit code",
				"ETXTBSY",
				26,
			),
			Entry("EWOULDBLOCK exit code",
				"EWOULDBLOCK",
				11,
			),
		)
	})
})

func randStringRunes(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
