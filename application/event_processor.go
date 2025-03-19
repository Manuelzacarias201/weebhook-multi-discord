package application

import (
	"fmt"
	"time"

	"pull_request_webhook/domain"
	"pull_request_webhook/ports"
)

type EventProcessor struct {
	notificationService ports.NotificationService
}

func NewEventProcessor(notificationService ports.NotificationService) *EventProcessor {
	return &EventProcessor{
		notificationService: notificationService,
	}
}

func (p *EventProcessor) ProcessPullRequestEvent(event domain.GitHubEvent) {
	var message string
	switch event.Action {
	case "opened", "reopened", "ready_for_review":
		message = fmt.Sprintf("📝 **Nuevo Pull Request**\n"+
			"**Título:** %s\n"+
			"**Autor:** @%s\n"+
			"**Enlace:** %s\n"+
			"**Fecha:** %s",
			event.PullRequest.Title,
			event.PullRequest.User.Login,
			event.PullRequest.HTMLURL,
			time.Now().Format("2006-01-02 15:04:05"))
		p.notificationService.SendDevelopmentNotification(message)
	case "closed":
		if event.PullRequest.Merged {
			message = fmt.Sprintf("✨ **Pull Request Fusionado**\n"+
				"**Título:** %s\n"+
				"**Enlace:** %s\n"+
				"**Fecha:** %s",
				event.PullRequest.Title,
				event.PullRequest.HTMLURL,
				time.Now().Format("2006-01-02 15:04:05"))
			p.notificationService.SendDevelopmentNotification(message)
		}
	}
}

func (p *EventProcessor) ProcessWorkflowEvent(event domain.GitHubEvent) {
	if event.WorkflowRun.Status == "completed" {
		status := "✅"
		if event.WorkflowRun.Conclusion != "success" {
			status = "❌"
		}
		message := fmt.Sprintf("%s **Resultado del Workflow**\n"+
			"**Nombre:** %s\n"+
			"**Estado:** %s\n"+
			"**Enlace:** %s\n"+
			"**Fecha:** %s",
			status,
			event.WorkflowRun.Name,
			event.WorkflowRun.Conclusion,
			event.WorkflowRun.HTMLURL,
			time.Now().Format("2006-01-02 15:04:05"))
		p.notificationService.SendTestingNotification(message)
	}
}
