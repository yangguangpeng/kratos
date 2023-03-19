package test

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _DcGoodsMgr struct {
	*_BaseMgr
}

// DcGoodsMgr open func
func DcGoodsMgr(db *gorm.DB) *_DcGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("DcGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DcGoodsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dc_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DcGoodsMgr) GetTableName() string {
	return "dc_goods"
}

// Reset 重置gorm会话
func (obj *_DcGoodsMgr) Reset() *_DcGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DcGoodsMgr) Get() (result DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DcGoodsMgr) Gets() (results []*DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DcGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DcGoodsMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGoodsName goods_name获取
func (obj *_DcGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithStatus status获取
func (obj *_DcGoodsMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_DcGoodsMgr) GetByOption(opts ...Option) (result DcGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DcGoodsMgr) GetByOptions(opts ...Option) (results []*DcGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DcGoodsMgr) GetFromID(id uint32) (result DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DcGoodsMgr) GetBatchFromID(ids []uint32) (results []*DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容
func (obj *_DcGoodsMgr) GetFromGoodsName(goodsName string) (results []*DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找
func (obj *_DcGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_DcGoodsMgr) GetFromStatus(status bool) (results []*DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_DcGoodsMgr) GetBatchFromStatus(statuss []bool) (results []*DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DcGoodsMgr) FetchByPrimaryKey(id uint32) (result DcGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DcGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}
