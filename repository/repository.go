package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"pikachu/config"
	"pikachu/model"
	"pikachu/util"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var zlog *util.Logger
var redisPrefix string

type dbLogger struct {
	*util.Logger
}

func (dl *dbLogger) LogMode(l logger.LogLevel) logger.Interface {
	return dl
}

func (dl *dbLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	dl.Logger.With(ctx).Info(msg, data)
}

func (dl *dbLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	dl.Logger.With(ctx).Warn(msg, data)
}

func (dl *dbLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	dl.Logger.With(ctx).Error(msg, data)
}

func (dl *dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		dl.Logger.With(ctx).Infow(err.Error(), "elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6), "rows", rows, "sql", sql)
	} else {
		dl.Logger.With(ctx).Infow("", "elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6), "rows", rows, "sql", sql)
	}
}

func init() {
	var err error
	zlog, err = util.NewLogger()
	if err != nil {
		log.Fatalf("InitLog module[service] err[%s]", err.Error())
		os.Exit(1)
	}
}

// Repository ...
type Repository struct {
	User UserRepository
}

// RedisRepository ...
type RedisRepository struct {
	User UserRepository
}

// Init ...
func Init(pikachu *config.ViperConfig) (*Repository, *RedisRepository, error) {
	mysqlConn, err := mysqlConnect(pikachu)
	if err != nil {
		return nil, nil, err
	}

	redisPrefix = pikachu.GetString("projectName")
	redisConn, err := redisConnect(pikachu)
	if err != nil {
		return nil, nil, err
	}

	userRepo := NewGormUserRepository(mysqlConn)

	redisUserRepo := NewRedisUserRepository(redisConn, userRepo)

	return &Repository{User: userRepo}, &RedisRepository{User: redisUserRepo}, nil
}

func mysqlConnect(pikachu *config.ViperConfig) (mysql *gorm.DB, err error) {
	mysql, err = gorm.Open(getDialector(pikachu), &gorm.Config{})

	return mysql, err
}

func getDialector(pikachu *config.ViperConfig) gorm.Dialector {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=UTC",
		pikachu.GetString("database.username"),
		pikachu.GetString("database.password"),
		pikachu.GetString("database.host"),
		pikachu.GetInt("database.port"),
		pikachu.GetString("database.dbname"),
	)

	return mysql.Open(dbURI)
}

func getConfig(pikachu *config.ViperConfig) (gConfig *gorm.Config) {
	dbLogger := &dbLogger{zlog}
	gConfig = &gorm.Config{
		Logger:                                   dbLogger,
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	return gConfig
}

// redisConnect ...
func redisConnect(pikachu *config.ViperConfig) (redisDB *redis.Client, err error) {
	host := fmt.Sprintf("%s:%d", pikachu.GetString("redis.host"), pikachu.GetInt("redis.port"))
	zlog.Infow("InitRedis", "redis_host", host)
	redisDB = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
	})
	if _, err := redisDB.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return redisDB, nil
}

// UserRepository ...
type UserRepository interface {
	NewUser(ctx context.Context, user *model.User) (ruser *model.User, err error)
	GetUser(ctx context.Context, uid string) (ruser *model.User, err error)
	GetUserByEmail(ctx context.Context, email string) (ruser *model.User, err error)
	UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error)
	DeleteUser(ctx context.Context, uid string) (err error)
}
