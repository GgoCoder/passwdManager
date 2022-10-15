package args

import (
	"flag"
	"fmt"
	"passwdManager/pkg/contanst"
	"passwdManager/pkg/typed"
)
func Parse()(*typed.PM, contanst.HandleAction, error){
	var website string
	var user string
	var passwd string
	var mode string

	flag.StringVar(&website, "n", "", "网站名")
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&passwd, "p", "", "密码")
	flag.StringVar(&mode, "m", "a", "工作模式, a = add, d = del, u = update, l = list, s = search")

	flag.Parse()
	fmt.Printf("website:%s, user: %s, passwd:%s\n", website, user, passwd)
	if user == ""{
		return nil, "" ,contanst.UserIsEmpty
	}
	if passwd == ""{
		return nil, "", contanst.PasswdIsEmpty
	}
	return &typed.PM{WebName:website, User:user, Passwd:passwd},getMode(mode), nil
}

func getMode(mode string)contanst.HandleAction{
	switch mode{
	case "a":
		return contanst.Add
	case "u":
		return contanst.Update
	case "d":
		return contanst.Delete
	case "l":
		return contanst.List
	case "s":
		return contanst.Search
	}
	return contanst.Add
}
