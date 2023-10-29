package config

import (
	"context"
	"fmt"
	"log"
	"reza/todolist-api/model"
	repository "reza/todolist-api/repository/gorm"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Setting ISettings
}

func (r *Postgres) Setup(s *Settings) {
	var dbConn *gorm.DB
	dbConn = NewPostgresConnection(s.ctx, &s.Config.Postgres)
	if s.PostgresSQLProvider == nil {
		repositoryPostgres := repository.NewRepository(dbConn)
		s.PostgresSQLProvider = repositoryPostgres
	} else {
		s.PostgresSQLProvider.Conn = dbConn
	}

	r.Setting.Setup(s)
}

// NewPostgresConnection ...
func NewPostgresConnection(ctx context.Context, s *model.PostgreSQLConfig) *gorm.DB {
	var err error
	connection := fmt.Sprintf("host= %s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", s.Host, s.Username, s.Password, s.Database, s.Port)
	log.Println(connection)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connection,
	}), &gorm.Config{})

	if err != nil {
		err = errors.Wrap(err, "failed connect postgres")
		panic(err)
	}

	return db
}
