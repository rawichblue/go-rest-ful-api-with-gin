package roledto

type ReqCreateRole struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsActived   bool   `json:"is_actived"`
}
type ReqUpdateRole struct {
	ReqCreateRole
}

type ReqSetPermission struct {
	RoleId        int64   `json:"role_id"`
	PermissionIds []int64 `json:"permission_ids"`
}

//	type ReqPermissionId struct {
//		Id int64 `uri:"id"`
//	}
type ReqPermissionId struct {
	Id int64 `uri:"id"`
}

type ReqGetRoleList struct {
	Page   int    `form:"page"`
	Size   int    `form:"size"`
	Search string `form:"search"`
}

type ReqStatusRole struct {
	IsActive bool `json:"is_active"`
}
