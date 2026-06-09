package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/SadikMR/travelshpere-sadik/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func newMockRemoteServer(t *testing.T) *httptest.Server {
	countries := []map[string]any{
		{
			"flags":      map[string]any{"png": "https://example.com/bd.png", "svg": "https://example.com/bd.svg"},
			"name":       map[string]any{"common": "Bangladesh"},
			"currencies": map[string]any{"BDT": map[string]any{"name": "Taka"}},
			"languages":  map[string]any{"ben": "Bengali"},
			"latlng":     []float64{23.8, 90.4},
			"capital":    []string{"Dhaka"},
			"region":     "Asia",
			"subregion":  "Southern Asia",
			"population": 170000000,
		},
		{
			"flags":      map[string]any{"png": "https://example.com/fr.png", "svg": "https://example.com/fr.svg"},
			"name":       map[string]any{"common": "France"},
			"currencies": map[string]any{"EUR": map[string]any{"name": "Euro"}},
			"languages":  map[string]any{"fra": "French"},
			"latlng":     []float64{46.0, 2.0},
			"capital":    []string{"Paris"},
			"region":     "Europe",
			"subregion":  "Western Europe",
			"population": 67000000,
		},
		{
			"flags":      map[string]any{"png": "https://example.com/de.png", "svg": "https://example.com/de.svg"},
			"name":       map[string]any{"common": "Germany"},
			"currencies": map[string]any{"EUR": map[string]any{"name": "Euro"}},
			"languages":  map[string]any{"deu": "German"},
			"latlng":     []float64{51.0, 10.0},
			"capital":    []string{"Berlin"},
			"region":     "Europe",
			"subregion":  "Western Europe",
			"population": 83000000,
		},
		{
			"flags":      map[string]any{"png": "https://example.com/au.png", "svg": "https://example.com/au.svg"},
			"name":       map[string]any{"common": "Australia"},
			"currencies": map[string]any{"AUD": map[string]any{"name": "Australian Dollar"}},
			"languages":  map[string]any{"eng": "English"},
			"latlng":     []float64{-25.0, 133.0},
			"capital":    []string{"Canberra"},
			"region":     "Oceania",
			"subregion":  "Australia and New Zealand",
			"population": 25000000,
		},
		{
			"flags":      map[string]any{"png": "https://example.com/jp.png", "svg": "https://example.com/jp.svg"},
			"name":       map[string]any{"common": "Japan"},
			"currencies": map[string]any{"JPY": map[string]any{"name": "Yen"}},
			"languages":  map[string]any{"jpn": "Japanese"},
			"latlng":     []float64{36.0, 138.0},
			"capital":    []string{"Tokyo"},
			"region":     "Asia",
			"subregion":  "Eastern Asia",
			"population": 125000000,
		},
		{
			"flags":      map[string]any{"png": "https://example.com/ca.png", "svg": "https://example.com/ca.svg"},
			"name":       map[string]any{"common": "Canada"},
			"currencies": map[string]any{"CAD": map[string]any{"name": "Canadian Dollar"}},
			"languages":  map[string]any{"eng": "English", "fra": "French"},
			"latlng":     []float64{56.0, -106.0},
			"capital":    []string{"Ottawa"},
			"region":     "Americas",
			"subregion":  "North America",
			"population": 38000000,
		},
		{
			"flags":      map[string]any{"png": "https://example.com/br.png", "svg": "https://example.com/br.svg"},
			"name":       map[string]any{"common": "Brazil"},
			"currencies": map[string]any{"BRL": map[string]any{"name": "Real"}},
			"languages":  map[string]any{"por": "Portuguese"},
			"latlng":     []float64{-10.0, -55.0},
			"capital":    []string{"Brasília"},
			"region":     "Americas",
			"subregion":  "South America",
			"population": 211000000,
		},
		{
			"flags":      map[string]any{"png": "https://example.com/in.png", "svg": "https://example.com/in.svg"},
			"name":       map[string]any{"common": "India"},
			"currencies": map[string]any{"INR": map[string]any{"name": "Rupee"}},
			"languages":  map[string]any{"hin": "Hindi", "eng": "English"},
			"latlng":     []float64{20.0, 77.0},
			"capital":    []string{"New Delhi"},
			"region":     "Asia",
			"subregion":  "Southern Asia",
			"population": 1380000000,
		},
		{
			"flags":      map[string]any{"png": "https://example.com/za.png", "svg": "https://example.com/za.svg"},
			"name":       map[string]any{"common": "South Africa"},
			"currencies": map[string]any{"ZAR": map[string]any{"name": "Rand"}},
			"languages":  map[string]any{"eng": "English", "afr": "Afrikaans"},
			"latlng":     []float64{-30.0, 25.0},
			"capital":    []string{"Pretoria"},
			"region":     "Africa",
			"subregion":  "Southern Africa",
			"population": 59000000,
		},
	}

	attractions := map[string]any{
		"type": "FeatureCollection",
		"features": []map[string]any{
			{
				"geometry": map[string]any{"coordinates": []float64{90.4, 23.8}},
				"properties": map[string]any{
					"xid":   "Q123",
					"name":  "Lalbagh Fort",
					"kinds": "historic,monuments",
					"rate":  3,
					"dist":  123.0,
				},
			},
			{
				"geometry": map[string]any{"coordinates": []float64{2.3, 48.8}},
				"properties": map[string]any{
					"xid":   "Q456",
					"name":  "Eiffel Tower",
					"kinds": "historic,monuments",
					"rate":  5,
					"dist":  45.0,
				},
			},
			{
				"geometry": map[string]any{"coordinates": []float64{-33.9, 151.2}},
				"properties": map[string]any{
					"xid":   "Q789",
					"name":  "Sydney Opera House",
					"kinds": "entertainment,architecture",
					"rate":  4,
					"dist":  70.0,
				},
			},
			{
				"geometry": map[string]any{"coordinates": []float64{55.7, 37.6}},
				"properties": map[string]any{
					"xid":   "Q101",
					"name":  "Red Square",
					"kinds": "historic,monuments",
					"rate":  4,
					"dist":  80.0,
				},
			},
			{
				"geometry": map[string]any{"coordinates": []float64{40.7, -74.0}},
				"properties": map[string]any{
					"xid":   "Q102",
					"name":  "Statue of Liberty",
					"kinds": "historic,monuments",
					"rate":  5,
					"dist":  12.0,
				},
			},
			{
				"geometry": map[string]any{"coordinates": []float64{35.7, 139.7}},
				"properties": map[string]any{
					"xid":   "Q103",
					"name":  "Tokyo Tower",
					"kinds": "historic,architecture",
					"rate":  4,
					"dist":  25.0,
				},
			},
		},
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(countries))
	})
	handler.HandleFunc("/radius", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(attractions))
	})

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	beego.AppConfig.Set("restcountriesBaseURL", server.URL)
	beego.AppConfig.Set("opentripBaseURL", server.URL)
	beego.AppConfig.Set("opentripApi", "testkey")

	return server
}

func newAuthenticatedCookie(t *testing.T, username string) *http.Cookie {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	sess, err := beego.GlobalSessions.SessionStart(w, req)
	require.NoError(t, err)
	require.NoError(t, sess.Set(context.Background(), "username", username))

	cookies := w.Result().Cookies()
	require.NotEmpty(t, cookies)

	return cookies[0]
}

func doRequest(t *testing.T, method, path string, body string, cookie *http.Cookie) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if cookie != nil {
		req.AddCookie(cookie)
	}

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	return w
}

func TestSSRRoutes(t *testing.T) {
	newMockRemoteServer(t)

	cases := []struct {
		name          string
		path          string
		wantStatus    int
		wantSubstring string
	}{
		{"Home page", "/", http.StatusOK, "TravelSphere"},
		{"Login page", "/login", http.StatusOK, "Login"},
		{"Countries page", "/countries", http.StatusOK, "Countries"},
		{"Country details", "/countries/bangladesh", http.StatusOK, "Bangladesh"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := doRequest(t, http.MethodGet, tc.path, "", nil)
			assert.Equal(t, tc.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tc.wantSubstring)
		})
	}
}

func TestSSRInvalidCountrySearch(t *testing.T) {
	w := doRequest(t, http.MethodGet, "/countries?region=Invalid", "", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid region")
}

func TestAPIRoutes(t *testing.T) {
	newMockRemoteServer(t)

	cases := []struct {
		name       string
		path       string
		wantStatus int
		assertBody func(t *testing.T, body string)
	}{
		{"List countries", "/api/countries", http.StatusOK, func(t *testing.T, body string) {
			assert.Contains(t, body, "Bangladesh")
		}},
		{"Search countries", "/api/countries/search?q=bang", http.StatusOK, func(t *testing.T, body string) {
			assert.Contains(t, body, "Bangladesh")
			assert.NotContains(t, body, "France")
		}},
		{"List attractions", "/api/attractions?lat=23.8&lon=90.4", http.StatusOK, func(t *testing.T, body string) {
			assert.Contains(t, body, "Lalbagh Fort")
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := doRequest(t, http.MethodGet, tc.path, "", nil)
			assert.Equal(t, tc.wantStatus, w.Code)
			assert.NotEmpty(t, w.Body.String())
			tc.assertBody(t, w.Body.String())
		})
	}
}

func TestWishlistFilterRedirect(t *testing.T) {
	w := doRequest(t, http.MethodGet, "/wishlist", "", nil)
	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/login", w.Header().Get("Location"))
}

func TestWishlistAPIWithSession(t *testing.T) {
	cookie := newAuthenticatedCookie(t, "sadik")

	w2 := doRequest(t, http.MethodPost, "/api/wishlist", `{"country_name":"Bangladesh","note":"my note"}`, cookie)
	assert.Equal(t, http.StatusCreated, w2.Code)

	var created map[string]any
	require.NoError(t, json.Unmarshal(w2.Body.Bytes(), &created))
	assert.Equal(t, "Bangladesh", created["country_name"])
	assert.Equal(t, "my note", created["note"])

	w3 := doRequest(t, http.MethodGet, "/api/wishlist", "", cookie)
	assert.Equal(t, http.StatusOK, w3.Code)

	var list []map[string]any
	require.NoError(t, json.Unmarshal(w3.Body.Bytes(), &list))
	assert.Len(t, list, 1)
	assert.Equal(t, "Bangladesh", list[0]["country_name"])
}
func TestAPIAuthRequiredReturnsUnauthorized(t *testing.T) {
	w := doRequest(t, http.MethodGet, "/api/wishlist", "", nil)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "login required")
}

func TestAuthLoginLogout(t *testing.T) {
	form := strings.NewReader("username=sadik")
	req := httptest.NewRequest(http.MethodPost, "/login", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, req)
	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/", w.Header().Get("Location"))
	require.NotEmpty(t, w.Result().Cookies())

	cookie := w.Result().Cookies()[0]
	w2 := doRequest(t, http.MethodGet, "/logout", "", cookie)
	assert.Equal(t, http.StatusFound, w2.Code)
	assert.Equal(t, "/login", w2.Header().Get("Location"))
}

func TestProtectedPagesWithAuth(t *testing.T) {
	newMockRemoteServer(t)
	cookie := newAuthenticatedCookie(t, "sadik")

	paths := []string{"/dashboard", "/wishlist", "/wishlist/rows"}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			w := doRequest(t, http.MethodGet, path, "", cookie)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.NotEmpty(t, w.Body.String())
		})
	}

	w := doRequest(t, http.MethodGet, "/api/dashboard/summary", "", cookie)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "totalWishlist")
}

func TestCountryAPIValidationAndDetail(t *testing.T) {
	newMockRemoteServer(t)

	w1 := doRequest(t, http.MethodGet, "/api/countries?region=Invalid", "", nil)
	assert.Equal(t, http.StatusBadRequest, w1.Code)
	assert.Contains(t, w1.Body.String(), "invalid region")

	w2 := doRequest(t, http.MethodGet, "/api/countries/bangladesh", "", nil)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.Contains(t, w2.Body.String(), "Bangladesh")

	w3 := doRequest(t, http.MethodGet, "/api/countries/nosuchcountry", "", nil)
	assert.Equal(t, http.StatusNotFound, w3.Code)
	assert.Contains(t, w3.Body.String(), "country not found")
}

func TestWishlistAPIUpdateDelete(t *testing.T) {
	cookie := newAuthenticatedCookie(t, "sadik")

	w := doRequest(t, http.MethodPost, "/api/wishlist", `{"country_name":"Bangladesh","note":"original note"}` , cookie)
	assert.Equal(t, http.StatusCreated, w.Code)

	var created map[string]any
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &created))
	id := int(created["id"].(float64))

	w2 := doRequest(t, http.MethodPut, "/api/wishlist/bad-id", `{"note":"updated note","status":"Visited"}` , cookie)
	assert.Equal(t, http.StatusBadRequest, w2.Code)

	w3 := doRequest(t, http.MethodPut, "/api/wishlist/"+strconv.Itoa(id), `{"note":"updated note","status":"Visited"}` , cookie)
	assert.Equal(t, http.StatusOK, w3.Code)

	var updated map[string]any
	require.NoError(t, json.Unmarshal(w3.Body.Bytes(), &updated))
	assert.Equal(t, "updated note", updated["note"])
	assert.Equal(t, "Visited", updated["status"])

	w4 := doRequest(t, http.MethodDelete, "/api/wishlist/"+strconv.Itoa(id), "", cookie)
	assert.Equal(t, http.StatusNoContent, w4.Code)
}

func TestAttractionAPIInvalidParams(t *testing.T) {
	w := doRequest(t, http.MethodGet, "/api/attractions?lat=bad&lon=bad", "", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "lat and lon query parameters are required")
}
