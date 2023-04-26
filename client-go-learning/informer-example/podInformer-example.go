package informer

import (
	"client-go-learning/getConfig-example"
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"time"
)

func PodInformerSimple() {
	clientSet, err := getConfig.NewClientSet("")
	if err != nil {
		panic(err)
	}
	informerFactory := informers.NewSharedInformerFactory(clientSet, 5*time.Second)
	podInformer := informerFactory.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(&cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			podObject := obj.(*v1.Pod).DeepCopy()
			fmt.Printf("get a new pod: %+v\n", podObject)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			newPodObj := newObj.(*v1.Pod).DeepCopy()
			oldPodObj := oldObj.(*v1.Pod).DeepCopy()
			fmt.Printf("a pod update from %+v to %+v \n", oldPodObj, newPodObj)
		},
		DeleteFunc: nil,
	})
	ctx := context.TODO()
	podInformer.Run(ctx.Done())
	cache.WaitForCacheSync(ctx.Done(), podInformer.HasSynced)
}

var (
	podResources = schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
)

func PodInformerDynamic() {
	clientSet, err := getConfig.NewDynamicClientSet("")
	if err != nil {
		panic(err)
	}
	dynamicInformerFactory := dynamicinformer.NewDynamicSharedInformerFactory(clientSet, 5*time.Second)
	podInformer := dynamicInformerFactory.ForResource(podResources).Informer()
	podInformer.AddEventHandler(&cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			cm := obj.(*unstructured.Unstructured)
			fmt.Printf("Informer event: Pod ADDED %s/%s\n", cm.GetNamespace(), cm.GetName())
		},
		UpdateFunc: nil,
		DeleteFunc: nil,
	})
	ctx := context.TODO()
	podInformer.Run(ctx.Done())
	cache.WaitForCacheSync(ctx.Done(), podInformer.HasSynced)
}
