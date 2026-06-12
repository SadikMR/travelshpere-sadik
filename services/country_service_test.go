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

// mockCountryServer returns a test server that serves country JSON in v5 API format.
func mockCountryServer() *httptest.Server {
	callCount := 0
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		callCount++
		// First call returns data, second call returns empty to stop pagination
		if callCount > 1 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"data": map[string]interface{}{
					"objects": []interface{}{},
				},
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"objects": []map[string]interface{}{
					{
						"names":      map[string]string{"common": "Bangladesh", "official": "People's Republic of Bangladesh"},
						"region":     "Asia",
						"subregion":  "Southern Asia",
						"population": 170000000,
						"capitals":   []map[string]interface{}{{"name": "Dhaka", "primary": true, "coordinates": map[string]float64{"lat": 23.7104, "lng": 90.4074}}},
						"geography":  map[string]interface{}{"coordinates": map[string]float64{"lat": 23.685, "lng": 90.356}},
						"flag":       map[string]string{"emoji": "🇧🇩", "url_png": "https://example.com/bd.png", "url_svg": ""},
						"currencies": []map[string]string{{"code": "BDT", "name": "Bangladeshi taka", "symbol": "৳"}},
						"languages":  []map[string]string{{"bcp47": "bn", "name": "Bengali", "native_name": "বাংলা"}},
					},
					{
						"names":      map[string]string{"common": "Japan", "official": "Japan"},
						"region":     "Asia",
						"subregion":  "Eastern Asia",
						"population": 125000000,
						"capitals":   []map[string]interface{}{{"name": "Tokyo", "primary": true, "coordinates": map[string]float64{"lat": 35.6762, "lng": 139.6503}}},
						"geography":  map[string]interface{}{"coordinates": map[string]float64{"lat": 36.0, "lng": 138.0}},
						"flag":       map[string]string{"emoji": "🇯🇵", "url_png": "https://example.com/jp.png", "url_svg": ""},
						"currencies": []map[string]string{{"code": "JPY", "name": "Japanese yen", "symbol": "¥"}},
						"languages":  []map[string]string{{"bcp47": "ja", "name": "Japanese", "native_name": "日本語"}},
					},
					{
						"names":      map[string]string{"common": "France", "official": "French Republic"},
						"region":     "Europe",
						"subregion":  "Western Europe",
						"population": 67000000,
						"capitals":   []map[string]interface{}{{"name": "Paris", "primary": true, "coordinates": map[string]float64{"lat": 48.8566, "lng": 2.3522}}},
						"geography":  map[string]interface{}{"coordinates": map[string]float64{"lat": 46.0, "lng": 2.0}},
						"flag":       map[string]string{"emoji": "🇫🇷", "url_png": "https://example.com/fr.png", "url_svg": ""},
						"currencies": []map[string]string{{"code": "EUR", "name": "Euro", "symbol": "€"}},
						"languages":  []map[string]string{{"bcp47": "fr", "name": "French", "native_name": "Français"}},
					},
				},
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
