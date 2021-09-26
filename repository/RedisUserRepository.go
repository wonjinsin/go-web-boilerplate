package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"pikachu/model"

	"github.com/go-redis/redis/v8"
)

type redisUserRepository struct {
	client   *redis.Client
	userRepo UserRepository
}

// NewRedisUserRepository ...
func NewRedisUserRepository(client *redis.Client, userRepo UserRepository) UserRepository {
	return &redisUserRepository{
		client:   client,
		userRepo: userRepo,
	}
}

// NewUser ...
func (r *redisUserRepository) NewUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	return r.userRepo.NewUser(ctx, user)
}

// GetUser ...
func (r *redisUserRepository) GetUser(ctx context.Context, uid string) (ruser *model.User, err error) {
	zlog.With(ctx).Infow("[New RedisRepository Request]", "uid", uid)
	userJSON, err := r.client.Get(ctx, fmt.Sprintf("%s:users:%s", redisPrefix, uid)).Bytes()
	if err == redis.Nil {
		zlog.With(ctx).Infow("GetUser Not Found", "uid", uid)
		if ruser, err = r.userRepo.GetUser(ctx, uid); err == nil {
			r.newUserToRedis(ctx, ruser)
		}
		return ruser, err
	} else if err != nil {
		zlog.With(ctx).Infow("GetUser Error", "uid", uid, "err", err)
		return nil, err
	}

	if err = json.Unmarshal(userJSON, &ruser); err != nil {
		return nil, err
	}
	return ruser, nil
}

// GetUserByEmail ...
func (r *redisUserRepository) GetUserByEmail(ctx context.Context, email string) (ruser *model.User, err error) {
	return r.userRepo.GetUserByEmail(ctx, email)
}

// UpdateUser ...
func (r *redisUserRepository) UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	return r.userRepo.UpdateUser(ctx, user)
}

// DeleteUser ...
func (r *redisUserRepository) DeleteUser(ctx context.Context, uid string) (err error) {
	return r.userRepo.DeleteUser(ctx, uid)
}

func (r *redisUserRepository) newUserToRedis(ctx context.Context, user *model.User) (err error) {
	zlog.With(ctx).Infow("[New RedisRepository Request]", "user", user)
	userJSON, err := json.Marshal(user)
	if err != nil {
		zlog.With(ctx).Errorw("newUserToRedis Error", "err", err)
		return err
	}

	if err = r.client.Set(ctx, fmt.Sprintf("pikachu:users:%s", user.UID), userJSON, 0).Err(); err != nil {
		zlog.With(ctx).Errorw("newUserToRedis Error", "err", err)
	}
	return err
}
