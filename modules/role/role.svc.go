package role

import (
	"app/models"
	"app/modules/response"
	roledto "app/modules/role/dto"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

type RoleService struct {
	db *bun.DB
}

func newService(db *bun.DB) *RoleService {
	return &RoleService{
		db: db,
	}
}

func (s *RoleService) Create(ctx context.Context, req roledto.ReqCreateRole) (*models.Role, error) {
	m := models.Role{
		Name:        req.Name,
		Description: req.Description,
		IsActived:   req.IsActived,
	}

	_, err := s.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *RoleService) SetPermission(ctx context.Context, req roledto.ReqSetPermission) error {

	for _, per := range req.PermissionIds {
		ex, err := s.db.NewSelect().TableExpr("permission").Where("id = ? AND is_active = ?", per, true).Exists(ctx)
		if err != nil {
			return err
		}

		if !ex {
			return errors.New("status not")
		}

	}

	_, err := s.db.NewDelete().TableExpr("role_permission").Where("role_id = ?", req.RoleId).Exec(ctx)

	if err != nil {
		return err
	}

	for _, per := range req.PermissionIds {
		rolePermission := models.RolePermission{
			RoleId:       req.RoleId,
			PermissionId: per,
		}

		_, err := s.db.NewInsert().Model(&rolePermission).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *RoleService) GetPermission(ctx context.Context, req roledto.ReqPermissionId) ([]int, error) {
	// []models.RolePermission
	// m := []models.RolePermission{}
	// err := s.db.NewSelect().Model(&m).Where("role_id = ?", req.Id).Scan(ctx)

	// return m, err

	var data []int

	err := s.db.NewSelect().TableExpr("role_permission").ColumnExpr("permission_id").Where("role_id = ?", req.Id).Scan(ctx, &data)

	return data, err
}

func (s *RoleService) DeleteRole(ctx context.Context, req roledto.ReqPermissionId) (*models.Role, error) {

	ex, err := s.db.NewSelect().TableExpr("employees").Where("role_id = ?", req.Id).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if ex {
		return nil, errors.New("have user used")
	}

	_, err = s.db.NewDelete().TableExpr("role_permission").Where("role_id = ?", req.Id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	_, err = s.db.NewDelete().TableExpr("role").Where("id = ?", req.Id).Exec(ctx)

	return nil, err
}

func (s *RoleService) GetRoleList(ctx context.Context, req roledto.ReqGetRoleList) ([]models.Role, *response.Paginate, error) {
	resp := []models.Role{}

	var offset int
	if req.Page > 1 {
		offset = (req.Page - 1) * req.Size
	} else {
		offset = 0
	}

	query := s.db.NewSelect().Model(&resp)

	if req.Search != "" {
		search := fmt.Sprintf("%%%s%%", req.Search)
		query.Where("name ILIKE ?", search)
	}

	Count, err := query.Count(ctx)

	if err != nil {
		return nil, nil, err
	}

	paginate := response.Paginate{
		Page:  int64(req.Page),
		Size:  int64(req.Size),
		Total: int64(Count),
	}

	err = query.Order("id ASC").Limit(req.Size).Offset(offset).Scan(ctx)

	log.Printf("data : %v", resp)
	if err != nil {
		return nil, nil, err
	}

	return resp, &paginate, nil
}

func (s *RoleService) Update(ctx context.Context, id roledto.ReqPermissionId, req roledto.ReqCreateRole) (*models.Role, error) {
	ex, err := s.db.NewSelect().Model((*models.Role)(nil)).Where("id = ?", id.Id).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("role not found")
	}

	m := models.Role{
		Id:          id.Id,
		Name:        req.Name,
		Description: req.Description,
		IsActived:   req.IsActived,
	}

	_, err = s.db.NewUpdate().Model(&m).
		Set("name = ?name").
		Set("description = ?description ").
		Set("is_actived = ?is_actived").
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)

	return &m, err
}

func (s *RoleService) UpdateStatus(ctx context.Context, id roledto.ReqPermissionId, req roledto.ReqStatusRole) (*models.Role, error) {

	ex, err := s.db.NewSelect().Model((*models.Role)(nil)).Where("id = ?", id.Id).Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errors.New("role not found")
	}

	m := models.Role{
		Id:        id.Id,
		IsActived: req.IsActive,
	}

	_, err = s.db.NewUpdate().Model(&m).
		Set("is_actived = ?", req.IsActive).
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)

	return &m, err
}
