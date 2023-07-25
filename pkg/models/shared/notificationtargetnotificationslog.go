// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type NotificationTargetNotificationsLogType string

const (
	NotificationTargetNotificationsLogTypeNotificationsLog NotificationTargetNotificationsLogType = "notifications_log"
)

func (e NotificationTargetNotificationsLogType) ToPointer() *NotificationTargetNotificationsLogType {
	return &e
}

func (e *NotificationTargetNotificationsLogType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "notifications_log":
		*e = NotificationTargetNotificationsLogType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NotificationTargetNotificationsLogType: %v", v)
	}
}

type NotificationTargetNotificationsLog struct {
	// Unique ID for this output
	ID string `json:"id"`
	// Directory in which to store the notification log
	LogsDir *string `json:"logsDir,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string                               `json:"systemFields,omitempty"`
	Type         NotificationTargetNotificationsLogType `json:"type"`
}

func (o *NotificationTargetNotificationsLog) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *NotificationTargetNotificationsLog) GetLogsDir() *string {
	if o == nil {
		return nil
	}
	return o.LogsDir
}

func (o *NotificationTargetNotificationsLog) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *NotificationTargetNotificationsLog) GetType() NotificationTargetNotificationsLogType {
	if o == nil {
		return NotificationTargetNotificationsLogType("")
	}
	return o.Type
}
