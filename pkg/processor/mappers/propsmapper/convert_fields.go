package propsmapper

import "time"

func convertFields[T any](val interface{}, defaultValue T) *T {
	result, ok := val.(T)
	if !ok {
		return &defaultValue
	}
	return &result
}

func factoryTime(date, layout string, defaultTime time.Time) *time.Time {
	newTime, err := time.Parse(layout, date)
	if err != nil {
		return &defaultTime
	}
	return &newTime
}
