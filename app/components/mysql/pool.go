package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"table-saas/config"
)

func getConnStr(conf *config.MysqlConfig) string  {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Extra,
	)
}

func GetMysqlPool() (*gorm.DB, error)  {
	dbCoon, err := gorm.Open("mysql", getConnStr(config.MysqlConf))
	if err != nil {
		logrus.Error("gorm.Open failure" + err.Error())
		return nil, err
	}
	dbCoon.DB().SetMaxIdleConns(10)
	dbCoon.DB().SetMaxOpenConns(100)
	dbCoon.SingularTable(true)

	// Enable Logger, show detailed log
	dbCoon.LogMode(true)
	dbCoon.SetLogger(log.New(os.Stdout, "\n", 0))

	return dbCoon, err
}
