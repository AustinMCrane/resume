package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AustinMCrane/resume/pkg/mocks"
	"github.com/AustinMCrane/resume/pkg/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestHandleResume_JSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResumeStore := mocks.NewMockResumeStore(ctrl)
	resumeServer := GetResumeServer(mockResumeStore)
	mockResumeStore.EXPECT().GetResume().Return(&model.Resume{}, nil).Times(1)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	resumeServer.HandleResume(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, "application/json", rec.Header().Get("Content-Type"))
}

func TestHandleResume_JSONWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResumeStore := mocks.NewMockResumeStore(ctrl)
	resumeServer := GetResumeServer(mockResumeStore)
	mockResumeStore.EXPECT().GetResume().Return(nil, errors.New("error")).Times(1)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	resumeServer.HandleResume(rec, req)
	require.Equal(t, http.StatusInternalServerError, rec.Code)
	require.Equal(t, "application/json", rec.Header().Get("Content-Type"))
}
