package k8s

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func CreateCoreClientFromKubeconfig(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*clientcmdapi.Config, error) {
		return clientcmd.Load([]byte(kubeconfig))
	})
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func CreateDynamicClientFromKubuconfig(kubeconfig string) (dynamic.Interface, error) {
	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*clientcmdapi.Config, error) {
		return clientcmd.Load([]byte(kubeconfig))
	})
	if err != nil {
		return nil, err
	}

	return dynamic.NewForConfig(config)
}
