package orm

import (
	"fmt"
	"sync"
	"time"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// HadeGorm 代表hade框架的orm实现
type HadeGorm struct {
    container framework.Container // 服务容器
    dbs       map[string]*gorm.DB // key为dsn, value为gorm.DB（连接池）

    lock *sync.RWMutex
}

// NewHadeGorm 代表实例化Gorm
func NewHadeGorm(params ...interface{}) (interface{}, error) {
    container := params[0].(framework.Container)
    dbs := make(map[string]*gorm.DB)
    lock := &sync.RWMutex{}
    return &HadeGorm{
        container: container,
        dbs:       dbs,
        lock:      lock,
    }, nil
}

// GetDB 获取DB实例
func (app *HadeGorm) GetDB(option ...contract.DBOption) (*gorm.DB, error) {
    // 读取默认配置
    config := GetBaseConfig(app.container)


    // option对opt进行修改
    for _, opt := range option {
        if err := opt(app.container, config); err != nil {
            return nil, err
        }
    }

    // 如果最终的config没有设置dsn,就生成dsn
    if config.Dsn == "" {
        dsn, err := config.FormatDsn()
        if err != nil {
            return nil, err
        }
        config.Dsn = dsn
    }
    fmt.Println(config.Dsn)
    // 判断是否已经实例化了gorm.DB
    app.lock.RLock()
    if db, ok := app.dbs[config.Dsn]; ok {
        app.lock.RUnlock()
        return db, nil
    }
    app.lock.RUnlock()
    
    // 没有实例化gorm.DB，那么就要进行实例化操作
    app.lock.Lock()
    defer app.lock.Unlock()

    // 实例化gorm.DB
    var db *gorm.DB
    var err error
    fmt.Println(2222)
    fmt.Println(config)
    switch config.Driver {
    case "mysql":
        db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
    case "postgres":
        db, err = gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
    case "sqlite":
        db, err = gorm.Open(sqlite.Open(config.Dsn), &gorm.Config{})
    case "sqlserver":
        db, err = gorm.Open(sqlserver.Open(config.Dsn), &gorm.Config{})
    case "clickhouse":
        db, err = gorm.Open(clickhouse.Open(config.Dsn), &gorm.Config{})
    }
    
    // 设置对应的连接池配置
    sqlDB, err := db.DB()
    if err != nil {
        return db, err
    }

    if config.ConnMaxIdle > 0 {
        sqlDB.SetMaxIdleConns(config.ConnMaxIdle)
    }
    if config.ConnMaxOpen > 0 {
        sqlDB.SetMaxOpenConns(config.ConnMaxOpen)
    }
    if config.ConnMaxLifetime != "" {
        liftTime, err := time.ParseDuration(config.ConnMaxLifetime)
        if err != nil {
            fmt.Println(err)
        } else {
            sqlDB.SetConnMaxLifetime(liftTime)
        }
    }

    if config.ConnMaxIdletime != "" {
        idleTime, err := time.ParseDuration(config.ConnMaxIdletime)
        if err != nil {
            fmt.Println(err)
        } else {
            sqlDB.SetConnMaxIdleTime(idleTime)
        }
    }
    // 挂载到map中，结束配置
    if err != nil {
        app.dbs[config.Dsn] = db
    }
    return db, err
}
