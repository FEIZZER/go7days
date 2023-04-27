package getClient

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
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

func DiscoveryClientUsage() {
	client, err := NewDiscoveryClient()
	if err != nil {
		panic(err)
	}
	apiGroups, apiResources, err := client.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}
	for _, group := range apiGroups {
		fmt.Printf("group: %+v \n", group)
	}
	for _, resourceList := range apiResources {
		group, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\ngroup:%+v :", group)

		for _, resource := range resourceList.APIResources {
			fmt.Printf("%+v ", resource)
		}
	}
}

func DynamicClientUsage() {
	client, err := NewDynamicClient()
	if err != nil {
		panic(err)
	}
	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
	list, err := client.Resource(gvr).Namespace("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("from kubernetes get pods:%+v \n", list.Items)
}
