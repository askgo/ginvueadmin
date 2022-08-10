// 自动生成模板Student
package student

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Student 结构体
type Student struct {
      global.GVA_MODEL
      Age  *int `json:"age" form:"age" gorm:"column:age;comment:;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Sex  *int `json:"sex" form:"sex" gorm:"column:sex;comment:;"`
}


// TableName Student 表名
func (Student) TableName() string {
  return "student"
}

