package getClient

import (
	"client-go-learning/getConfig-example"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

func NewRestClient() (*rest.RESTClient, error) {
	config, err := getConfig.GetKubernetesConfig("")
	if err != nil {
		return nil, err
	}
	// 配置 API 路径
	config.APIPath = "api"
	// 配置 资源版本
	config.GroupVersion = &corev1.SchemeGroupVersion
	// 配置 解码器
	config.NegotiatedSerializer = scheme.Codecs
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		return nil, err
	}
	return restClient, nil
}

func NewClientSet() (*kubernetes.Clientset, error) {
	config, err := getConfig.GetKubernetesConfig("")
	if err != nil {
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}
