package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Artifact struct {
	ArtifactID   int64 `gorm:"primary_key"`
	VersionMajor int
	VersionMinor int
	VersionPatch int
	ReleaseStage string
	Conditions   postgres.Jsonb
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func (af *Artifact) BeforeSave() error {
	fmt.Println("before save artifact")
	return nil
}

func (af *Artifact) Prepare() error {
	af.VersionMajor = int(af.VersionMajor)
	af.VersionMinor = int(af.VersionMinor)
	af.VersionPatch = int(af.VersionPatch)
	return nil
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
	if err := db.Debug().Table("artifacts").Where("artifact_id = ?", af.ArtifactID).First(artifact).Error; err != nil {
		return nil, err
	}

	return artifact, nil
}

func (af *Artifact) GetAllArtifacts(db *gorm.DB) (*[]Artifact, error) {
	artifacts := []Artifact{}
	if err := db.Debug().Table("artifacts").Find(&artifacts).Error; err != nil {
		return &[]Artifact{}, err
	}

	return &artifacts, nil
}
