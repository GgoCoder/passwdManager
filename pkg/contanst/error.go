package contanst

import "fmt"

var(
	UserIsEmpty = fmt.Errorf("user name is empty")
	PasswdIsEmpty = fmt.Errorf("passwd is empty")
	WebNameIsEmpty = fmt.Errorf("web name is empty")
)
