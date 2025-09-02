package model

type AccountLogin struct {
	//不能有空格
	/**
	json:"AccountName"：指定JSON序列化时使用"AccountName"作为字段名

	binding:"..."：应该是数据验证标签（但存在拼写错误）

	required：表示该字段是必需的

	min = 8和max = 32：指定字段长度的最小和最大值
	*/
	AccountName string `json:"AccountName" binding:"required,min=16,max=32"`
	Password    string `json:"Password" binding:"required,min=16,max=32"`
}
