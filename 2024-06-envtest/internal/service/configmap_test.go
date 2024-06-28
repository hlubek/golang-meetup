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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("ApplicationCreator", func() {

	BeforeEach(func() {
		ctx = context.Background()
		logger := logrus.New()
		logger.SetOutput(io.Discard)
		ctx = log.ToContext(ctx, &logrus.Entry{
			Logger: logger,
		})

		configMapService = service.NewConfigMapService(clientSet)
	})

	Context("CreateConfigMap", func() {
		It("should not find a configmap with name cm-test", func() {
			_, err := clientSet.CoreV1().ConfigMaps(testingNamespace).Get(ctx, "cm-test", metav1.GetOptions{})
			Expect(err).To(HaveOccurred())
		})
		It("should create a configmap with name cm-test", func() {
			err := configMapService.CreateConfigMap(ctx, "cm-test")
			Expect(err).ToNot(HaveOccurred(), "creation should not error for not existing cm")
			cmTestConfigMap, err := clientSet.CoreV1().ConfigMaps(testingNamespace).Get(ctx, "cm-test", metav1.GetOptions{})
			Expect(err).ToNot(HaveOccurred())
			Expect(cmTestConfigMap.Name).To(Equal("cm-test"))
		})
	})
})
