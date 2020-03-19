package controllers

import (
	"net/http"
	"net/url"

	"lsc-ecosys/ota-poc/otaku/api/models"
	"lsc-ecosys/ota-poc/otaku/api/responses"
	"lsc-ecosys/ota-poc/otaku/utils"
)

// RELEASESTABLE defines constant value of string of stable release
const RELEASESTABLE string = "STABLE"

// RELEASEBETA defines constant value of string of beta release
const RELEASEBETA string = "BETA"

// ReleaseTag struct is used for marking of the closest version of upgrade
type ReleaseTag struct {
	Index int8
	Bound bool
}

// ReleaseArtifact struct defines organized structure of versions for returing result
type ReleaseArtifact struct {
	Version        string
	AndroidVersion string
	IosVersion     string
}

// GetAllArtifacts returns all artifacts directly
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

// CheckUpdateByVersion should check if need to upgrade and return corrsponding artifact
func (a *App) CheckUpdateByVersion(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "data": nil}
	var err error
	var closestArtifactVersion map[string]*ReleaseArtifact
	var params url.Values
	var version string
	var versionCompared int8
	params, err = url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	version = params.Get("version")
	closestArtifactVersion, err = GetClosestArtifactVersion(version, a)
	versionCompared, err = utils.CompareSemVers(version, closestArtifactVersion[RELEASESTABLE].Version)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// BETA Decision flow
	if beta, ok := closestArtifactVersion[RELEASEBETA]; ok {
		betaCompared, err := utils.CompareSemVers(beta.Version, closestArtifactVersion[RELEASESTABLE].Version)
		versionCompared, err = utils.CompareSemVers(version, closestArtifactVersion[RELEASESTABLE].Version)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		// not valid if beta version is older than stable
		if betaCompared > 0 {
			// TODO: design a mechanism which guarantes only specifc devices authorized
			// to use beta release can consume a quota.
			// need to update
			if versionCompared == 2 {
				return
			}
			if versionCompared < 0 {
				// consume a quota
				return
			}
		}
	}

	// always return stable version

}

// GetClosestArtifactVersion finds out the closest version of artifact for each release stages
func GetClosestArtifactVersion(version string, a *App) (map[string]*ReleaseArtifact, error) {
	var err error
	var vArray = []int8{}
	var artifacts = &[]models.Artifact{}
	var releases = map[string]*ReleaseArtifact{
		RELEASESTABLE: &ReleaseArtifact{Version: "", AndroidVersion: "", IosVersion: ""},
	}
	vArray, err = utils.ConvertVersionToInt8Array(version)

	if err != nil {
		return releases, err
	}

	for len(vArray) < 3 {
		vArray = append(vArray, 0)
	}
	artifacts, err = models.GetAllArtifacts(a.DB)
	if err != nil {
		return releases, err
	}

	// two types of release
	var relTag = map[string]*ReleaseTag{
		RELEASESTABLE: &ReleaseTag{-1, false},
		RELEASEBETA:   &ReleaseTag{-1, false},
	}

	var count = 0
	var lengthOfAfs = len(*artifacts)

	// find out the closest version for upgrade
	for i := 0; i < lengthOfAfs; i++ {
		var release = (*artifacts)[i].ReleaseStage
		//
		if relTag[release].Bound == true {
			continue
		}

		if (*artifacts)[i].VersionMajor < vArray[0] {
			relTag[release].Bound = true
		} else if (*artifacts)[i].VersionMajor == vArray[0] {
			if (*artifacts)[i].VersionMinor < vArray[1] {
				relTag[release].Bound = true
			} else if (*artifacts)[i].VersionMinor == vArray[1] {
				if (*artifacts)[i].VersionPatch <= vArray[2] {
					relTag[release].Bound = true
				}
			}
		}

		if relTag[release].Bound == true {
			// when the first one artifact's version are older than compared version
			if relTag[release].Index == -1 {
				relTag[release].Index = int8(i)
			}
			count++
			// stop when we find 2 qualified versions of artifacts for upgrade
			if count == 2 {
				break
			}
		} else {
			// mark as latest qualified index
			relTag[release].Index = int8(i)
		}
	}

	for rel := range relTag {
		idx := relTag[rel].Index
		if idx != -1 {
			releases[rel].Version = utils.ConvertIntSemVerToString((*artifacts)[idx].VersionMajor, (*artifacts)[idx].VersionMinor, (*artifacts)[idx].VersionPatch)
			// TODO: connect to db schema
			releases[rel].AndroidVersion = ""
			releases[rel].IosVersion = ""
		}
	}

	return releases, nil
}

//func GetPreSignedUrl(rollout *models.Rollout{})
