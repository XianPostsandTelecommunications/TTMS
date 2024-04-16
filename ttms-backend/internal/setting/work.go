/**
 * @Author: lenovo
 * @Description:
 * @File:  work
 * @Version: 1.0.0
 * @Date: 2023/05/29 11:21
 */

package setting

import (
	"mognolia/internal/global"
	"mognolia/internal/pkg/goroutine/work"
)

type worker struct {
}

func (w worker) Init() {
	global.Worker = work.Init(work.Config{
		TaskChanCapacity:   global.Settings.Work.TaskChanCapacity,
		WorkerChanCapacity: global.Settings.Work.WorkerChanCapacity,
		WorkerNum:          global.Settings.Work.WorkerNum,
	})
}
