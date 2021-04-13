package command

import (
	"context"
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	k8sServiceKind          = "Service"
	k8sDeploymentKind       = "Deployment"
	k8sNamespaceKind        = "Namespace"
	k8sStatefulSetKind      = "StatefulSet"
	k8sConfigMapKind        = "ConfigMap"
	k8sPersistentVolumeKind = "PersistentVolume"
)

func createK8SObject(clientset *kubernetes.Clientset, object runtime.Object, groupVersionKind *schema.GroupVersionKind, namespace string) error {
	var err error
	switch groupVersionKind.Kind {
	case k8sServiceKind:
		err = createService(object.(*v1.Service), clientset, namespace)
	case k8sDeploymentKind:
		err = createDeployment(object, clientset, namespace)
	case k8sNamespaceKind:
		err = createNameSpace(object.(*v1.Namespace), clientset)
	case k8sStatefulSetKind:
		err = createStatefulSet(object, clientset, namespace)
	case k8sConfigMapKind:
		err = createConfigMap(object.(*v1.ConfigMap), clientset, namespace)
	case k8sPersistentVolumeKind:
		err = createPersistentVolume(object, clientset)
	default:
		return fmt.Errorf("Create K8S Object failed, unknonwn resource kind, %v ", groupVersionKind)

	}
	return err
}

func createNameSpace(namespace *v1.Namespace, clientSet *kubernetes.Clientset) error {
	_, err := clientSet.CoreV1().Namespaces().Get(context.TODO(), namespace.Name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		_, err := clientSet.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
		return err
	}
	return nil
}

func createDeployment(obj runtime.Object, clientSet *kubernetes.Clientset, namespaces string) error {
	deployment := obj.(*appsV1.Deployment)
	_, err := clientSet.AppsV1().Deployments(namespaces).Get(context.TODO(), deployment.Name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		_, err = clientSet.AppsV1().Deployments(namespaces).Create(context.TODO(), deployment, metav1.CreateOptions{})
	} else {
		_, err = clientSet.AppsV1().Deployments(namespaces).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	}
	return err
}

func createStatefulSet(obj runtime.Object, clientSet *kubernetes.Clientset, namespaces string) error {
	statefulSet := obj.(*appsV1.StatefulSet)
	_, err := clientSet.AppsV1().StatefulSets(namespaces).Get(context.TODO(), statefulSet.Name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		_, err = clientSet.AppsV1().StatefulSets(namespaces).Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	} else {
		_, err = clientSet.AppsV1().StatefulSets(namespaces).Update(context.TODO(), statefulSet, metav1.UpdateOptions{})
	}
	return err
}

func createService(service *v1.Service, clientSet *kubernetes.Clientset, namespaces string) error {
	_, err := clientSet.CoreV1().Services(namespaces).Get(context.TODO(), service.Name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		_, err = clientSet.CoreV1().Services(namespaces).Create(context.TODO(), service, metav1.CreateOptions{})
	}
	return err
}

func createConfigMap(configMap *v1.ConfigMap, clientSet *kubernetes.Clientset, namespaces string) error {
	_, err := clientSet.CoreV1().ConfigMaps(namespaces).Get(context.TODO(), configMap.Name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		_, err = clientSet.CoreV1().ConfigMaps(namespaces).Create(context.TODO(), configMap, metav1.CreateOptions{})
	} else {
		_, err = clientSet.CoreV1().ConfigMaps(namespaces).Update(context.TODO(), configMap, metav1.UpdateOptions{})
	}
	return err

}

func createPersistentVolume(obj runtime.Object, clientSet *kubernetes.Clientset) error {

	pv := obj.(*v1.PersistentVolume)
	_, err := clientSet.CoreV1().PersistentVolumes().Get(context.TODO(), pv.Name, metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		_, err = clientSet.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
	} else {
		_, err = clientSet.CoreV1().PersistentVolumes().Update(context.TODO(), pv, metav1.UpdateOptions{})
	}
	return err
}

func decodeToK8SObject(content []byte) (runtime.Object, *schema.GroupVersionKind, error) {
	return scheme.Codecs.UniversalDeserializer().Decode(content, nil, nil)
}
