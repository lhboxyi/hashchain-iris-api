package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IStuService interface {
	GetStuHello() string
	GetStuHello2() string
}

type Stu struct {
	Id            uint `gorm:"primary_key"`
	Name, Address string
}

func (stu *Stu) GetStuHello() string {
	fmt.Println("stu实现了接口服务")
	return ""
}

/**
存储任务信息结构体
*/
type Task struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	NodeIp     string             `json:"node_ip" bson:"node_ip"`         //节点ip
	JobName    string             `json:"job_name" bson:"job_name"`       //job名称
	StartDate  string             `json:"start_date" bson:"start_date"`   //job运行开始时间
	EndDate    string             `json:"end_date" bson:"end_date"`       //job运行结束时间
	RunStatus  string             `json:"run_status" bson:"run_status"`   //job的运行状态：online stopped errored
	LogPath    string             `json:"log_path" bson:"log_path"`       //job运行日志路径
	DeleteFlag int8               `json:"delete_flag" bson:"delete_flag"` //删除标记 0:未删除 1:删除
}

//func (stu Stu) GetStuInfo() {
//	fmt.Println("调用无参方法")
//}
//
//func (stu Stu) SumStu(num int) (sum int) {
//	for i := 0; i <= num; i++ {
//		sum += i
//	}
//	fmt.Println("调用有参方法！！")
//	return
//}
