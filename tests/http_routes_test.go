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
			"flag":       map[string]any{"emoji": "🇧🇩", "url_png": "https://example.com/bd.png", "url_svg": "https://example.com/bd.svg"},
			"names":      map[string]any{"common": "Bangladesh", "official": "People's Republic of Bangladesh"},
			"currencies": []map[string]any{{"code": "BDT", "name": "Taka", "symbol": "৳"}},
			"languages":  []map[string]any{{"bcp47": "bn", "name": "Bengali", "native_name": "বাংলা"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": 23.8, "lng": 90.4}},
			"capitals":   []map[string]any{{"name": "Dhaka", "primary": true, "coordinates": map[string]any{"lat": 23.7104, "lng": 90.4074}}},
			"region":     "Asia",
			"subregion":  "Southern Asia",
			"population": 170000000,
		},
		{
			"flag":       map[string]any{"emoji": "🇫🇷", "url_png": "https://example.com/fr.png", "url_svg": "https://example.com/fr.svg"},
			"names":      map[string]any{"common": "France", "official": "French Republic"},
			"currencies": []map[string]any{{"code": "EUR", "name": "Euro", "symbol": "€"}},
			"languages":  []map[string]any{{"bcp47": "fr", "name": "French", "native_name": "Français"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": 46.0, "lng": 2.0}},
			"capitals":   []map[string]any{{"name": "Paris", "primary": true, "coordinates": map[string]any{"lat": 48.8566, "lng": 2.3522}}},
			"region":     "Europe",
			"subregion":  "Western Europe",
			"population": 67000000,
		},
		{
			"flag":       map[string]any{"emoji": "🇩🇪", "url_png": "https://example.com/de.png", "url_svg": "https://example.com/de.svg"},
			"names":      map[string]any{"common": "Germany", "official": "Federal Republic of Germany"},
			"currencies": []map[string]any{{"code": "EUR", "name": "Euro", "symbol": "€"}},
			"languages":  []map[string]any{{"bcp47": "de", "name": "German", "native_name": "Deutsch"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": 51.0, "lng": 10.0}},
			"capitals":   []map[string]any{{"name": "Berlin", "primary": true, "coordinates": map[string]any{"lat": 52.52, "lng": 13.405}}},
			"region":     "Europe",
			"subregion":  "Western Europe",
			"population": 83000000,
		},
		{
			"flag":       map[string]any{"emoji": "🇦🇺", "url_png": "https://example.com/au.png", "url_svg": "https://example.com/au.svg"},
			"names":      map[string]any{"common": "Australia", "official": "Commonwealth of Australia"},
			"currencies": []map[string]any{{"code": "AUD", "name": "Australian Dollar", "symbol": "$"}},
			"languages":  []map[string]any{{"bcp47": "en", "name": "English", "native_name": "English"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": -25.0, "lng": 133.0}},
			"capitals":   []map[string]any{{"name": "Canberra", "primary": true, "coordinates": map[string]any{"lat": -35.2809, "lng": 149.1300}}},
			"region":     "Oceania",
			"subregion":  "Australia and New Zealand",
			"population": 25000000,
		},
		{
			"flag":       map[string]any{"emoji": "🇯🇵", "url_png": "https://example.com/jp.png", "url_svg": "https://example.com/jp.svg"},
			"names":      map[string]any{"common": "Japan", "official": "Japan"},
			"currencies": []map[string]any{{"code": "JPY", "name": "Yen", "symbol": "¥"}},
			"languages":  []map[string]any{{"bcp47": "ja", "name": "Japanese", "native_name": "日本語"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": 36.0, "lng": 138.0}},
			"capitals":   []map[string]any{{"name": "Tokyo", "primary": true, "coordinates": map[string]any{"lat": 35.6762, "lng": 139.6503}}},
			"region":     "Asia",
			"subregion":  "Eastern Asia",
			"population": 125000000,
		},
		{
			"flag":       map[string]any{"emoji": "🇨🇦", "url_png": "https://example.com/ca.png", "url_svg": "https://example.com/ca.svg"},
			"names":      map[string]any{"common": "Canada", "official": "Canada"},
			"currencies": []map[string]any{{"code": "CAD", "name": "Canadian Dollar", "symbol": "$"}},
			"languages":  []map[string]any{{"bcp47": "en", "name": "English", "native_name": "English"}, {"bcp47": "fr", "name": "French", "native_name": "Français"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": 56.0, "lng": -106.0}},
			"capitals":   []map[string]any{{"name": "Ottawa", "primary": true, "coordinates": map[string]any{"lat": 45.4215, "lng": -75.6972}}},
			"region":     "Americas",
			"subregion":  "North America",
			"population": 38000000,
		},
		{
			"flag":       map[string]any{"emoji": "🇧🇷", "url_png": "https://example.com/br.png", "url_svg": "https://example.com/br.svg"},
			"names":      map[string]any{"common": "Brazil", "official": "Federative Republic of Brazil"},
			"currencies": []map[string]any{{"code": "BRL", "name": "Real", "symbol": "R$"}},
			"languages":  []map[string]any{{"bcp47": "pt", "name": "Portuguese", "native_name": "Português"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": -10.0, "lng": -55.0}},
			"capitals":   []map[string]any{{"name": "Brasília", "primary": true, "coordinates": map[string]any{"lat": -15.7975, "lng": -47.8919}}},
			"region":     "Americas",
			"subregion":  "South America",
			"population": 211000000,
		},
		{
			"flag":       map[string]any{"emoji": "🇮🇳", "url_png": "https://example.com/in.png", "url_svg": "https://example.com/in.svg"},
			"names":      map[string]any{"common": "India", "official": "Republic of India"},
			"currencies": []map[string]any{{"code": "INR", "name": "Rupee", "symbol": "₹"}},
			"languages":  []map[string]any{{"bcp47": "hi", "name": "Hindi", "native_name": "हिन्दी"}, {"bcp47": "en", "name": "English", "native_name": "English"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": 20.0, "lng": 77.0}},
			"capitals":   []map[string]any{{"name": "New Delhi", "primary": true, "coordinates": map[string]any{"lat": 28.6139, "lng": 77.2090}}},
			"region":     "Asia",
			"subregion":  "Southern Asia",
			"population": 1380000000,
		},
		{
			"flag":       map[string]any{"emoji": "🇿🇦", "url_png": "https://example.com/za.png", "url_svg": "https://example.com/za.svg"},
			"names":      map[string]any{"common": "South Africa", "official": "Republic of South Africa"},
			"currencies": []map[string]any{{"code": "ZAR", "name": "Rand", "symbol": "R"}},
			"languages":  []map[string]any{{"bcp47": "en", "name": "English", "native_name": "English"}, {"bcp47": "af", "name": "Afrikaans", "native_name": "Afrikaans"}},
			"geography":  map[string]any{"coordinates": map[string]any{"lat": -30.0, "lng": 25.0}},
			"capitals":   []map[string]any{{"name": "Pretoria", "primary": true, "coordinates": map[string]any{"lat": -25.7479, "lng": 28.2293}}},
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
	// v5 API: client hits baseURL directly with ?limit=&offset= params
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if this is a /name path (for GetCountryByName)
		if strings.HasPrefix(r.URL.Path, "/name") {
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{"objects": countries},
			}))
			return
		}
		// Check if this is a /radius path (for attractions)
		if strings.HasPrefix(r.URL.Path, "/radius") {
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(attractions))
			return
		}
		// Default: paginated countries endpoint
		// Use offset param to determine pagination — offset 0 returns data, offset > 0 returns empty
		w.Header().Set("Content-Type", "application/json")
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		if offset > 0 {
			// Return empty to stop pagination loop
			require.NoError(t, json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{"objects": []any{}},
			}))
			return
		}
		require.NoError(t, json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{"objects": countries},
		}))
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

func TestAuthLoginFormValidation(t *testing.T) {
	form := strings.NewReader("username=")
	req := httptest.NewRequest(http.MethodPost, "/login", form)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Username is required")
	assert.Contains(t, w.Body.String(), "Login")
}

func TestSSRUnknownCountryDetailsReturnsNotFound(t *testing.T) {
	newMockRemoteServer(t)

	w := doRequest(t, http.MethodGet, "/countries/nosuchcountry", "", nil)
	assert.Equal(t, http.StatusNotFound, w.Code)
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

	w := doRequest(t, http.MethodPost, "/api/wishlist", `{"country_name":"Bangladesh","note":"original note"}`, cookie)
	assert.Equal(t, http.StatusCreated, w.Code)

	var created map[string]any
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &created))
	id := int(created["id"].(float64))

	w2 := doRequest(t, http.MethodPut, "/api/wishlist/bad-id", `{"note":"updated note","status":"Visited"}`, cookie)
	assert.Equal(t, http.StatusBadRequest, w2.Code)

	w3 := doRequest(t, http.MethodPut, "/api/wishlist/"+strconv.Itoa(id), `{"note":"updated note","status":"Visited"}`, cookie)
	assert.Equal(t, http.StatusOK, w3.Code)

	var updated map[string]any
	require.NoError(t, json.Unmarshal(w3.Body.Bytes(), &updated))
	assert.Equal(t, "updated note", updated["note"])
	assert.Equal(t, "Visited", updated["status"])

	w4 := doRequest(t, http.MethodDelete, "/api/wishlist/"+strconv.Itoa(id), "", cookie)
	assert.Equal(t, http.StatusNoContent, w4.Code)
}

func TestWishlistAPIInvalidJsonPayload(t *testing.T) {
	cookie := newAuthenticatedCookie(t, "sadik-2")

	w := doRequest(t, http.MethodPut, "/api/wishlist/1", `{bad json`, cookie)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid payload")
}

func TestWishlistAPIDeleteInvalidID(t *testing.T) {
	cookie := newAuthenticatedCookie(t, "sadik-3")

	w := doRequest(t, http.MethodDelete, "/api/wishlist/bad-id", "", cookie)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid id")
}

func TestAttractionAPIInvalidParams(t *testing.T) {
	w := doRequest(t, http.MethodGet, "/api/attractions?lat=bad&lon=bad", "", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "lat and lon query parameters are required")
}
