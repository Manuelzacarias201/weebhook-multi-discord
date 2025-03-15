package application

import "fmt"
//pull-02
func GenerateMessageToDiscord(action, base, titulo, repoFullName, user, urlPullRequest string, merged bool) string {
	var mensaje string

	switch action {
	case "opened":
		mensaje = "Nuevo pull request abierto"
	case "closed":
		if merged {
			mensaje = "Pull request fusionado (merged)"
		} else {
			mensaje = "Pull request cerrado"
		}
	case "reopened":
		mensaje = "Pull request reabierto"
	case "synchronize":
		mensaje = "Pull request actualizado"
	default:
		mensaje = fmt.Sprintf("Acción de Pull Request: %s", action)
	}

	return fmt.Sprintf("%s a la rama %s en el repositorio %s\nTítulo: %s\nAutor: %s\nDetalles: %s",
		mensaje, base, repoFullName, titulo, user, urlPullRequest)
}
