package constants

const (
	StatusError = "error"
	Windows     = "windows"
	Linux       = "linux"
	Darwin      = "darwin"

	//项目用的mongo表名
	NodeTable = "c_nodes"
	JobTable  = "c_jobs"
	TaskTable = "c_tasks"

	//源代码上传方式
	GitTag = "git"
	ZipTag = "zip"

	// 自定义返回值状态码
	SuccessCode = "S200"
	FailCode    = "F500"
)
