package config

// Config for the app
var Config = struct {
	AppName string

	PubSub struct {
		EventName struct {
			NewIncomingRequest string `required:"true"`
		}
	}
}{}
