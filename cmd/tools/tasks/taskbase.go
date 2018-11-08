package tasks

import "gitlab.33.cn/chain33/chain33/common/log/log15"

var (
	mlog = log15.New("module", "task")
)

type TaskBase struct {
	NextTask Task
}

func (this *TaskBase) Execute() error {
	return nil
}

func (this *TaskBase) SetNext(t Task) {
	this.NextTask = t
}

func (this *TaskBase) Next() Task {
	return this.NextTask
}
