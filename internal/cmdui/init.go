package cmdui

import (
	"fmt"
	"passwdManager/internal/manager"
	"passwdManager/pkg/contanst"
	"passwdManager/pkg/typed"

	"github.com/abiosoft/ishell/v2"
)
var shell = ishell.New()
func CmdUIStart() {
	shell.Println("welcome to passwd manager!")

	shell.AddCmd(&ishell.Cmd{
		Name: "add",
		Help: "add one passwd ",
		Func :func (c *ishell.Context)  {
			if manager.IsLogin(){
				c.ShowPrompt(false)
				defer c.ShowPrompt(true)

				c.Print("web name: ")
				webName := c.ReadLine()
				c.Print("user name: ")
				userName := c.ReadLine()
				c.Print("passwd: ")
				passwd := c.ReadLine()

				pm := &typed.PM{WebName:webName, User:userName, Passwd:passwd}
				manager.Start(pm, contanst.Add)
			}else{
				c.Println("please login")
			}
		
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "please login",
		Func :func (c *ishell.Context)  {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Print("user name: ")
			userName := c.ReadLine()
			c.Print("password: ")
			passwd := c.ReadPassword()
			if manager.Login(userName, passwd){
				c.Println("login successfully!")
			}else{
				c.Println("failed to login")
			}

		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "ls",
		Help: "list all passwd",
		Func :func (c *ishell.Context)  {
			if manager.IsLogin(){
				pms := manager.GetPMRoot().ListPasswd()
				c.Println("password count: ", len(pms))
				createUI(pms)
			}else{
				c.Println("please login")
			}
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "test",
		Help: "add test passwd, [{'百度','user1','12345'}, {'阿里','user2','54321'}]",
		Func :func (c *ishell.Context)  {	
			pm := &typed.PM{WebName:"百度", User:"user1", Passwd:"12345"}
			manager.Start(pm, contanst.Add)			
			pm1 := &typed.PM{WebName:"阿里", User:"user2", Passwd:"54321"}
			manager.Start(pm1, contanst.Add)
		},
	})

	shell.Run()
}

func createUI(pms []typed.PM){
	fmt.Println("*序号\t*网站名\t*用户名\t*密码*\t创建时间\t*上次使用时间**")
	for _, pm := range pms{
		fmt.Printf("%s\t%s\t  %s\t  %s\t  %s\t  %s\t\n", pm.ID, pm.WebName, pm.User, pm.Passwd, pm.CreateTime, pm.LastUseTime)
	}
}


