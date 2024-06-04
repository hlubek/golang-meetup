package test

import (
	"fmt"
	"os"
	"runtime"

	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

func SetupLoggerAndAssetPath() {

	if _, assetsIsSet := os.LookupEnv("KUBEBUILDER_ASSETS"); !assetsIsSet {
		assetPath := fmt.Sprintf("../../bin/k8s/1.27.1-%s-%s", runtime.GOOS, runtime.GOARCH)
		logf.Log.Info(fmt.Sprintf("KUBEBUILDER_ASSETS is not set, setting it to %s", assetPath))
		Expect(os.Setenv("KUBEBUILDER_ASSETS", assetPath)).To(Succeed())
	}
	//Expect(os.Setenv("KUBEBUILDER_ATTACH_CONTROL_PLANE_OUTPUT", "true")).To(Succeed())
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
