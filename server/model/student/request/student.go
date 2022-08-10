package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/student"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type StudentSearch struct{
    student.Student
    request.PageInfo
}
