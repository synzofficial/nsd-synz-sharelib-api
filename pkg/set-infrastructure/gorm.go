package setinfrastructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

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
