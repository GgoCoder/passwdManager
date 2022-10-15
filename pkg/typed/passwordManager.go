package typed

import "time"

//password manager
type PM struct{
	ID string
	WebName string
	User string
	Passwd  string
	CreateTime time.Time
	//UpdateTime time.Time
	LastUseTime time.Time
}


