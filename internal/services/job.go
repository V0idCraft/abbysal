package services

import (
	"context"
	"log/slog"

	"github.com/V0idCraft/abyssal/internal/models"
	"github.com/andygrunwald/go-jira"
)

var _ models.JobExecutor = (*jobBaseExecutor)(nil)

// jobBaseExecutor is the base struct for all jobs, it implements the JobExecutor interface.
type jobBaseExecutor struct {
	// Kind is the type of job, it represents the action that the job will perform.
	// Example: "list", "create", "update", "delete"
	kind models.ExecutorKind
	// Status is the current status of the job, it represents the current state of the job.
	Status string `json:"status"`
	// Title is the title of the job, it represents the title of the job.
	Title string `json:"title"`
	// Description is the description of the job, it represents the description of the job.
	Description string `json:"description"`
	// Metadata is the metadata of the job, it represents the metadata of the job.
	Metadata interface{} `json:"metadata"`

	next models.JobExecutor

	pipelineID string

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

func (j *jobBaseExecutor) SetNext(next models.JobExecutor) {
	j.next = next
}

func (j *jobBaseExecutor) SetTitle(title string) {
	j.Title = title
}

func (j *jobBaseExecutor) GetTitle() string {
	return j.Title
}

func (j *jobBaseExecutor) SetDescription(description string) {
	j.Description = description
}

func (j *jobBaseExecutor) GetDescription() string {
	return j.Description
}

func (j *jobBaseExecutor) SetMetadata(metadata interface{}) {
	j.Metadata = metadata
}

func (j *jobBaseExecutor) GetMetadata() interface{} {
	return j.Metadata
}

func (j *jobBaseExecutor) GetPipelineID() string {
	return j.pipelineID
}

func (j *jobBaseExecutor) SetPipelineID(id string) {
	j.pipelineID = id
}

func (j *jobBaseExecutor) GetKind() models.ExecutorKind {
	return j.kind
}
