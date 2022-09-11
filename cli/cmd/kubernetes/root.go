//: Copyright Herb Guo
//: Licensed under the terms of the Apache 2.0 License. See LICENSE file in the project root for terms.

package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
)

var clientSet *kubernetes.Clientset
var restClientSet *rest.RESTClient

func Connect(clientGetter genericclioptions.RESTClientGetter) error {
	restConfig, err := clientGetter.ToRESTConfig()
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	restConfig.APIPath = "api"
	restConfig.ContentConfig = rest.ContentConfig{
		GroupVersion:         &corev1.SchemeGroupVersion,
		NegotiatedSerializer: scheme.Codecs,
	}
	restClientset, err := rest.RESTClientFor(restConfig)
	if err != nil {
		return err
	}

	clientSet = clientset
	restClientSet = restClientset

	return nil

	// homeDirectory, err := homedir.Dir()
	// if err != nil {
	// 	return err
	// }
	// kubeconfig := filepath.Join(homeDirectory, ".kube", "config")
	// config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	// if err != nil {
	// 	return err
	// }

	// // create the clientset
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	return err
	// }

	// clientSet = clientset
	// return nil
}
