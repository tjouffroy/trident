// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AntiRansomwareState Anti-ransomware state.<br>`disabled` Anti-ransomware monitoring is disabled on the volume.  This is the default state in a POST operation.<br>`disable_in_progress` Anti-ransomware monitoring is being disabled and a cleanup operation is in effect. Valid in GET operation.<br>`dry_run` Anti-ransomware monitoring is enabled in the evaluation mode.<br>`enabled` Anti-ransomware monitoring is active on the volume.<br>`paused` Anti-ransomware monitoring is paused on the volume.<br>`enable_paused` Anti-ransomware monitoring is paused on the volume from its earlier enabled state. Valid in GET operation. <br>`dry_run_paused` Anti-ransomware monitoring is paused on the volume from its earlier dry_run state. Valid in GET operation. <br>For POST, the valid Anti-ransomware states are only `disabled`, `enabled` and `dry_run`, whereas for PATCH, `paused` is also valid along with the three valid states for POST.
//
// swagger:model anti_ransomware_state
type AntiRansomwareState string

func NewAntiRansomwareState(value AntiRansomwareState) *AntiRansomwareState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AntiRansomwareState.
func (m AntiRansomwareState) Pointer() *AntiRansomwareState {
	return &m
}

const (

	// AntiRansomwareStateDisabled captures enum value "disabled"
	AntiRansomwareStateDisabled AntiRansomwareState = "disabled"

	// AntiRansomwareStateDisableInProgress captures enum value "disable_in_progress"
	AntiRansomwareStateDisableInProgress AntiRansomwareState = "disable_in_progress"

	// AntiRansomwareStateDryRun captures enum value "dry_run"
	AntiRansomwareStateDryRun AntiRansomwareState = "dry_run"

	// AntiRansomwareStateEnabled captures enum value "enabled"
	AntiRansomwareStateEnabled AntiRansomwareState = "enabled"

	// AntiRansomwareStatePaused captures enum value "paused"
	AntiRansomwareStatePaused AntiRansomwareState = "paused"

	// AntiRansomwareStateEnablePaused captures enum value "enable_paused"
	AntiRansomwareStateEnablePaused AntiRansomwareState = "enable_paused"

	// AntiRansomwareStateDryRunPaused captures enum value "dry_run_paused"
	AntiRansomwareStateDryRunPaused AntiRansomwareState = "dry_run_paused"
)

// for schema
var antiRansomwareStateEnum []interface{}

func init() {
	var res []AntiRansomwareState
	if err := json.Unmarshal([]byte(`["disabled","disable_in_progress","dry_run","enabled","paused","enable_paused","dry_run_paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		antiRansomwareStateEnum = append(antiRansomwareStateEnum, v)
	}
}

func (m AntiRansomwareState) validateAntiRansomwareStateEnum(path, location string, value AntiRansomwareState) error {
	if err := validate.EnumCase(path, location, value, antiRansomwareStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this anti ransomware state
func (m AntiRansomwareState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAntiRansomwareStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this anti ransomware state based on context it is used
func (m AntiRansomwareState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
