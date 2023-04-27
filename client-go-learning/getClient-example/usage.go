package getClient

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func RESTClientUsage() {
	client, err := NewRestClient()
	if err != nil {
		panic(err)
	}
	podList := &corev1.PodList{}
	err = client.Get().
		Namespace("default").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec).
		Do(context.Background()).
		Into(podList)
	if err != nil {
		panic(err)
	}
	fmt.Printf("have %+v pod in default\n", len(podList.Items))
	fmt.Printf("from k8s get pods: %+v \n", podList.Items)
}
