//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	conditionsv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/third_party/conditions/apis/conditions/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIBinding) DeepCopyInto(out *APIBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIBinding.
func (in *APIBinding) DeepCopy() *APIBinding {
	if in == nil {
		return nil
	}
	out := new(APIBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIBindingList) DeepCopyInto(out *APIBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIBindingList.
func (in *APIBindingList) DeepCopy() *APIBindingList {
	if in == nil {
		return nil
	}
	out := new(APIBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIBindingSpec) DeepCopyInto(out *APIBindingSpec) {
	*out = *in
	in.Reference.DeepCopyInto(&out.Reference)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIBindingSpec.
func (in *APIBindingSpec) DeepCopy() *APIBindingSpec {
	if in == nil {
		return nil
	}
	out := new(APIBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIBindingStatus) DeepCopyInto(out *APIBindingStatus) {
	*out = *in
	if in.BoundAPIExport != nil {
		in, out := &in.BoundAPIExport, &out.BoundAPIExport
		*out = new(ExportReference)
		(*in).DeepCopyInto(*out)
	}
	if in.BoundResources != nil {
		in, out := &in.BoundResources, &out.BoundResources
		*out = make([]BoundAPIResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIBindingStatus.
func (in *APIBindingStatus) DeepCopy() *APIBindingStatus {
	if in == nil {
		return nil
	}
	out := new(APIBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIExport) DeepCopyInto(out *APIExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIExport.
func (in *APIExport) DeepCopy() *APIExport {
	if in == nil {
		return nil
	}
	out := new(APIExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIExportList) DeepCopyInto(out *APIExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIExportList.
func (in *APIExportList) DeepCopy() *APIExportList {
	if in == nil {
		return nil
	}
	out := new(APIExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIExportSpec) DeepCopyInto(out *APIExportSpec) {
	*out = *in
	if in.LatestResourceSchemas != nil {
		in, out := &in.LatestResourceSchemas, &out.LatestResourceSchemas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity)
		(*in).DeepCopyInto(*out)
	}
	if in.MaximalPermissionPolicy != nil {
		in, out := &in.MaximalPermissionPolicy, &out.MaximalPermissionPolicy
		*out = new(MaximalPermissionPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIExportSpec.
func (in *APIExportSpec) DeepCopy() *APIExportSpec {
	if in == nil {
		return nil
	}
	out := new(APIExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIExportStatus) DeepCopyInto(out *APIExportStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualWorkspaces != nil {
		in, out := &in.VirtualWorkspaces, &out.VirtualWorkspaces
		*out = make([]VirtualWorkspace, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIExportStatus.
func (in *APIExportStatus) DeepCopy() *APIExportStatus {
	if in == nil {
		return nil
	}
	out := new(APIExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceSchema) DeepCopyInto(out *APIResourceSchema) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceSchema.
func (in *APIResourceSchema) DeepCopy() *APIResourceSchema {
	if in == nil {
		return nil
	}
	out := new(APIResourceSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIResourceSchema) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceSchemaList) DeepCopyInto(out *APIResourceSchemaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIResourceSchema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceSchemaList.
func (in *APIResourceSchemaList) DeepCopy() *APIResourceSchemaList {
	if in == nil {
		return nil
	}
	out := new(APIResourceSchemaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIResourceSchemaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceSchemaSpec) DeepCopyInto(out *APIResourceSchemaSpec) {
	*out = *in
	in.Names.DeepCopyInto(&out.Names)
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]APIResourceVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceSchemaSpec.
func (in *APIResourceSchemaSpec) DeepCopy() *APIResourceSchemaSpec {
	if in == nil {
		return nil
	}
	out := new(APIResourceSchemaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceVersion) DeepCopyInto(out *APIResourceVersion) {
	*out = *in
	if in.DeprecationWarning != nil {
		in, out := &in.DeprecationWarning, &out.DeprecationWarning
		*out = new(string)
		**out = **in
	}
	in.Schema.DeepCopyInto(&out.Schema)
	in.Subresources.DeepCopyInto(&out.Subresources)
	if in.AdditionalPrinterColumns != nil {
		in, out := &in.AdditionalPrinterColumns, &out.AdditionalPrinterColumns
		*out = make([]v1.CustomResourceColumnDefinition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceVersion.
func (in *APIResourceVersion) DeepCopy() *APIResourceVersion {
	if in == nil {
		return nil
	}
	out := new(APIResourceVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoundAPIResource) DeepCopyInto(out *BoundAPIResource) {
	*out = *in
	out.Schema = in.Schema
	if in.StorageVersions != nil {
		in, out := &in.StorageVersions, &out.StorageVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoundAPIResource.
func (in *BoundAPIResource) DeepCopy() *BoundAPIResource {
	if in == nil {
		return nil
	}
	out := new(BoundAPIResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoundAPIResourceSchema) DeepCopyInto(out *BoundAPIResourceSchema) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoundAPIResourceSchema.
func (in *BoundAPIResourceSchema) DeepCopy() *BoundAPIResourceSchema {
	if in == nil {
		return nil
	}
	out := new(BoundAPIResourceSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportReference) DeepCopyInto(out *ExportReference) {
	*out = *in
	if in.Workspace != nil {
		in, out := &in.Workspace, &out.Workspace
		*out = new(WorkspaceExportReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportReference.
func (in *ExportReference) DeepCopy() *ExportReference {
	if in == nil {
		return nil
	}
	out := new(ExportReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity) DeepCopyInto(out *Identity) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.SecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in *Identity) DeepCopy() *Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalAPIExportPolicy) DeepCopyInto(out *LocalAPIExportPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalAPIExportPolicy.
func (in *LocalAPIExportPolicy) DeepCopy() *LocalAPIExportPolicy {
	if in == nil {
		return nil
	}
	out := new(LocalAPIExportPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaximalPermissionPolicy) DeepCopyInto(out *MaximalPermissionPolicy) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalAPIExportPolicy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaximalPermissionPolicy.
func (in *MaximalPermissionPolicy) DeepCopy() *MaximalPermissionPolicy {
	if in == nil {
		return nil
	}
	out := new(MaximalPermissionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualWorkspace) DeepCopyInto(out *VirtualWorkspace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualWorkspace.
func (in *VirtualWorkspace) DeepCopy() *VirtualWorkspace {
	if in == nil {
		return nil
	}
	out := new(VirtualWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceExportReference) DeepCopyInto(out *WorkspaceExportReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceExportReference.
func (in *WorkspaceExportReference) DeepCopy() *WorkspaceExportReference {
	if in == nil {
		return nil
	}
	out := new(WorkspaceExportReference)
	in.DeepCopyInto(out)
	return out
}
