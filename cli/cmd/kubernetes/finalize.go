/**
 * @Author guohb65
 * @Date 9/9/2022 下午 6:28
 **/

package kubernetes

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func Finalize(namespaceName string) error {
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}
	result := &corev1.Namespace{}
	return restClientSet.Put().Resource("namespaces").Name(namespace.Name).
		VersionedParams(&metav1.UpdateOptions{}, scheme.ParameterCodec).
		SubResource("finalize").Body(namespace).
		Do(context.Background()).
		Into(result)
}
