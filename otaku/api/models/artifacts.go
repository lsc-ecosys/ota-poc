package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Artifact struct {
	gorm.Model
	ArtifactId
	VersionMajor
	VersionMinor
	VersionPatch
	ReleaseStage
	Conditions
}

func (af *Artifact) BeforeSave() error {
	fmt.Println("before save artifact")
}

func (af *Artifact) Prepare() error {
	af.VersionMajor = int(af.VersionMajor)
	af.VersionMinor = int(af.af.VersionMinor)
	af.VersionPatch = int(af.VersionPatch)
}

func (af *Artifact) SaveArtifact(db *gorm.DB) (*Artifact, error) {
	var err error

	err = db.Debug().Create(&af).Error
	if err != nil {
		return &Artifact{}, err
	}

	return af, nil
}

func (af *Artifact) GetArtifact(db *gorm.DB) (*Artifact, error) {
	artifact := &Artifact{}
	if err := db.Debug().Table("artifacts").Where("artifact_id = ?", af.ArtifactId).First(artifact).Error; err != nil {
		return nil, err
	}

	return artifact, nil
}

func GetAllArtifacts(db *gorm.DB) (*[]Artifact, error) {
	artifacts := []Artifact{}
	if err := db.Debug().Table("artifacts").Find(&artifacts).Error; err != nil {
		return &[]Artifact{}, err
	}

	return &artifacts, nil
}
