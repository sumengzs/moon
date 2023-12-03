package repository

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"prometheus-manager/app/prom_server/internal/biz/bo"
	"prometheus-manager/pkg/helper/valueobj"
)

var _ PromDictRepo = (*UnimplementedPromDictRepo)(nil)

type (
	PromDictRepo interface {
		mustEmbedUnimplemented()
		// CreateDict 创建字典
		CreateDict(ctx context.Context, dict *bo.DictBO) (*bo.DictBO, error)
		// UpdateDictById 通过id更新字典
		UpdateDictById(ctx context.Context, id uint, dict *bo.DictBO) (*bo.DictBO, error)
		// BatchUpdateDictStatusByIds 通过id批量更新字典状态
		BatchUpdateDictStatusByIds(ctx context.Context, status valueobj.Status, ids []uint) error
		// DeleteDictByIds 通过id删除字典
		DeleteDictByIds(ctx context.Context, id ...uint) error
		// GetDictById 通过id获取字典详情
		GetDictById(ctx context.Context, id uint) (*bo.DictBO, error)
		// ListDict 获取字典列表
		ListDict(ctx context.Context, pgInfo query.Pagination, scopes ...query.ScopeMethod) ([]*bo.DictBO, error)
	}

	UnimplementedPromDictRepo struct{}
)

func (UnimplementedPromDictRepo) mustEmbedUnimplemented() {}

func (UnimplementedPromDictRepo) CreateDict(_ context.Context, _ *bo.DictBO) (*bo.DictBO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDict not implemented")
}

func (UnimplementedPromDictRepo) UpdateDictById(_ context.Context, _ uint, _ *bo.DictBO) (*bo.DictBO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDictById not implemented")
}

func (UnimplementedPromDictRepo) BatchUpdateDictStatusByIds(_ context.Context, _ valueobj.Status, _ []uint) error {
	return status.Errorf(codes.Unimplemented, "method BatchUpdateDictStatusByIds not implemented")
}

func (UnimplementedPromDictRepo) DeleteDictByIds(_ context.Context, _ ...uint) error {
	return status.Errorf(codes.Unimplemented, "method DeleteDictByIds not implemented")
}

func (UnimplementedPromDictRepo) GetDictById(_ context.Context, _ uint) (*bo.DictBO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDictById not implemented")
}

func (UnimplementedPromDictRepo) ListDict(_ context.Context, _ query.Pagination, _ ...query.ScopeMethod) ([]*bo.DictBO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDict not implemented")
}
