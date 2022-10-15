package cmdui

import (
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
				for _, pm := range pms{
					c.Println(pm)
				}
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


