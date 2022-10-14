package models

import (
	"fmt"
	"testing"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/alexadastra/hisho/hisho-core-service/pkg/api"
)

func TestNewChangeStatusEvent(t *testing.T) {
	type args struct {
		from   api.Status
		to     api.Status
		reason string
	}
	tests := []struct {
		name        string
		args        args
		expected    models.ChangeStatusEvent
		wantErr     bool
		expectedErr error
	}{
		{
			name: "from undefined",
			args: args{
				from: api.Status_STATUS_UNDEFINED,
				to:   api.Status_STATUS_CLOSED,
			},
			expected:    "",
			wantErr:     true,
			expectedErr: fmt.Errorf(models.UndefinedFromErrorMessage),
		},
		{
			name: "to undefined",
			args: args{
				from: api.Status_STATUS_NOT_DONE,
				to:   api.Status_STATUS_UNDEFINED,
			},
			expected:    "",
			wantErr:     true,
			expectedErr: fmt.Errorf(models.UndefinedToErrorMessage),
		},
		{
			name: "undone->done",
			args: args{
				from: api.Status_STATUS_NOT_DONE,
				to:   api.Status_STATUS_DONE,
			},
			expected: models.SetDoneEvent,
		},
		{
			name: "undone->closed without reason",
			args: args{
				from: api.Status_STATUS_NOT_DONE,
				to:   api.Status_STATUS_CLOSED,
			},
			expected:    "",
			wantErr:     true,
			expectedErr: fmt.Errorf(models.ClosingWithoutReasonErrorMessage),
		},
		{
			name: "undone->closed with reason",
			args: args{
				from:   api.Status_STATUS_NOT_DONE,
				to:     api.Status_STATUS_CLOSED,
				reason: "irrelevant",
			},
			expected: models.CloseEvent,
		},
		{
			name: "undone->undone",
			args: args{
				from: api.Status_STATUS_NOT_DONE,
				to:   api.Status_STATUS_NOT_DONE,
			},
			expected: models.LeaveUnchangedEvent,
		},
		{
			name: "done->done",
			args: args{
				from: api.Status_STATUS_DONE,
				to:   api.Status_STATUS_DONE,
			},
			expected: models.LeaveUnchangedEvent,
		},
		{
			name: "done->closed without reason",
			args: args{
				from: api.Status_STATUS_DONE,
				to:   api.Status_STATUS_CLOSED,
			},
			expected:    "",
			wantErr:     true,
			expectedErr: fmt.Errorf(models.ClosingWithoutReasonErrorMessage),
		},
		{
			name: "done->closed with reason",
			args: args{
				from:   api.Status_STATUS_DONE,
				to:     api.Status_STATUS_CLOSED,
				reason: "irrelevant",
			},
			expected: models.CloseEvent,
		},
		{
			name: "done->undone",
			args: args{
				from: api.Status_STATUS_DONE,
				to:   api.Status_STATUS_NOT_DONE,
			},
			expected: models.ReopenEvent,
		},
		{
			name: "closed->done",
			args: args{
				from: api.Status_STATUS_CLOSED,
				to:   api.Status_STATUS_DONE,
			},
			expected: models.SetDoneEvent,
		},
		{
			name: "closed->closed",
			args: args{
				from: api.Status_STATUS_CLOSED,
				to:   api.Status_STATUS_CLOSED,
			},
			expected: models.LeaveUnchangedEvent,
		},
		{
			name: "closed->undone",
			args: args{
				from: api.Status_STATUS_CLOSED,
				to:   api.Status_STATUS_NOT_DONE,
			},
			expected: models.ReopenEvent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := models.NewChangeStatusEvent(tt.args.from, tt.args.to, tt.args.reason)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewChangeStatusEvent() actual_error = `%s`, wantErr %v", err.Error(), tt.wantErr)
					return
				}
				if err.Error() != tt.expectedErr.Error() {
					t.Errorf("NewChangeStatusEvent() actual_error = `%s`, expected_error =  `%s`", err.Error(), tt.expectedErr.Error())
					return
				}
			}
			if got != tt.expected {
				t.Errorf("NewChangeStatusEvent() = %v, want %v", got, tt.expected)
			}
		})
	}
}
