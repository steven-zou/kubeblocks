/*
Copyright ApeCloud Inc.

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

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	// which backupPolicy to perform this backup
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern:=`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`
	BackupPolicyName string `json:"backupPolicyName"`

	// Backup Type. full or incremental or snapshot. if unset, default is full.
	// +kubebuilder:validation:Enum={full,incremental,snapshot}
	// +kubebuilder:default=full
	BackupType BackupType `json:"backupType"`

	// if backupType is incremental, parentBackupName is required.
	// +optional
	ParentBackupName string `json:"parentBackupName,omitempty"`

	// TTL is a time.Duration-parsable string describing how long
	// the Backup should be retained for.
	// +optional
	TTL *metav1.Duration `json:"ttl,omitempty"`
}

// BackupPhase The current phase. Valid values are New, InProgress, Completed, Failed.
// +enum
type BackupPhase string

// These are the valid statuses of Backup.
const (
	BackupNew BackupPhase = "New"

	BackupInProgress BackupPhase = "InProgress"

	BackupCompleted BackupPhase = "Completed"

	BackupFailed BackupPhase = "Failed"
)

// BackupType the job type, marked job is full or incremental.
// +enum
type BackupType string

const (
	BackupTypeFull BackupType = "full"

	BackupTypeIncremental BackupType = "incremental"

	BackupTypeSnapshot BackupType = "snapshot"
)

// BackupStatus defines the observed state of Backup
type BackupStatus struct {
	// +optional
	Phase BackupPhase `json:"phase,omitempty"`

	// record parentBackupName if backupType is incremental.
	// +optional
	ParentBackupName string `json:"parentBackupName,omitempty"`

	// The date and time when the Backup is eligible for garbage collection.
	// 'null' means the Backup is NOT be cleaned except delete manual.
	// +optional
	Expiration *metav1.Time `json:"expiration,omitempty"`

	// Date/time when the backup started being processed.
	// +optional
	StartTimestamp *metav1.Time `json:"startTimestamp,omitempty"`

	// Date/time when the backup finished being processed.
	// +optional
	CompletionTimestamp *metav1.Time `json:"completionTimestamp,omitempty"`

	// The duration time of backup execution.
	// When converted to a string, the form is "1h2m0.5s".
	// +optional
	Duration *metav1.Duration `json:"duration,omitempty"`

	// backup total size
	// string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	// +optional
	TotalSize string `json:"totalSize,omitempty"`

	// backup total size
	// string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	// +optional
	UploadTotalSize string `json:"uploadTotalSize,omitempty"`

	// checksum of backup file, generated by md5 or sha1 or sha256
	// +optional
	CheckSum string `json:"checkSum,omitempty"`

	// backup check point, for incremental backup.
	// +optional
	CheckPoint string `json:"CheckPoint,omitempty"`

	// the reason if backup failed.
	// +optional
	FailureReason string `json:"failureReason,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories={kubeblocks},scope=Namespaced
// +kubebuilder:printcolumn:name="STATUS",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="COMPLETIONTIME",type=date,JSONPath=`.status.completionTimestamp`
// +kubebuilder:printcolumn:name="TOTALSIZE",type=string,JSONPath=`.status.totalSize`
// +kubebuilder:printcolumn:name="DURATION",type=string,JSONPath=`.status.duration`

// Backup is the Schema for the backups API (defined by User)
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSpec   `json:"spec,omitempty"`
	Status BackupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BackupList contains a list of Backup
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
