package mysql_locks

import (
	"context"
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

var GlobalSqlDbMysqlLockFactoryBeanFactory *storage_lock_factory.StorageLockFactoryBeanFactory[*sql.DB, *sql.DB]

// CreateLockBySqlDb 从一个sql.DB中创建一个分布式锁
func CreateLockBySqlDb(ctx context.Context, db *sql.DB, lockId string) (*storage_lock.StorageLock, error) {

	factory, err := GlobalSqlDbMysqlLockFactoryBeanFactory.GetOrInit(ctx, db, func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[*sql.DB], error) {
		connectionManager := storage.NewFixedSqlDBConnectionManager(db)
		options := mysql_storage.NewMySQLStorageOptions().SetConnectionManager(connectionManager)
		mysqlStorage, err := mysql_storage.NewMySQLStorage(ctx, options)
		if err != nil {
			return nil, err
		}
		return storage_lock_factory.NewStorageLockFactory[*sql.DB](mysqlStorage, connectionManager), nil
	})
	if err != nil {
		return nil, err
	}
	return factory.CreateLock(lockId)
}
