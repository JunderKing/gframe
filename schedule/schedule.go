package schedule

import (
	"github.com/robfig/cron"
)

func Setup() {
	c := cron.New()
	setTask(c, "* * * * * *", ApiSign)
	c.Start()
}

func setTask(c *cron.Cron, spec string, handler func()) {
	if err := c.AddFunc(spec, handler); err != nil {
		panic(err)
	}
}