package worker

import (
	"github.com/knands42/lorecrafter/internal/config"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

type Worker struct {
	cfg  config.Config
	repo sqlc.Querier
}

func NewWorker(cfg config.Config, repo sqlc.Querier) *Worker {

	return &Worker{
		cfg:  cfg,
		repo: repo,
	}
}

func (w *Worker) StartWorkers() {

}
