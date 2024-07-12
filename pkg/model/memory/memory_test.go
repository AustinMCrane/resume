package memory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResumeStore(t *testing.T) {
	resumeStore := GetResumeStore()
	require.NotNil(t, resumeStore)

	resume, err := resumeStore.GetResume()
	require.NoError(t, err)
	require.NotNil(t, resume)
}
