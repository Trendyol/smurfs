package plugin

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// Plugin describes a plugin manifest file.
type Plugin struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata"`

	Spec Spec `json:"spec"`
}

type Spec struct {
	Name             string   `json:"name"`
	Version          string   `json:"version"`
	Description      string   `json:"description"`
	ShortDescription string   `json:"shortDescription"`
	Runnable         Runnable `json:"runnable"`
}

type Runnable struct {
	// Archive address of the plugin
	URI string `json:"uri" yaml:"uri"`

	// Sha256 of the plugin to check integrity
	Sha256 string          `json:"sha256" yaml:"sha256"`
	Files  []FileOperation `json:"files"  yaml:"files"`

	// Entrypoint is the path to the binary in the archive
	Entrypoint string `json:"name" yaml:"name"`

	Bin string `json:"bin" yaml:"bin"`
}

// FileOperation todo: do not use it
type FileOperation struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Receipt is a record of a plugin installation.
type Receipt struct {
	Plugin      `yaml:",inline" json:",inline"`
	InstalledAt time.Time `json:"installedAt" yaml:"installedAt"`
}
