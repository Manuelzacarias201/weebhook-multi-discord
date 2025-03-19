package main

import (
	"os"

	"pull_request_webhook/adapters"
	"pull_request_webhook/application"
	"pull_request_webhook/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		panic("Error al cargar el archivo .env")
	}

	// Configurar el puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configuración de notificaciones
	notificationSettings := domain.NewNotificationSettings(
		os.Getenv("DISCORD_DEV_CHANNEL"),
		os.Getenv("DISCORD_TEST_CHANNEL"),
		os.Getenv("DISCORD_GENERAL_CHANNEL"),
	)

	// Inicialización de servicios
	notificationService := adapters.NewDiscordNotifier(notificationSettings)
	eventProcessor := application.NewEventProcessor(notificationService)
	ginHandler := adapters.NewGinHandler(eventProcessor)

	// Inicializar el router
	router := gin.Default()

	// Configurar rutas
	ginHandler.SetupRoutes(router)

	// Iniciar el servidor
	if err := router.Run(":" + port); err != nil {
		panic("Error al iniciar el servidor: " + err.Error())
	}
}
