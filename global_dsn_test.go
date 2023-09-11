package mysql_locks

import (
	"context"
	storage_lock_test_helper "github.com/storage-lock/go-storage-lock-test-helper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const MYSQL_DSN_ENV_NAME = "STORAGE_LOCK_MYSQL_DSN"

func TestNewMysqlLockByDsn(t *testing.T) {
	mysqlDsn := os.Getenv(MYSQL_DSN_ENV_NAME)
	assert.NotEmpty(t, mysqlDsn)

	factory, err := GetMysqlLockFactoryByDsn(context.Background(), mysqlDsn)
	assert.Nil(t, err)

	storage_lock_test_helper.PlayerNum = 10
	storage_lock_test_helper.EveryOnePlayTimes = 100
	storage_lock_test_helper.TestStorageLock(t, factory)
}
