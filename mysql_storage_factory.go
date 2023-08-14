package mysql_locks

import (
	"context"
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

type LockFactory struct {
	*storage_lock_factory.StorageLockFactory[*sql.DB]
}

// NewLockFactory 从DSN创建锁
func NewLockFactory(ctx context.Context, dsn string) (*LockFactory, error) {
	connectionManager := mysql_storage.NewMySQLConnectionManagerFromDSN(dsn)
	return NewLockFactoryFromConnectionManager(ctx, connectionManager)
}

// NewLockFactoryFromSqlDB 从给定的sql.DB创建mysql锁
func NewLockFactoryFromSqlDB(ctx context.Context, db *sql.DB) (*LockFactory, error) {
	connectionManager := storage.NewFixedSqlDBConnectionManager(db)
	return NewLockFactoryFromConnectionManager(ctx, connectionManager)
}

// NewLockFactoryFromConnectionManager 从连接管理器创建锁
func NewLockFactoryFromConnectionManager(ctx context.Context, connectionManager storage.ConnectionManager[*sql.DB]) (*LockFactory, error) {
	options := mysql_storage.NewMySQLStorageOptions()
	options.SetConnectionManager(connectionManager)
	mysqlStorage, err := mysql_storage.NewMySQLStorage(ctx, options)
	if err != nil {
		return nil, err
	}
	factory := storage_lock_factory.NewStorageLockFactory(mysqlStorage, connectionManager)
	if err != nil {
		return nil, err
	}
	return &LockFactory{
		StorageLockFactory: factory,
	}, nil
}
