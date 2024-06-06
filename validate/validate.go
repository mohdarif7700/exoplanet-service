package validate

import (
	"errors"
	"exoplanet-service/dto"
)

func ValidateExoplanet(exoplanet *dto.Exoplanet) error {
	if exoplanet.Name == "" {
		return errors.New("name is required")
	}
	if exoplanet.Description == "" {
		return errors.New("description is required")
	}
	if exoplanet.Distance < 10 || exoplanet.Distance > 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius < 0.1 || exoplanet.Radius > 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.Type != dto.GasGiant && exoplanet.Type != dto.Terrestrial {
		return errors.New("type must be either 'GasGiant' or 'Terrestrial'")
	}
	if exoplanet.Type == dto.Terrestrial && (exoplanet.Mass < 0.1 || exoplanet.Mass > 10) {
		return errors.New("mass must be between 0.1 and 10 Earth-mass units for Terrestrial planets")
	}
	return nil
}
