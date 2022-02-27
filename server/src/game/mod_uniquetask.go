package game

import "sync"

type TaskInfo struct {
	TaskId int
	State  int
}

type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo
	Locker     *sync.RWMutex
}

func (self *ModUniqueTask) IsTaskFinish(taskId int) bool {
	if taskId == 10001 || taskId == 10002 || taskId == 10003 {

		return true
	}

	task, ok := self.MyTaskInfo[taskId]
	if !ok {
		return false
	}
	return task.State == TASK_STATE_FINISH
}
