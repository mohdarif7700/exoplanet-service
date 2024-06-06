package dto

import "errors"

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`       // light years
	Radius      float64       `json:"radius"`         // Earth-radius units
	Mass        float64       `json:"mass,omitempty"` // Earth-mass units, only for Terrestrial
	Type        ExoplanetType `json:"type"`
}

var (
	ErrExoplanetNotFound = errors.New("exoplanet not found")
)
