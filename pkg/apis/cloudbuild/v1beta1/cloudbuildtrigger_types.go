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

type TriggerArtifacts struct {
	/* A list of images to be pushed upon the successful completion of all build steps.

	The images will be pushed using the builder service account's credentials.

	The digests of the pushed images will be stored in the Build resource's results field.

	If any of the images fail to be pushed, the build is marked FAILURE. */
	// +optional
	Images []string `json:"images,omitempty"`

	/* A list of objects to be uploaded to Cloud Storage upon successful completion of all build steps.

	Files in the workspace matching specified paths globs will be uploaded to the
	Cloud Storage location using the builder service account's credentials.

	The location and generation of the uploaded objects will be stored in the Build resource's results field.

	If any objects fail to be pushed, the build is marked FAILURE. */
	// +optional
	Objects *TriggerObjects `json:"objects,omitempty"`
}

type TriggerBuild struct {
	/* Artifacts produced by the build that should be uploaded upon successful completion of all build steps. */
	// +optional
	Artifacts *TriggerArtifacts `json:"artifacts,omitempty"`

	/* A list of images to be pushed upon the successful completion of all build steps.
	The images are pushed using the builder service account's credentials.
	The digests of the pushed images will be stored in the Build resource's results field.
	If any of the images fail to be pushed, the build status is marked FAILURE. */
	// +optional
	Images []string `json:"images,omitempty"`

	/* Google Cloud Storage bucket where logs should be written. Logs file
	names will be of the format ${logsBucket}/log-${build_id}.txt. */
	// +optional
	LogsBucketRef *v1alpha1.ResourceRef `json:"logsBucketRef,omitempty"`

	/* Special options for this build. */
	// +optional
	Options *TriggerOptions `json:"options,omitempty"`

	/* TTL in queue for this build. If provided and the build is enqueued longer than this value,
	the build will expire and the build status will be EXPIRED.
	The TTL starts ticking from createTime.
	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	QueueTtl *string `json:"queueTtl,omitempty"`

	/* Secrets to decrypt using Cloud Key Management Service. */
	// +optional
	Secret []TriggerSecret `json:"secret,omitempty"`

	/* The location of the source files to build.

	One of 'storageSource' or 'repoSource' must be provided. */
	// +optional
	Source *TriggerSource `json:"source,omitempty"`

	/* The operations to be performed on the workspace. */
	Step []TriggerStep `json:"step"`

	/* Substitutions data for Build resource. */
	// +optional
	Substitutions map[string]string `json:"substitutions,omitempty"`

	/* Tags for annotation of a Build. These are not docker tags. */
	// +optional
	Tags []string `json:"tags,omitempty"`

	/* Amount of time that this build should be allowed to run, to second granularity.
	If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.
	This timeout must be equal to or greater than the sum of the timeouts for build steps within the build.
	The expected format is the number of seconds followed by s.
	Default time is ten minutes (600s). */
	// +optional
	Timeout *string `json:"timeout,omitempty"`
}

type TriggerGithub struct {
	/* Name of the repository. For example: The name for
	https://github.com/googlecloudplatform/cloud-builders is "cloud-builders". */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Owner of the repository. For example: The owner for
	https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform". */
	// +optional
	Owner *string `json:"owner,omitempty"`

	/* filter to match changes in pull requests.  Specify only one of pullRequest or push. */
	// +optional
	PullRequest *TriggerPullRequest `json:"pullRequest,omitempty"`

	/* filter to match changes in refs, like branches or tags.  Specify only one of pullRequest or push. */
	// +optional
	Push *TriggerPush `json:"push,omitempty"`
}

type TriggerObjects struct {
	/* Cloud Storage bucket and optional object path, in the form "gs://bucket/path/to/somewhere/".

	Files in the workspace matching any path pattern will be uploaded to Cloud Storage with
	this location as a prefix. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Path globs used to match files in the build's workspace. */
	// +optional
	Paths []string `json:"paths,omitempty"`

	/* Output only. Stores timing information for pushing all artifact objects. */
	// +optional
	Timing []TriggerTiming `json:"timing,omitempty"`
}

type TriggerOptions struct {
	/* Requested disk size for the VM that runs the build. Note that this is NOT "disk free";
	some of the space will be used by the operating system and build utilities.
	Also note that this is the minimum disk size that will be allocated for the build --
	the build may run with a larger disk than requested. At present, the maximum disk size
	is 1000GB; builds that request more than the maximum are rejected with an error. */
	// +optional
	DiskSizeGb *int `json:"diskSizeGb,omitempty"`

	/* Option to specify whether or not to apply bash style string operations to the substitutions.

	NOTE this is always enabled for triggered builds and cannot be overridden in the build configuration file. */
	// +optional
	DynamicSubstitutions *bool `json:"dynamicSubstitutions,omitempty"`

	/* A list of global environment variable definitions that will exist for all build steps
	in this build. If a variable is defined in both globally and in a build step,
	the variable will use the build step value.

	The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE". */
	// +optional
	Env []string `json:"env,omitempty"`

	/* Option to define build log streaming behavior to Google Cloud Storage. Possible values: ["STREAM_DEFAULT", "STREAM_ON", "STREAM_OFF"] */
	// +optional
	LogStreamingOption *string `json:"logStreamingOption,omitempty"`

	/* Option to specify the logging mode, which determines if and where build logs are stored. Possible values: ["LOGGING_UNSPECIFIED", "LEGACY", "GCS_ONLY", "STACKDRIVER_ONLY", "NONE"] */
	// +optional
	Logging *string `json:"logging,omitempty"`

	/* Compute Engine machine type on which to run the build. Possible values: ["UNSPECIFIED", "N1_HIGHCPU_8", "N1_HIGHCPU_32", "E2_HIGHCPU_8", "E2_HIGHCPU_32"] */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/* Requested verifiability options. Possible values: ["NOT_VERIFIED", "VERIFIED"] */
	// +optional
	RequestedVerifyOption *string `json:"requestedVerifyOption,omitempty"`

	/* A list of global environment variables, which are encrypted using a Cloud Key Management
	Service crypto key. These values must be specified in the build's Secret. These variables
	will be available to all build steps in this build. */
	// +optional
	SecretEnv []string `json:"secretEnv,omitempty"`

	/* Requested hash for SourceProvenance. Possible values: ["NONE", "SHA256", "MD5"] */
	// +optional
	SourceProvenanceHash []string `json:"sourceProvenanceHash,omitempty"`

	/* Option to specify behavior when there is an error in the substitution checks.

	NOTE this is always set to ALLOW_LOOSE for triggered builds and cannot be overridden
	in the build configuration file. Possible values: ["MUST_MATCH", "ALLOW_LOOSE"] */
	// +optional
	SubstitutionOption *string `json:"substitutionOption,omitempty"`

	/* Global list of volumes to mount for ALL build steps

	Each volume is created as an empty volume prior to starting the build process.
	Upon completion of the build, volumes and their contents are discarded. Global
	volume names and paths cannot conflict with the volumes defined a build step.

	Using a global volume in a build with only one step is not valid as it is indicative
	of a build request with an incorrect configuration. */
	// +optional
	Volumes []TriggerVolumes `json:"volumes,omitempty"`

	/* Option to specify a WorkerPool for the build. Format projects/{project}/workerPools/{workerPool}

	This field is experimental. */
	// +optional
	WorkerPool *string `json:"workerPool,omitempty"`
}

type TriggerPubsubConfig struct {
	/* Service account that will make the push request. */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* Potential issues with the underlying Pub/Sub subscription configuration.
	Only populated on get requests. */
	// +optional
	State *string `json:"state,omitempty"`

	/* Output only. Name of the subscription. */
	// +optional
	Subscription *string `json:"subscription,omitempty"`

	/* The name of the topic from which this subscription
	is receiving messages. */
	TopicRef v1alpha1.ResourceRef `json:"topicRef"`
}

type TriggerPullRequest struct {
	/* Regex of branches to match. */
	Branch string `json:"branch"`

	/* Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator. Possible values: ["COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"] */
	// +optional
	CommentControl *string `json:"commentControl,omitempty"`

	/* If true, branches that do NOT match the git_ref will trigger a build. */
	// +optional
	InvertRegex *bool `json:"invertRegex,omitempty"`
}

type TriggerPush struct {
	/* Regex of branches to match.  Specify only one of branch or tag. */
	// +optional
	Branch *string `json:"branch,omitempty"`

	/* When true, only trigger a build if the revision regex does NOT match the git_ref regex. */
	// +optional
	InvertRegex *bool `json:"invertRegex,omitempty"`

	/* Regex of tags to match.  Specify only one of branch or tag. */
	// +optional
	Tag *string `json:"tag,omitempty"`
}

type TriggerRepoSource struct {
	/* Regex matching branches to build. Exactly one a of branch name, tag, or commit SHA must be provided.
	The syntax of the regular expressions accepted is the syntax accepted by RE2 and
	described at https://github.com/google/re2/wiki/Syntax */
	// +optional
	BranchName *string `json:"branchName,omitempty"`

	/* Explicit commit SHA to build. Exactly one a of branch name, tag, or commit SHA must be provided. */
	// +optional
	CommitSha *string `json:"commitSha,omitempty"`

	/* Directory, relative to the source root, in which to run the build.
	This must be a relative path. If a step's dir is specified and is an absolute path,
	this value is ignored for that step's execution. */
	// +optional
	Dir *string `json:"dir,omitempty"`

	/* Only trigger a build if the revision regex does NOT match the revision regex. */
	// +optional
	InvertRegex *bool `json:"invertRegex,omitempty"`

	/* ID of the project that owns the Cloud Source Repository.
	If omitted, the project ID requesting the build is assumed. */
	// +optional
	ProjectId *string `json:"projectId,omitempty"`

	/* The desired Cloud Source Repository. If omitted, "default" is
	assumed. */
	RepoRef v1alpha1.ResourceRef `json:"repoRef"`

	/* Substitutions to use in a triggered build. Should only be used with triggers.run */
	// +optional
	Substitutions map[string]string `json:"substitutions,omitempty"`

	/* Regex matching tags to build. Exactly one a of branch name, tag, or commit SHA must be provided.
	The syntax of the regular expressions accepted is the syntax accepted by RE2 and
	described at https://github.com/google/re2/wiki/Syntax */
	// +optional
	TagName *string `json:"tagName,omitempty"`
}

type TriggerSecret struct {
	/* KMS crypto key to use to decrypt these envs. */
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef"`

	/* Map of environment variable name to its encrypted value.
	Secret environment variables must be unique across all of a build's secrets,
	and must be used by at least one build step. Values can be at most 64 KB in size.
	There can be at most 100 secret values across all of a build's secrets. */
	// +optional
	SecretEnv map[string]string `json:"secretEnv,omitempty"`
}

type TriggerSource struct {
	/* Location of the source in a Google Cloud Source Repository. */
	// +optional
	RepoSource *TriggerRepoSource `json:"repoSource,omitempty"`

	/* Location of the source in an archive file in Google Cloud Storage. */
	// +optional
	StorageSource *TriggerStorageSource `json:"storageSource,omitempty"`
}

type TriggerStep struct {
	/* A list of arguments that will be presented to the step when it is started.

	If the image used to run the step's container has an entrypoint, the args
	are used as arguments to that entrypoint. If the image does not define an
	entrypoint, the first element in args is used as the entrypoint, and the
	remainder will be used as arguments. */
	// +optional
	Args []string `json:"args,omitempty"`

	/* Working directory to use when running this step's container.

	If this value is a relative path, it is relative to the build's working
	directory. If this value is absolute, it may be outside the build's working
	directory, in which case the contents of the path may not be persisted
	across build step executions, unless a 'volume' for that path is specified.

	If the build specifies a 'RepoSource' with 'dir' and a step with a
	'dir',
	which specifies an absolute path, the 'RepoSource' 'dir' is ignored
	for the step's execution. */
	// +optional
	Dir *string `json:"dir,omitempty"`

	/* Entrypoint to be used instead of the build step image's
	default entrypoint.
	If unset, the image's default entrypoint is used */
	// +optional
	Entrypoint *string `json:"entrypoint,omitempty"`

	/* A list of environment variable definitions to be used when
	running a step.

	The elements are of the form "KEY=VALUE" for the environment variable
	"KEY" being given the value "VALUE". */
	// +optional
	Env []string `json:"env,omitempty"`

	/* Unique identifier for this build step, used in 'wait_for' to
	reference this build step as a dependency. */
	// +optional
	Id *string `json:"id,omitempty"`

	/* The name of the container image that will run this particular build step.

	If the image is available in the host's Docker daemon's cache, it will be
	run directly. If not, the host will attempt to pull the image first, using
	the builder service account's credentials if necessary.

	The Docker daemon's cache will already have the latest versions of all of
	the officially supported build steps (see https://github.com/GoogleCloudPlatform/cloud-builders
	for images and examples).
	The Docker daemon will also have cached many of the layers for some popular
	images, like "ubuntu", "debian", but they will be refreshed at the time
	you attempt to use them.

	If you built an image in a previous build step, it will be stored in the
	host's Docker daemon's cache and is available to use as the name for a
	later build step. */
	Name string `json:"name"`

	/* A list of environment variables which are encrypted using
	a Cloud Key
	Management Service crypto key. These values must be specified in
	the build's 'Secret'. */
	// +optional
	SecretEnv []string `json:"secretEnv,omitempty"`

	/* Time limit for executing this build step. If not defined,
	the step has no
	time limit and will be allowed to continue to run until either it
	completes or the build itself times out. */
	// +optional
	Timeout *string `json:"timeout,omitempty"`

	/* Output only. Stores timing information for executing this
	build step. */
	// +optional
	Timing *string `json:"timing,omitempty"`

	/* List of volumes to mount into the build step.

	Each volume is created as an empty volume prior to execution of the
	build step. Upon completion of the build, volumes and their contents
	are discarded.

	Using a named volume in only one step is not valid as it is
	indicative of a build request with an incorrect configuration. */
	// +optional
	Volumes []TriggerVolumes `json:"volumes,omitempty"`

	/* The ID(s) of the step(s) that this build step depends on.

	This build step will not start until all the build steps in 'wait_for'
	have completed successfully. If 'wait_for' is empty, this build step
	will start when all previous build steps in the 'Build.Steps' list
	have completed successfully. */
	// +optional
	WaitFor []string `json:"waitFor,omitempty"`
}

type TriggerStorageSource struct {
	/* Google Cloud Storage bucket containing the source. */
	BucketRef v1alpha1.ResourceRef `json:"bucketRef"`

	/* Google Cloud Storage generation for the object.
	If the generation is omitted, the latest generation will be used */
	// +optional
	Generation *string `json:"generation,omitempty"`

	/* Google Cloud Storage object containing the source.
	This object must be a gzipped archive file (.tar.gz) containing source to build. */
	Object string `json:"object"`
}

type TriggerTiming struct {
	/* End of time span.

	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	EndTime *string `json:"endTime,omitempty"`

	/* Start of time span.

	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	StartTime *string `json:"startTime,omitempty"`
}

type TriggerTriggerTemplate struct {
	/* Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
	This field is a regular expression. */
	// +optional
	BranchName *string `json:"branchName,omitempty"`

	/* Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided. */
	// +optional
	CommitSha *string `json:"commitSha,omitempty"`

	/* Directory, relative to the source root, in which to run the build.

	This must be a relative path. If a step's dir is specified and
	is an absolute path, this value is ignored for that step's
	execution. */
	// +optional
	Dir *string `json:"dir,omitempty"`

	/* Only trigger a build if the revision regex does NOT match the revision regex. */
	// +optional
	InvertRegex *bool `json:"invertRegex,omitempty"`

	/* The Cloud Source Repository to build. If omitted, the repo with
	name "default" is assumed. */
	// +optional
	RepoRef *v1alpha1.ResourceRef `json:"repoRef,omitempty"`

	/* Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
	This field is a regular expression. */
	// +optional
	TagName *string `json:"tagName,omitempty"`
}

type TriggerVolumes struct {
	/* Name of the volume to mount.

	Volume names must be unique per build step and must be valid names for
	Docker volumes. Each named volume must be used by at least two build steps. */
	Name string `json:"name"`

	/* Path at which to mount the volume.

	Paths must be absolute and cannot conflict with other volume paths on
	the same build step or with certain reserved volume paths. */
	Path string `json:"path"`
}

type TriggerWebhookConfig struct {
	/* The secret required */
	SecretRef v1alpha1.ResourceRef `json:"secretRef"`

	/* Potential issues with the underlying Pub/Sub subscription configuration.
	Only populated on get requests. */
	// +optional
	State *string `json:"state,omitempty"`
}

type CloudBuildTriggerSpec struct {
	/* Contents of the build template. Either a filename or build template must be provided. */
	// +optional
	Build *TriggerBuild `json:"build,omitempty"`

	/* Human-readable description of the trigger. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Whether the trigger is disabled or not. If true, the trigger will never result in a build. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided. */
	// +optional
	Filename *string `json:"filename,omitempty"`

	/* Describes the configuration of a trigger that creates a build whenever a GitHub event is received.

	One of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided. */
	// +optional
	Github *TriggerGithub `json:"github,omitempty"`

	/* ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	extended with support for '**'.

	If ignoredFiles and changed files are both empty, then they are not
	used to determine whether or not to trigger a build.

	If ignoredFiles is not empty, then we ignore any files that match any
	of the ignored_file globs. If the change has no files that are outside
	of the ignoredFiles globs, then we do not trigger a build. */
	// +optional
	IgnoredFiles []string `json:"ignoredFiles,omitempty"`

	/* ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	extended with support for '**'.

	If any of the files altered in the commit pass the ignoredFiles filter
	and includedFiles is empty, then as far as this filter is concerned, we
	should trigger the build.

	If any of the files altered in the commit pass the ignoredFiles filter
	and includedFiles is not empty, then we make sure that at least one of
	those files matches a includedFiles glob. If not, then we do not trigger
	a build. */
	// +optional
	IncludedFiles []string `json:"includedFiles,omitempty"`

	/* PubsubConfig describes the configuration of a trigger that creates
	a build whenever a Pub/Sub message is published.

	One of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided. */
	// +optional
	PubsubConfig *TriggerPubsubConfig `json:"pubsubConfig,omitempty"`

	/* Substitutions data for Build resource. */
	// +optional
	Substitutions map[string]string `json:"substitutions,omitempty"`

	/* Tags for annotation of a BuildTrigger */
	// +optional
	Tags []string `json:"tags,omitempty"`

	/* Template describing the types of source changes to trigger a build.

	Branch and tag names in trigger templates are interpreted as regular
	expressions. Any branch or tag change that matches that regular
	expression will trigger a build.

	One of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided. */
	// +optional
	TriggerTemplate *TriggerTriggerTemplate `json:"triggerTemplate,omitempty"`

	/* WebhookConfig describes the configuration of a trigger that creates
	a build whenever a webhook is sent to a trigger's webhook URL.

	One of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided. */
	// +optional
	WebhookConfig *TriggerWebhookConfig `json:"webhookConfig,omitempty"`
}

type CloudBuildTriggerStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudBuildTrigger's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Time when the trigger was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* The unique identifier for the trigger. */
	TriggerId string `json:"triggerId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudBuildTrigger is the Schema for the cloudbuild API
// +k8s:openapi-gen=true
type CloudBuildTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudBuildTriggerSpec   `json:"spec,omitempty"`
	Status CloudBuildTriggerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudBuildTriggerList contains a list of CloudBuildTrigger
type CloudBuildTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudBuildTrigger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudBuildTrigger{}, &CloudBuildTriggerList{})
}
