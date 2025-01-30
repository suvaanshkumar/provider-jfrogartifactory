// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LocalNpmRepositoryInitParameters struct {

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory.
	// This may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `json:"archiveBrowsingEnabled,omitempty" tf:"archive_browsing_enabled,omitempty"`

	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `json:"blackedOut,omitempty" tf:"blacked_out,omitempty"`

	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `json:"cdnRedirect,omitempty" tf:"cdn_redirect,omitempty"`

	// Public description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `json:"downloadDirect,omitempty" tf:"download_direct,omitempty"`

	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no artifacts are excluded.
	ExcludesPattern *string `json:"excludesPattern,omitempty" tf:"excludes_pattern,omitempty"`

	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `json:"includesPattern,omitempty" tf:"includes_pattern,omitempty"`

	// Internal description.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `json:"priorityResolution,omitempty" tf:"priority_resolution,omitempty"`

	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The attribute should only be used if the repository is already assigned to the existing project.
	// +listType=set
	ProjectEnvironments []*string `json:"projectEnvironments,omitempty" tf:"project_environments,omitempty"`

	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// List of property set name
	// +listType=set
	PropertySets []*string `json:"propertySets,omitempty" tf:"property_sets,omitempty"`

	// Repository layout key for the local repository
	RepoLayoutRef *string `json:"repoLayoutRef,omitempty" tf:"repo_layout_ref,omitempty"`

	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
	XrayIndex *bool `json:"xrayIndex,omitempty" tf:"xray_index,omitempty"`
}

type LocalNpmRepositoryObservation struct {

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory.
	// This may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `json:"archiveBrowsingEnabled,omitempty" tf:"archive_browsing_enabled,omitempty"`

	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `json:"blackedOut,omitempty" tf:"blacked_out,omitempty"`

	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `json:"cdnRedirect,omitempty" tf:"cdn_redirect,omitempty"`

	// Public description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `json:"downloadDirect,omitempty" tf:"download_direct,omitempty"`

	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no artifacts are excluded.
	ExcludesPattern *string `json:"excludesPattern,omitempty" tf:"excludes_pattern,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `json:"includesPattern,omitempty" tf:"includes_pattern,omitempty"`

	// Internal description.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	PackageType *string `json:"packageType,omitempty" tf:"package_type,omitempty"`

	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `json:"priorityResolution,omitempty" tf:"priority_resolution,omitempty"`

	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The attribute should only be used if the repository is already assigned to the existing project.
	// +listType=set
	ProjectEnvironments []*string `json:"projectEnvironments,omitempty" tf:"project_environments,omitempty"`

	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// List of property set name
	// +listType=set
	PropertySets []*string `json:"propertySets,omitempty" tf:"property_sets,omitempty"`

	// Repository layout key for the local repository
	RepoLayoutRef *string `json:"repoLayoutRef,omitempty" tf:"repo_layout_ref,omitempty"`

	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
	XrayIndex *bool `json:"xrayIndex,omitempty" tf:"xray_index,omitempty"`
}

type LocalNpmRepositoryParameters struct {

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory.
	// This may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).
	// +kubebuilder:validation:Optional
	ArchiveBrowsingEnabled *bool `json:"archiveBrowsingEnabled,omitempty" tf:"archive_browsing_enabled,omitempty"`

	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	// +kubebuilder:validation:Optional
	BlackedOut *bool `json:"blackedOut,omitempty" tf:"blacked_out,omitempty"`

	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	// +kubebuilder:validation:Optional
	CdnRedirect *bool `json:"cdnRedirect,omitempty" tf:"cdn_redirect,omitempty"`

	// Public description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.
	// +kubebuilder:validation:Optional
	DownloadDirect *bool `json:"downloadDirect,omitempty" tf:"download_direct,omitempty"`

	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no artifacts are excluded.
	// +kubebuilder:validation:Optional
	ExcludesPattern *string `json:"excludesPattern,omitempty" tf:"excludes_pattern,omitempty"`

	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	// +kubebuilder:validation:Optional
	IncludesPattern *string `json:"includesPattern,omitempty" tf:"includes_pattern,omitempty"`

	// Internal description.
	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	// +kubebuilder:validation:Optional
	PriorityResolution *bool `json:"priorityResolution,omitempty" tf:"priority_resolution,omitempty"`

	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The attribute should only be used if the repository is already assigned to the existing project.
	// +kubebuilder:validation:Optional
	// +listType=set
	ProjectEnvironments []*string `json:"projectEnvironments,omitempty" tf:"project_environments,omitempty"`

	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	// +kubebuilder:validation:Optional
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// List of property set name
	// +kubebuilder:validation:Optional
	// +listType=set
	PropertySets []*string `json:"propertySets,omitempty" tf:"property_sets,omitempty"`

	// Repository layout key for the local repository
	// +kubebuilder:validation:Optional
	RepoLayoutRef *string `json:"repoLayoutRef,omitempty" tf:"repo_layout_ref,omitempty"`

	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
	// +kubebuilder:validation:Optional
	XrayIndex *bool `json:"xrayIndex,omitempty" tf:"xray_index,omitempty"`
}

// LocalNpmRepositorySpec defines the desired state of LocalNpmRepository
type LocalNpmRepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocalNpmRepositoryParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LocalNpmRepositoryInitParameters `json:"initProvider,omitempty"`
}

// LocalNpmRepositoryStatus defines the observed state of LocalNpmRepository.
type LocalNpmRepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocalNpmRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LocalNpmRepository is the Schema for the LocalNpmRepositorys API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,artfactory}
type LocalNpmRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocalNpmRepositorySpec   `json:"spec"`
	Status            LocalNpmRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocalNpmRepositoryList contains a list of LocalNpmRepositorys
type LocalNpmRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocalNpmRepository `json:"items"`
}

// Repository type metadata.
var (
	LocalNpmRepository_Kind             = "LocalNpmRepository"
	LocalNpmRepository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocalNpmRepository_Kind}.String()
	LocalNpmRepository_KindAPIVersion   = LocalNpmRepository_Kind + "." + CRDGroupVersion.String()
	LocalNpmRepository_GroupVersionKind = CRDGroupVersion.WithKind(LocalNpmRepository_Kind)
)

func init() {
	SchemeBuilder.Register(&LocalNpmRepository{}, &LocalNpmRepositoryList{})
}
