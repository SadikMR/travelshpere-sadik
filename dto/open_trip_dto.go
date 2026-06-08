package dto

type OpenTripMapResponseDTO struct {
	Type     string               `json:"type"`
	Features []OpenTripFeatureDTO `json:"features"`
}

type OpenTripFeatureDTO struct {
	Geometry   OpenTripGeometryDTO `json:"geometry"`
	Properties OpenTripPropertyDTO `json:"properties"`
}

type OpenTripGeometryDTO struct {
	Coordinates []float64 `json:"coordinates"`
}

type OpenTripPropertyDTO struct {
	Xid   string  `json:"xid"`
	Name  string  `json:"name"`
	Kinds string  `json:"kinds"`
	Rate  int     `json:"rate"`
	Dist  float64 `json:"dist"`
}
