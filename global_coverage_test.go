package mysql_locks

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	"github.com/stretchr/testify/assert"
)

func TestNewMysqlLockBySqlDb(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer mockDB.Close()

	// Init 建表
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS").
		WillReturnResult(sqlmock.NewResult(0, 0))

	ctx := context.Background()
	factory, err := GetMysqlLockFactoryBySqlDb(ctx, mockDB)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	// 缓存命中
	factory2, err := GetMysqlLockFactoryBySqlDb(ctx, mockDB)
	assert.Nil(t, err)
	assert.Equal(t, factory, factory2)

	lock, err := NewMysqlLockBySqlDb(ctx, mockDB, "test-mysql-lock-sqldb")
	assert.Nil(t, err)
	assert.NotNil(t, lock)

	options := storage_lock.NewStorageLockOptionsWithLockId("test-mysql-lock-sqldb-opt")
	lock2, err := NewMysqlLockBySqlDbWithOptions(ctx, mockDB, options)
	assert.Nil(t, err)
	assert.NotNil(t, lock2)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestNewMysqlLockByDsn_Local(t *testing.T) {
	dsn := "root:rootpass@tcp(127.0.0.1:13307)/storage_lock_test"
	ctx := context.Background()

	probe, err := GetMysqlLockFactoryByDsn(ctx, dsn)
	if err != nil {
		t.Skipf("本机 sl-mysql 不可用，跳过: %v", err)
	}
	assert.NotNil(t, probe)

	// 缓存命中
	factory2, err := GetMysqlLockFactoryByDsn(ctx, dsn)
	assert.Nil(t, err)
	assert.Equal(t, probe, factory2)

	lock, err := NewMysqlLockByDsn(ctx, dsn, "test-mysql-lock-dsn")
	assert.Nil(t, err)
	assert.NotNil(t, lock)

	options := storage_lock.NewStorageLockOptionsWithLockId("test-mysql-lock-dsn-opt")
	lock2, err := NewMysqlLockByDsnWithOptions(ctx, dsn, options)
	assert.Nil(t, err)
	assert.NotNil(t, lock2)

	assert.Nil(t, lock.Lock(ctx, "owner-1"))
	assert.Nil(t, lock.UnLock(ctx, "owner-1"))
}

func TestNewMysqlLockBySqlDb_Errors(t *testing.T) {
	old := newMysqlStorage
	newMysqlStorage = func(ctx context.Context, options *mysql_storage.MysqlStorageOptions) (*mysql_storage.MysqlStorage, error) {
		return nil, errors.New("boom")
	}
	defer func() { newMysqlStorage = old }()

	mockDB, _, err := sqlmock.New()
	assert.Nil(t, err)
	defer mockDB.Close()

	ctx := context.Background()
	_, err = GetMysqlLockFactoryBySqlDb(ctx, mockDB)
	assert.NotNil(t, err)
	_, err = NewMysqlLockBySqlDb(ctx, mockDB, "x")
	assert.NotNil(t, err)
	_, err = NewMysqlLockBySqlDbWithOptions(ctx, mockDB, storage_lock.NewStorageLockOptionsWithLockId("x"))
	assert.NotNil(t, err)
}
