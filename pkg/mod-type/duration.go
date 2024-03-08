package modtype

import "time"

type Duration time.Duration

func (m *Duration) UnmarshalFlag(value string) error {
	return m.unmarshal(value)
}

func (m *Duration) unmarshal(value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*m = Duration(duration)
	return nil
}

func (m *Duration) TimeDuration() time.Duration {
	return time.Duration(*m)
}

func (m *Duration) Decode(s string) error {
	return m.unmarshal(s)
}
