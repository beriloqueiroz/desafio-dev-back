package interfaces

type MessageGateway interface {
	MessageByLocation(city string, state string) (*string, error)
}
