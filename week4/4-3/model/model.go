package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// TableName 指定表名为 "Product"

/**
具体功能说明：
覆盖默认命名策略：

默认情况下，GORM 会自动将结构体名称转换为复数形式作为表名

此方法强制指定表名为 "Product"（单数形式），而不是默认的 "products"

方法签名说明：

func (Product) TableName() string：这是一个值接收者方法，属于 Product 结构体

方法返回一个字符串，即希望映射的数据库表名

GORM 的约定：

GORM 会检查模型是否实现了 TableName() 方法

如果存在此方法，GORM 将使用其返回值作为表名

如果不存在，GORM 将使用默认的命名策略（复数化）
*/
/*func (Product) TableName() string {

	return "Product_test"
}*/
