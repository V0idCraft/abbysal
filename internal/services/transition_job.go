package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/V0idCraft/abyssal/internal/chain"
	"github.com/V0idCraft/abyssal/internal/models"
	"github.com/andygrunwald/go-jira"
)

var _ chain.Executor = (*transitionIssueJobExecutor)(nil)

type transitionIssueJobExecutor struct {
	jobBaseExecutor
}

func (c *transitionIssueJobExecutor) Execute(ctx context.Context) error {
	if c.GetKind() == models.ExecutorKindTransition {

		metadata, ok := c.GetMetadata().(models.TransitionIssueMetadata)

		if !ok {
			return fmt.Errorf("metadata is not of type TransitionIssueMetadata")
		}

		issues, ok := ctx.Value(models.CtxDataKeyListIssueData).(*models.ListIssueData)

		if !ok {
			return fmt.Errorf("listIssueData not found in context")
		}

		firstIssueKey := issues.Issues[0]

		transitions, _, err := c.client.Issue.GetTransitions(firstIssueKey)

		if err != nil {
			return err
		}

		transitionID := ""

		for _, transition := range transitions {
			if transition.To.Name == metadata.TransitionTo {
				transitionID = transition.ID
				break
			}
		}

		if transitionID == "" {
			return fmt.Errorf("transition not found")
		}

		for _, issueKey := range issues.Issues {
			_, err = c.client.Issue.DoTransition(issueKey, transitionID)

			if err != nil {
				return err
			}

			fmt.Printf("Issue %s transitioned to %s\n", issueKey, metadata.TransitionTo)
		}

	}

	if c.next != nil {
		return c.next.Execute(ctx)
	}

	return nil

}

func NewTransitionIssueJobExecutor(job models.Job, client *jira.Client, logger *slog.Logger) *transitionIssueJobExecutor {
	return &transitionIssueJobExecutor{
		jobBaseExecutor: jobBaseExecutor{
			client: client,
			logger: logger,
			Job:    job,
		},
	}
}
