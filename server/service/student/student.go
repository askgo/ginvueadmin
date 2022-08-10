package student

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/student"
	studentReq "github.com/flipped-aurora/gin-vue-admin/server/model/student/request"
)

type StudentService struct {
}

// CreateStudent 创建Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) CreateStudent(student student.Student) (err error) {
	err = global.GVA_DB.Create(&student).Error
	return err
}

// DeleteStudent 删除Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) DeleteStudent(student student.Student) (err error) {
	err = global.GVA_DB.Delete(&student).Error
	return err
}

// DeleteStudentByIds 批量删除Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) DeleteStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]student.Student{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateStudent 更新Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) UpdateStudent(student student.Student) (err error) {
	err = global.GVA_DB.Save(&student).Error
	return err
}

// GetStudent 根据id获取Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) GetStudent(id uint) (student student.Student, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

// GetStudentInfoList 分页获取Student记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) GetStudentInfoList(info studentReq.StudentSearch) (list []student.Student, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&student.Student{})
	var students []student.Student
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&students).Error
	return students, total, err
}
