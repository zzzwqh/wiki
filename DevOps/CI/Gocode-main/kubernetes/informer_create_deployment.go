package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// kubeconfig 是个字符串类型，就是文件路径
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 使用 clientcmd.BuildConfigFromFlags() 函数创建 *rest.Config 对象，该函数需要传入 kubeconfig 文件路径
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error building Kubernetes configuration: %s", err.Error())
	}

	// 使用 *rest.Config 对象创建 kubernetes.Clientset 对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %s", err.Error())
	}

	// 定义一个 Deployment
	replicas := int32(3)
	image := "nginx:latest"
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "nginx",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// create an Informer factory for Deployments
	factory := cache.NewSharedIndexInformer(
		cache.NewListWatchFromClient(clientset.AppsV1().RESTClient(), "deployments", metav1.NamespaceDefault, nil),
		&appsv1.Deployment{},
		time.Second*30,
		cache.Indexers{},
	)

	// start the Informer factory
	stopCh := make(chan struct{})
	defer close(stopCh)
	go factory.Run(stopCh)

	// wait for the cache to sync
	if !cache.WaitForCacheSync(stopCh, factory.HasSynced) {
		runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}

	// create the deployment
	result, err := clientset.AppsV1().Deployments("default").Create(context.Background(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating deployment: %s", err.Error())
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

// homeDir returns the user's home directory.
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}
