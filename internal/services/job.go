package services

import (
	"context"
	"log/slog"

	"github.com/V0idCraft/abyssal/internal/chain"
	"github.com/V0idCraft/abyssal/internal/models"
	"github.com/andygrunwald/go-jira"
)

var _ chain.Executor = (*jobBaseExecutor)(nil)

// jobBaseExecutor is the base struct for all jobs, it implements the JobExecutor interface.
type jobBaseExecutor struct {
	models.Job

	next chain.Executor

	client *jira.Client
	logger *slog.Logger
}

func (j *jobBaseExecutor) Execute(ctx context.Context) error {
	err := j.next.Execute(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (j *jobBaseExecutor) SetNext(next chain.Executor) {
	j.next = next
}
