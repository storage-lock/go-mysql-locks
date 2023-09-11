package main

import (
	"context"
	"fmt"
	mysql_locks "github.com/storage-lock/go-mysql-locks"
)

func main() {

	// 第一步：创建一把分布式锁
	mysqlDsn := "root:UeGqAm8CxYGldMDLoNNt@tcp(127.0.0.1:3306)/storage_lock_test"
	lockId := "test-lock"
	lock, err := mysql_locks.NewMysqlLockByDsn(context.Background(), mysqlDsn, lockId)
	if err != nil {
		panic(err)
	}

	// 第二步：尝试加锁
	ownerId := "owner-A"
	err = lock.Lock(context.Background(), ownerId)
	if err != nil {
		panic(err)
	}
	// 加锁成功之后要记得释放锁
	defer func() {
		err := lock.UnLock(context.Background(), ownerId)
		if err != nil {
			panic(err)
		}
	}()

	// 临界区，操作临界资源
	fmt.Println("Lock success")

}
