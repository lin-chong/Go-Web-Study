// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package queries

import (
	"time"
)

// 用户
type User struct {
	ID int64
	// 用户名
	Username string
	// 用户密码
	Password string
	// 注册时间
	CreatedAt time.Time
}
