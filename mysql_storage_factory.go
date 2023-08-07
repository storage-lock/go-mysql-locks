package mysql_locks

import (
	"context"
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

type MysqlLockFactory struct {
	*storage_lock_factory.StorageLockFactory[*sql.DB]
}

func NewMysqlLockFactory(ctx context.Context, dsn string) (*MysqlLockFactory, error) {
	connectionManager := mysql_storage.NewMySQLConnectionManagerFromDSN(dsn)
	return NewMysqlLockFactoryFromConnectionManager(ctx, connectionManager)
}

func NewMysqlLockFactoryFromConnectionManager(ctx context.Context, connectionManager storage.ConnectionManager[*sql.DB]) (*MysqlLockFactory, error) {
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
	return &MysqlLockFactory{
		StorageLockFactory: factory,
	}, nil
}
