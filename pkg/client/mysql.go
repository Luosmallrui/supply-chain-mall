package client

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"supply-chain-mall/config"
)

func NewMySQLClient(config *config.Config) *gorm.DB {
	//file := logger2.CreateFileWriter(conf.Log.LogFilePath("slow-sql.log"))

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
		//	SlowThreshold:             5 * time.Millisecond,
		//	LogLevel:                  logger.Warn,
		//	IgnoreRecordNotFoundError: true,
		//}),
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?&parseTime=True&loc=Local",
			config.DB.User,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
			config.DB.Name,
		),
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), gormConfig)

	if err != nil {
		panic(fmt.Errorf("mysql connect error :%v", err))
	}

	if db.Error != nil {
		panic(fmt.Errorf("database error :%v", err))
	}

	_, _ = db.DB()

	fmt.Println("successful to connect db")
	return db
}
