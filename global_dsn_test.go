package mysql_locks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalDsnLockFactoryBeanFactory_InitiallyInitialized(t *testing.T) {
	assert.NotNil(t, dsnStorageLockFactoryBeanFactory)
}

func TestNewMysqlLockByDsn_InvalidDsn(t *testing.T) {
	_, err := NewMysqlLockByDsn(context.Background(), "invalid-dsn", "test-lock")
	assert.Error(t, err)
}

func TestGetMysqlLockFactoryByDsn_InvalidDsn(t *testing.T) {
	_, err := GetMysqlLockFactoryByDsn(context.Background(), "invalid-dsn")
	assert.Error(t, err)
}

func TestNewMysqlLockByDsnWithOptions_InvalidDsn(t *testing.T) {
	_, err := NewMysqlLockByDsnWithOptions(context.Background(), "invalid-dsn", nil)
	assert.Error(t, err)
}

