package when

import (
	"regexp"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func lastWeek(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)\s*(last week)\s*(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			if c.Duration == 0 || overwrite {
				c.Duration = -(7 * 24 * time.Hour)
			}

			return true, nil
		},
	}
}

func lastMonth(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)\s*(last month)\s*(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			if c.Month == nil || overwrite {
				month := int(ref.Month()) - 1
				if month <= 0 {
					month += 12
				}
				c.Month = pointer.ToInt(month)
			}

			return true, nil
		},
	}
}
