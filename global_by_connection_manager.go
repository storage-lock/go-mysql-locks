package mysql_locks

import (
	"context"
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

var GlobalConnectionManagerMysqlLockFactoryBeanFactory *storage_lock_factory.StorageLockFactoryBeanFactory[storage.ConnectionManager[*sql.DB], *sql.DB]

// CreateLockByConnectionManager 从一个连接管理器中创建一个分布式锁
func CreateLockByConnectionManager(ctx context.Context, connectionManager storage.ConnectionManager[*sql.DB], lockId string) (*storage_lock.StorageLock, error) {
	factory, err := GlobalConnectionManagerMysqlLockFactoryBeanFactory.GetOrInit(ctx, connectionManager, func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[*sql.DB], error) {
		options := mysql_storage.NewMySQLStorageOptions()
		options.SetConnectionManager(connectionManager)
		mysqlStorage, err := mysql_storage.NewMySQLStorage(ctx, options)
		if err != nil {
			return nil, err
		}
		return storage_lock_factory.NewStorageLockFactory(mysqlStorage, connectionManager), nil
	})
	if err != nil {
		return nil, err
	}
	return factory.CreateLock(lockId)
}
