// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package douyin_query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"Quantum/internal/model/douyin_model/model"
)

func newUserFavorVideo(db *gorm.DB, opts ...gen.DOOption) userFavorVideo {
	_userFavorVideo := userFavorVideo{}

	_userFavorVideo.userFavorVideoDo.UseDB(db, opts...)
	_userFavorVideo.userFavorVideoDo.UseModel(&model.UserFavorVideo{})

	tableName := _userFavorVideo.userFavorVideoDo.TableName()
	_userFavorVideo.ALL = field.NewAsterisk(tableName)
	_userFavorVideo.UserInfoID = field.NewInt64(tableName, "user_info_id")
	_userFavorVideo.VideoID = field.NewInt64(tableName, "video_id")

	_userFavorVideo.fillFieldMap()

	return _userFavorVideo
}

type userFavorVideo struct {
	userFavorVideoDo

	ALL        field.Asterisk
	UserInfoID field.Int64
	VideoID    field.Int64

	fieldMap map[string]field.Expr
}

func (u userFavorVideo) Table(newTableName string) *userFavorVideo {
	u.userFavorVideoDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userFavorVideo) As(alias string) *userFavorVideo {
	u.userFavorVideoDo.DO = *(u.userFavorVideoDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userFavorVideo) updateTableName(table string) *userFavorVideo {
	u.ALL = field.NewAsterisk(table)
	u.UserInfoID = field.NewInt64(table, "user_info_id")
	u.VideoID = field.NewInt64(table, "video_id")

	u.fillFieldMap()

	return u
}

func (u *userFavorVideo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userFavorVideo) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 2)
	u.fieldMap["user_info_id"] = u.UserInfoID
	u.fieldMap["video_id"] = u.VideoID
}

func (u userFavorVideo) clone(db *gorm.DB) userFavorVideo {
	u.userFavorVideoDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userFavorVideo) replaceDB(db *gorm.DB) userFavorVideo {
	u.userFavorVideoDo.ReplaceDB(db)
	return u
}

type userFavorVideoDo struct{ gen.DO }

type IUserFavorVideoDo interface {
	gen.SubQuery
	Debug() IUserFavorVideoDo
	WithContext(ctx context.Context) IUserFavorVideoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserFavorVideoDo
	WriteDB() IUserFavorVideoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserFavorVideoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserFavorVideoDo
	Not(conds ...gen.Condition) IUserFavorVideoDo
	Or(conds ...gen.Condition) IUserFavorVideoDo
	Select(conds ...field.Expr) IUserFavorVideoDo
	Where(conds ...gen.Condition) IUserFavorVideoDo
	Order(conds ...field.Expr) IUserFavorVideoDo
	Distinct(cols ...field.Expr) IUserFavorVideoDo
	Omit(cols ...field.Expr) IUserFavorVideoDo
	Join(table schema.Tabler, on ...field.Expr) IUserFavorVideoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserFavorVideoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserFavorVideoDo
	Group(cols ...field.Expr) IUserFavorVideoDo
	Having(conds ...gen.Condition) IUserFavorVideoDo
	Limit(limit int) IUserFavorVideoDo
	Offset(offset int) IUserFavorVideoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserFavorVideoDo
	Unscoped() IUserFavorVideoDo
	Create(values ...*model.UserFavorVideo) error
	CreateInBatches(values []*model.UserFavorVideo, batchSize int) error
	Save(values ...*model.UserFavorVideo) error
	First() (*model.UserFavorVideo, error)
	Take() (*model.UserFavorVideo, error)
	Last() (*model.UserFavorVideo, error)
	Find() ([]*model.UserFavorVideo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFavorVideo, err error)
	FindInBatches(result *[]*model.UserFavorVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserFavorVideo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserFavorVideoDo
	Assign(attrs ...field.AssignExpr) IUserFavorVideoDo
	Joins(fields ...field.RelationField) IUserFavorVideoDo
	Preload(fields ...field.RelationField) IUserFavorVideoDo
	FirstOrInit() (*model.UserFavorVideo, error)
	FirstOrCreate() (*model.UserFavorVideo, error)
	FindByPage(offset int, limit int) (result []*model.UserFavorVideo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserFavorVideoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userFavorVideoDo) Debug() IUserFavorVideoDo {
	return u.withDO(u.DO.Debug())
}

func (u userFavorVideoDo) WithContext(ctx context.Context) IUserFavorVideoDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userFavorVideoDo) ReadDB() IUserFavorVideoDo {
	return u.Clauses(dbresolver.Read)
}

func (u userFavorVideoDo) WriteDB() IUserFavorVideoDo {
	return u.Clauses(dbresolver.Write)
}

func (u userFavorVideoDo) Session(config *gorm.Session) IUserFavorVideoDo {
	return u.withDO(u.DO.Session(config))
}

func (u userFavorVideoDo) Clauses(conds ...clause.Expression) IUserFavorVideoDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userFavorVideoDo) Returning(value interface{}, columns ...string) IUserFavorVideoDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userFavorVideoDo) Not(conds ...gen.Condition) IUserFavorVideoDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userFavorVideoDo) Or(conds ...gen.Condition) IUserFavorVideoDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userFavorVideoDo) Select(conds ...field.Expr) IUserFavorVideoDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userFavorVideoDo) Where(conds ...gen.Condition) IUserFavorVideoDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userFavorVideoDo) Order(conds ...field.Expr) IUserFavorVideoDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userFavorVideoDo) Distinct(cols ...field.Expr) IUserFavorVideoDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userFavorVideoDo) Omit(cols ...field.Expr) IUserFavorVideoDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userFavorVideoDo) Join(table schema.Tabler, on ...field.Expr) IUserFavorVideoDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userFavorVideoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserFavorVideoDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userFavorVideoDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserFavorVideoDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userFavorVideoDo) Group(cols ...field.Expr) IUserFavorVideoDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userFavorVideoDo) Having(conds ...gen.Condition) IUserFavorVideoDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userFavorVideoDo) Limit(limit int) IUserFavorVideoDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userFavorVideoDo) Offset(offset int) IUserFavorVideoDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userFavorVideoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserFavorVideoDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userFavorVideoDo) Unscoped() IUserFavorVideoDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userFavorVideoDo) Create(values ...*model.UserFavorVideo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userFavorVideoDo) CreateInBatches(values []*model.UserFavorVideo, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userFavorVideoDo) Save(values ...*model.UserFavorVideo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userFavorVideoDo) First() (*model.UserFavorVideo, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorVideo), nil
	}
}

func (u userFavorVideoDo) Take() (*model.UserFavorVideo, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorVideo), nil
	}
}

func (u userFavorVideoDo) Last() (*model.UserFavorVideo, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorVideo), nil
	}
}

func (u userFavorVideoDo) Find() ([]*model.UserFavorVideo, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserFavorVideo), err
}

func (u userFavorVideoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFavorVideo, err error) {
	buf := make([]*model.UserFavorVideo, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userFavorVideoDo) FindInBatches(result *[]*model.UserFavorVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userFavorVideoDo) Attrs(attrs ...field.AssignExpr) IUserFavorVideoDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userFavorVideoDo) Assign(attrs ...field.AssignExpr) IUserFavorVideoDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userFavorVideoDo) Joins(fields ...field.RelationField) IUserFavorVideoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userFavorVideoDo) Preload(fields ...field.RelationField) IUserFavorVideoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userFavorVideoDo) FirstOrInit() (*model.UserFavorVideo, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorVideo), nil
	}
}

func (u userFavorVideoDo) FirstOrCreate() (*model.UserFavorVideo, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFavorVideo), nil
	}
}

func (u userFavorVideoDo) FindByPage(offset int, limit int) (result []*model.UserFavorVideo, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userFavorVideoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userFavorVideoDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userFavorVideoDo) Delete(models ...*model.UserFavorVideo) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userFavorVideoDo) withDO(do gen.Dao) *userFavorVideoDo {
	u.DO = *do.(*gen.DO)
	return u
}
