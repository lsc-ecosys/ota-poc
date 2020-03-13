package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"lsc-ecosys/ota-poc/otaku/api/models"
	"lsc-ecosys/ota-poc/otaku/api/responses"
)

func (a *App) GetAllRollouts(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "data": nil}

	rollouts, err := models.GetAllRollouts(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	resp["data"] = rollouts
	responses.JSON(w, http.StatusOK, resp)

	return
}

func (a *App) CreateRollout(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "data": nil}

	roModel := models.Rollout{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &roModel)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	rollout, err := roModel.SaveRollout(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["data"] = rollout
	responses.JSON(w, http.StatusCreated, resp)
	return

}
