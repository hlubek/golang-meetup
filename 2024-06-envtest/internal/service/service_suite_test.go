//go:build integration

package service_test

import (
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"golang-meetup.kiel/envtest/internal/test"
)

var (
	testEnv   *envtest.Environment
	clientSet kubernetes.Interface
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = BeforeSuite(func() {
	test.SetupLoggerAndAssetPath()

	testEnv = test.InitializeTestEnvWithCRDs()

	clientSet = test.InitializeClientSet(testEnv)
})

var _ = AfterSuite(func() {
	defer GinkgoRecover()
	gexec.KillAndWait(8 * time.Second)

	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
	Expect(os.Unsetenv("KUBEBUILDER_ASSETS")).To(Succeed())
})
