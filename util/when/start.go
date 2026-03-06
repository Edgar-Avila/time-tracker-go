package when

import (
	"regexp"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func startOfMonth(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)\s*(start of (?:the )?month)\s*(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			if c.Day == nil || overwrite {
				c.Day = pointer.ToInt(1)
			}

			if c.Hour == nil || overwrite {
				c.Hour = pointer.ToInt(0)
			}

			if c.Minute == nil || overwrite {
				c.Minute = pointer.ToInt(0)
			}

			if c.Second == nil || overwrite {
				c.Second = pointer.ToInt(0)
			}

			return true, nil
		},
	}
}

func startOfWeek(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)\s*(start of (?:the )?week)\s*(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			weekday := int(ref.Weekday())
			if weekday == 0 {
				weekday = 7 // Sunday -> 7
			}

			offset := -(weekday - 1)

			if c.Duration == 0 || overwrite {
				c.Duration = time.Duration(offset) * 24 * time.Hour
			}

			if c.Hour == nil || overwrite {
				c.Hour = pointer.ToInt(0)
			}

			if c.Minute == nil || overwrite {
				c.Minute = pointer.ToInt(0)
			}

			if c.Second == nil || overwrite {
				c.Second = pointer.ToInt(0)
			}

			return true, nil
		},
	}
}

func startOfYear(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)\s*(start of (?:the )?year)\s*(?:\W|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			if c.Month == nil || overwrite {
				c.Month = pointer.ToInt(1)
			}

			if c.Day == nil || overwrite {
				c.Day = pointer.ToInt(1)
			}

			if c.Hour == nil || overwrite {
				c.Hour = pointer.ToInt(0)
			}

			if c.Minute == nil || overwrite {
				c.Minute = pointer.ToInt(0)
			}

			if c.Second == nil || overwrite {
				c.Second = pointer.ToInt(0)
			}

			return true, nil
		},
	}
}
