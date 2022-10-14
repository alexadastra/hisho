package models

import (
	"time"

	"github.com/pkg/errors"

	"github.com/alexadastra/hisho/hisho-core-service/pkg/api"
)

const (
	// SetDoneEvent : setting status from (not_done, closed) -> done
	SetDoneEvent ChangeStatusEvent = "set done"
	// CloseEvent : setting status from (not_done, done) -> closed
	CloseEvent ChangeStatusEvent = "close"
	// ReopenEvent : setting status from (closed, done) -> not_done
	ReopenEvent ChangeStatusEvent = "reopen"
	// LeaveUnchangedEvent : leaving status the same as it was
	LeaveUnchangedEvent ChangeStatusEvent = "leave unchanged"
	// UndefinedFromErrorMessage is `undefined 'from' status``
	UndefinedFromErrorMessage = "undefined 'from' status"
	// UndefinedToErrorMessage is `undefined 'to' status`
	UndefinedToErrorMessage = "undefined 'to' status"
	// UnsupportedStatusChangeErrorMessage is `unsupported status change from %s to %s`
	UnsupportedStatusChangeErrorMessage = "unsupported status change from %s to %s"
	// ClosingWithoutReasonErrorMessage is `trying to close task without reason`
	ClosingWithoutReasonErrorMessage = "trying to close task without reason"
)

// ChangeStatusEvent represents update action on task status (set from one state to another)
type ChangeStatusEvent string

// NewChangeStatusEvent creates new ChangeStatusEvent from request fields
func NewChangeStatusEvent(from, to api.Status, reason string) (ChangeStatusEvent, error) {
	if from == api.Status_STATUS_UNDEFINED {
		return "", errors.New(UndefinedFromErrorMessage)
	}
	if to == api.Status_STATUS_UNDEFINED {
		return "", errors.New(UndefinedToErrorMessage)
	}
	switch from {
	case api.Status_STATUS_NOT_DONE:
		return processFromNotDone(to, reason)
	case api.Status_STATUS_DONE:
		return processFromDone(to, reason)
	case api.Status_STATUS_CLOSED:
		return processFromClosed(to, reason)
	default:
		return "", errors.Errorf(UnsupportedStatusChangeErrorMessage, from.String(), to.String())
	}
}

func processFromNotDone(to api.Status, reason string) (ChangeStatusEvent, error) {
	switch to {
	case api.Status_STATUS_NOT_DONE:
		return LeaveUnchangedEvent, nil
	case api.Status_STATUS_DONE:
		return SetDoneEvent, nil
	case api.Status_STATUS_CLOSED:
		if reason == "" {
			return "", errors.New(ClosingWithoutReasonErrorMessage)
		}
		return CloseEvent, nil
	default:
		return "", errors.Errorf(UnsupportedStatusChangeErrorMessage, api.Status_STATUS_NOT_DONE.String(), to.String())
	}
}

func processFromDone(to api.Status, reason string) (ChangeStatusEvent, error) {
	switch to {
	case api.Status_STATUS_DONE:
		return LeaveUnchangedEvent, nil
	case api.Status_STATUS_NOT_DONE:
		return ReopenEvent, nil
	case api.Status_STATUS_CLOSED:
		if reason == "" {
			return "", errors.New(ClosingWithoutReasonErrorMessage)
		}
		return CloseEvent, nil
	default:
		return "", errors.Errorf(UnsupportedStatusChangeErrorMessage, api.Status_STATUS_DONE.String(), to.String())
	}
}

func processFromClosed(to api.Status, reason string) (ChangeStatusEvent, error) {
	switch to {
	case api.Status_STATUS_DONE:
		return SetDoneEvent, nil
	case api.Status_STATUS_NOT_DONE:
		return ReopenEvent, nil
	case api.Status_STATUS_CLOSED:
		return LeaveUnchangedEvent, nil
	default:
		return "", errors.Errorf(UnsupportedStatusChangeErrorMessage, api.Status_STATUS_CLOSED.String(), to.String())
	}
}

func newStatus(closedAt *time.Time, closedReason string) api.Status {
	if closedAt == nil {
		return api.Status_STATUS_NOT_DONE
	}
	if closedReason == "" {
		return api.Status_STATUS_DONE
	}
	return api.Status_STATUS_CLOSED
}
