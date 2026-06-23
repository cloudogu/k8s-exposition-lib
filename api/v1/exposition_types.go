/*
Copyright 2026.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ConditionReady indicates that the exposition is active and all routes/ports are mapped.
	ConditionReady = "Ready"
	// ReasonPortsAllocated indicates that all requested ports have been allocated.
	ReasonPortsAllocated = "PortsAllocated"
)

// RegexRewrite defines a complex URL rewrite using a regular expression.
type RegexRewrite struct {
	// Pattern is the regular expression pattern to match against the request path.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Pattern string `json:"pattern"`

	// Replacement is the replacement string, which may include capture group references (e.g. $1).
	// +kubebuilder:validation:Required
	Replacement string `json:"replacement"`
}

// Rewrite defines optional rewrite rules for HTTP routes.
type Rewrite struct {
	// StripPrefix removes the specified prefix from the request path (e.g. Traefik StripPrefix).
	// +optional
	StripPrefix *string `json:"stripPrefix,omitempty"`

	// Regex defines a complex rewrite using a regular expression (e.g. Traefik ReplacePathRegex).
	// +optional
	Regex *RegexRewrite `json:"regex,omitempty"`
}

// HTTPEntry defines an HTTP interface (Layer 7) for the exposition.
type HTTPEntry struct {
	// Name is a unique identifier for this HTTP route.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Name string `json:"name"`

	// Service is the name of the Kubernetes Service to route traffic to.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?([.][a-z0-9]([-a-z0-9]*[a-z0-9])?)*$
	Service string `json:"service"`

	// Port is the port of the referenced Kubernetes Service.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port"`

	// Path is the URL path under which the application should be reachable.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=^/.*$
	Path string `json:"path"`

	// Rewrite defines optional rewrite rules for this route.
	// +optional
	Rewrite *Rewrite `json:"rewrite,omitempty"`
}

// TCPEntry defines a TCP interface (Layer 4) for the exposition.
type TCPEntry struct {
	// Name is a unique identifier for this TCP route.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Name string `json:"name"`

	// Service is the name of the Kubernetes Service to route traffic to.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?([.][a-z0-9]([-a-z0-9]*[a-z0-9])?)*$
	Service string `json:"service"`

	// Port is the port of the referenced Kubernetes Service.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port"`

	// RequestedExternalPort is the desired external port. In Multi-CES environments,
	// this may be overridden or dynamically assigned by the operator.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +optional
	RequestedExternalPort *int32 `json:"requestedExternalPort,omitempty"`

	// Protocol is a protocol hint (e.g. "ssh") for documentation or firewall rules.
	// +kubebuilder:validation:MinLength=1
	// +optional
	Protocol *string `json:"protocol,omitempty"`
}

// UDPEntry defines a UDP interface (Layer 4) for the exposition.
type UDPEntry struct {
	// Name is a unique identifier for this UDP route.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Name string `json:"name"`

	// Service is the name of the Kubernetes Service to route traffic to.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?([.][a-z0-9]([-a-z0-9]*[a-z0-9])?)*$
	Service string `json:"service"`

	// Port is the port of the referenced Kubernetes Service.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port"`

	// RequestedExternalPort is the desired external port. In Multi-CES environments,
	// this may be overridden or dynamically assigned by the operator.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +optional
	RequestedExternalPort *int32 `json:"requestedExternalPort,omitempty"`

	// Protocol is a protocol hint for documentation or firewall rules.
	// +kubebuilder:validation:MinLength=1
	// +optional
	Protocol *string `json:"protocol,omitempty"`
}

// AllocatedPort represents an actually assigned external port for a TCP or UDP entry.
type AllocatedPort struct {
	// Name matches the name of the corresponding TCP or UDP entry.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=^[a-z0-9]([-a-z0-9]*[a-z0-9])?$
	Name string `json:"name"`

	// ExternalPort is the actually assigned external port.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	ExternalPort int32 `json:"externalPort"`
}

// ExpositionSpec defines the desired state of Exposition.
type ExpositionSpec struct {
	// HTTP defines the HTTP interfaces (Layer 7) for this exposition.
	// +listType=map
	// +listMapKey=name
	// +optional
	HTTP []HTTPEntry `json:"http,omitempty"`

	// TCP defines the TCP interfaces (Layer 4) for this exposition.
	// +listType=map
	// +listMapKey=name
	// +optional
	TCP []TCPEntry `json:"tcp,omitempty"`

	// UDP defines the UDP interfaces (Layer 4) for this exposition.
	// +listType=map
	// +listMapKey=name
	// +optional
	UDP []UDPEntry `json:"udp,omitempty"`
}

// ExpositionStatus defines the observed state of Exposition.
type ExpositionStatus struct {
	// AllocatedPorts contains the actually assigned external ports for TCP and UDP entries.
	// +listType=map
	// +listMapKey=name
	// +optional
	AllocatedPorts []AllocatedPort `json:"allocatedPorts,omitempty"`

	// Conditions represent the current state of the Exposition resource.
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=exp
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Valid",type="string",JSONPath=".status.conditions[?(@.type == 'Valid')].status",description="Whether the exposition is valid"
// +kubebuilder:printcolumn:name="IngressesReady",type="string",JSONPath=".status.conditions[?(@.type == 'IngressesReady')].status",description="Whether the ingresses are ready"
// +kubebuilder:printcolumn:name="HTTP-Paths",type="string",JSONPath=".spec.http[*].path",description="Configured HTTP paths"
// +kubebuilder:printcolumn:name="External-Ports",type="string",JSONPath=".status.allocatedPorts[*].externalPort",description="Allocated external TCP/UDP ports"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the resource"

// Exposition is the Schema for the expositions API.
type Exposition struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of Exposition.
	// +required
	Spec ExpositionSpec `json:"spec"`

	// status defines the observed state of Exposition.
	// +optional
	Status ExpositionStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ExpositionList contains a list of Exposition.
type ExpositionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []Exposition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Exposition{}, &ExpositionList{})
}
