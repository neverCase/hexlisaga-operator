package helixsaga

import (
	"fmt"
	"time"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"

	helixSagaV1 "github.com/Shanghai-Lunara/helixsaga-operator/pkg/apis/helixsaga/v1"
	helixSagaClientSet "github.com/Shanghai-Lunara/helixsaga-operator/pkg/generated/helixsaga/clientset/versioned"
	k8sCoreV1 "github.com/nevercase/k8s-controller-custom-resource/core/v1"
)

func NewStatefulSetAndService(ks k8sCoreV1.KubernetesResource, client helixSagaClientSet.Interface, hs *helixSagaV1.HelixSaga, spec *helixSagaV1.HelixSagaAppSpec, wo *WatchOption) error {
	//if *spec.Replicas == 0 {
	//	if err := ks.StatefulSet().Delete(hs.Namespace, spec.Name); err != nil {
	//		klog.V(2).Info(err)
	//	}
	//	if err := ks.Service().Delete(hs.Namespace, k8sCoreV1.GetStatefulSetName(spec.Name)); err != nil {
	//		klog.V(2).Info(err)
	//	}
	//	return nil
	//}
	var err error
	wo.StatefulSet, err = ks.StatefulSet().Get(hs.Namespace, spec.Name)
	if err != nil {
		klog.Info("statefulSet err:", err)
		if !errors.IsNotFound(err) {
			return err
		}
		klog.Info("new statefulSet")
		if wo.StatefulSet, err = ks.StatefulSet().Create(hs.Namespace, NewStatefulSet(hs, spec)); err != nil {
			return err
		}
	}
	klog.Info("rds:", *spec.Replicas)
	klog.Info("statefulSet:", *wo.StatefulSet.Spec.Replicas)
	if ok := compareStatefulSet(wo.StatefulSet, spec); ok {
		if wo.StatefulSet, err = ks.StatefulSet().Update(hs.Namespace, NewStatefulSet(hs, spec)); err != nil {
			klog.V(2).Info(err)
			return err
		}
	}
	if len(spec.ServicePorts) == 0 {
		if err = ks.Service().Delete(hs.Namespace, k8sCoreV1.GetServiceName(spec.Name)); err != nil {
			return err
		}
	} else {
		svc, err := ks.Service().Get(hs.Namespace, k8sCoreV1.GetServiceName(spec.Name))
		if err != nil {
			klog.Info("service err:", err)
			if !errors.IsNotFound(err) {
				return err
			}
			klog.Info("new service check")
			if _, err = ks.Service().Create(hs.Namespace, NewService(hs, spec)); err != nil {
				return err
			}
		} else {
			tmpSvc := NewService(hs, spec)
			if ok := compareService(svc, tmpSvc); ok {
				svc.Labels = tmpSvc.Labels
				svc.Spec.Type = tmpSvc.Spec.Type
				svc.Spec.Ports = tmpSvc.Spec.Ports
				svc.Spec.Selector = tmpSvc.Spec.Selector
				if _, err = ks.Service().Update(hs.Namespace, svc); err != nil {
					return err
				}
			}
		}
	}
	if err = updateStatus(hs, client, wo.StatefulSet, spec.Name); err != nil {
		return err
	}
	return nil
}

func compareStatefulSet(original *appsV1.StatefulSet, updateSpec *helixSagaV1.HelixSagaAppSpec) bool {
	if updateSpec.Replicas != nil && *updateSpec.Replicas != *original.Spec.Replicas {
		return true
	}
	if updateSpec.Image != original.Spec.Template.Spec.Containers[0].Image {
		return true
	}
	// compare Affinity
	// compare Tolerations
	return false
}

func compareService(s1 *coreV1.Service, s2 *coreV1.Service) bool {
	if s1.Spec.Type != s2.Spec.Type {
		return true
	}
	if len(s1.Spec.Ports) != len(s2.Spec.Ports) {
		return true
	}
	for _, v := range s1.Spec.Ports {
		exist := false
		for _, v2 := range s2.Spec.Ports {
			if v.Port == v2.Port {
				exist = true
				if v.Name != v2.Name {
					return true
				}
				if v.NodePort != v2.NodePort && v2.NodePort != 0 {
					return true
				}
				if v.Protocol != v2.Protocol {
					return true
				}
				if v.TargetPort.Type != v2.TargetPort.Type {
					return true
				}
				if v.TargetPort.IntVal != v2.TargetPort.IntVal {
					return true
				}
				if v.TargetPort.StrVal != v2.TargetPort.StrVal {
					return true
				}
				break
			}
		}
		if !exist {
			return true
		}
	}
	return false
}

func updateStatus(foo *helixSagaV1.HelixSaga, clientSet helixSagaClientSet.Interface, ss *appsV1.StatefulSet, name string) error {
	// NEVER modify objects from the store. It's a read-only, local cache.
	// You can use DeepCopy() to make a deep copy of original object and modify this copy
	// Or create a copy manually for better performance
	fooCopy := foo.DeepCopy()
	t := make([]helixSagaV1.HelixSagaApp, 0)
	for _, v := range fooCopy.Spec.Applications {
		if v.Spec.Name == name {
			v.Status.ObservedGeneration = ss.Status.ObservedGeneration
			v.Status.Replicas = ss.Status.Replicas
			v.Status.ReadyReplicas = ss.Status.ReadyReplicas
			v.Status.CurrentReplicas = ss.Status.CurrentReplicas
			v.Status.UpdatedReplicas = ss.Status.UpdatedReplicas
			v.Status.CurrentRevision = ss.Status.CurrentRevision
			v.Status.UpdateRevision = ss.Status.UpdateRevision
			v.Status.CollisionCount = ss.Status.CollisionCount
		}
		t = append(t, v)
	}
	fooCopy.Spec.Applications = t

	// If the CustomResourceSubResources feature gate is not enabled,
	// we must use Update instead of UpdateStatus to update the Status block of the RedisOperator resource.
	// UpdateStatus will not allow changes to the Spec of the resource,
	// which is ideal for ensuring nothing other than resource status has been updated.
	_, err := clientSet.NevercaseV1().HelixSagas(foo.Namespace).Update(fooCopy)
	return err
}

func DeleteStatefulSetAndService(ks k8sCoreV1.KubernetesResource, namespace string, name string) error {
	if err := ks.StatefulSet().Delete(namespace, name); err != nil {
		klog.V(2).Info(err)
		return err
	}
	if err := ks.Service().Delete(namespace, k8sCoreV1.GetServiceName(name)); err != nil {
		klog.V(2).Info(err)
		return err
	}
	return nil
}

const (
	ErrorPodsHadNotBeenClosed = "namespace:%s crdName:%s image:%s error: pods hadn't been closed completed"
)

func RetryPatchHelixSaga(ki kubernetes.Interface, clientSet helixSagaClientSet.Interface, namespace, crdName, image string, replicas map[string]int32) (map[string]int32, error) {
	var res = make(map[string]int32, 0)
	var defaultConfig = wait.Backoff{
		Steps:    10000,
		Duration: 5 * time.Millisecond,
		Factor:   1.0,
		Jitter:   0.1,
	}
	err := retry.RetryOnConflict(defaultConfig, func() error {
		hs, err := clientSet.NevercaseV1().HelixSagas(namespace).Get(crdName, metav1.GetOptions{})
		if err != nil {
			klog.V(2).Info(err)
			return errors.NewConflict(schema.GroupResource{Resource: "test"}, "RetryPatchHelixSaga", err)
		}
		if len(replicas) > 0 {
			if pl, err := ListPodByLabels(ki, namespace, crdName, ""); err != nil {
				klog.V(2).Info(err)
				return errors.NewConflict(schema.GroupResource{Resource: "test"}, "RetryPatchHelixSaga", err)
			} else {
				klog.Infof("namespace:%s crdName:%s image:%s pods-numbers:%d", namespace, crdName, image, len(pl.Items))
				if len(pl.Items) > 0 {
					policyMap := make(map[string]helixSagaV1.WatchPolicy, 0)
					for _, v := range hs.Spec.Applications {
						policyMap[v.Spec.Name] = v.Spec.WatchPolicy
					}
					for _, v := range pl.Items {
						if len(v.Spec.Containers) > 0 {
							if v.Spec.Containers[0].Image == image {
								if policy, ok := policyMap[v.Spec.Containers[0].Name]; ok {
									if policy == helixSagaV1.WatchPolicyAuto {
										klog.Infof("check namespace:%s crdName:%s image:%s container-name:%d", namespace, crdName, image, v.Spec.Containers[0].Name)
										err = fmt.Errorf(ErrorPodsHadNotBeenClosed, namespace, crdName, image)
										klog.V(2).Info(err)
										return errors.NewConflict(schema.GroupResource{Resource: "test"}, "RetryPatchHelixSaga", err)
									}
								}
							}
						}
					}
				}
			}
		}
		exist := false
		apps := make([]helixSagaV1.HelixSagaApp, 0)
		for _, v := range hs.Spec.Applications {
			if v.Spec.Image == image {
				var a int32
				if v.Spec.WatchPolicy == helixSagaV1.WatchPolicyAuto {
					if t, ok := replicas[v.Spec.Name]; ok {
						a = t
					} else {
						res[v.Spec.Name] = *v.Spec.Replicas
					}
					v.Spec.Replicas = &a
				}
				exist = true
				klog.Infof("Patch change crd-name:%s image:%s specName:%s replicas:%d", crdName, image, v.Spec.Name, *v.Spec.Replicas)
			}
			apps = append(apps, v)
		}
		hs.Spec.Applications = apps
		if !exist {
			err = fmt.Errorf("error: the crd-name:%s image:%s was not found", crdName, image)
			klog.V(2).Info(err)
			defaultConfig.Steps = 0
			return err
		}
		if _, err = clientSet.NevercaseV1().HelixSagas(namespace).Update(hs); err != nil {
			klog.V(2).Info(err)
			return errors.NewConflict(schema.GroupResource{Resource: "test"}, "RetryPatchHelixSaga", err)
		}
		return nil
	})
	if errors.IsConflict(err) {
		err = fmt.Errorf("RetryUpdateHelixSaga UpdateMaxRetries(%d) has reached. The RetryUpdateHelixSaga will retry later for owner namespace:%s crdName:%s image:%s",
			defaultConfig.Steps, namespace, crdName, image)
		klog.V(2).Info(err)
	}
	return res, err
}
