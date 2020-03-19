package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Artifact struct {
	ArtifactID   int64 `gorm:"primary_key"`
	VersionMajor int8
	VersionMinor int8
	VersionPatch int8
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
	af.VersionMajor = int8(af.VersionMajor)
	af.VersionMinor = int8(af.VersionMinor)
	af.VersionPatch = int8(af.VersionPatch)
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

func (af *Artifact) GetArtifactByID(id int, db *gorm.DB) (*Artifact, error) {
	artifact := &Artifact{}
	if err := db.Debug().Table("artifacts").Where("artifact_id = ?", id).First(artifact).Error; err != nil {
		return nil, err
	}

	return artifact, nil
}

func GetAllArtifacts(db *gorm.DB) (*[]Artifact, error) {
	artifacts := []Artifact{}
	if err := db.Debug().Order("version_major DESC, version_minor DESC, version_patch DESC").Table("artifacts").Find(&artifacts).Error; err != nil {
		return &[]Artifact{}, err
	}

	return &artifacts, nil
}
