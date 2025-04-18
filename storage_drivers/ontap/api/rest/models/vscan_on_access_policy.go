// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VscanOnAccessPolicy An On-Access policy that defines the scope of an On-Access scan. Use On-Access scanning to check for viruses when clients open, read, rename, or close files over CIFS. By default, ONTAP creates an On-Access policy named "default_CIFS" and enables it for all the SVMs in a cluster.
//
// swagger:model vscan_on_access_policy
type VscanOnAccessPolicy struct {

	// Status of the On-Access Vscan policy
	Enabled *bool `json:"enabled,omitempty"`

	// Specifies if scanning is mandatory. File access is denied if there are no external virus-scanning servers available for virus scanning.
	Mandatory *bool `json:"mandatory,omitempty"`

	// On-Access policy name
	// Example: on-access-test
	// Max Length: 256
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// scope
	Scope *VscanOnAccessPolicyInlineScope `json:"scope,omitempty"`
}

// Validate validates this vscan on access policy
func (m *VscanOnAccessPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnAccessPolicy) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	return nil
}

func (m *VscanOnAccessPolicy) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vscan on access policy based on the context it is used
func (m *VscanOnAccessPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnAccessPolicy) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {
		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VscanOnAccessPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VscanOnAccessPolicy) UnmarshalBinary(b []byte) error {
	var res VscanOnAccessPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VscanOnAccessPolicyInlineScope vscan on access policy inline scope
//
// swagger:model vscan_on_access_policy_inline_scope
type VscanOnAccessPolicyInlineScope struct {

	// List of file extensions for which scanning is not performed.
	// Example: ["mp*","txt"]
	// Max Items: 300
	// Min Items: 1
	ExcludeExtensions []*string `json:"exclude_extensions,omitempty"`

	// List of file paths for which scanning must not be performed.
	// Example: ["\\dir1\\dir2\\name","\\vol\\a b","\\vol\\a,b\\"]
	// Max Items: 100
	// Min Items: 1
	ExcludePaths []*string `json:"exclude_paths,omitempty"`

	// List of file extensions to be scanned.
	// Example: ["mp*","txt"]
	// Max Items: 300
	// Min Items: 1
	IncludeExtensions []*string `json:"include_extensions,omitempty"`

	// Maximum file size, in bytes, allowed for scanning.
	// Example: 2147483648
	// Maximum: 1.099511627776e+12
	// Minimum: 1024
	MaxFileSize *int64 `json:"max_file_size,omitempty"`

	// Scan only files opened with execute-access.
	OnlyExecuteAccess *bool `json:"only_execute_access,omitempty"`

	// Specifies whether or not read-only volume can be scanned.
	ScanReadonlyVolumes *bool `json:"scan_readonly_volumes,omitempty"`

	// Specifies whether or not files without any extension can be scanned.
	ScanWithoutExtension *bool `json:"scan_without_extension,omitempty"`
}

// Validate validates this vscan on access policy inline scope
func (m *VscanOnAccessPolicyInlineScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExcludeExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludePaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxFileSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VscanOnAccessPolicyInlineScope) validateExcludeExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeExtensions) { // not required
		return nil
	}

	iExcludeExtensionsSize := int64(len(m.ExcludeExtensions))

	if err := validate.MinItems("scope"+"."+"exclude_extensions", "body", iExcludeExtensionsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("scope"+"."+"exclude_extensions", "body", iExcludeExtensionsSize, 300); err != nil {
		return err
	}

	return nil
}

func (m *VscanOnAccessPolicyInlineScope) validateExcludePaths(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludePaths) { // not required
		return nil
	}

	iExcludePathsSize := int64(len(m.ExcludePaths))

	if err := validate.MinItems("scope"+"."+"exclude_paths", "body", iExcludePathsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("scope"+"."+"exclude_paths", "body", iExcludePathsSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(m.ExcludePaths); i++ {
		if swag.IsZero(m.ExcludePaths[i]) { // not required
			continue
		}

		if err := validate.MinLength("scope"+"."+"exclude_paths"+"."+strconv.Itoa(i), "body", *m.ExcludePaths[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength("scope"+"."+"exclude_paths"+"."+strconv.Itoa(i), "body", *m.ExcludePaths[i], 255); err != nil {
			return err
		}

	}

	return nil
}

func (m *VscanOnAccessPolicyInlineScope) validateIncludeExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludeExtensions) { // not required
		return nil
	}

	iIncludeExtensionsSize := int64(len(m.IncludeExtensions))

	if err := validate.MinItems("scope"+"."+"include_extensions", "body", iIncludeExtensionsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("scope"+"."+"include_extensions", "body", iIncludeExtensionsSize, 300); err != nil {
		return err
	}

	return nil
}

func (m *VscanOnAccessPolicyInlineScope) validateMaxFileSize(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxFileSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("scope"+"."+"max_file_size", "body", *m.MaxFileSize, 1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("scope"+"."+"max_file_size", "body", *m.MaxFileSize, 1.099511627776e+12, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vscan on access policy inline scope based on context it is used
func (m *VscanOnAccessPolicyInlineScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VscanOnAccessPolicyInlineScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VscanOnAccessPolicyInlineScope) UnmarshalBinary(b []byte) error {
	var res VscanOnAccessPolicyInlineScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
