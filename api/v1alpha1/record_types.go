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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	RecordScheduled   = RecordStatusPhase("Scheduled")
	RecordPreparation = RecordStatusPhase("Preparation")
	RecordRecording   = RecordStatusPhase("Recording")
	RecordFinished    = RecordStatusPhase("Finished")
	RecordCanceled    = RecordStatusPhase("Canceled")
	RecordFailed      = RecordStatusPhase("Failed")
	RecordUnknown     = RecordStatusPhase("Unknown")
)

//+kubebuilder:validation:Enum=Scheduled;Recording;Finished;Canceled;Failed;Unknown

type RecordStatusPhase string

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	RecordName  string        `json:"recordName"`
	StartTime   RecordTime    `json:"startTime,omitempty"`
	EndTime     RecordTime    `json:"endTime"`
	Channel     RecordChannel `json:"channel"`
	SaveSetting SaveSetting   `json:"saveSetting"`
	//+kubebuilder:default=false
	Suspend bool `json:"suspend"`
}

type RecordTime struct {
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=9999
	Years uint16 `json:"years"`
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=12
	Months uint16 `json:"months"`
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=31
	Days uint16 `json:"days"`
	//+kubebuilder:validation:Minimum=0
	//+kubebuilder:validation:Maximum=23
	Hours uint16 `json:"hours"`
	//+kubebuilder:validation:Minimum=0
	//+kubebuilder:validation:Maximum=59
	Minutes uint16 `json:"minutes"`
}

type RecordChannel struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

type SaveSetting struct {
	FileName          string        `json:"fileName"`
	TemporarySavePath string        `json:"temporarySavePath,omitempty"`
	SavePath          string        `json:"savePath"`
	Volume            corev1.Volume `json:"volume,omitempty"`
}

// RecordStatus defines the observed state of Record
type RecordStatus struct {
	Phase   RecordStatusPhase `json:"phase"`
	Reason  string            `json:"reason,omitempty"`
	Message string            `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=rc

// Record is the Schema for the records API
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecordSpec   `json:"spec,omitempty"`
	Status RecordStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RecordList contains a list of Record
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}
