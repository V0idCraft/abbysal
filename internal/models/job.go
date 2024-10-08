package models

import (
	"context"
)

type CtxDataKey string
type ExecutorKind string

const (
	CtxDataKeyListIssueData = CtxDataKey("listIssueData")
)

const (
	ExecutorKindList       = ExecutorKind("list")
	ExecutorKindTransition = ExecutorKind("transition")
	ExecutorKindWorkLog    = ExecutorKind("worklog")
)

// JobExecutor is the Job interface that all jobs must implement, this is the base interface for all jobs.
type JobExecutor interface {
	// Responsibility of chain pattern
	Execute(ctx context.Context) error
	SetNext(JobExecutor)

	SetTitle(string)
	GetTitle() string
	SetDescription(string)
	GetDescription() string
	SetMetadata(interface{})
	GetMetadata() interface{}
	GetPipelineID() string
	SetPipelineID(string)
	GetKind() ExecutorKind
}

type ListIssueMetadata struct {
	// Jql is the JQL query that will be used to list the issues.
	Jql string `json:"jql"`
}

type ListIssueData struct {
	Issues []string `json:"issues"`
}

type TransitionIssueMetadata struct {
	TransitionTo string
}

type WorkLogIssueMetadata struct {
	TimeSpent string `json:"timeSpent"`
}
