package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAttractionsSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"type": "FeatureCollection",
			"features": []map[string]interface{}{
				{
					"geometry": map[string]interface{}{
						"coordinates": []float64{90.388, 23.719},
					},
					"properties": map[string]interface{}{
						"name":  "Lalbagh Fort",
						"kinds": "historic",
						"rate":  7,
						"dist":  1200.5,
					},
				},
			},
		})
	}))
	defer server.Close()

	beego.AppConfig.Set("opentripApi", "testkey")
	beego.AppConfig.Set("opentripBaseURL", server.URL)

	attractions, err := GetAttractions(23.8, 90.4)

	require.NoError(t, err)
	assert.Len(t, attractions, 1)
	assert.Equal(t, "Lalbagh Fort", attractions[0].Name)
}

func TestGetAttractionsAPIError(t *testing.T) {
	beego.AppConfig.Set("opentripApi", "testkey")
	beego.AppConfig.Set("opentripBaseURL", "http://localhost:1")

	_, err := GetAttractions(23.8, 90.4)

	assert.Error(t, err)
}
