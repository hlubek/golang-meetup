package test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/api"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/scheme"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

func SetupLoggerAndAssetPath() {

	if _, assetsIsSet := os.LookupEnv("KUBEBUILDER_ASSETS"); !assetsIsSet {
		assetPath := fmt.Sprintf("../../bin/k8s/1.27.1-%s-%s", runtime.GOOS, runtime.GOARCH)
		logf.Log.Info(fmt.Sprintf("KUBEBUILDER_ASSETS is not set, setting it to %s", assetPath))
		Expect(os.Setenv("KUBEBUILDER_ASSETS", assetPath)).To(Succeed())
	}
	Expect(os.Setenv("KUBEBUILDER_ATTACH_CONTROL_PLANE_OUTPUT", "true")).To(Succeed())
}

func InitializeTestEnvWithCRDs() *envtest.Environment {
	By("bootstrapping test environment")
	err := certmanagerv1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())
	testEnv := &envtest.Environment{

		Scheme:                scheme.Scheme,
		ErrorIfCRDPathMissing: true,
		CRDDirectoryPaths:     []string{filepath.Join("..", "..", "external", "cert-manager", "crds")},
	}

	return testEnv
}

func InitializeClientSet(testEnv *envtest.Environment) *kubernetes.Clientset {
	restCfg, err := testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(restCfg).NotTo(BeNil())

	// clientSet is defined in this file globally
	clientSet, err := kubernetes.NewForConfig(restCfg)
	Expect(err).NotTo(HaveOccurred())
	Expect(clientSet).NotTo(BeNil())

	return clientSet
}
