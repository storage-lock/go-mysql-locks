package mysql_locks

import (
	"context"
	"database/sql"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

var (
	globalDsnLockFactory                 = storage_lock_factory.NewStorageLockFactoryBeanFactory[string, *sql.DB]()
	globalSqlDbLockFactory               = storage_lock_factory.NewStorageLockFactoryBeanFactory[*sql.DB, *sql.DB]()
	globalConnectionManagerDbLockFactory = storage_lock_factory.NewStorageLockFactoryBeanFactory[storage.ConnectionManager[*sql.DB], *sql.DB]()
)

// NewLock 从DSN创建锁
func NewLock(ctx context.Context, dsn, lockId string) (*storage_lock.StorageLock, error) {
	init, err := globalDsnLockFactory.GetOrInit(ctx, dsn, func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[*sql.DB], error) {
		factory, err := NewLockFactory(ctx, dsn)
		if err != nil {
			return nil, err
		}
		return factory.StorageLockFactory, nil
	})
	if err != nil {
		return nil, err
	}
	return init.CreateLock(lockId)
}

// NewLockFromSqlDB 从给定的sql.DB创建mysql锁
func NewLockFromSqlDB(ctx context.Context, db *sql.DB, lockId string) (*storage_lock.StorageLock, error) {
	init, err := globalSqlDbLockFactory.GetOrInit(ctx, db, func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[*sql.DB], error) {
		factory, err := NewLockFactoryFromSqlDB(ctx, db)
		if err != nil {
			return nil, err
		}
		return factory.StorageLockFactory, nil
	})
	if err != nil {
		return nil, err
	}
	return init.CreateLock(lockId)
}

// NewLockFromConnectionManager 从连接管理器创建锁
func NewLockFromConnectionManager(ctx context.Context, connectionManager storage.ConnectionManager[*sql.DB], lockId string) (*storage_lock.StorageLock, error) {
	init, err := globalConnectionManagerDbLockFactory.GetOrInit(ctx, connectionManager, func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[*sql.DB], error) {
		factory, err := NewLockFactoryFromConnectionManager(ctx, connectionManager)
		if err != nil {
			return nil, err
		}
		return factory.StorageLockFactory, nil
	})
	if err != nil {
		return nil, err
	}
	return init.CreateLock(lockId)
}
