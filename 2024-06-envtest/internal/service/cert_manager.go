package service

import (
	"context"
	"encoding/json"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"

	cmmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	rest "k8s.io/client-go/rest"
)

type CertManagerService struct {
	clientSet rest.Interface
}

func NewCertManagerService(clientSet kubernetes.Interface) *CertManagerService {
	return &CertManagerService{
		clientSet: clientSet.CoreV1().RESTClient(),
	}
}

func (c *CertManagerService) CreateCert(ctx context.Context, certName string, domain string) error {

	certificate := certmanagerv1.Certificate{
		TypeMeta: v1.TypeMeta{
			APIVersion: "cert-manager.io/v1",
			Kind:       "Certificate",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: certName,
		},
		Spec: certmanagerv1.CertificateSpec{
			DNSNames:   []string{domain},
			SecretName: certName + "-tls",
			IssuerRef: cmmeta.ObjectReference{
				Name: "letsencrypt-staging",
				Kind: "ClusterIssuer",
			},
		},
	}

	body, err := json.Marshal(certificate)
	if err != nil {
		return err
	}
	result := c.clientSet.
		Post().
		AbsPath("/apis/cert-manager.io/v1").
		Namespace("default").
		Resource("certificates").
		Body(body).
		Do(ctx)

	return result.Error()
}
