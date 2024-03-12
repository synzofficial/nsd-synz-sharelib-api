package setinfrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type GormConfig struct {
	DbHost        string `envconfig:"DB_HOST"`
	DbPort        string `envconfig:"DB_PORT"`
	DbName        string `envconfig:"DB_NAME"`
	DbUser        string `envconfig:"DB_USER"`
	DbPassword    string `envconfig:"DB_PASSWORD"`
	DbReplicaHost string `envconfig:"DB_REPLICA_HOST"`
}

func NewGorm_MySQL(cfg GormConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb3&parseTime=True&loc=Local",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewGormWithReplica_MySQL(cfg GormConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb3&parseTime=True&loc=Local",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName)
	replicaDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb3&parseTime=True&loc=Local",
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbReplicaHost,
		cfg.DbPort,
		cfg.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{
			mysql.Open(replicaDsn),
		},
		TraceResolverMode: true,
	})); err != nil {
		return nil, err
	}

	// test ping
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

/*
func NewGORM(dsn string) (*gorm.DB, error) {
	// Connect to database
	fmt.Printf("Connecting to database: %s\n", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewGORMWithReplica(dsn, replicaDsn string) (*gorm.DB, error) {
	// Connect to database
	fmt.Printf("Connecting to database: %s\n", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		return nil, err
	}

	// dbresolver will automatically select replicas (replicaDsn) for read operation
	// and select master (dsn) for write operation
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{
			postgres.Open(replicaDsn),
		},
		TraceResolverMode: true,
	}))

	if err != nil {
		return nil, err
	}

	return db, nil
}
*/
