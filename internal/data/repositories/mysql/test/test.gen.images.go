package test

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ImagesMgr struct {
	*_BaseMgr
}

// ImagesMgr open func
func ImagesMgr(db *gorm.DB) *_ImagesMgr {
	if db == nil {
		panic(fmt.Errorf("ImagesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImagesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("images"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImagesMgr) GetTableName() string {
	return "images"
}

// Reset 重置gorm会话
func (obj *_ImagesMgr) Reset() *_ImagesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ImagesMgr) Get() (result Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImagesMgr) Gets() (results []*Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ImagesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Images{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ImagesMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithURL url获取
func (obj *_ImagesMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithImageableID imageable_id获取
func (obj *_ImagesMgr) WithImageableID(imageableID uint32) Option {
	return optionFunc(func(o *options) { o.query["imageable_id"] = imageableID })
}

// WithImageableType imageable_type获取
func (obj *_ImagesMgr) WithImageableType(imageableType string) Option {
	return optionFunc(func(o *options) { o.query["imageable_type"] = imageableType })
}

// GetByOption 功能选项模式获取
func (obj *_ImagesMgr) GetByOption(opts ...Option) (result Images, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ImagesMgr) GetByOptions(opts ...Option) (results []*Images, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_ImagesMgr) GetFromID(id uint32) (result Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ImagesMgr) GetBatchFromID(ids []uint32) (results []*Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_ImagesMgr) GetFromURL(url string) (results []*Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找
func (obj *_ImagesMgr) GetBatchFromURL(urls []string) (results []*Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromImageableID 通过imageable_id获取内容
func (obj *_ImagesMgr) GetFromImageableID(imageableID uint32) (results []*Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`imageable_id` = ?", imageableID).Find(&results).Error

	return
}

// GetBatchFromImageableID 批量查找
func (obj *_ImagesMgr) GetBatchFromImageableID(imageableIDs []uint32) (results []*Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`imageable_id` IN (?)", imageableIDs).Find(&results).Error

	return
}

// GetFromImageableType 通过imageable_type获取内容
func (obj *_ImagesMgr) GetFromImageableType(imageableType string) (results []*Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`imageable_type` = ?", imageableType).Find(&results).Error

	return
}

// GetBatchFromImageableType 批量查找
func (obj *_ImagesMgr) GetBatchFromImageableType(imageableTypes []string) (results []*Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`imageable_type` IN (?)", imageableTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ImagesMgr) FetchByPrimaryKey(id uint32) (result Images, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Images{}).Where("`id` = ?", id).First(&result).Error

	return
}
