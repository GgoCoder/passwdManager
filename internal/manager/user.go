package manager

import "sync"

type UserRoot struct{
	sync.RWMutex
	Users map[string]int
}

 var userRoot = &UserRoot{
	Users: make(map[string]int),
}

func GetUserRoot()*UserRoot{
	return userRoot
}

func Login(user, passwd string)bool{
	userRoot.Lock()
	defer userRoot.Unlock()
	if user == "123456" && passwd == "123456"{
		userRoot.Users[user] = 1
		return true
	}
	return false
}

func IsLogin()bool{
	return len(userRoot.Users)>0
}
