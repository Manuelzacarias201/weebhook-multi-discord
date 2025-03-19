package ports

// NotificationService define la interfaz para enviar notificaciones
type NotificationService interface {
	SendDevelopmentNotification(message string) error
	SendTestingNotification(message string) error
	SendGeneralNotification(message string) error
}
