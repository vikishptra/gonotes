package withgorm

import (
	"context"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"vikishptra/domain_todo-core/model/entity"
	"vikishptra/domain_todo-core/model/errorenum"
	"vikishptra/domain_todo-core/model/vo"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
)

type Gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	Db      *gorm.DB
	Todo    entity.Todo
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *Gateway {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	database := os.Getenv("MYSQLDATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, database)

	Db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = Db.AutoMigrate(entity.Todo{}).Error
	if err != nil {
		panic(err)
	}

	return &Gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		Db:      Db,
	}
}

func (r *Gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")
	if err := r.Db.Save(obj).Error; err != nil {
		panic(err)
	}

	return nil
}

func (r *Gateway) FindOneChecked(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	var todo entity.Todo
	if err := r.Db.First(&todo, "id = ?", todoID); err.RecordNotFound() {
		return nil, errorenum.DataNull
	}

	return &todo, nil
}

func (r *Gateway) GetAllTodoByPagination(ctx context.Context, page int, size int) ([]*entity.Todo, int64, int, error) {
	r.log.Info(ctx, "called")
	var todo []*entity.Todo
	var count int64
	flag := false

	if page == 0 && size == 0 {
		if err := r.Db.
			Model(entity.Todo{}).
			Count(&count).
			Find(&todo); err.RowsAffected == 0 {
			return nil, 0, 0, errorenum.DataNull
		}
		return todo, count, 1, nil
	} else if page < 0 {
		if err := r.Db.
			Model(entity.Todo{}).
			Limit(size).Offset((page - 1) * size).
			Find(&todo); err.RowsAffected == 0 {
			return nil, 0, 0, errorenum.DataNull
		}

		todos := r.Db.Model(entity.Todo{}).Count(&count).Limit(size).Offset((page - 1) * size).Find(&todo)
		if todos.RowsAffected < int64(size) || 0 > int64(size) {
			return todo, todos.RowsAffected, 1, nil
		}
		return todo, int64(size), 1, nil
	} else if size < 0 {
		if err := r.Db.
			Model(entity.Todo{}).
			Count(&count).
			Find(&todo); err.RowsAffected == 0 {
			return nil, 0, 0, errorenum.DataNull
		}
		return todo, count, 1, nil
	} else if err := r.Db.
		Model(entity.Todo{}).
		Count(&count).
		Limit(size).Offset((page - 1) * size).
		Find(&todo); err.RowsAffected == 0 {
		return nil, count, page, errorenum.DataNull
	}

	if int64(size) > count && int64(size) != count {
		flag = true
	}
	if !flag {
		count = int64(size)
	}

	return todo, count, page, nil
}

func (r *Gateway) DeleteOneTodoByID(ctx context.Context, todoID string) error {
	r.log.Info(ctx, "called")
	var Todo entity.Todo
	if err := r.Db.Where("id = ? ", todoID).Delete(Todo); err.RecordNotFound() {
		return errorenum.DataNull
	}

	return nil
}

func (r *Gateway) FindMessageTodoEmpty(ctx context.Context, todo *entity.Todo) error {
	r.log.Info(ctx, "called")
	var todos entity.Todo
	if err := r.Db.Where("title = ?", todo.Title).First(&todos); !err.RecordNotFound() {
		return errorenum.ObjSame
	}
	return nil
}

func (r *Gateway) GetTodoByID(ctx context.Context, todoID vo.TodoID) ([]*entity.Todo, error) {

	r.log.Info(ctx, "called")
	var todos []*entity.Todo
	if err := r.Db.First(&todos, "id = ?", todoID); err.RecordNotFound() {
		return nil, errorenum.DataNull
	}
	return todos, nil
}
