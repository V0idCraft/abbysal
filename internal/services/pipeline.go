package services

import (
	"context"
	"log/slog"

	"github.com/V0idCraft/abyssal/internal/models"
)

type PipelineService struct {
	logger *slog.Logger
}

func NewPipelineService(s *slog.Logger) *PipelineService {
	return &PipelineService{
		logger: s,
	}
}

func (s *PipelineService) Run(ctx context.Context, pipeline *models.Pipeline) error {

	if len(pipeline.GetJobs()) == 0 {
		s.logger.Warn("No jobs found in pipeline, skipping execution")
		return nil
	}
	jobs := pipeline.GetJobs()
	pipeline.Status = "running"

	for i := 0; i < len(jobs)-1; i++ {
		jobs[i].SetNext(jobs[i+1])
	}
	s.logger.Info("Executing Pipeline", slog.String("ID", pipeline.ID), slog.String("Name", pipeline.Title))
	err := jobs[0].Execute(ctx)

	if err != nil {
		pipeline.Status = "failed"
		s.logger.Error("Error while executing jobs", slog.Any("error", err))
		return err
	}

	pipeline.Status = "success"
	s.logger.Info("Pipeline executed successfully")
	return nil
}
