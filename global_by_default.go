package mysql_locks

import (
	"context"
	"database/sql"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

var GlobalDefaultMysqlLockFactoryBeanFactory *storage_lock_factory.StorageLockFactoryBeanFactory[string, *sql.DB]

func InitDefaultMysqlLockFactory(ctx context.Context, dsn string) error {

}

func CreateMysqlStorageLock(lockId string) (*storage_lock.StorageLock, error) {
	
}
