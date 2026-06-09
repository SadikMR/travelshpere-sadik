package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
			"flags": map[string]any{"png": "https://example.com/bd.png", "svg": "https://example.com/bd.svg"},
			"name": map[string]any{"common": "Bangladesh"},
			"currencies": map[string]any{"BDT": map[string]any{"name": "Taka"}},
			"languages": map[string]any{"ben": "Bengali"},
			"latlng": []float64{23.8, 90.4},
			"capital": []string{"Dhaka"},
			"region": "Asia",
			"subregion": "Southern Asia",
			"population": 170000000,
		},
		{
			"flags": map[string]any{"png": "https://example.com/fr.png", "svg": "https://example.com/fr.svg"},
			"name": map[string]any{"common": "France"},
			"currencies": map[string]any{"EUR": map[string]any{"name": "Euro"}},
			"languages": map[string]any{"fra": "French"},
			"latlng": []float64{46.0, 2.0},
			"capital": []string{"Paris"},
			"region": "Europe",
			"subregion": "Western Europe",
			"population": 67000000,
		},
	}

	attractions := map[string]any{
		"type": "FeatureCollection",
		"features": []map[string]any{
			{
				"geometry": map[string]any{"coordinates": []float64{90.4, 23.8}},
				"properties": map[string]any{
					"xid": "Q123",
					"name": "Lalbagh Fort",
					"kinds": "historic,monuments",
					"rate": 3,
					"dist": 123.0,
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

	w2 := doRequest(t, http.MethodPost, "/api/wishlist", `{"country_name":"Bangladesh","note":"my note"}` , cookie)
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
