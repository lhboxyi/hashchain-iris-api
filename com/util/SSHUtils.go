package util

import (
	"github.com/sirupsen/logrus"
	"iris-api/com/constants"
	"os/exec"
	"runtime"
)

/**
执行shell命令
参数：cmdStr:需要执行的命令，cwd:需要执行的目录
返回值：错误信息和命令执行结果
*/
func ExecuteShellCmd(cmdStr string, cwd string) (err error, res string) {
	logrus.Infof("cmd：%s,cwd: %s", cmdStr, cwd)
	// 生成执行命令
	var cmd *exec.Cmd
	if runtime.GOOS == constants.Windows {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("sh", "-c", cmdStr)
	}
	// 工作目录
	cmd.Dir = cwd

	bytes, err := cmd.Output()
	if err != nil {
		return err, ""
	}
	return nil, string(bytes)
}
