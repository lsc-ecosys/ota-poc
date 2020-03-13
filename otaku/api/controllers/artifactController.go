package controllers

import (
	"net/http"

	"lsc-ecosys/ota-poc/otaku/api/models"
	"lsc-ecosys/ota-poc/otaku/api/responses"
	"lsc-ecosys/ota-poc/otaku/utils"
)

type releaseStage struct {
	Index int8
	Bound bool
}

// Get all artifacts directly
func (a *App) GetAllArtifacts(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "data": nil}

	afs, err := models.GetAllArtifacts(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	resp["data"] = afs
	responses.JSON(w, http.StatusOK, resp)

	return
}

func DeriveUpdatableArtifacts(version string, a *App) (*models.Rollout, error) {
	var err error
	var vArray = []int8{}
	var artifacts = &[]models.Artifact{}

	vArray, err = utils.ConvertVersionToInt8Array(version)

	if err != nil {
		return &models.Rollout{}, err
	}

	for len(vArray) < 3 {
		vArray = append(vArray, 0)
	}
	artifacts, err = models.GetAllArtifacts(a.DB)
	if err != nil {
		return &models.Rollout{}, err
	}

	// two types of release
	var releaseTag = map[string]releaseStage{
		"STABLE": {-1, false},
		"BETA":   {-1, false},
	}

	var count = 0
	var lengthOfAfs = len(*artifacts)

	for i := 0; i < lengthOfAfs; i++ {
		var release = (*artifacts)[i].ReleaseStage
		if releaseTag[release].Bound == true {
			continue
		}
	}

}

//
func (a *App) CheckUpdateByVersion(w http.ResponseWriter, r *http.Request) {
	//var resp = map[string]interface{}{"status": "success", "data": nil}
	return
}
