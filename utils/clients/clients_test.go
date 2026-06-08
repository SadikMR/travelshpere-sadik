package clients

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCountriesSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]interface{}{
			{"name": map[string]string{"common": "Bangladesh"}, "region": "Asia"},
		})
	}))
	defer server.Close()

	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	countries, err := GetCountries()

	require.NoError(t, err)
	assert.Len(t, countries, 1)
	assert.Equal(t, "Bangladesh", countries[0].Name.Common)
}

func TestGetCountriesHTTPError(t *testing.T) {
	beego.AppConfig.Set("restcountriesBaseURL", "http://localhost:1")

	_, err := GetCountries()

	assert.Error(t, err)
}

func TestGetCountriesInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	}))
	defer server.Close()

	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	_, err := GetCountries()

	assert.Error(t, err)
}

func TestGetCountryByNameSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]interface{}{
			{"name": map[string]string{"common": "Japan"}, "region": "Asia"},
		})
	}))
	defer server.Close()

	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	countries, err := GetCountryByName("Japan")

	require.NoError(t, err)
	assert.Len(t, countries, 1)
}

func TestGetCountryByNameHTTPError(t *testing.T) {
	beego.AppConfig.Set("restcountriesBaseURL", "http://localhost:1")

	_, err := GetCountryByName("Japan")

	assert.Error(t, err)
}

func TestGetCountryByNameInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bad"))
	}))
	defer server.Close()

	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	_, err := GetCountryByName("Japan")

	assert.Error(t, err)
}

func TestGetAttractionsSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"type":     "FeatureCollection",
			"features": []interface{}{},
		})
	}))
	defer server.Close()

	beego.AppConfig.Set("opentripBaseURL", server.URL)
	beego.AppConfig.Set("opentripApi", "testkey123")

	result, err := GetAttractions(23.8, 90.4)

	require.NoError(t, err)
	assert.Equal(t, "FeatureCollection", result.Type)
}

func TestGetAttractionsMissingAPIKey(t *testing.T) {
	beego.AppConfig.Set("opentripApi", "")

	_, err := GetAttractions(23.8, 90.4)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "api key missing")
}

func TestGetAttractionsHTTPError(t *testing.T) {
	beego.AppConfig.Set("opentripApi", "testkey")
	beego.AppConfig.Set("opentripBaseURL", "http://localhost:1")

	_, err := GetAttractions(23.8, 90.4)

	assert.Error(t, err)
}

func TestGetAttractionsNon200(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	beego.AppConfig.Set("opentripApi", "testkey")
	beego.AppConfig.Set("opentripBaseURL", server.URL)

	_, err := GetAttractions(23.8, 90.4)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "500")
}

func TestGetAttractionsInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	}))
	defer server.Close()

	beego.AppConfig.Set("opentripApi", "testkey")
	beego.AppConfig.Set("opentripBaseURL", server.URL)

	_, err := GetAttractions(23.8, 90.4)

	assert.Error(t, err)
}
