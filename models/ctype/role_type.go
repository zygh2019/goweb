package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin   Role = 1
	PermissionUser    Role = 2
	PermissionVisitor Role = 3
	PermissionDisUser Role = 4
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	switch r {

	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisitor:
		return "访客"
	case PermissionDisUser:
		return "禁止的用户"
	default:
		return "unknown"
	}
}
