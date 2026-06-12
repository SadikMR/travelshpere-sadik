package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/SadikMR/travelshpere-sadik/dto"
)

func GetAttractions(
	lat float64,
	lon float64,
) (dto.OpenTripMapResponseDTO, error) {

	apiKey, _ := beego.AppConfig.String("opentripApi")

	if apiKey == "" {
		return dto.OpenTripMapResponseDTO{},
			fmt.Errorf("opentrip api key missing")
	}

	opentripBaseURL, _ := beego.AppConfig.String("opentripBaseURL")

	url := fmt.Sprintf(
		"%s/radius?radius=10000&lat=%f&lon=%f&kinds=museums,historic,architecture,monuments,cultural&format=geojson&apikey=%s",
		opentripBaseURL,
		lat,
		lon,
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return dto.OpenTripMapResponseDTO{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dto.OpenTripMapResponseDTO{},
			fmt.Errorf(
				"opentripmap api failed: %d",
				resp.StatusCode,
			)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.OpenTripMapResponseDTO{}, err
	}

	var result dto.OpenTripMapResponseDTO

	if err := json.Unmarshal(body, &result); err != nil {
		return dto.OpenTripMapResponseDTO{}, err
	}

	return result, nil
}
