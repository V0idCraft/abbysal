package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/V0idCraft/abyssal/internal/models"
	"github.com/andygrunwald/go-jira"
)

var _ models.JobExecutor = (*listIssueExecutor)(nil)

type listIssueExecutor struct {
	jobBaseExecutor
}

func (l *listIssueExecutor) Execute(ctx context.Context) error {

	if l.GetKind() != models.ExecutorKindList {
		if l.next != nil {
			return l.next.Execute(ctx)
		}
		return nil
	}

	l.logger.Info("Executing list issue job")
	metadata, ok := l.Metadata.(models.ListIssueMetadata)

	if !ok {
		return fmt.Errorf("metadata is not of type ListIssueMetadata")
	}

	issues, _, err := l.client.Issue.Search(metadata.Jql, nil)

	if err != nil {
		l.logger.Error("Error while searching for issues", slog.Any("error", err))
		return err
	}

	issueKeys := make([]string, len(issues))

	for index, issue := range issues {
		issueKeys[index] = issue.Key
	}

	l.logger.Info("Issues found", slog.Any("issues", issueKeys))

	data := &models.ListIssueData{
		Issues: issueKeys,
	}

	newCtx := context.WithValue(ctx, models.CtxDataKeyListIssueData, data)
	l.logger.Info("List issue job executed successfully")
	if l.next != nil {
		return l.next.Execute(newCtx)
	}
	return nil

}

func (l *listIssueExecutor) GetKind() models.ExecutorKind {
	return models.ExecutorKindList
}

func NewListIssueExecutor(client *jira.Client, logger *slog.Logger) *listIssueExecutor {
	return &listIssueExecutor{
		jobBaseExecutor: jobBaseExecutor{
			client: client,
			logger: logger,
		},
	}
}
