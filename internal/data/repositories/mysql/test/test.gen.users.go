package test

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UsersMgr struct {
	*_BaseMgr
}

// UsersMgr open func
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Reset 重置gorm会话
func (obj *_UsersMgr) Reset() *_UsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Users{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UsersMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_UsersMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAge age获取
func (obj *_UsersMgr) WithAge(age int8) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithNumberid numberID获取
func (obj *_UsersMgr) WithNumberid(numberid int) Option {
	return optionFunc(func(o *options) { o.query["numberID"] = numberid })
}

// WithSex sex获取
func (obj *_UsersMgr) WithSex(sex int8) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// GetByOption 功能选项模式获取
func (obj *_UsersMgr) GetByOption(opts ...Option) (result Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UsersMgr) GetFromID(id uint32) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UsersMgr) GetBatchFromID(ids []uint32) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_UsersMgr) GetFromName(name string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_UsersMgr) GetBatchFromName(names []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容
func (obj *_UsersMgr) GetFromAge(age int8) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`age` = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量查找
func (obj *_UsersMgr) GetBatchFromAge(ages []int8) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`age` IN (?)", ages).Find(&results).Error

	return
}

// GetFromNumberid 通过numberID获取内容
func (obj *_UsersMgr) GetFromNumberid(numberid int) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`numberID` = ?", numberid).First(&result).Error

	return
}

// GetBatchFromNumberid 批量查找
func (obj *_UsersMgr) GetBatchFromNumberid(numberids []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`numberID` IN (?)", numberids).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容
func (obj *_UsersMgr) GetFromSex(sex int8) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找
func (obj *_UsersMgr) GetBatchFromSex(sexs []int8) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UsersMgr) FetchByPrimaryKey(id uint32) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByNumberid primary or index 获取唯一内容
func (obj *_UsersMgr) FetchUniqueByNumberid(numberid int) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`numberID` = ?", numberid).First(&result).Error

	return
}

// FetchIndexByAge  获取多个内容
func (obj *_UsersMgr) FetchIndexByAge(age int8) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`age` = ?", age).Find(&results).Error

	return
}

// FetchIndexBySex  获取多个内容
func (obj *_UsersMgr) FetchIndexBySex(sex int8) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}
