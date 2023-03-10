//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	"github.com/fluxcd/source-controller/api/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceTransformer) DeepCopyInto(out *SourceTransformer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceTransformer.
func (in *SourceTransformer) DeepCopy() *SourceTransformer {
	if in == nil {
		return nil
	}
	out := new(SourceTransformer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceTransformer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceTransformerList) DeepCopyInto(out *SourceTransformerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceTransformer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceTransformerList.
func (in *SourceTransformerList) DeepCopy() *SourceTransformerList {
	if in == nil {
		return nil
	}
	out := new(SourceTransformerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceTransformerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceTransformerSpec) DeepCopyInto(out *SourceTransformerSpec) {
	*out = *in
	out.SourceRef = in.SourceRef
	out.Transformer = in.Transformer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceTransformerSpec.
func (in *SourceTransformerSpec) DeepCopy() *SourceTransformerSpec {
	if in == nil {
		return nil
	}
	out := new(SourceTransformerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceTransformerStatus) DeepCopyInto(out *SourceTransformerStatus) {
	*out = *in
	if in.Artifact != nil {
		in, out := &in.Artifact, &out.Artifact
		*out = new(v1beta1.Artifact)
		(*in).DeepCopyInto(*out)
	}
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceTransformerStatus.
func (in *SourceTransformerStatus) DeepCopy() *SourceTransformerStatus {
	if in == nil {
		return nil
	}
	out := new(SourceTransformerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transformer) DeepCopyInto(out *Transformer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transformer.
func (in *Transformer) DeepCopy() *Transformer {
	if in == nil {
		return nil
	}
	out := new(Transformer)
	in.DeepCopyInto(out)
	return out
}
