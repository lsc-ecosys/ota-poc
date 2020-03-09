package ota

import (
	"context"
)

// Condition defines a structure for flexible control of artifact management.
// e.g. limitation
type Condition map[string]interface{}

// Artifact structure
// use semver (semantic version) to manage version control of OTA
type Artifact struct {
	ArtifactID   string
	ReleaseStage string
	VersionMajor int8
	VersionMinor int8
	VersionPatch int8
	Conditions   Condition
}

// ArtifactRepository represents methods for artifact persistence
type ArtifactRepository interface {
	Save(context.Context, Artifact) (Artifact, error)

	Update(context.Context, Artifact) error
}

// ArtifactCache owns cached content for artifact
type ArtifactCache interface{}
