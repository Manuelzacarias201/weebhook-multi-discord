package domain

type DiscordChannels struct {
	DevelopmentChannel string
	TestingChannel     string
	GeneralChannel     string
}

func NewDiscordChannels(dev, test, general string) *DiscordChannels {
	return &DiscordChannels{
		DevelopmentChannel: dev,
		TestingChannel:     test,
		GeneralChannel:     general,
	}
}
