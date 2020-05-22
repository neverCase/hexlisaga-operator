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
func (in *AppNotificationSpec) DeepCopyInto(out *AppNotificationSpec) {
	*out = *in
	in.DispatchSpec.DeepCopyInto(&out.DispatchSpec)
	in.LogicSpec.DeepCopyInto(&out.LogicSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppNotificationSpec.
func (in *AppNotificationSpec) DeepCopy() *AppNotificationSpec {
	if in == nil {
		return nil
	}
	out := new(AppNotificationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppNotificationStatus) DeepCopyInto(out *AppNotificationStatus) {
	*out = *in
	out.DispatchStatus = in.DispatchStatus
	out.LogicStatus = in.LogicStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppNotificationStatus.
func (in *AppNotificationStatus) DeepCopy() *AppNotificationStatus {
	if in == nil {
		return nil
	}
	out := new(AppNotificationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CampaignSpec) DeepCopyInto(out *CampaignSpec) {
	*out = *in
	in.GatewaySpec.DeepCopyInto(&out.GatewaySpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CampaignSpec.
func (in *CampaignSpec) DeepCopy() *CampaignSpec {
	if in == nil {
		return nil
	}
	out := new(CampaignSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CampaignStatus) DeepCopyInto(out *CampaignStatus) {
	*out = *in
	out.GatewayStatus = in.GatewayStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CampaignStatus.
func (in *CampaignStatus) DeepCopy() *CampaignStatus {
	if in == nil {
		return nil
	}
	out := new(CampaignStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuildWarSpec) DeepCopyInto(out *GuildWarSpec) {
	*out = *in
	in.RegisterSpec.DeepCopyInto(&out.RegisterSpec)
	in.GatewaySpec.DeepCopyInto(&out.GatewaySpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuildWarSpec.
func (in *GuildWarSpec) DeepCopy() *GuildWarSpec {
	if in == nil {
		return nil
	}
	out := new(GuildWarSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuildWarStatus) DeepCopyInto(out *GuildWarStatus) {
	*out = *in
	out.RegisterStatus = in.RegisterStatus
	out.GatewayStatus = in.GatewayStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuildWarStatus.
func (in *GuildWarStatus) DeepCopy() *GuildWarStatus {
	if in == nil {
		return nil
	}
	out := new(GuildWarStatus)
	in.DeepCopyInto(out)
	return out
}

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
func (in *HelixSagaCore) DeepCopyInto(out *HelixSagaCore) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaCore.
func (in *HelixSagaCore) DeepCopy() *HelixSagaCore {
	if in == nil {
		return nil
	}
	out := new(HelixSagaCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaCoreSpec) DeepCopyInto(out *HelixSagaCoreSpec) {
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaCoreSpec.
func (in *HelixSagaCoreSpec) DeepCopy() *HelixSagaCoreSpec {
	if in == nil {
		return nil
	}
	out := new(HelixSagaCoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaCoreStatus) DeepCopyInto(out *HelixSagaCoreStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaCoreStatus.
func (in *HelixSagaCoreStatus) DeepCopy() *HelixSagaCoreStatus {
	if in == nil {
		return nil
	}
	out := new(HelixSagaCoreStatus)
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
	if in.NginxPhpFpm != nil {
		in, out := &in.NginxPhpFpm, &out.NginxPhpFpm
		*out = make([]HelixSagaCore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PhpSwoole != nil {
		in, out := &in.PhpSwoole, &out.PhpSwoole
		*out = make([]HelixSagaCore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PhpWorkerman != nil {
		in, out := &in.PhpWorkerman, &out.PhpWorkerman
		*out = make([]PhpWorkermanSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CampaignSpec.DeepCopyInto(&out.CampaignSpec)
	in.GuildWarSpec.DeepCopyInto(&out.GuildWarSpec)
	in.AppNotificationSpec.DeepCopyInto(&out.AppNotificationSpec)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelixSagaStatus) DeepCopyInto(out *HelixSagaStatus) {
	*out = *in
	out.VersionStatus = in.VersionStatus
	out.ApiStatus = in.ApiStatus
	out.GameStatus = in.GameStatus
	out.PayNotifyStatus = in.PayNotifyStatus
	out.GmtStatus = in.GmtStatus
	out.FriendStatus = in.FriendStatus
	out.QueueStatus = in.QueueStatus
	out.RankStatus = in.RankStatus
	out.ChatStatus = in.ChatStatus
	out.HeartStatus = in.HeartStatus
	out.CampaignStatus = in.CampaignStatus
	out.GuildWarStatus = in.GuildWarStatus
	out.AppNotificationStatus = in.AppNotificationStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelixSagaStatus.
func (in *HelixSagaStatus) DeepCopy() *HelixSagaStatus {
	if in == nil {
		return nil
	}
	out := new(HelixSagaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhpWorkermanSpec) DeepCopyInto(out *PhpWorkermanSpec) {
	*out = *in
	in.RegisterSpec.DeepCopyInto(&out.RegisterSpec)
	in.GatewaySpec.DeepCopyInto(&out.GatewaySpec)
	in.BusinessWorkerSpec.DeepCopyInto(&out.BusinessWorkerSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhpWorkermanSpec.
func (in *PhpWorkermanSpec) DeepCopy() *PhpWorkermanSpec {
	if in == nil {
		return nil
	}
	out := new(PhpWorkermanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhpWorkermanStatus) DeepCopyInto(out *PhpWorkermanStatus) {
	*out = *in
	out.RegisterStatus = in.RegisterStatus
	out.GatewayStatus = in.GatewayStatus
	out.BusinessWorkerStatus = in.BusinessWorkerStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhpWorkermanStatus.
func (in *PhpWorkermanStatus) DeepCopy() *PhpWorkermanStatus {
	if in == nil {
		return nil
	}
	out := new(PhpWorkermanStatus)
	in.DeepCopyInto(out)
	return out
}
