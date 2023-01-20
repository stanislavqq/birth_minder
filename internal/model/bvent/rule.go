package bevent

import "time"

type Rule struct {
	day   int
	month int
}

type Rules []Rule

func NewTimeRule(interval time.Duration) Rule {
	now := time.Now()
	timeTarget := now.Add(interval)
	return Rule{
		day:   timeTarget.Day(),
		month: int(timeTarget.Month()),
	}
}
