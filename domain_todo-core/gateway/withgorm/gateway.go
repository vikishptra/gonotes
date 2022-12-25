package withgorm

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

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
	Db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/notes?charset=utf8&parseTime=True")

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

	var testt entity.Todo
	test := r.Db.Where("message = ?", obj.Message).First(&testt)

	if !test.RecordNotFound() {
		return errorenum.ObjSame
	}

	err := r.Db.Save(obj).Error
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *Gateway) FindOneChecked(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	var test entity.Todo
	if err := r.Db.First(&test, "id = ?", todoID); err.RecordNotFound() {
		return nil, errorenum.DataNull
	}
	test.SetTrue()
	r.Db.Save(&test)

	return &test, nil
}

func (r *Gateway) GetAllTodo(ctx context.Context) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")
	var todo []*entity.Todo
	if test := r.Db.Find(&todo); test.RowsAffected != 1 {
		return nil, 0, errorenum.DataNull
	}
	return todo, 0, nil
}
