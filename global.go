package mysql_locks

import (
	"context"
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"sync"
)

var (
	GlobalMysqlLockFactory      *MysqlLockFactory
	globalMysqlLockFactoryOnce  sync.Once
	globalMysqlLockFactoryError error
)

func InitGlobalMysqlLockFactory(ctx context.Context, dsn string) error {
	connectionManager := mysql_storage.NewMySQLConnectionManagerFromDSN(dsn)
	return InitGlobalMysqlLockFactoryFromConnectionManager(ctx, connectionManager)
}

func InitGlobalMysqlLockFactoryFromConnectionManager(ctx context.Context, connectionManager storage.ConnectionManager[*sql.DB]) error {
	options := mysql_storage.NewMySQLStorageOptions()
	options.SetConnectionManager(connectionManager)
	mysqlStorage, err := mysql_storage.NewMySQLStorage(ctx, options)
	if err != nil {
		return err
	}
	factory := storage_lock_factory.NewStorageLockFactory(mysqlStorage, connectionManager)
	if err != nil {
		return err
	}
	GlobalMysqlLockFactory = &MysqlLockFactory{
		StorageLockFactory: factory,
	}
	return nil
}

// TODO 似乎不能这么简单的once，因为每次传入的dsn都可能会不同
func CreateMysqlLock(ctx context.Context, dsn string, lockId string) (*storage_lock.StorageLock, error) {
	globalMysqlLockFactoryOnce.Do(func() {
		globalMysqlLockFactoryError = InitGlobalMysqlLockFactory(ctx, dsn)
	})
	if globalMysqlLockFactoryError != nil {
		return nil, globalMysqlLockFactoryError
	}
	return GlobalMysqlLockFactory.CreateLock(lockId)
}

//func CreateMysqlLockFromConnectionManager(ctx context.Context, connectionManager storage.ConnectionManager[*sql.DB], lockId string) (*storage_lock.StorageLock, error) {
//
//}
