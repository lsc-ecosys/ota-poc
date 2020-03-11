package controllers

import (
	"net/http"

	"lsc-ecosys/ota-poc/otaku/api/models"
	"lsc-ecosys/ota-poc/otaku/api/responses"
)

// Get all artifacts directly
func (a *App) GetAllArtifacts(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "data": nil}

	afModel := models.Artifact{}
	afs, err := afModel.GetAllArtifacts(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["data"] = afs
	response.JSON(w, http.StatusOK, resp)

	return
}
