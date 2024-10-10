package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/V0idCraft/abyssal/internal/config"
	"github.com/V0idCraft/abyssal/internal/factories"
	"github.com/V0idCraft/abyssal/internal/models"
	"github.com/V0idCraft/abyssal/internal/services"
	"github.com/andygrunwald/go-jira"
)

func getFirstWorkableDateOfMonth() time.Time {

	now := time.Now()
	year, month, _ := now.Date()

	firstDayOnMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	if firstDayOnMonth.Weekday() == time.Saturday {
		return firstDayOnMonth.AddDate(0, 0, 2)
	}

	if firstDayOnMonth.Weekday() == time.Sunday {
		return firstDayOnMonth.AddDate(0, 0, 1)
	}

	return firstDayOnMonth

}

func getDateRange() []time.Time {
	firstDay := getFirstWorkableDateOfMonth()
	today := time.Now()
	dates := []time.Time{}
	for firstDay.Before(today) || firstDay.Equal(today) {
		if firstDay.Weekday() == time.Saturday {
			firstDay = firstDay.AddDate(0, 0, 2)
		} else if firstDay.Weekday() == time.Sunday {
			firstDay = firstDay.AddDate(0, 0, 1)
		}
		dates = append(dates, firstDay)
		firstDay = firstDay.AddDate(0, 0, 1)
	}

	return dates
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	cfg, err := config.InitConfig()

	if err != nil {
		panic(err)
	}

	tp := jira.BasicAuthTransport{
		Username: cfg.JiraUsername,
		Password: cfg.JiraToken,
	}

	client, err := jira.NewClient(tp.Client(), cfg.JiraHost)

	if err != nil {
		panic(err)
	}

	pipeline := &models.Pipeline{
		ID:          "1",
		Title:       "Full Hour Report",
		Description: "Transitions issues from To Do to Done adding 1d of worklog",
		Status:      "pending",
	}
	ranges := getDateRange()

	summaries := []string{}
	for _, date := range ranges {
		summaries = append(summaries, fmt.Sprintf("'Hour report Jose Gil Lizardo %s'", date.Format("2006-01-02")))
	}

	summaryQuery := fmt.Sprintf("summary ~ %s", strings.Join(summaries, " or summary ~ "))

	listJob := models.NewJob(models.ExecutorKindList)
	listJob.SetTitle("List issues")
	listJob.SetDescription("List issues from the Jira API")
	listJob.SetMetadata(models.ListIssueMetadata{
		Jql: fmt.Sprintf("project = 'NW' AND assignee = currentUser() and (%s) AND status in ('To Do')", summaryQuery),
	})

	listExecutor := factories.NewExecutorFactory(listJob, client, logger)

	transitionJob := models.NewJob(models.ExecutorKindTransition)
	transitionJob.SetTitle("Transition issues")
	transitionJob.SetDescription("Transition issues from the Jira API")
	transitionJob.SetMetadata(models.TransitionIssueMetadata{
		TransitionTo: "In Progress",
	})

	transitionExecutor := factories.NewExecutorFactory(transitionJob, client, logger)

	workLogJob := models.NewJob(models.ExecutorKindWorkLog)
	workLogJob.SetTitle("WorkLog issues")
	workLogJob.SetDescription("WorkLog issues from the Jira API")
	workLogJob.SetMetadata(models.WorkLogIssueMetadata{
		TimeSpent: "1d",
	})

	workExecutor := factories.NewExecutorFactory(workLogJob, client, logger)

	transitionDoneJob := models.NewJob(models.ExecutorKindTransition)
	transitionDoneJob.SetTitle("Transition issues")
	transitionDoneJob.SetDescription("Transition issues from the Jira API")
	transitionDoneJob.SetMetadata(models.TransitionIssueMetadata{
		TransitionTo: "Done",
	})

	transitionDoneExecutor := factories.NewExecutorFactory(transitionDoneJob, client, logger)

	pipeline.Add(listExecutor)
	pipeline.Add(transitionExecutor)
	pipeline.Add(workExecutor)
	pipeline.Add(transitionDoneExecutor)

	pipelineSvc := services.NewPipelineService(logger)

	mainContext := context.Background()

	err = pipelineSvc.Run(mainContext, pipeline)

	if err != nil {
		panic(err)
	}

}
