package test

// DcGoods [...]
type DcGoods struct {
	ID        uint32 `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null"`
	GoodsName string `gorm:"column:goods_name;type:varchar(255);not null"`
	Status    bool   `gorm:"column:status;type:tinyint(1);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *DcGoods) TableName() string {
	return "dc_goods"
}

// Images [...]
type Images struct {
	ID            uint32 `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null"`
	URL           string `gorm:"column:url;type:varchar(255);not null"`
	ImageableID   uint32 `gorm:"column:imageable_id;type:int(10) unsigned;not null"`
	ImageableType string `gorm:"column:imageable_type;type:varchar(255);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *Images) TableName() string {
	return "images"
}

// Posts [...]
type Posts struct {
	ID   uint32 `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null"`
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *Posts) TableName() string {
	return "posts"
}

// Users [...]
type Users struct {
	ID       uint32 `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null"`
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	Age      int8   `gorm:"index:age;column:age;type:tinyint(2);not null"`
	Numberid int    `gorm:"unique;column:numberID;type:int(11);not null"`
	Sex      int8   `gorm:"index:sex;column:sex;type:tinyint(4);default:null"`
}

// TableName get sql table name.获取数据库表名
func (m *Users) TableName() string {
	return "users"
}
