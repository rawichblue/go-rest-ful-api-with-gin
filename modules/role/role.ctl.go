package role

import (
	"app/modules/response"
	roledto "app/modules/role/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleSvc *RoleService
}

func newController(roleSvcService *RoleService) *RoleController {
	return &RoleController{
		roleSvc: roleSvcService,
	}
}

func (c *RoleController) CreateRole(ctx *gin.Context) {
	req := roledto.ReqCreateRole{}
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if req.Name == "" {
		response.BadRequest(ctx, "ใส่ชื่อมาไอเวร")
		return
	}

	data, err := c.roleSvc.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	// if err != nil {
	// 	response.BadRequest(ctx, err.Error())
	// 	return
	// }
	// response.Success(ctx, data)

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func (c *RoleController) SetPermission(ctx *gin.Context) {
	req := roledto.ReqSetPermission{}
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if req.RoleId == 0 {
		response.BadRequest(ctx, "ใส่เลขมาด้วย จ้า")
		return
	}

	err := c.roleSvc.SetPermission(ctx, req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// ถ้าไม่มีขือ
func (c *RoleController) GetPermission(ctx *gin.Context) {
	req := roledto.ReqPermissionId{}
	if err := ctx.BindUri(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if req.Id == 0 {
		response.BadRequest(ctx, "ใส่เลขมาด้วย จ้า")
		return
	}

	data, err := c.roleSvc.GetPermission(ctx, req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	response.Success(ctx, data)
}

func (c *RoleController) DeleteRole(ctx *gin.Context) {
	id := roledto.ReqPermissionId{}
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := c.roleSvc.DeleteRole(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func (c *RoleController) GetRoleList(ctx *gin.Context) {
	var req roledto.ReqGetRoleList

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.BadRequest(ctx, "Invalid request data")
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	role, paginate, err := c.roleSvc.GetRoleList(ctx.Request.Context(), req)
	if err != nil {
		response.InternalError(ctx, err.Error())
		return
	}
	response.SuccessWithPaginate(ctx, role, paginate)
}

func (c *RoleController) UpdateRole(ctx *gin.Context) {
	id := roledto.ReqPermissionId{}
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req := roledto.ReqCreateRole{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := c.roleSvc.Update(ctx, id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func (ctl *RoleController) PermissionChangeStatus(c *gin.Context) {
	id := roledto.ReqPermissionId{}

	if err := c.BindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	status := roledto.ReqStatusRole{}
	if err := c.Bind(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := ctl.roleSvc.UpdateStatus(c, id, status)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)
}
