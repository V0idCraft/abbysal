package models

type Pipeline struct {
	ID          string
	Title       string
	Description string
	Status      string
	jobs        []JobExecutor
}

func (p *Pipeline) Add(job JobExecutor) {
	job.SetPipelineID(p.ID)
	p.jobs = append(p.jobs, job)
}

func (p *Pipeline) GetJobs() []JobExecutor {
	return p.jobs
}
