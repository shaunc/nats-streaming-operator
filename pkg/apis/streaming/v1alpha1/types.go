package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NatsStreamingClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []NatsStreamingCluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NatsStreamingCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NatsStreamingClusterSpec   `json:"spec"`
	Status            NatsStreamingClusterStatus `json:"status,omitempty"`
}

type NatsStreamingClusterSpec struct {
	// Size is the number of nodes in the NATS Streaming cluster.
	Size int32 `json:"size"`

	// NatsService is the Kubernetes service to which the
	// NATS Streaming nodes will connect.
	NatsService string `json:"natsSvc"`
}

type NatsStreamingClusterStatus struct {
	// TODO
}
