package jobs

import (
	"go_cron/models"
	"go_cron/pkg/logging"
	"os/exec"
	"time"
)

func InitJobs() {
	list, _ := models.ListTaskTotal()
	for _, task := range list {
		job, err := NewJobFromTask(&task)
		if err != nil {
			logging.Error("InitJobs:", err.Error())
			//beego.Error("InitJobs:", err.Error())
			continue
		}
		AddJob(task.CronSpec, job)
	}
}

func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		logging.Warn("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid)
		//beego.Warn(fmt.Sprintf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid))
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			logging.Error("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err)
			//beego.Error(fmt.Sprintf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err))
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
