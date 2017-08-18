package params_test

import (
	"encoding/json"
	"testing"

	"github.com/dailyburn/sailthru-go/params"
	"github.com/stretchr/testify/assert"
)

func TestMarshalScheduleTimeEmpty(t *testing.T) {
	schedule := &params.ScheduleTime{}
	bytes, err := schedule.MarshalJSON()
	assert.Equal(t, []byte(`""`), bytes)
	assert.NoError(t, err)
}

func TestMarshalScheduleTimeBlank(t *testing.T) {
	schedule := &params.ScheduleTime{StartTime: ""}
	bytes, err := json.Marshal(schedule)
	assert.Equal(t, []byte(`""`), bytes)
	assert.NoError(t, err)
}

func TestMarshalScheduleTimeStartTimeOnly(t *testing.T) {
	schedule := &params.ScheduleTime{
		StartTime: "Tue, 27 Jul 2010 12:10:00 -0400",
	}
	bytes, err := json.Marshal(schedule)
	expected := []byte(`"Tue, 27 Jul 2010 12:10:00 -0400"`)
	assert.Equal(t, expected, bytes)
	assert.NoError(t, err)
}

func TestMarshalScheduleTimeStartTimeAndEndTime(t *testing.T) {
	schedule := &params.ScheduleTime{
		StartTime: "Tue, 27 Jul 2010 12:10:00 -0400",
		EndTime:   "Wed, 28 Jul 2010 12:10:00 -0400",
	}
	bytes, err := json.Marshal(schedule)
	expected := []byte(`{"start_time":"Tue, 27 Jul 2010 12:10:00 -0400","end_time":"Wed, 28 Jul 2010 12:10:00 -0400"}`)
	assert.Equal(t, expected, bytes)
	assert.NoError(t, err)
}
