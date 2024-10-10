package factories

import (
	"log/slog"

	"github.com/V0idCraft/abyssal/internal/chain"
	"github.com/V0idCraft/abyssal/internal/models"
	"github.com/V0idCraft/abyssal/internal/services"
	"github.com/andygrunwald/go-jira"
)

func NewExecutorFactory(job models.Job, client *jira.Client, logger *slog.Logger) chain.Executor {
	switch job.GetKind() {
	case models.ExecutorKindList:
		return services.NewListIssueExecutor(job, client, logger)
	case models.ExecutorKindWorkLog:
		return services.NewWorkLogIssueExecutor(job, client, logger)
	case models.ExecutorKindTransition:
		return services.NewTransitionIssueJobExecutor(job, client, logger)
	default:
		return nil
	}
}
