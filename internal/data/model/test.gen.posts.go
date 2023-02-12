package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PostsMgr struct {
	*_BaseMgr
}

// PostsMgr open func
func PostsMgr(db *gorm.DB) *_PostsMgr {
	if db == nil {
		panic(fmt.Errorf("PostsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PostsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("posts"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PostsMgr) GetTableName() string {
	return "posts"
}

// Reset 重置gorm会话
func (obj *_PostsMgr) Reset() *_PostsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PostsMgr) Get() (result Posts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PostsMgr) Gets() (results []*Posts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PostsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Posts{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PostsMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_PostsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_PostsMgr) GetByOption(opts ...Option) (result Posts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PostsMgr) GetByOptions(opts ...Option) (results []*Posts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_PostsMgr) GetFromID(id uint32) (result Posts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PostsMgr) GetBatchFromID(ids []uint32) (results []*Posts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_PostsMgr) GetFromName(name string) (results []*Posts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_PostsMgr) GetBatchFromName(names []string) (results []*Posts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PostsMgr) FetchByPrimaryKey(id uint32) (result Posts, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Posts{}).Where("`id` = ?", id).First(&result).Error

	return
}
