package models

type CtxDataKey string
type ExecutorKind string

const (
	CtxDataKeyListIssueData = CtxDataKey("listIssueData")
)

const (
	ExecutorUnknown        = ExecutorKind("unknown")
	ExecutorKindList       = ExecutorKind("list")
	ExecutorKindTransition = ExecutorKind("transition")
	ExecutorKindWorkLog    = ExecutorKind("worklog")
)

// Job is the Job interface that all jobs must implement, this is the base interface for all jobs.
type Job interface {
	// Responsibility of chain pattern
	// Execute(ctx context.Context) error
	// SetNext(JobExecutor)

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

var _ Job = (*JobBase)(nil)

type JobBase struct {
	// Status is the current status of the job, it represents the current state of the job.
	Status string `json:"status"`
	// Title is the title of the job, it represents the title of the job.
	Title string `json:"title"`
	// Description is the description of the job, it represents the description of the job.
	Description string `json:"description"`
	// Metadata is the metadata of the job, it represents the metadata of the job.
	Metadata interface{} `json:"metadata"`

	pipelineID string
}

// GetDescription implements Job.
func (j *JobBase) GetDescription() string {
	return j.Description
}

// GetKind implements Job.
func (j *JobBase) GetKind() ExecutorKind {
	return ExecutorUnknown
}

// GetMetadata implements Job.
func (j *JobBase) GetMetadata() interface{} {
	return j.Metadata
}

// GetPipelineID implements Job.
func (j *JobBase) GetPipelineID() string {
	return j.pipelineID
}

// GetTitle implements Job.
func (j *JobBase) GetTitle() string {
	return j.Title
}

// SetDescription implements Job.
func (j *JobBase) SetDescription(description string) {
	j.Description = description
}

// SetMetadata implements Job.
func (j *JobBase) SetMetadata(metadata interface{}) {
	j.Metadata = metadata
}

// SetPipelineID implements Job.
func (j *JobBase) SetPipelineID(pipelineID string) {
	j.pipelineID = pipelineID
}

// SetTitle implements Job.
func (j *JobBase) SetTitle(title string) {
	j.Title = title
}

var _ Job = (*ListJob)(nil)

type ListJob struct {
	JobBase
}

func (l *ListJob) GetKind() ExecutorKind {
	return ExecutorKindList
}

var _ Job = (*TransitionJob)(nil)

type TransitionJob struct {
	JobBase
}

func (t *TransitionJob) GetKind() ExecutorKind {
	return ExecutorKindTransition
}

var _ Job = (*WorkLogJob)(nil)

type WorkLogJob struct {
	JobBase
}

func (w *WorkLogJob) GetKind() ExecutorKind {
	return ExecutorKindWorkLog
}

func NewJob(kind ExecutorKind) Job {
	switch kind {
	case ExecutorKindList:
		return &ListJob{}
	case ExecutorKindTransition:
		return &TransitionJob{}
	case ExecutorKindWorkLog:
		return &WorkLogJob{}
	default:
		return nil
	}
}
