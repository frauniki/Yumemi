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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MirakurunSpec defines the desired state of Mirakurun
type MirakurunSpec struct {
	Endpoint string `json:"endpoint"`
	//+kubebuilder:default=true
	IsDefault bool `json:"isDefault"`
}

type Tuner struct {
	Name    string   `json:"name"`
	Types   []string `json:"types"`
	IsReady bool     `json:"isReady"`
}

type Channel struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Type        string `json:"type"`
	Channel     string `json:"channel"`
}

// MirakurunStatus defines the observed state of Mirakurun
type MirakurunStatus struct {
	Tuners          []Tuner      `json:"tuners,omitempty"`
	Channels        []Channel    `json:"channels,omitempty"`
	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`
	Reason          string       `json:"reason,omitempty"`
	Message         string       `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=mk

// Mirakurun is the Schema for the mirakuruns API
type Mirakurun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MirakurunSpec   `json:"spec,omitempty"`
	Status MirakurunStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MirakurunList contains a list of Mirakurun
type MirakurunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mirakurun `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mirakurun{}, &MirakurunList{})
}
