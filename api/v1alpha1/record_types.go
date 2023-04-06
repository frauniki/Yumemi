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
	Scheduled = RecordStatusPhase("Scheduled")
	Recording = RecordStatusPhase("Recording")
	Finished  = RecordStatusPhase("Finished")
	Canceled  = RecordStatusPhase("Canceled")
	Failed    = RecordStatusPhase("Failed")
	Unknown   = RecordStatusPhase("Unknown")
)

//+kubebuilder:validation:Enum=Scheduled;Recording;Finished;Canceled;Failed;Unknown

type RecordStatusPhase string

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	RecordName  string        `json:"recordName"`
	StartTime   *metav1.Time  `json:"startTime,omitempty"`
	EndTime     *metav1.Time  `json:"endTime"`
	Channel     RecordChannel `json:"channel"`
	SaveSetting SaveSetting   `json:"saveSetting"`
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
	//+kubebuilder:default=Scheduled
	Phase   *RecordStatusPhase `json:"phase"`
	Reason  string             `json:"reason,omitempty"`
	Message string             `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

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
