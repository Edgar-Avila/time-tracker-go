package when

import "github.com/olebedev/when/rules"

var All = []rules.Rule{
	lastWeek(rules.Override),
	lastMonth(rules.Override),
	startOfWeek(rules.Override),
	startOfMonth(rules.Override),
}
