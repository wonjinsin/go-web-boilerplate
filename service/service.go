package service

import (
	"context"
	"log"
	"os"
	"pikachu/repository"
	"pikachu/util"

	"pikachu/model"
)

var zlog *util.Logger

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[service] err[%s]", err.Error())
		os.Exit(1)
	}
}

// Service ...
type Service struct {
	User UserService
}

// Init ...
func Init(redis *repository.RedisRepository) (*Service, error) {
	userSvc := NewUserService(redis.User)
	return &Service{User: userSvc}, nil
}

// UserService ...
type UserService interface {
	NewUser(ctx context.Context, user *model.User) (ruser *model.User, err error)
	GetUser(ctx context.Context, id string) (ruser *model.User, err error)
	GetUserByEmail(ctx context.Context, email string) (ruser *model.User, err error)
	UpdateUser(ctx context.Context, uid string, user *model.User) (ruser *model.User, err error)
	DeleteUser(ctx context.Context, id string) (err error)
}
