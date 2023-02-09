package manager

import (
	"passwdManager/pkg/typed"
	"passwdManager/pkg/util"
	"sync"
)

type PasswdManagerRoot struct{
	sync.RWMutex
	WebName2PM map[string]*typed.PM
	PasswdCount int32
}

var passwdManagerRoot = &PasswdManagerRoot{
	WebName2PM: make(map[string]*typed.PM),
}

func GetPMRoot()*PasswdManagerRoot{
	return passwdManagerRoot
}

func(r *PasswdManagerRoot) DeletePasswd(obj *typed.PM){
	passwdManagerRoot.Lock()
	defer passwdManagerRoot.Unlock()
	delete(passwdManagerRoot.WebName2PM, obj.User)
	passwdManagerRoot.PasswdCount --
}

func (r *PasswdManagerRoot)UpdatePasswd( newObj *typed.PM){
	passwdManagerRoot.Lock()
	defer passwdManagerRoot.Unlock()
	passwdManagerRoot.WebName2PM[newObj.WebName] = newObj
	newObj.LastUseTime = util.GetCurrectTime()
}

func(r *PasswdManagerRoot)AddPasswd(obj *typed.PM){
	passwdManagerRoot.Lock()
	defer passwdManagerRoot.Unlock()
	passwdManagerRoot.PasswdCount ++
	obj.CreateTime = util.GetCurrectTime()
	obj.LastUseTime = util.GetCurrectTime()
	passwdManagerRoot.WebName2PM[obj.WebName] = obj
}

func(r *PasswdManagerRoot)ListPasswd()[]typed.PM{
	passwdManagerRoot.Lock()
	defer passwdManagerRoot.Unlock()
	var allPasswds []typed.PM
	for _, v := range passwdManagerRoot.WebName2PM{
		allPasswds = append(allPasswds, *v)
	}
	return allPasswds

}

func(r *PasswdManagerRoot)GetPasswdByWebName(name string)typed.PM{
	passwdManagerRoot.Lock()
	defer passwdManagerRoot.Unlock()
	if v, ok := passwdManagerRoot.WebName2PM[name];ok{
		tmp := v
		passwdManagerRoot.WebName2PM[name].LastUseTime = util.GetCurrectTime()
		return *tmp
	}
	return typed.PM{}
}

func(r *PasswdManagerRoot) GetPasswdByUserName()typed.PM{
	passwdManagerRoot.Lock()
	defer passwdManagerRoot.Unlock()
	return typed.PM{}
}
