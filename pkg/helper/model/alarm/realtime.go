package alarm

import (
	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/gorm"
	"prometheus-manager/pkg/util/types"
)

const (
	RealtimeAssociationReplaceIntervenes = "AlarmIntervenes"
	RealtimeAssociationUpgradeInfo       = "AlarmUpgradeInfo"
	RealtimeAssociationSuppressInfo      = "AlarmSuppressInfo"
	RealtimeAssociationBeenNotifyMembers = "BeenNotifyMembers"
	RealtimeAssociationBeenChatGroups    = "BeenChatGroups"
)

// RealtimeLike 查询关键字
func RealtimeLike(keyword string) query.ScopeMethod {
	if keyword == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	return query.WhereLikeKeyword(keyword+"%", "note", "instance")
}

// RealtimeEventAtDesc 事件时间倒序
func RealtimeEventAtDesc() query.ScopeMethod {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("event_at DESC")
	}
}

// InIds 查询ID列表
func InIds[T types.Int](ids ...T) query.ScopeMethod {
	return query.WhereInColumn("id", ids)
}
