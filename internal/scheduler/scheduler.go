package scheduler

import (
	"context"
	"database/sql"
	customerRepository "github.com/je117er/ocg-final-project/internal/customer/repository"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	db       *sql.DB
	cron     *cron.Cron
	ctx      context.Context
	dataChan chan<- *models.SentMail
}

var logger = utils.SugarLog()

func NewScheduler(ctx context.Context, db *sql.DB, dataChan chan<- *models.SentMail) *Scheduler {
	return &Scheduler{
		db:       db,
		cron:     cron.New(cron.WithSeconds()),
		ctx:      ctx,
		dataChan: dataChan,
	}
}

func (scheduler *Scheduler) Start() {
	logger.Info("Start scheduler...")
	scheduler.cron.AddFunc("*/10 * * * * *", scheduler.schedulerJob)
	scheduler.cron.Start()
}

func (scheduler *Scheduler) schedulerJob() {
	logger.Info("Scan to email....")
	customerRepo := customerRepository.NewCustomerRepository(scheduler.db)
	emails, _ := customerRepo.GetUnSendEmails(scheduler.ctx, 10)
	for _, v := range emails {
		scheduler.dataChan <- v
	}
}
