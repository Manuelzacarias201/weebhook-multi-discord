package domain

type NotificationType int

const (
	DevNotification NotificationType = iota
	TestNotification
	GeneralNotification
)

type NotificationConfig struct {
	WebhookURL string
	ChannelID  string
	Enabled    bool
}

type NotificationSettings struct {
	Dev     NotificationConfig
	Test    NotificationConfig
	General NotificationConfig
}

func NewNotificationSettings(devURL, testURL, generalURL string) *NotificationSettings {
	return &NotificationSettings{
		Dev: NotificationConfig{
			WebhookURL: devURL,
			Enabled:    true,
		},
		Test: NotificationConfig{
			WebhookURL: testURL,
			Enabled:    true,
		},
		General: NotificationConfig{
			WebhookURL: generalURL,
			Enabled:    true,
		},
	}
}
