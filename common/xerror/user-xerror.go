package xerror

import "errors"

var (
	UserNotExistError = errors.New("用户不存在")
	UserExistError    = errors.New("改用户存在")
	LoginFailError    = errors.New("登录失败，账号或秘密不正确")
	CreateFailError   = errors.New("用户创建失败正确")
)
