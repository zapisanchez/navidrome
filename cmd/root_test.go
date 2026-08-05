package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("mainContext", func() {
	It("does not cancel when it receives SIGHUP", func() {
		signal.Reset(syscall.SIGHUP)
		DeferCleanup(signal.Ignore, syscall.SIGHUP)

		ctx, cancel := mainContext(context.Background())
		DeferCleanup(cancel)

		Expect(signal.Ignored(syscall.SIGHUP)).To(BeTrue())
		process, err := os.FindProcess(os.Getpid())
		Expect(err).ToNot(HaveOccurred())
		Expect(process.Signal(syscall.SIGHUP)).To(Succeed())
		Consistently(ctx.Done(), 200*time.Millisecond).ShouldNot(Receive())
	})
})
