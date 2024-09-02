package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Println("Podname", pod.Name)
		fmt.Println("nodename:", pod.Spec.NodeName)
		//fmt.Println("conditions", pod.Status.Conditions)
		fmt.Println("Podlabels:")
		for key := range pod.Labels {
			fmt.Println(key, "=", pod.Labels[key])
		}
		fmt.Println("status:", pod.Status.Conditions[0].Status)
		fmt.Println("Ready:", pod.Status.ContainerStatuses[0].Ready)
		fmt.Println("Started:", pod.Status.ContainerStatuses[0].Started)
		fmt.Println("重启次数:", pod.Status.ContainerStatuses[0].RestartCount)
	}
	pv, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pv := range pv.Items {
		fmt.Println("PV Name: ", pv.Name)
		fmt.Println("PV Storage: ", pv.Spec.StorageClassName)
		fmt.Println("PV Capacity:", pv.Spec.Capacity)
	}
}
