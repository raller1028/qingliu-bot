/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-03-28 15:10:50
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-08-20 14:22:56
 * @FilePath: /feishu-bot/service/service.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"github.com/chyroc/lark"
)

var LarkCli *lark.Lark

func InitLark() {
	LarkCli = lark.New(
		lark.WithAppCredential("cli_a19174e22678100e", "VSuQjM1MbY0JK1OTdXMVLfhNzbSeJy0g"),
		lark.WithEventCallbackVerify("abc123456!@#$%^", "CnTfMDBn67fpUGBfYY6KKbLxzG0xojXR"),
	)
}

var Service Repository
var PXCOOKIE = ""

type Repository interface {
	Lark() LarkService
	// Github() GithubService
	Px4() Px4Service
	// Packages() PackageService
}

// func NewService(db *gorm.DB) Repository {
func NewService() Repository {
	return &store{
		lark: NewLarkService(),
		// github:   NewGithubService(),
		px4: NewPx4Service(),
		// packages: NewPackageService(db),
	}
}

type store struct {
	lark LarkService
	// github   GithubService
	px4 Px4Service
	// packages PackageService
}

// func (c *store) Packages() PackageService {
// 	return c.packages
// }

func (c *store) Lark() LarkService {
	return c.lark
}

//	func (c *store) Github() GithubService {
//		return c.github
//	}
func (c *store) Px4() Px4Service {
	return c.px4
}
