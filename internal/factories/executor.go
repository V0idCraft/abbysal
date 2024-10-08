package factories

import (
	"log/slog"

	"github.com/V0idCraft/abyssal/internal/models"
	"github.com/V0idCraft/abyssal/internal/services"
	"github.com/andygrunwald/go-jira"
)

func NewExecutorFactory(kind models.ExecutorKind, client *jira.Client, logger *slog.Logger) models.JobExecutor {
	switch kind {
	case models.ExecutorKindList:
		return services.NewListIssueExecutor(client, logger)
	case models.ExecutorKindWorkLog:
		return services.NewWorkLogIssueExecutor(client, logger)
	case models.ExecutorKindTransition:
		return services.NewTransitionIssueJobExecutor(client, logger)
	default:
		return nil
	}
}
