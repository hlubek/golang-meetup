package service

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapService struct {
	clientSet kubernetes.Interface
}

func NewConfigMapService(clientSet kubernetes.Interface) *ConfigMapService {
	return &ConfigMapService{
		clientSet,
	}
}

func (c *ConfigMapService) CreateConfigMap(ctx context.Context, configMapName string) error {
	cm := v1.ConfigMap(configMapName, "default")
	o := metav1.ApplyOptions{
		FieldManager: "application/apply-patch",
	}
	_, err := c.clientSet.CoreV1().ConfigMaps("default").Apply(ctx, cm, o)
	return err
}
