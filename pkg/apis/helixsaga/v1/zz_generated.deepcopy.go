// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSaga) DeepCopyInto(out *HelixSaga) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSaga.
func (in *HelixSaga) DeepCopy() *HelixSaga {
	if in == nil {
		return nil
	}
	out := new(HelixSaga)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelixSaga) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaApp) DeepCopyInto(out *HelixSagaApp) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaApp.
func (in *HelixSagaApp) DeepCopy() *HelixSagaApp {
	if in == nil {
		return nil
	}
	out := new(HelixSagaApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaAppSpec) DeepCopyInto(out *HelixSagaAppSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContainerPorts != nil {
		in, out := &in.ContainerPorts, &out.ContainerPorts
		*out = make([]corev1.ContainerPort, len(*in))
		copy(*out, *in)
	}
	if in.ServicePorts != nil {
		in, out := &in.ServicePorts, &out.ServicePorts
		*out = make([]corev1.ServicePort, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaAppSpec.
func (in *HelixSagaAppSpec) DeepCopy() *HelixSagaAppSpec {
	if in == nil {
		return nil
	}
	out := new(HelixSagaAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaAppStatus) DeepCopyInto(out *HelixSagaAppStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaAppStatus.
func (in *HelixSagaAppStatus) DeepCopy() *HelixSagaAppStatus {
	if in == nil {
		return nil
	}
	out := new(HelixSagaAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaConfigMap) DeepCopyInto(out *HelixSagaConfigMap) {
	*out = *in
	in.Volume.DeepCopyInto(&out.Volume)
	in.VolumeMount.DeepCopyInto(&out.VolumeMount)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaConfigMap.
func (in *HelixSagaConfigMap) DeepCopy() *HelixSagaConfigMap {
	if in == nil {
		return nil
	}
	out := new(HelixSagaConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaList) DeepCopyInto(out *HelixSagaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelixSaga, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaList.
func (in *HelixSagaList) DeepCopy() *HelixSagaList {
	if in == nil {
		return nil
	}
	out := new(HelixSagaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelixSagaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaSpec) DeepCopyInto(out *HelixSagaSpec) {
	*out = *in
	in.ConfigMap.DeepCopyInto(&out.ConfigMap)
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = make([]HelixSagaApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaSpec.
func (in *HelixSagaSpec) DeepCopy() *HelixSagaSpec {
	if in == nil {
		return nil
	}
	out := new(HelixSagaSpec)
	in.DeepCopyInto(out)
	return out
}
