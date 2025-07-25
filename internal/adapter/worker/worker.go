package worker

import (
	"github.com/hibiken/asynq"
	"github.com/knands42/lorecrafter/internal/adapter/worker/background_jobs"
	"github.com/knands42/lorecrafter/internal/config"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"log"
)

type Worker struct {
	cfg       config.Config
	repo      sqlc.Querier
	scheduler *asynq.Scheduler
	server    *asynq.Server
	mux       *asynq.ServeMux
}

func NewWorker(cfg config.Config, repo sqlc.Querier, redisOpt asynq.RedisClientOpt) *Worker {
	scheduler := asynq.NewScheduler(redisOpt, &asynq.SchedulerOpts{})
	server := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
	})
	mux := asynq.NewServeMux()

	return &Worker{
		cfg:       cfg,
		repo:      repo,
		scheduler: scheduler,
		server:    server,
		mux:       mux,
	}
}

func (w *Worker) RegisterBackgroundWorkers() {
	err := background_jobs.RegisterCheckExpiredCampaignInvitationsTask(w.scheduler)
	if err != nil {
		log.Fatalf("Failed to register CheckExpiredCampaignInvitationsTask with err: %v", err)
	}
	w.mux.HandleFunc(
		background_jobs.TypeCheckExpiredCampaignInvitations,
		background_jobs.HandleCheckExpiredCampaignInvitationsTaskWrapper(w.repo),
	)

	err = w.server.Start(w.mux)
	if err != nil {
		log.Fatalf("Worker failed to start: %v", err)
	}
}
