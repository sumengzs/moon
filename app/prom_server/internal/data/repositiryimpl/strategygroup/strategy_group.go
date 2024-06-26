package strategygroup

import (
	"bytes"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/aide-family/moon/app/prom_server/internal/biz/bo"
	"github.com/aide-family/moon/app/prom_server/internal/biz/do"
	"github.com/aide-family/moon/app/prom_server/internal/biz/do/basescopes"
	"github.com/aide-family/moon/app/prom_server/internal/biz/repository"
	"github.com/aide-family/moon/app/prom_server/internal/biz/vobj"
	"github.com/aide-family/moon/app/prom_server/internal/data"
	"github.com/aide-family/moon/pkg/after"
	"github.com/aide-family/moon/pkg/helper/prom"
	"github.com/aide-family/moon/pkg/util/slices"
)

var _ repository.StrategyGroupRepo = (*strategyGroupRepoImpl)(nil)

type (
	strategyGroupRepoImpl struct {
		repository.UnimplementedStrategyGroupRepo

		changeGroupChannel chan<- uint32
		removeGroupChannel chan<- bo.RemoveStrategyGroupBO

		data *data.Data
		log  *log.Helper
	}
)

type StrategyCount struct {
	Count   int64  `json:"count"`
	GroupId uint32 `json:"group_id"`
}

func (l *strategyGroupRepoImpl) GetByParams(ctx context.Context, scopes ...basescopes.ScopeMethod) ([]*bo.StrategyGroupBO, error) {
	var list []*do.PromStrategyGroup
	if err := l.data.DB().WithContext(ctx).Scopes(scopes...).Find(&list).Error; err != nil {
		return nil, err
	}
	return slices.To(list, func(item *do.PromStrategyGroup) *bo.StrategyGroupBO { return bo.StrategyGroupModelToBO(item) }), nil
}

func (l *strategyGroupRepoImpl) UpdateStrategyCount(ctx context.Context, ids ...uint32) error {
	if len(ids) == 0 {
		return nil
	}

	tx := basescopes.GetTx(ctx, l.data.DB())
	strategyDB := tx.Model(&do.PromStrategy{}).WithContext(ctx)
	var strategyCountList []StrategyCount
	if err := strategyDB.Scopes(do.StrategyInGroupIds(ids...)).Select("count(id) as count, group_id").Group("group_id").Scan(&strategyCountList).Error; err != nil {
		return err
	}

	db := tx.Model(&do.PromStrategyGroup{}).WithContext(ctx)
	switch len(strategyCountList) {
	case 0:
		return db.Scopes(basescopes.InIds(ids...)).Update("strategy_count", 0).Error
	case 1:
		return db.Scopes(basescopes.InIds(ids...)).Update("strategy_count", strategyCountList[0].Count).Error
	default:
		var caseSet bytes.Buffer
		caseSet.WriteString("CASE id ")
		for _, strategyCount := range strategyCountList {
			caseSet.WriteString("WHEN ")
			caseSet.WriteString(fmt.Sprintf("%d", strategyCount.GroupId))
			caseSet.WriteString(" THEN ")
			caseSet.WriteString(fmt.Sprintf("%d", strategyCount.Count))
			caseSet.WriteString(" ")
		}
		caseSet.WriteString("END")

		return db.Scopes(basescopes.InIds(ids...)).Update("strategy_count", gorm.Expr(caseSet.String())).Error
	}
}

func (l *strategyGroupRepoImpl) UpdateEnableStrategyCount(ctx context.Context, ids ...uint32) error {
	if len(ids) == 0 {
		return nil
	}

	tx := basescopes.GetTx(ctx, l.data.DB())
	strategyDB := tx.Model(&do.PromStrategy{}).WithContext(ctx)
	var strategyCountList []StrategyCount
	wheres := []basescopes.ScopeMethod{
		do.StrategyInGroupIds(ids...),
		basescopes.StatusEQ(vobj.StatusEnabled),
	}
	if err := strategyDB.Scopes(wheres...).Select("count(id) as count, group_id").Group("group_id").Scan(&strategyCountList).Error; err != nil {
		return err
	}

	db := tx.Model(&do.PromStrategyGroup{}).WithContext(ctx)
	switch len(strategyCountList) {
	case 0:
		return db.Scopes(basescopes.InIds(ids...)).Update("enable_strategy_count", 0).Error
	case 1:
		prom.WorkingStrategyCounter.WithLabelValues("prom-server").Inc()
		return db.Scopes(basescopes.InIds(ids...)).Update("enable_strategy_count", strategyCountList[0].Count).Error
	default:
		var caseSet bytes.Buffer
		caseSet.WriteString("CASE id ")
		total := float64(0)
		for _, strategyCount := range strategyCountList {
			caseSet.WriteString("WHEN ")
			caseSet.WriteString(fmt.Sprintf("%d", strategyCount.GroupId))
			caseSet.WriteString(" THEN ")
			caseSet.WriteString(fmt.Sprintf("%d", strategyCount.Count))
			caseSet.WriteString(" ")
			total += float64(strategyCount.Count)
		}
		caseSet.WriteString("END")
		prom.WorkingStrategyCounter.WithLabelValues("prom-server").Add(total)
		return db.Scopes(basescopes.InIds(ids...)).Update("enable_strategy_count", gorm.Expr(caseSet.String())).Error
	}
}

func (l *strategyGroupRepoImpl) Create(ctx context.Context, strategyGroup *bo.StrategyGroupBO) (*bo.StrategyGroupBO, error) {
	strategyGroupModel := strategyGroup.ToModel()
	// 默认不开启
	strategyGroupModel.Status = vobj.StatusDisabled
	if err := l.data.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: basescopes.BaseFieldID.String()}},
		UpdateAll: true,
	}).Create(strategyGroupModel).Error; err != nil {
		return nil, err
	}
	defer func() {
		_ = l.UpdateStrategyCount(context.Background(), strategyGroup.Id)
	}()
	return bo.StrategyGroupModelToBO(strategyGroupModel), nil
}

// BatchCreate 批量创建
func (l *strategyGroupRepoImpl) BatchCreate(ctx context.Context, strategyGroups []*bo.StrategyGroupBO) ([]*bo.StrategyGroupBO, error) {
	strategyGroupModelList := slices.To(strategyGroups, func(strategyGroup *bo.StrategyGroupBO) *do.PromStrategyGroup {
		item := strategyGroup.ToModel()
		item.Status = vobj.StatusDisabled
		return item
	})

	err := l.data.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).
			Omit(do.PromGroupPreloadFieldPromStrategies).
			Clauses(clause.OnConflict{DoNothing: true}).
			CreateInBatches(strategyGroupModelList, 10).Error; err != nil {
			return err
		}
		strategyList := make([]*do.PromStrategy, 0, 20)
		for _, strategyGroupModel := range strategyGroupModelList {
			for _, strategy := range strategyGroupModel.GetPromStrategies() {
				strategyTmp := strategy
				strategyTmp.GroupID = strategyGroupModel.ID
				strategyList = append(strategyList, strategy)
			}
		}
		return tx.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(strategyList, 20).Error
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = l.UpdateStrategyCount(context.Background(), slices.To(strategyGroupModelList, func(item *do.PromStrategyGroup) uint32 { return item.ID })...)
	}()

	return slices.To(strategyGroupModelList, func(strategyGroupModel *do.PromStrategyGroup) *bo.StrategyGroupBO {
		return bo.StrategyGroupModelToBO(strategyGroupModel)
	}), nil
}

func (l *strategyGroupRepoImpl) UpdateById(ctx context.Context, id uint32, strategyGroup *bo.StrategyGroupBO) (*bo.StrategyGroupBO, error) {
	strategyGroupModel := strategyGroup.ToModel()
	categories := strategyGroup.GetCategories()
	categoriesDos := slices.To(categories, func(item *bo.DictBO) *do.SysDict { return item.ToModel() })
	err := l.data.DB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Scopes(basescopes.InIds(id)).Updates(strategyGroupModel).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(strategyGroupModel).Association(do.PromGroupPreloadFieldCategories).Replace(categoriesDos); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	go func() {
		defer after.Recover(l.log)
		l.changeGroupChannel <- strategyGroupModel.ID
	}()
	return bo.StrategyGroupModelToBO(strategyGroupModel), nil
}

func (l *strategyGroupRepoImpl) BatchUpdateStatus(ctx context.Context, status vobj.Status, ids []uint32) error {
	if err := l.data.DB().WithContext(ctx).Scopes(basescopes.InIds(ids...)).Updates(&do.PromStrategyGroup{Status: status}).Error; err != nil {
		return err
	}
	go func() {
		defer after.Recover(l.log)
		if status.IsEnabled() {
			for _, id := range ids {
				l.changeGroupChannel <- id
			}
		} else {
			for _, id := range ids {
				l.removeGroupChannel <- bo.RemoveStrategyGroupBO{
					Id: id,
				}
			}
		}
	}()
	return nil
}

func (l *strategyGroupRepoImpl) DeleteByIds(ctx context.Context, ids ...uint32) error {
	if err := l.data.DB().WithContext(ctx).Scopes(basescopes.InIds(ids...)).Delete(&do.PromStrategyGroup{}).Error; err != nil {
		return err
	}
	go func() {
		defer after.Recover(l.log)
		for _, id := range ids {
			l.removeGroupChannel <- bo.RemoveStrategyGroupBO{
				Id: id,
			}
		}
	}()
	return nil
}

func (l *strategyGroupRepoImpl) GetById(ctx context.Context, id uint32) (*bo.StrategyGroupBO, error) {
	var first do.PromStrategyGroup
	if err := l.data.DB().WithContext(ctx).Scopes(do.StrategyGroupPreloadPromStrategies(), do.StrategyGroupPreloadCategories(true)).First(&first, id).Error; err != nil {
		return nil, err
	}

	return bo.StrategyGroupModelToBO(&first), nil
}

func (l *strategyGroupRepoImpl) List(ctx context.Context, pgInfo bo.Pagination, scopes ...basescopes.ScopeMethod) ([]*bo.StrategyGroupBO, error) {
	var strategyModelList []*do.PromStrategyGroup
	if err := l.data.DB().WithContext(ctx).Scopes(append(scopes, bo.Page(pgInfo))...).Find(&strategyModelList).Error; err != nil {
		return nil, err
	}
	if pgInfo != nil {
		var total int64
		if err := l.data.DB().WithContext(ctx).Model(&do.PromStrategyGroup{}).WithContext(ctx).Scopes(scopes...).Count(&total).Error; err != nil {
			return nil, err
		}
		pgInfo.SetTotal(total)
	}
	list := slices.To(strategyModelList, func(m *do.PromStrategyGroup) *bo.StrategyGroupBO {
		return bo.StrategyGroupModelToBO(m)
	})
	return list, nil
}

func (l *strategyGroupRepoImpl) ListAllLimit(ctx context.Context, limit int, scopes ...basescopes.ScopeMethod) ([]*bo.StrategyGroupBO, error) {
	var strategyModelList []*do.PromStrategyGroup
	if err := l.data.DB().WithContext(ctx).Scopes(scopes...).Limit(limit).Find(&strategyModelList).Error; err != nil {
		return nil, err
	}
	list := slices.To(strategyModelList, func(m *do.PromStrategyGroup) *bo.StrategyGroupBO {
		return bo.StrategyGroupModelToBO(m)
	})
	return list, nil
}

func NewStrategyGroupRepo(
	data *data.Data,
	changeGroupChannel chan<- uint32,
	removeGroupChannel chan<- bo.RemoveStrategyGroupBO,
	logger log.Logger,
) repository.StrategyGroupRepo {
	return &strategyGroupRepoImpl{
		data:               data,
		log:                log.NewHelper(logger),
		changeGroupChannel: changeGroupChannel,
		removeGroupChannel: removeGroupChannel,
	}
}
