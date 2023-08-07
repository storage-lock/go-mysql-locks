package mysql_locks

//import (
//	"context"
//	mysql_storage "github.com/storage-lock/go-mysql-storage"
//	"github.com/storage-lock/go-storage"
//	storage_lock "github.com/storage-lock/go-storage-lock"
//)
//
//// NewMySQLStorageLock 高层API，使用默认配置快速创建基于MySQL的分布式锁
//func NewMySQLStorageLock(ctx context.Context, lockId string, dsn string) (*storage_lock.StorageLock, error) {
//	connectionManager := mysql_storage.NewMySQLConnectionManagerFromDSN(dsn)
//	storageOptions := &mysql_storage.MySQLStorageOptions{
//		ConnectionManager: connectionManager,
//		TableName:         storage.DefaultStorageTableName,
//	}
//
//	s, err := mysql_storage.NewMySQLStorage(ctx, storageOptions)
//	if err != nil {
//		return nil, err
//	}
//
//	lockOptions := &storage_lock.StorageLockOptions{
//		LockId:               lockId,
//		LeaseExpireAfter:     storage_lock.DefaultLeaseExpireAfter,
//		LeaseRefreshInterval: storage_lock.DefaultLeaseRefreshInterval,
//		//VersionMissRetryTimes: storage_lock.DefaultVersionMissRetryTimes,
//	}
//	return storage_lock.NewStorageLockWithOptions(s, lockOptions)
//}
