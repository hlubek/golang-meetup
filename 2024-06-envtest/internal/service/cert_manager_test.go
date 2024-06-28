//go:build integration

package service_test

import (
	"context"
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/sirupsen/logrus"
	"golang-meetup.kiel/envtest/internal/log"
	"golang-meetup.kiel/envtest/internal/service"
)

var (
	configMapService   *service.ConfigMapService
	certmanagerService *service.CertManagerService
	ctx                context.Context
	testingNamespace   string = "default"
)
var _ = Describe("ApplicationCreator", func() {

	BeforeEach(func() {
		ctx = context.Background()
		logger := logrus.New()
		logger.SetOutput(io.Discard)
		ctx = log.ToContext(ctx, &logrus.Entry{
			Logger: logger,
		})

		certmanagerService = service.NewCertManagerService(clientSet)
	})

	Context("CreateConfigMap", func() {
		It("should create a certificate with name cert-test", func() {
			err := certmanagerService.CreateCert(ctx, "cert-test", "golang-meetup.kiel")
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
