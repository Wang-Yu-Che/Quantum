// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserRelation = "user_relations"

// UserRelation mapped from table <user_relations>
type UserRelation struct {
	UserInfoID int64 `gorm:"column:user_info_id;type:bigint;primaryKey" json:"user_info_id"`
	FollowID   int64 `gorm:"column:follow_id;type:bigint;primaryKey" json:"follow_id"`
}

// TableName UserRelation's table name
func (*UserRelation) TableName() string {
	return TableNameUserRelation
}
