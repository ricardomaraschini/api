// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	time "time"

	operatorv1 "github.com/openshift/api/operator/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Config) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigList) DeepCopyInto(out *ConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Config, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigList.
func (in *ConfigList) DeepCopy() *ConfigList {
	if in == nil {
		return nil
	}
	out := new(ConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePruner) DeepCopyInto(out *ImagePruner) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePruner.
func (in *ImagePruner) DeepCopy() *ImagePruner {
	if in == nil {
		return nil
	}
	out := new(ImagePruner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePruner) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrunerList) DeepCopyInto(out *ImagePrunerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImagePruner, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrunerList.
func (in *ImagePrunerList) DeepCopy() *ImagePrunerList {
	if in == nil {
		return nil
	}
	out := new(ImagePrunerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePrunerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrunerSpec) DeepCopyInto(out *ImagePrunerSpec) {
	*out = *in
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	if in.KeepTagRevisions != nil {
		in, out := &in.KeepTagRevisions, &out.KeepTagRevisions
		*out = new(int)
		**out = **in
	}
	if in.KeepYoungerThan != nil {
		in, out := &in.KeepYoungerThan, &out.KeepYoungerThan
		*out = new(time.Duration)
		**out = **in
	}
	if in.KeepYoungerThanDuration != nil {
		in, out := &in.KeepYoungerThanDuration, &out.KeepYoungerThanDuration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrunerSpec.
func (in *ImagePrunerSpec) DeepCopy() *ImagePrunerSpec {
	if in == nil {
		return nil
	}
	out := new(ImagePrunerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrunerStatus) DeepCopyInto(out *ImagePrunerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]operatorv1.OperatorCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrunerStatus.
func (in *ImagePrunerStatus) DeepCopy() *ImagePrunerStatus {
	if in == nil {
		return nil
	}
	out := new(ImagePrunerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigProxy) DeepCopyInto(out *ImageRegistryConfigProxy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigProxy.
func (in *ImageRegistryConfigProxy) DeepCopy() *ImageRegistryConfigProxy {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigRequests) DeepCopyInto(out *ImageRegistryConfigRequests) {
	*out = *in
	out.Read = in.Read
	out.Write = in.Write
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigRequests.
func (in *ImageRegistryConfigRequests) DeepCopy() *ImageRegistryConfigRequests {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigRequestsLimits) DeepCopyInto(out *ImageRegistryConfigRequestsLimits) {
	*out = *in
	out.MaxWaitInQueue = in.MaxWaitInQueue
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigRequestsLimits.
func (in *ImageRegistryConfigRequestsLimits) DeepCopy() *ImageRegistryConfigRequestsLimits {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigRequestsLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigRoute) DeepCopyInto(out *ImageRegistryConfigRoute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigRoute.
func (in *ImageRegistryConfigRoute) DeepCopy() *ImageRegistryConfigRoute {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigStorage) DeepCopyInto(out *ImageRegistryConfigStorage) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(ImageRegistryConfigStorageEmptyDir)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(ImageRegistryConfigStorageS3)
		(*in).DeepCopyInto(*out)
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		*out = new(ImageRegistryConfigStorageGCS)
		**out = **in
	}
	if in.Swift != nil {
		in, out := &in.Swift, &out.Swift
		*out = new(ImageRegistryConfigStorageSwift)
		**out = **in
	}
	if in.PVC != nil {
		in, out := &in.PVC, &out.PVC
		*out = new(ImageRegistryConfigStoragePVC)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(ImageRegistryConfigStorageAzure)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigStorage.
func (in *ImageRegistryConfigStorage) DeepCopy() *ImageRegistryConfigStorage {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigStorageAzure) DeepCopyInto(out *ImageRegistryConfigStorageAzure) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigStorageAzure.
func (in *ImageRegistryConfigStorageAzure) DeepCopy() *ImageRegistryConfigStorageAzure {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigStorageAzure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigStorageEmptyDir) DeepCopyInto(out *ImageRegistryConfigStorageEmptyDir) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigStorageEmptyDir.
func (in *ImageRegistryConfigStorageEmptyDir) DeepCopy() *ImageRegistryConfigStorageEmptyDir {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigStorageEmptyDir)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigStorageGCS) DeepCopyInto(out *ImageRegistryConfigStorageGCS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigStorageGCS.
func (in *ImageRegistryConfigStorageGCS) DeepCopy() *ImageRegistryConfigStorageGCS {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigStorageGCS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigStoragePVC) DeepCopyInto(out *ImageRegistryConfigStoragePVC) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigStoragePVC.
func (in *ImageRegistryConfigStoragePVC) DeepCopy() *ImageRegistryConfigStoragePVC {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigStoragePVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigStorageS3) DeepCopyInto(out *ImageRegistryConfigStorageS3) {
	*out = *in
	if in.CloudFront != nil {
		in, out := &in.CloudFront, &out.CloudFront
		*out = new(ImageRegistryConfigStorageS3CloudFront)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigStorageS3.
func (in *ImageRegistryConfigStorageS3) DeepCopy() *ImageRegistryConfigStorageS3 {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigStorageS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigStorageS3CloudFront) DeepCopyInto(out *ImageRegistryConfigStorageS3CloudFront) {
	*out = *in
	in.PrivateKey.DeepCopyInto(&out.PrivateKey)
	out.Duration = in.Duration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigStorageS3CloudFront.
func (in *ImageRegistryConfigStorageS3CloudFront) DeepCopy() *ImageRegistryConfigStorageS3CloudFront {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigStorageS3CloudFront)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryConfigStorageSwift) DeepCopyInto(out *ImageRegistryConfigStorageSwift) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryConfigStorageSwift.
func (in *ImageRegistryConfigStorageSwift) DeepCopy() *ImageRegistryConfigStorageSwift {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryConfigStorageSwift)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistrySpec) DeepCopyInto(out *ImageRegistrySpec) {
	*out = *in
	in.OperatorSpec.DeepCopyInto(&out.OperatorSpec)
	out.Proxy = in.Proxy
	in.Storage.DeepCopyInto(&out.Storage)
	out.Requests = in.Requests
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]ImageRegistryConfigRoute, len(*in))
		copy(*out, *in)
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(int64)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistrySpec.
func (in *ImageRegistrySpec) DeepCopy() *ImageRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(ImageRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistryStatus) DeepCopyInto(out *ImageRegistryStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	in.Storage.DeepCopyInto(&out.Storage)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistryStatus.
func (in *ImageRegistryStatus) DeepCopy() *ImageRegistryStatus {
	if in == nil {
		return nil
	}
	out := new(ImageRegistryStatus)
	in.DeepCopyInto(out)
	return out
}
