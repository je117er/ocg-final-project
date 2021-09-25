package scheduler

import (
	"context"
	"database/sql"
	customerRepository "github.com/je117er/ocg-final-project/internal/customer/repository"
	"github.com/je117er/ocg-final-project/internal/rabbitmq"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	db       *sql.DB
	cron     *cron.Cron
	ctx      context.Context
	producer *rabbitmq.Producer
}

var logger = utils.SugarLog()

func NewScheduler(ctx context.Context, db *sql.DB, producer *rabbitmq.Producer) *Scheduler {
	return &Scheduler{
		db:       db,
		cron:     cron.New(cron.WithSeconds()),
		ctx:      ctx,
		producer: producer,
	}
}

func (scheduler *Scheduler) Start() {
	logger.Info("Start scheduler...")
	scheduler.cron.AddFunc(" 0 * * * *", scheduler.schedulerJob)
	scheduler.producer.Start()
	scheduler.cron.Start()
}

func (scheduler *Scheduler) schedulerJob() {
	logger.Info("Scan to email....")
	customerRepo := customerRepository.NewCustomerRepository(scheduler.db)
	emails, _ := customerRepo.GetUnSendEmails(scheduler.ctx, 5, 10)
	for _, v := range emails {
		scheduler.producer.Public(v)
	}
}
