// +build !ignore_autogenerated

/*
Copyright 2021 The Scribe authors.

This file may be used, at your option, according to either the GNU AGPL 3.0 or
the Apache V2 license.

---
This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option) any
later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along
with this program.  If not, see <https://www.gnu.org/licenses/>.

---
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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/operator-framework/operator-lib/status"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestination) DeepCopyInto(out *ReplicationDestination) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ReplicationDestinationStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestination.
func (in *ReplicationDestination) DeepCopy() *ReplicationDestination {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationDestination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationExternalSpec) DeepCopyInto(out *ReplicationDestinationExternalSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationExternalSpec.
func (in *ReplicationDestinationExternalSpec) DeepCopy() *ReplicationDestinationExternalSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationExternalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationList) DeepCopyInto(out *ReplicationDestinationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationList.
func (in *ReplicationDestinationList) DeepCopy() *ReplicationDestinationList {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationDestinationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationRcloneSpec) DeepCopyInto(out *ReplicationDestinationRcloneSpec) {
	*out = *in
	in.ReplicationDestinationVolumeOptions.DeepCopyInto(&out.ReplicationDestinationVolumeOptions)
	if in.RcloneConfigSection != nil {
		in, out := &in.RcloneConfigSection, &out.RcloneConfigSection
		*out = new(string)
		**out = **in
	}
	if in.RcloneDestPath != nil {
		in, out := &in.RcloneDestPath, &out.RcloneDestPath
		*out = new(string)
		**out = **in
	}
	if in.RcloneConfig != nil {
		in, out := &in.RcloneConfig, &out.RcloneConfig
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationRcloneSpec.
func (in *ReplicationDestinationRcloneSpec) DeepCopy() *ReplicationDestinationRcloneSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationRcloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationResticSpec) DeepCopyInto(out *ReplicationDestinationResticSpec) {
	*out = *in
	in.ReplicationDestinationVolumeOptions.DeepCopyInto(&out.ReplicationDestinationVolumeOptions)
	if in.CacheCapacity != nil {
		in, out := &in.CacheCapacity, &out.CacheCapacity
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.CacheStorageClassName != nil {
		in, out := &in.CacheStorageClassName, &out.CacheStorageClassName
		*out = new(string)
		**out = **in
	}
	if in.CacheAccessModes != nil {
		in, out := &in.CacheAccessModes, &out.CacheAccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationResticSpec.
func (in *ReplicationDestinationResticSpec) DeepCopy() *ReplicationDestinationResticSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationResticSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationRsyncSpec) DeepCopyInto(out *ReplicationDestinationRsyncSpec) {
	*out = *in
	in.ReplicationDestinationVolumeOptions.DeepCopyInto(&out.ReplicationDestinationVolumeOptions)
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = new(string)
		**out = **in
	}
	if in.ServiceType != nil {
		in, out := &in.ServiceType, &out.ServiceType
		*out = new(v1.ServiceType)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationRsyncSpec.
func (in *ReplicationDestinationRsyncSpec) DeepCopy() *ReplicationDestinationRsyncSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationRsyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationRsyncStatus) DeepCopyInto(out *ReplicationDestinationRsyncStatus) {
	*out = *in
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = new(string)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationRsyncStatus.
func (in *ReplicationDestinationRsyncStatus) DeepCopy() *ReplicationDestinationRsyncStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationRsyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationSpec) DeepCopyInto(out *ReplicationDestinationSpec) {
	*out = *in
	if in.Trigger != nil {
		in, out := &in.Trigger, &out.Trigger
		*out = new(ReplicationDestinationTriggerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Rsync != nil {
		in, out := &in.Rsync, &out.Rsync
		*out = new(ReplicationDestinationRsyncSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Rclone != nil {
		in, out := &in.Rclone, &out.Rclone
		*out = new(ReplicationDestinationRcloneSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Restic != nil {
		in, out := &in.Restic, &out.Restic
		*out = new(ReplicationDestinationResticSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ReplicationDestinationExternalSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationSpec.
func (in *ReplicationDestinationSpec) DeepCopy() *ReplicationDestinationSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationStatus) DeepCopyInto(out *ReplicationDestinationStatus) {
	*out = *in
	if in.LastSyncTime != nil {
		in, out := &in.LastSyncTime, &out.LastSyncTime
		*out = (*in).DeepCopy()
	}
	if in.LastSyncDuration != nil {
		in, out := &in.LastSyncDuration, &out.LastSyncDuration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.NextSyncTime != nil {
		in, out := &in.NextSyncTime, &out.NextSyncTime
		*out = (*in).DeepCopy()
	}
	if in.LatestImage != nil {
		in, out := &in.LatestImage, &out.LatestImage
		*out = new(v1.TypedLocalObjectReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Rsync != nil {
		in, out := &in.Rsync, &out.Rsync
		*out = new(ReplicationDestinationRsyncStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationStatus.
func (in *ReplicationDestinationStatus) DeepCopy() *ReplicationDestinationStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationTriggerSpec) DeepCopyInto(out *ReplicationDestinationTriggerSpec) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationTriggerSpec.
func (in *ReplicationDestinationTriggerSpec) DeepCopy() *ReplicationDestinationTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationDestinationVolumeOptions) DeepCopyInto(out *ReplicationDestinationVolumeOptions) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.VolumeSnapshotClassName != nil {
		in, out := &in.VolumeSnapshotClassName, &out.VolumeSnapshotClassName
		*out = new(string)
		**out = **in
	}
	if in.DestinationPVC != nil {
		in, out := &in.DestinationPVC, &out.DestinationPVC
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationDestinationVolumeOptions.
func (in *ReplicationDestinationVolumeOptions) DeepCopy() *ReplicationDestinationVolumeOptions {
	if in == nil {
		return nil
	}
	out := new(ReplicationDestinationVolumeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSource) DeepCopyInto(out *ReplicationSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ReplicationSourceStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSource.
func (in *ReplicationSource) DeepCopy() *ReplicationSource {
	if in == nil {
		return nil
	}
	out := new(ReplicationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceExternalSpec) DeepCopyInto(out *ReplicationSourceExternalSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceExternalSpec.
func (in *ReplicationSourceExternalSpec) DeepCopy() *ReplicationSourceExternalSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceExternalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceList) DeepCopyInto(out *ReplicationSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceList.
func (in *ReplicationSourceList) DeepCopy() *ReplicationSourceList {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceRcloneSpec) DeepCopyInto(out *ReplicationSourceRcloneSpec) {
	*out = *in
	in.ReplicationSourceVolumeOptions.DeepCopyInto(&out.ReplicationSourceVolumeOptions)
	if in.RcloneConfigSection != nil {
		in, out := &in.RcloneConfigSection, &out.RcloneConfigSection
		*out = new(string)
		**out = **in
	}
	if in.RcloneDestPath != nil {
		in, out := &in.RcloneDestPath, &out.RcloneDestPath
		*out = new(string)
		**out = **in
	}
	if in.RcloneConfig != nil {
		in, out := &in.RcloneConfig, &out.RcloneConfig
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceRcloneSpec.
func (in *ReplicationSourceRcloneSpec) DeepCopy() *ReplicationSourceRcloneSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceRcloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceResticSpec) DeepCopyInto(out *ReplicationSourceResticSpec) {
	*out = *in
	in.ReplicationSourceVolumeOptions.DeepCopyInto(&out.ReplicationSourceVolumeOptions)
	if in.PruneIntervalDays != nil {
		in, out := &in.PruneIntervalDays, &out.PruneIntervalDays
		*out = new(int32)
		**out = **in
	}
	if in.Retain != nil {
		in, out := &in.Retain, &out.Retain
		*out = new(ResticRetainPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.CacheCapacity != nil {
		in, out := &in.CacheCapacity, &out.CacheCapacity
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.CacheStorageClassName != nil {
		in, out := &in.CacheStorageClassName, &out.CacheStorageClassName
		*out = new(string)
		**out = **in
	}
	if in.CacheAccessModes != nil {
		in, out := &in.CacheAccessModes, &out.CacheAccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceResticSpec.
func (in *ReplicationSourceResticSpec) DeepCopy() *ReplicationSourceResticSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceResticSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceResticStatus) DeepCopyInto(out *ReplicationSourceResticStatus) {
	*out = *in
	if in.LastPruned != nil {
		in, out := &in.LastPruned, &out.LastPruned
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceResticStatus.
func (in *ReplicationSourceResticStatus) DeepCopy() *ReplicationSourceResticStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceResticStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceRsyncSpec) DeepCopyInto(out *ReplicationSourceRsyncSpec) {
	*out = *in
	in.ReplicationSourceVolumeOptions.DeepCopyInto(&out.ReplicationSourceVolumeOptions)
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = new(string)
		**out = **in
	}
	if in.ServiceType != nil {
		in, out := &in.ServiceType, &out.ServiceType
		*out = new(v1.ServiceType)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceRsyncSpec.
func (in *ReplicationSourceRsyncSpec) DeepCopy() *ReplicationSourceRsyncSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceRsyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceRsyncStatus) DeepCopyInto(out *ReplicationSourceRsyncStatus) {
	*out = *in
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = new(string)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceRsyncStatus.
func (in *ReplicationSourceRsyncStatus) DeepCopy() *ReplicationSourceRsyncStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceRsyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceSpec) DeepCopyInto(out *ReplicationSourceSpec) {
	*out = *in
	if in.Trigger != nil {
		in, out := &in.Trigger, &out.Trigger
		*out = new(ReplicationSourceTriggerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Rsync != nil {
		in, out := &in.Rsync, &out.Rsync
		*out = new(ReplicationSourceRsyncSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Rclone != nil {
		in, out := &in.Rclone, &out.Rclone
		*out = new(ReplicationSourceRcloneSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Restic != nil {
		in, out := &in.Restic, &out.Restic
		*out = new(ReplicationSourceResticSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ReplicationSourceExternalSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceSpec.
func (in *ReplicationSourceSpec) DeepCopy() *ReplicationSourceSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceStatus) DeepCopyInto(out *ReplicationSourceStatus) {
	*out = *in
	if in.LastSyncTime != nil {
		in, out := &in.LastSyncTime, &out.LastSyncTime
		*out = (*in).DeepCopy()
	}
	if in.LastSyncDuration != nil {
		in, out := &in.LastSyncDuration, &out.LastSyncDuration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.NextSyncTime != nil {
		in, out := &in.NextSyncTime, &out.NextSyncTime
		*out = (*in).DeepCopy()
	}
	if in.Rsync != nil {
		in, out := &in.Rsync, &out.Rsync
		*out = new(ReplicationSourceRsyncStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Restic != nil {
		in, out := &in.Restic, &out.Restic
		*out = new(ReplicationSourceResticStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceStatus.
func (in *ReplicationSourceStatus) DeepCopy() *ReplicationSourceStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceTriggerSpec) DeepCopyInto(out *ReplicationSourceTriggerSpec) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceTriggerSpec.
func (in *ReplicationSourceTriggerSpec) DeepCopy() *ReplicationSourceTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSourceVolumeOptions) DeepCopyInto(out *ReplicationSourceVolumeOptions) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.VolumeSnapshotClassName != nil {
		in, out := &in.VolumeSnapshotClassName, &out.VolumeSnapshotClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSourceVolumeOptions.
func (in *ReplicationSourceVolumeOptions) DeepCopy() *ReplicationSourceVolumeOptions {
	if in == nil {
		return nil
	}
	out := new(ReplicationSourceVolumeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResticRetainPolicy) DeepCopyInto(out *ResticRetainPolicy) {
	*out = *in
	if in.Hourly != nil {
		in, out := &in.Hourly, &out.Hourly
		*out = new(int32)
		**out = **in
	}
	if in.Daily != nil {
		in, out := &in.Daily, &out.Daily
		*out = new(int32)
		**out = **in
	}
	if in.Weekly != nil {
		in, out := &in.Weekly, &out.Weekly
		*out = new(int32)
		**out = **in
	}
	if in.Monthly != nil {
		in, out := &in.Monthly, &out.Monthly
		*out = new(int32)
		**out = **in
	}
	if in.Yearly != nil {
		in, out := &in.Yearly, &out.Yearly
		*out = new(int32)
		**out = **in
	}
	if in.Within != nil {
		in, out := &in.Within, &out.Within
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResticRetainPolicy.
func (in *ResticRetainPolicy) DeepCopy() *ResticRetainPolicy {
	if in == nil {
		return nil
	}
	out := new(ResticRetainPolicy)
	in.DeepCopyInto(out)
	return out
}
