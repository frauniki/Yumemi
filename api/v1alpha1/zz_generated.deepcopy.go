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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Channel) DeepCopyInto(out *Channel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Channel.
func (in *Channel) DeepCopy() *Channel {
	if in == nil {
		return nil
	}
	out := new(Channel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mirakurun) DeepCopyInto(out *Mirakurun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mirakurun.
func (in *Mirakurun) DeepCopy() *Mirakurun {
	if in == nil {
		return nil
	}
	out := new(Mirakurun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mirakurun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirakurunList) DeepCopyInto(out *MirakurunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mirakurun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirakurunList.
func (in *MirakurunList) DeepCopy() *MirakurunList {
	if in == nil {
		return nil
	}
	out := new(MirakurunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MirakurunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirakurunSpec) DeepCopyInto(out *MirakurunSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirakurunSpec.
func (in *MirakurunSpec) DeepCopy() *MirakurunSpec {
	if in == nil {
		return nil
	}
	out := new(MirakurunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirakurunStatus) DeepCopyInto(out *MirakurunStatus) {
	*out = *in
	if in.Tuners != nil {
		in, out := &in.Tuners, &out.Tuners
		*out = make([]Tuner, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = make([]Channel, len(*in))
		copy(*out, *in)
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirakurunStatus.
func (in *MirakurunStatus) DeepCopy() *MirakurunStatus {
	if in == nil {
		return nil
	}
	out := new(MirakurunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Record) DeepCopyInto(out *Record) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Record.
func (in *Record) DeepCopy() *Record {
	if in == nil {
		return nil
	}
	out := new(Record)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Record) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordChannel) DeepCopyInto(out *RecordChannel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordChannel.
func (in *RecordChannel) DeepCopy() *RecordChannel {
	if in == nil {
		return nil
	}
	out := new(RecordChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordList) DeepCopyInto(out *RecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Record, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordList.
func (in *RecordList) DeepCopy() *RecordList {
	if in == nil {
		return nil
	}
	out := new(RecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSpec) DeepCopyInto(out *RecordSpec) {
	*out = *in
	out.StartTime = in.StartTime
	out.EndTime = in.EndTime
	out.Channel = in.Channel
	in.SaveSetting.DeepCopyInto(&out.SaveSetting)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSpec.
func (in *RecordSpec) DeepCopy() *RecordSpec {
	if in == nil {
		return nil
	}
	out := new(RecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordStatus) DeepCopyInto(out *RecordStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordStatus.
func (in *RecordStatus) DeepCopy() *RecordStatus {
	if in == nil {
		return nil
	}
	out := new(RecordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordTime) DeepCopyInto(out *RecordTime) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordTime.
func (in *RecordTime) DeepCopy() *RecordTime {
	if in == nil {
		return nil
	}
	out := new(RecordTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SaveSetting) DeepCopyInto(out *SaveSetting) {
	*out = *in
	in.Volume.DeepCopyInto(&out.Volume)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SaveSetting.
func (in *SaveSetting) DeepCopy() *SaveSetting {
	if in == nil {
		return nil
	}
	out := new(SaveSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tuner) DeepCopyInto(out *Tuner) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tuner.
func (in *Tuner) DeepCopy() *Tuner {
	if in == nil {
		return nil
	}
	out := new(Tuner)
	in.DeepCopyInto(out)
	return out
}
