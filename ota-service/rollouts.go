package ota

// RolloutCondition defines extra flexible condition in key-value structure
type RolloutCondition map[string]interface{}

// Rollout structure
type Rollout struct {
	ArtifactVersion string
	DownloadURL     string
	Available       bool
	Quota           int
	Reserved        int
	Upgraded        int
	Conditions      RolloutCondition
}

type PatchRepository interface{}
