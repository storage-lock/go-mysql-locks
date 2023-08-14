package mysql_locks

import (
	"context"
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

var GlobalDsnMysqlLockFactoryBeanFactory *storage_lock_factory.StorageLockFactoryBeanFactory[string, *sql.DB]

func CreateMysqlLock(ctx context.Context, dsn string, resourceID string) (*storage_lock.StorageLock, error) {
	factory, err := GlobalDsnMysqlLockFactoryBeanFactory.GetOrInit(ctx, dsn, func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[*sql.DB], error) {
		connectionManager := mysql_storage.NewMySQLConnectionManagerFromDSN(dsn)
		options := mysql_storage.NewMySQLStorageOptions()
		options.SetConnectionManager(connectionManager)
		mysqlStorage, err := mysql_storage.NewMySQLStorage(ctx, options)
		if err != nil {
			return nil, err
		}
		factory := storage_lock_factory.NewStorageLockFactory[*sql.DB](mysqlStorage, connectionManager)
		return factory, nil
	})
	if err != nil {
		return nil, err
	}
	return factory.CreateLock(resourceID)
}
