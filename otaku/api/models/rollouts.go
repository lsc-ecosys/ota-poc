package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Rollout struct {
	ArtifactVersion string `gorm:"primary_key;size:255"`
	DownloadURL     string `gorm:"not null;size:255"`
	Available       bool   `gorm:"not null;default false"`
	Quota           int    `gorm:"default 0"`
	Reserved        int    `gorm:"default 0"`
	Upgraded        int    `gorm:"default 0"`
	Conditions      postgres.Jsonb
	Memo            string `gorm:"size:255"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func (ro *Rollout) BeforeSave() error {
	if ro.Quota < 0 || ro.Reserved < 0 || ro.Upgraded < 0 {
		return errors.New("Quota, Reserved or Upgraded should not be negative")
	}
	return nil
}

func (ro *Rollout) Prepare() error {
	ro.UpdatedAt = time.Now()
	if ro.CreatedAt.IsZero() {
		ro.CreatedAt = time.Now()
	}
	return nil
}

func (ro *Rollout) SaveRollout(db *gorm.DB) (*Rollout, error) {
	var err error

	err = db.Debug().Create(&ro).Error
	if err != nil {
		return &Rollout{}, err
	}

	return ro, nil
}

func (ro *Rollout) GetAllRollouts(db *gorm.DB) (*[]Rollout, error) {
	var rollouts = []Rollout{}
	if err := db.Debug().Table("rollouts").Find(&rollouts).Error; err != nil {
		return &[]Rollout{}, err
	}

	return &rollouts, nil
}
