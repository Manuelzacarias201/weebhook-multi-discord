package application

import "fmt"

func GenerateMessageToDiscordForActions(action, workflow, conclusion, repoFullName, url string) string {
	var mensaje string

	if action == "" {
		action = "Desconocido"
	}
	if workflow == "" {
		workflow = "Desconocido"
	}
	if conclusion == "" {
		conclusion = "Desconocido"
	}

	switch action {
	case "completed":
		if conclusion == "success" {
			mensaje = "Workflow completado con éxito"
		} else if conclusion == "failure" {
			mensaje = "Workflow fallido"
		} else {
			mensaje = fmt.Sprintf("Workflow con estado: %s", conclusion)
		}
	case "requested":
		mensaje = "Workflow iniciado"
	default:
		mensaje = fmt.Sprintf("Evento de workflow: %s", action)
	}


	return fmt.Sprintf("%s en el repositorio %s\nWorkflow: %s\nDetalles: %s\n\n---\n", 
		mensaje, repoFullName, workflow, url)
}
