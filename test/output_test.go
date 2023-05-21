package test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func TestMock(t *testing.T) {
	mockServerURL := os.Getenv("MOCK_SERVER_URL")
	client := http.Client{}
	t.Run("healthcheck", func(t *testing.T) {
		req, err := http.NewRequest("GET", mockServerURL+"/healthcheck", nil)
		require.NoError(t, err)
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer func() {
			require.NoError(t, resp.Body.Close())
		}()
		assert.Equal(t, 200, resp.StatusCode)
		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, "OK 200", string(b))
	})

	t.Run("json", func(t *testing.T) {
		req, err := http.NewRequest("GET", mockServerURL+"/api/v1/json", nil)
		require.NoError(t, err)
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer func() {
			require.NoError(t, resp.Body.Close())
		}()
		assert.Equal(t, 200, resp.StatusCode)
		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%s\n", `{"key":"value"}`), string(b))
	})
}
