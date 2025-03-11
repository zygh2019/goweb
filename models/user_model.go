package models

import "awesomeProject1/models/ctype"

type User struct {
	Username string     ` json:"username"`
	Password string     ` json:"password"`
	role     ctype.Role ` json:"role"`
	list     []string   ` json:"-"`
}
