package main

import (
	"coredemo/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	time.Sleep(15 * time.Second)
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
