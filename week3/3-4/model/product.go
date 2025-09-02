package model

// 入参校验
type Product struct {
	ID   int    `uri:"id" json:"id" binding:required`
	Name string `uri:"name" json:"name" binding:required`
}
