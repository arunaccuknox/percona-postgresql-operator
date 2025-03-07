//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 Copyright 2021 - 2023 Crunchy Data Solutions, Inc.
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

package v2beta1

import (
	"github.com/percona/percona-postgresql-operator/pkg/apis/postgres-operator.crunchydata.com/v1beta1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PGBouncerSpec) DeepCopyInto(out *PGBouncerSpec) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(v1beta1.Metadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Config.DeepCopyInto(&out.Config)
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomTLSSecret != nil {
		in, out := &in.CustomTLSSecret, &out.CustomTLSSecret
		*out = new(v1.SecretProjection)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.PriorityClassName != nil {
		in, out := &in.PriorityClassName, &out.PriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ServiceExpose != nil {
		in, out := &in.ServiceExpose, &out.ServiceExpose
		*out = new(ServiceExpose)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PGBouncerSpec.
func (in *PGBouncerSpec) DeepCopy() *PGBouncerSpec {
	if in == nil {
		return nil
	}
	out := new(PGBouncerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PGBouncerStatus) DeepCopyInto(out *PGBouncerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PGBouncerStatus.
func (in *PGBouncerStatus) DeepCopy() *PGBouncerStatus {
	if in == nil {
		return nil
	}
	out := new(PGBouncerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PGInstanceSetSpec) DeepCopyInto(out *PGInstanceSetSpec) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(v1beta1.Metadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PriorityClassName != nil {
		in, out := &in.PriorityClassName, &out.PriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WALVolumeClaimSpec != nil {
		in, out := &in.WALVolumeClaimSpec, &out.WALVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	in.DataVolumeClaimSpec.DeepCopyInto(&out.DataVolumeClaimSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PGInstanceSetSpec.
func (in *PGInstanceSetSpec) DeepCopy() *PGInstanceSetSpec {
	if in == nil {
		return nil
	}
	out := new(PGInstanceSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PGInstanceSets) DeepCopyInto(out *PGInstanceSets) {
	{
		in := &in
		*out = make(PGInstanceSets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PGInstanceSets.
func (in PGInstanceSets) DeepCopy() PGInstanceSets {
	if in == nil {
		return nil
	}
	out := new(PGInstanceSets)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PGProxySpec) DeepCopyInto(out *PGProxySpec) {
	*out = *in
	if in.PGBouncer != nil {
		in, out := &in.PGBouncer, &out.PGBouncer
		*out = new(PGBouncerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PGProxySpec.
func (in *PGProxySpec) DeepCopy() *PGProxySpec {
	if in == nil {
		return nil
	}
	out := new(PGProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PMMSpec) DeepCopyInto(out *PMMSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PMMSpec.
func (in *PMMSpec) DeepCopy() *PMMSpec {
	if in == nil {
		return nil
	}
	out := new(PMMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGBackup) DeepCopyInto(out *PerconaPGBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGBackup.
func (in *PerconaPGBackup) DeepCopy() *PerconaPGBackup {
	if in == nil {
		return nil
	}
	out := new(PerconaPGBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaPGBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGBackupList) DeepCopyInto(out *PerconaPGBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaPGBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGBackupList.
func (in *PerconaPGBackupList) DeepCopy() *PerconaPGBackupList {
	if in == nil {
		return nil
	}
	out := new(PerconaPGBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaPGBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGBackupSpec) DeepCopyInto(out *PerconaPGBackupSpec) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGBackupSpec.
func (in *PerconaPGBackupSpec) DeepCopy() *PerconaPGBackupSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaPGBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGBackupStatus) DeepCopyInto(out *PerconaPGBackupStatus) {
	*out = *in
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGBackupStatus.
func (in *PerconaPGBackupStatus) DeepCopy() *PerconaPGBackupStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaPGBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGCluster) DeepCopyInto(out *PerconaPGCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGCluster.
func (in *PerconaPGCluster) DeepCopy() *PerconaPGCluster {
	if in == nil {
		return nil
	}
	out := new(PerconaPGCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaPGCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGClusterList) DeepCopyInto(out *PerconaPGClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaPGCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGClusterList.
func (in *PerconaPGClusterList) DeepCopy() *PerconaPGClusterList {
	if in == nil {
		return nil
	}
	out := new(PerconaPGClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaPGClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGClusterSpec) DeepCopyInto(out *PerconaPGClusterSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(ServiceExpose)
		(*in).DeepCopyInto(*out)
	}
	in.Secrets.DeepCopyInto(&out.Secrets)
	if in.Standby != nil {
		in, out := &in.Standby, &out.Standby
		*out = new(v1beta1.PostgresStandbySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenShift != nil {
		in, out := &in.OpenShift, &out.OpenShift
		*out = new(bool)
		**out = **in
	}
	if in.Patroni != nil {
		in, out := &in.Patroni, &out.Patroni
		*out = new(v1beta1.PatroniSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]v1beta1.PostgresUserSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DatabaseInitSQL != nil {
		in, out := &in.DatabaseInitSQL, &out.DatabaseInitSQL
		*out = new(v1beta1.DatabaseInitSQL)
		**out = **in
	}
	if in.Pause != nil {
		in, out := &in.Pause, &out.Pause
		*out = new(bool)
		**out = **in
	}
	if in.Unmanaged != nil {
		in, out := &in.Unmanaged, &out.Unmanaged
		*out = new(bool)
		**out = **in
	}
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = new(v1beta1.DataSource)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceSets != nil {
		in, out := &in.InstanceSets, &out.InstanceSets
		*out = make(PGInstanceSets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(PGProxySpec)
		(*in).DeepCopyInto(*out)
	}
	in.Backups.DeepCopyInto(&out.Backups)
	if in.PMM != nil {
		in, out := &in.PMM, &out.PMM
		*out = new(PMMSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGClusterSpec.
func (in *PerconaPGClusterSpec) DeepCopy() *PerconaPGClusterSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaPGClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGClusterStatus) DeepCopyInto(out *PerconaPGClusterStatus) {
	*out = *in
	in.Postgres.DeepCopyInto(&out.Postgres)
	out.PGBouncer = in.PGBouncer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGClusterStatus.
func (in *PerconaPGClusterStatus) DeepCopy() *PerconaPGClusterStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaPGClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGRestore) DeepCopyInto(out *PerconaPGRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGRestore.
func (in *PerconaPGRestore) DeepCopy() *PerconaPGRestore {
	if in == nil {
		return nil
	}
	out := new(PerconaPGRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaPGRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGRestoreList) DeepCopyInto(out *PerconaPGRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaPGRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGRestoreList.
func (in *PerconaPGRestoreList) DeepCopy() *PerconaPGRestoreList {
	if in == nil {
		return nil
	}
	out := new(PerconaPGRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaPGRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGRestoreSpec) DeepCopyInto(out *PerconaPGRestoreSpec) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGRestoreSpec.
func (in *PerconaPGRestoreSpec) DeepCopy() *PerconaPGRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaPGRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaPGRestoreStatus) DeepCopyInto(out *PerconaPGRestoreStatus) {
	*out = *in
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaPGRestoreStatus.
func (in *PerconaPGRestoreStatus) DeepCopy() *PerconaPGRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaPGRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresInstanceSetStatus) DeepCopyInto(out *PostgresInstanceSetStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresInstanceSetStatus.
func (in *PostgresInstanceSetStatus) DeepCopy() *PostgresInstanceSetStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresInstanceSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresStatus) DeepCopyInto(out *PostgresStatus) {
	*out = *in
	if in.InstanceSets != nil {
		in, out := &in.InstanceSets, &out.InstanceSets
		*out = make([]PostgresInstanceSetStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresStatus.
func (in *PostgresStatus) DeepCopy() *PostgresStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsSpec) DeepCopyInto(out *SecretsSpec) {
	*out = *in
	if in.CustomTLSSecret != nil {
		in, out := &in.CustomTLSSecret, &out.CustomTLSSecret
		*out = new(v1.SecretProjection)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomReplicationClientTLSSecret != nil {
		in, out := &in.CustomReplicationClientTLSSecret, &out.CustomReplicationClientTLSSecret
		*out = new(v1.SecretProjection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsSpec.
func (in *SecretsSpec) DeepCopy() *SecretsSpec {
	if in == nil {
		return nil
	}
	out := new(SecretsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExpose) DeepCopyInto(out *ServiceExpose) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExpose.
func (in *ServiceExpose) DeepCopy() *ServiceExpose {
	if in == nil {
		return nil
	}
	out := new(ServiceExpose)
	in.DeepCopyInto(out)
	return out
}
