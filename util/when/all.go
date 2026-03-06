package when

import (
	owhen "github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	en "github.com/olebedev/when/rules/en"
)

var All = []rules.Rule{
	lastWeek(rules.Override),
	lastMonth(rules.Override),
	startOfWeek(rules.Override),
	startOfMonth(rules.Override),
	startOfYear(rules.Override),
}

// New returns a when.When configured with the standard English rules
// and the project's custom rules (All).
func New() *owhen.Parser {
	w := owhen.New(nil)
	w.Add(en.All...)
	w.Add(All...)
	return w
}
