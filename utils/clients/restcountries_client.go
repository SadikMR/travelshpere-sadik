package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/SadikMR/travelshpere-sadik/dto"
)

// GetCountries retrieves all countries using the v5 API and auto-paginates in the background.
func GetCountries() ([]dto.RestCountryDTO, error) {
	baseURL, _ := beego.AppConfig.String("restcountriesBaseURL")
	apiKey, _ := beego.AppConfig.String("restcountriesAPIKey")

	var allCountries []dto.RestCountryDTO
	limit := 20 // v5 maximum limit
	offset := 0
	client := &http.Client{}

	for {
		// Build the pagination URL targeting the root endpoint
		apiURL := fmt.Sprintf(
			"%s?limit=%d&offset=%d&response_fields=names,capitals,currencies,languages,flag,geography,region,subregion,population",
			baseURL,
			limit,
			offset,
		)

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Add("Authorization", "Bearer "+apiKey)
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("request failed: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("API returned status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		}

		// 1. Corrected JSON:API wrapper: The array is nested inside data.objects
		var response struct {
			Data struct {
				Objects []dto.RestCountryDTO `json:"objects"`
			} `json:"data"`
		}

		// 2. Decode the response
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to decode response page: %w", err)
		}
		resp.Body.Close()

		// 3. If the Objects array is empty, we have fetched all countries. Break the loop.
		if len(response.Data.Objects) == 0 {
			break
		}

		// 4. Append to our master list
		allCountries = append(allCountries, response.Data.Objects...)
		offset += limit
	}

	return allCountries, nil
}

// GetCountryByName retrieves a country by name using the v5 API.
func GetCountryByName(name string) ([]dto.RestCountryDTO, error) {
	baseURL, _ := beego.AppConfig.String("restcountriesBaseURL")
	apiKey, _ := beego.AppConfig.String("restcountriesAPIKey")

	apiURL := fmt.Sprintf(
		"%s/name?q=%s&response_fields=names,capitals,currencies,languages,flag,geography,region,subregion,population",
		baseURL,
		url.QueryEscape(name),
	)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// 1. Corrected JSON:API wrapper here as well
	var response struct {
		Data struct {
			Objects []dto.RestCountryDTO `json:"objects"`
		} `json:"data"`
	}

	// 2. Decode the response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// 3. Return the unwrapped Objects array
	return response.Data.Objects, nil
}
