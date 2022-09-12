package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toTimestampb(t *time.Time) *timestamppb.Timestamp {
	if t != nil {
		return timestamppb.New(*t)
	}
	return nil
}

func fromTimestampb(timestamp *timestamppb.Timestamp) *time.Time {
	if timestamp != nil {
		tmp := timestamp.AsTime()
		return &tmp
	}
	return nil
}
