package server

import (
	"encoding/json"
	"net/http"

	"github.com/AustinMCrane/resume/pkg/model"
)

type Resume struct {
	ResumeStore model.ResumeStore
}

func (ctx *Resume) HandleResume(w http.ResponseWriter, r *http.Request) {
	resume, err := ctx.ResumeStore.GetResume()
	if err != nil {
		HttpJSONError(w, err)
		return
	}

	data, _ := json.MarshalIndent(resume, "", "    ")
	HttpJSONResponse(http.StatusOK, w, data)
	return
}

func GetResumeServer(resumeStore model.ResumeStore) Resume {
	return Resume{
		ResumeStore: resumeStore,
	}
}
