package driven

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"

	"github.com/satori/go.uuid"
)

type (
	Driven struct {
		DB *gorm.DB
	}

	DatabaseSettings struct {
		Type     string
		User     string
		Password string
		Host     string
		Name     string
	}

	Settings struct {
		Database DatabaseSettings
	}
)

// CloseDB closes database connection (unnecessary)
func (d Driven) CloseDB() {
	defer d.DB.Close()
}

func (d Driven) Setup(settings Settings) {

	db, err := connectDB(settings)

	if err != nil {
		log.Fatalf("Gorm err: %v", err)
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)

	d.DB = db
}

func connectDB(settings Settings) (*gorm.DB, error) {
	connect := func() (*gorm.DB, error) {
		db, err := gorm.Open(
				settings.Database.Type,
				fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
					settings.Database.User,
					settings.Database.Password,
					settings.Database.Host,
					settings.Database.Name))


		if err != nil {
			return nil, err
		}

		return db, nil
	}

	return retry(connect, 20)
}

func retry(connect func() (*gorm.DB, error), retries int) (*gorm.DB, error) {
	i := 0
	for  {
		db, err := connect()
		if err == nil {
			return db, nil
		}

		if i >= retries {
			return nil, errors.New("database unavailable")
		}

		i++
		time.Sleep(1 * time.Second)
	}
}

// updateTimeStampForCreateCallback will set `CreatedAt`, `ModifiedAt` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if idField, ok := scope.FieldByName("Id"); ok {
			if idField.IsBlank {
				idField.Set(uuid.NewV4())
			}
		}

		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}

// deleteCallback will set `DeletedAt` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
