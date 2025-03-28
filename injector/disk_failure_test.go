// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.

package injector_test

import (
	"github.com/DataDog/chaos-controller/command"
	"os"
	"strconv"

	"github.com/DataDog/chaos-controller/api"
	v1beta1 "github.com/DataDog/chaos-controller/api/v1beta1"
	"github.com/DataDog/chaos-controller/container"
	. "github.com/DataDog/chaos-controller/injector"
	"github.com/DataDog/chaos-controller/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Disk Failure", func() {
	var (
		config         DiskFailureInjectorConfig
		level          types.DisruptionLevel
		proc           *os.Process
		inj            Injector
		spec           v1beta1.DiskFailureSpec
		cmdFactoryMock *command.FactoryMock
		containerMock  *container.ContainerMock
	)

	const PID = 1

	BeforeEach(func() {
		proc = &os.Process{Pid: PID}

		containerMock = container.NewContainerMock(GinkgoT())

		cmd := command.NewCmdMock(GinkgoT())
		cmd.EXPECT().DryRun().Return(false).Maybe()
		cmd.EXPECT().Start().Return(nil).Maybe()
		cmd.EXPECT().Wait().Return(nil).Maybe()
		cmd.EXPECT().PID().Return(41).Maybe()
		cmdFactoryMock = command.NewFactoryMock(GinkgoT())
		cmdFactoryMock.EXPECT().NewCmd(mock.Anything, mock.Anything, mock.Anything).Return(cmd).Maybe()

		config = DiskFailureInjectorConfig{
			Config: Config{
				Log:         log,
				MetricsSink: ms,
				Disruption: api.DisruptionArgs{
					Level: level,
				},
				TargetContainer: containerMock,
			},
			CmdFactory: cmdFactoryMock,
		}

		spec = v1beta1.DiskFailureSpec{
			Paths: []string{"/"},
		}
	})

	Describe("injection", func() {
		JustBeforeEach(func() {
			// instantiate lately so config can be updated in BeforeEach
			var err error
			inj, err = NewDiskFailureInjector(spec, config)

			Expect(err).ToNot(HaveOccurred())

			Expect(inj.Inject()).To(Succeed())
		})

		Context("with a pod level", func() {
			BeforeEach(func() {
				config.Disruption.Level = types.DisruptionLevelPod

				containerMock.EXPECT().PID().Return(PID).Once()
			})

			It("should start the eBPF Disk failure program", func() {
				cmdFactoryMock.AssertCalled(GinkgoT(), "NewCmd", mock.Anything, EBPFDiskFailureCmd, []string{
					"-p", strconv.Itoa(proc.Pid),
					"-f", "/",
				})
			})

			Context("with multiple valid paths", func() {
				BeforeEach(func() {
					spec.Paths = []string{"/test", "/toto"}
				})

				It("should run two eBPF program per paths", func() {
					cmdFactoryMock.AssertCalled(GinkgoT(), "NewCmd", mock.Anything, EBPFDiskFailureCmd, []string{
						"-p", strconv.Itoa(proc.Pid),
						"-f", "/test",
					})
					cmdFactoryMock.AssertCalled(GinkgoT(), "NewCmd", mock.Anything, EBPFDiskFailureCmd, []string{
						"-p", strconv.Itoa(proc.Pid),
						"-f", "/toto",
					})
				})
			})

			Context("with custom OpenatSyscall exit code", func() {
				BeforeEach(func() {
					spec.OpenatSyscall = &v1beta1.OpenatSyscallSpec{ExitCode: "EACCES"}
				})

				It("should start with a valid exit code", func() {
					cmdFactoryMock.AssertCalled(GinkgoT(), "NewCmd", mock.Anything, EBPFDiskFailureCmd, []string{
						"-p", strconv.Itoa(proc.Pid),
						"-f", "/",
						"-c", "13",
					})
				})
			})

			Context("with an empty custom OpenatSyscall exit code", func() {
				BeforeEach(func() {
					spec.OpenatSyscall = &v1beta1.OpenatSyscallSpec{}
				})

				It("should start with a valid exit code", func() {
					cmdFactoryMock.AssertCalled(GinkgoT(), "NewCmd", mock.Anything, EBPFDiskFailureCmd, []string{
						"-p", strconv.Itoa(proc.Pid),
						"-f", "/",
					})
				})
			})
		})

		Context("with a node level", func() {
			BeforeEach(func() {
				config.Disruption.Level = types.DisruptionLevelNode
			})

			It("should start the eBPF Disk failure program", func() {
				containerMock.AssertNumberOfCalls(GinkgoT(), "PID", 0)
				cmdFactoryMock.AssertCalled(GinkgoT(), "NewCmd", mock.Anything, EBPFDiskFailureCmd, []string{
					"-p", strconv.Itoa(0),
					"-f", "/",
				})
			})

			Context("with custom OpenatSyscall exit code", func() {
				BeforeEach(func() {
					spec.OpenatSyscall = &v1beta1.OpenatSyscallSpec{ExitCode: "EEXIST"}
				})

				It("should start with a valid exit code", func() {
					cmdFactoryMock.AssertCalled(GinkgoT(), "NewCmd", mock.Anything, EBPFDiskFailureCmd, []string{
						"-p", strconv.Itoa(0),
						"-f", "/",
						"-c", "17",
					})
				})
			})

			Context("with an empty custom OpenatSyscall exit code", func() {
				BeforeEach(func() {
					spec.OpenatSyscall = &v1beta1.OpenatSyscallSpec{}
				})

				It("should start with a valid exit code", func() {
					cmdFactoryMock.AssertCalled(GinkgoT(), "NewCmd", mock.Anything, EBPFDiskFailureCmd, []string{
						"-p", strconv.Itoa(0),
						"-f", "/",
					})
				})
			})
		})
	})
})
