package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/V0idCraft/abyssal/internal/models"
	"github.com/andygrunwald/go-jira"
)

var _ models.JobExecutor = (*workLogIssueJob)(nil)

type workLogIssueJob struct {
	jobBaseExecutor
}

func (w *workLogIssueJob) Execute(ctx context.Context) error {

	if w.kind == models.ExecutorKindWorkLog {
		issues, ok := ctx.Value(models.CtxDataKeyListIssueData).(*models.ListIssueData)

		if !ok {
			return fmt.Errorf("listIssueData not found in context")
		}

		issueKey := issues.Issues[0]

		metadata := w.Metadata.(models.WorkLogIssueMetadata)

		worklog := jira.WorklogRecord{
			Comment:   "Worklog from the abyssal CLI",
			TimeSpent: metadata.TimeSpent,
		}
		_, _, err := w.client.Issue.AddWorklogRecord(issueKey, &worklog)
		if err != nil {
			return err
		}
		fmt.Printf("Worklog added to issue %s\n", issueKey)
	}

	if w.next != nil {
		return w.next.Execute(ctx)
	}

	return nil

}
func (l *workLogIssueJob) GetKind() models.ExecutorKind {
	return models.ExecutorKindWorkLog
}

func NewWorkLogIssueExecutor(client *jira.Client, logger *slog.Logger) *workLogIssueJob {
	return &workLogIssueJob{
		jobBaseExecutor: jobBaseExecutor{
			client: client,
			logger: logger,
		},
	}
}
