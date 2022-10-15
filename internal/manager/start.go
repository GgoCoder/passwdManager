package manager

import (
	"passwdManager/pkg/contanst"
	"passwdManager/pkg/typed"
)

func Start(obj *typed.PM, mode contanst.HandleAction)interface{}{
	var result interface{}

	switch mode{
	case contanst.Add:
		passwdManagerRoot.AddPasswd(obj)
	case contanst.Delete:
		passwdManagerRoot.DeletePasswd(obj)
	case contanst.Update:
		passwdManagerRoot.UpdatePasswd(obj)
	case contanst.List:
		result = passwdManagerRoot.ListPasswd()
	case contanst.Search:
		result = passwdManagerRoot.GetPasswdByWebName(obj.WebName)
	}
	return result

}
