// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FeaturemembershipBinauthz struct {
	/* Whether binauthz is enabled in this cluster. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type FeaturemembershipConfigSync struct {
	/*  */
	// +optional
	Git *FeaturemembershipGit `json:"git,omitempty"`

	/* Specifies whether the Config Sync Repo is in "hierarchical" or "unstructured" mode. */
	// +optional
	SourceFormat *string `json:"sourceFormat,omitempty"`
}

type FeaturemembershipConfigmanagement struct {
	/* Binauthz conifguration for the cluster. */
	// +optional
	Binauthz *FeaturemembershipBinauthz `json:"binauthz,omitempty"`

	/* Config Sync configuration for the cluster. */
	// +optional
	ConfigSync *FeaturemembershipConfigSync `json:"configSync,omitempty"`

	/* Hierarchy Controller configuration for the cluster. */
	// +optional
	HierarchyController *FeaturemembershipHierarchyController `json:"hierarchyController,omitempty"`

	/* Policy Controller configuration for the cluster. */
	// +optional
	PolicyController *FeaturemembershipPolicyController `json:"policyController,omitempty"`

	/* Version of ACM installed. */
	// +optional
	Version *string `json:"version,omitempty"`
}

type FeaturemembershipGit struct {
	/* URL for the HTTPS proxy to be used when communicating with the Git repo. */
	// +optional
	HttpsProxy *string `json:"httpsProxy,omitempty"`

	/* The path within the Git repository that represents the top level of the repo to sync. Default: the root directory of the repository. */
	// +optional
	PolicyDir *string `json:"policyDir,omitempty"`

	/* Type of secret configured for access to the Git repo. */
	// +optional
	SecretType *string `json:"secretType,omitempty"`

	/* The branch of the repository to sync from. Default: master. */
	// +optional
	SyncBranch *string `json:"syncBranch,omitempty"`

	/* The URL of the Git repository to use as the source of truth. */
	// +optional
	SyncRepo *string `json:"syncRepo,omitempty"`

	/* Git revision (tag or hash) to check out. Default HEAD. */
	// +optional
	SyncRev *string `json:"syncRev,omitempty"`

	/* Period in seconds between consecutive syncs. Default: 15. */
	// +optional
	SyncWaitSecs *string `json:"syncWaitSecs,omitempty"`
}

type FeaturemembershipHierarchyController struct {
	/* Whether hierarchical resource quota is enabled in this cluster. */
	// +optional
	EnableHierarchicalResourceQuota *bool `json:"enableHierarchicalResourceQuota,omitempty"`

	/* Whether pod tree labels are enabled in this cluster. */
	// +optional
	EnablePodTreeLabels *bool `json:"enablePodTreeLabels,omitempty"`

	/* Whether Hierarchy Controller is enabled in this cluster. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type FeaturemembershipPolicyController struct {
	/* Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether. */
	// +optional
	AuditIntervalSeconds *string `json:"auditIntervalSeconds,omitempty"`

	/* Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster. */
	// +optional
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty"`

	/* Logs all denies and dry run failures. */
	// +optional
	LogDeniesEnabled *bool `json:"logDeniesEnabled,omitempty"`

	/* Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated. */
	// +optional
	ReferentialRulesEnabled *bool `json:"referentialRulesEnabled,omitempty"`

	/* Installs the default template library along with Policy Controller. */
	// +optional
	TemplateLibraryInstalled *bool `json:"templateLibraryInstalled,omitempty"`
}

type GKEHubFeatureMembershipSpec struct {
	/* Config Management-specific spec. */
	// +optional
	Configmanagement *FeaturemembershipConfigmanagement `json:"configmanagement,omitempty"`

	/*  */
	FeatureRef v1alpha1.ResourceRef `json:"featureRef"`

	/* The location of the feature */
	Location string `json:"location"`

	/*  */
	MembershipRef v1alpha1.ResourceRef `json:"membershipRef"`

	/* The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`
}

type GKEHubFeatureMembershipStatus struct {
	/* Conditions represent the latest available observations of the
	   GKEHubFeatureMembership's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GKEHubFeatureMembership is the Schema for the gkehub API
// +k8s:openapi-gen=true
type GKEHubFeatureMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEHubFeatureMembershipSpec   `json:"spec,omitempty"`
	Status GKEHubFeatureMembershipStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GKEHubFeatureMembershipList contains a list of GKEHubFeatureMembership
type GKEHubFeatureMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEHubFeatureMembership `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEHubFeatureMembership{}, &GKEHubFeatureMembershipList{})
}
