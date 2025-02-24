package datetime_test

import (
	"backend/internal/datetime"
	"testing"
	"time"
)

func TestFromTime(t *testing.T) {

	knownTime := time.Date(2024, 6, 14, 12, 34, 56, 789000000, time.UTC)
	expectedMinute := knownTime.Minute()

	// Execution: Convert the known time to a datetime object.
	actualDatetime := datetime.Must(datetime.Time(knownTime))

	// Assertions: Check if the conversion was successful and the minute matches.
	if actualDatetime.ToTime().IsZero() {
		t.Fatalf("Expected non-zero datetime, got zero: %v", actualDatetime)
	}

	if expectedMinute != actualDatetime.ToTime().Minute() {
		t.Errorf("Expected minute %d, got %d", expectedMinute, actualDatetime.ToTime().Minute())
	}
}

func TestFormDate(t *testing.T) {
	mockDate := time.Date(2024, 6, 14, 0, 0, 0, 0, datetime.GetHongKongTimeLocation())

	mockDateString := mockDate.Format(datetime.FormatDateTime)

	testDateString := datetime.MustDateString("2024-06-14")

	if mockDateString != testDateString {
		t.Errorf("Expected date %s, got %s", mockDate, testDateString)
	}
}
