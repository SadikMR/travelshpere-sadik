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

// mockCountryServer returns a test server that serves country JSON.
func mockCountryServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]interface{}{
			{
				"name":       map[string]string{"common": "Bangladesh"},
				"region":     "Asia",
				"subregion":  "Southern Asia",
				"population": 170000000,
				"capital":    []string{"Dhaka"},
				"latlng":     []float64{23.685, 90.356},
				"flags":      map[string]string{"png": "https://example.com/bd.png"},
			},
			{
				"name":       map[string]string{"common": "Japan"},
				"region":     "Asia",
				"subregion":  "Eastern Asia",
				"population": 125000000,
				"capital":    []string{"Tokyo"},
				"latlng":     []float64{36.0, 138.0},
				"flags":      map[string]string{"png": "https://example.com/jp.png"},
			},
			{
				"name":       map[string]string{"common": "France"},
				"region":     "Europe",
				"subregion":  "Western Europe",
				"population": 67000000,
				"capital":    []string{"Paris"},
				"latlng":     []float64{46.0, 2.0},
				"flags":      map[string]string{"png": "https://example.com/fr.png"},
			},
		})
	}))
}

func TestGetCountriesNoFilter(t *testing.T) {
	server := mockCountryServer()
	defer server.Close()
	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	svc := CountryService{}
	countries, err := svc.GetCountries("", "")

	require.NoError(t, err)
	assert.Len(t, countries, 3)
}

func TestGetCountriesSearchFilter(t *testing.T) {
	server := mockCountryServer()
	defer server.Close()
	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	svc := CountryService{}
	countries, err := svc.GetCountries("bang", "")

	require.NoError(t, err)
	assert.Len(t, countries, 1)
	assert.Equal(t, "Bangladesh", countries[0].Name)
}

func TestGetCountriesRegionFilter(t *testing.T) {
	server := mockCountryServer()
	defer server.Close()
	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	svc := CountryService{}
	countries, err := svc.GetCountries("", "Europe")

	require.NoError(t, err)
	assert.Len(t, countries, 1)
	assert.Equal(t, "France", countries[0].Name)
}

func TestGetCountriesSorted(t *testing.T) {
	server := mockCountryServer()
	defer server.Close()
	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	svc := CountryService{}
	countries, err := svc.GetCountries("", "")

	require.NoError(t, err)
	assert.Equal(t, "Bangladesh", countries[0].Name)
	assert.Equal(t, "France", countries[1].Name)
	assert.Equal(t, "Japan", countries[2].Name)
}

func TestGetCountriesAPIError(t *testing.T) {
	beego.AppConfig.Set("restcountriesBaseURL", "http://localhost:1")

	svc := CountryService{}
	_, err := svc.GetCountries("", "")

	assert.Error(t, err)
}

func TestGetCountryBySlugFound(t *testing.T) {
	server := mockCountryServer()
	defer server.Close()
	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	svc := CountryService{}
	country, err := svc.GetCountryBySlug("bangladesh")

	require.NoError(t, err)
	assert.Equal(t, "Bangladesh", country.Name)
}

func TestGetCountryBySlugNotFound(t *testing.T) {
	server := mockCountryServer()
	defer server.Close()
	beego.AppConfig.Set("restcountriesBaseURL", server.URL)

	svc := CountryService{}
	_, err := svc.GetCountryBySlug("atlantis")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestGetCountryBySlugAPIError(t *testing.T) {
	beego.AppConfig.Set("restcountriesBaseURL", "http://localhost:1")

	svc := CountryService{}
	_, err := svc.GetCountryBySlug("japan")

	assert.Error(t, err)
}
