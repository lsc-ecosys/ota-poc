package controllers

import (
	"net/http"

	"github.com/lsc-ecosys/ota-poc/otaku/api/models"
	"github.com/lsc-ecosys/ota-poc/otaku/api/responses"
)

// Get all artifacts directly
func (a *App) GetAllArtifacts(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "data": nil}

	afModel := models.Artifact{}
	afs, err := afModel.GetAllArtifacts(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	resp["data"] = afs
	responses.JSON(w, http.StatusOK, resp)

	return
}
