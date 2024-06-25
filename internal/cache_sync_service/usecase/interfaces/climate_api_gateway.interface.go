package interfaces

type ClimateAPIGateway interface {
	ClimateAndWaveByLocation(city string, state string) (*string, error)
}
