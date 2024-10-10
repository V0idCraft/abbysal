package models

import "github.com/V0idCraft/abyssal/internal/chain"

type Pipeline struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      string           `json:"status"`
	jobs        []chain.Executor `json:"-"`
}

func (p *Pipeline) Add(job chain.Executor) {
	p.jobs = append(p.jobs, job)
}

func (p *Pipeline) GetJobs() []chain.Executor {
	return p.jobs
}
