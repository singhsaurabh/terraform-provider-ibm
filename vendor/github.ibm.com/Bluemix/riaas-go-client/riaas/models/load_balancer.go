// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoadBalancer load balancer
// swagger:model LoadBalancer
type LoadBalancer struct {

	// The date and time that this load balancer was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The CRN for this load balancer
	Crn string `json:"crn,omitempty"`

	// Fully qualified domain name assigned to this load balancer
	Hostname string `json:"hostname,omitempty"`

	// The load balancer's canonical URL.
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The load balancer's unique identifier
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The type of this load balancer, public or private
	IsPublic bool `json:"is_public,omitempty"`

	// The listeners of this load balancer
	Listeners []*ListenerReference `json:"listeners"`

	// The load balancer's user-defined name. Load balancer names must be unique. within the scope of a user account. Load balancer names are part of FQDN auto-assigned to the load balancer.
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// The operating status of this load balancer
	// Enum: [online offline]
	OperatingStatus string `json:"operating_status,omitempty"`

	// The pools of this load balancer
	Pools []*PoolReference `json:"pools"`

	// The private IP addresses assigned to this load balancer.
	PrivateIps []*LoadBalancerIP `json:"private_ips"`

	// The provisioning status of this load balancer
	// Enum: [active create_pending update_pending delete_pending maintenance_pending]
	ProvisioningStatus string `json:"provisioning_status,omitempty"`

	// The public IP addresses assigned to this load balancer. These are applicable only for public load balancers.
	PublicIps []*LoadBalancerIP `json:"public_ips"`

	// resource group
	ResourceGroup *ResourceReference `json:"resource_group,omitempty"`

	// SubnetReference
	Subnets []*LoadBalancerSubnetsItems0 `json:"subnets"`
}

// Validate validates this load balancer
func (m *LoadBalancer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListeners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePools(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisioningStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancer) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoadBalancer) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *LoadBalancer) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoadBalancer) validateListeners(formats strfmt.Registry) error {

	if swag.IsZero(m.Listeners) { // not required
		return nil
	}

	for i := 0; i < len(m.Listeners); i++ {
		if swag.IsZero(m.Listeners[i]) { // not required
			continue
		}

		if m.Listeners[i] != nil {
			if err := m.Listeners[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listeners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadBalancer) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

var loadBalancerTypeOperatingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","offline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loadBalancerTypeOperatingStatusPropEnum = append(loadBalancerTypeOperatingStatusPropEnum, v)
	}
}

const (

	// LoadBalancerOperatingStatusOnline captures enum value "online"
	LoadBalancerOperatingStatusOnline string = "online"

	// LoadBalancerOperatingStatusOffline captures enum value "offline"
	LoadBalancerOperatingStatusOffline string = "offline"
)

// prop value enum
func (m *LoadBalancer) validateOperatingStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, loadBalancerTypeOperatingStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LoadBalancer) validateOperatingStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.OperatingStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatingStatusEnum("operating_status", "body", m.OperatingStatus); err != nil {
		return err
	}

	return nil
}

func (m *LoadBalancer) validatePools(formats strfmt.Registry) error {

	if swag.IsZero(m.Pools) { // not required
		return nil
	}

	for i := 0; i < len(m.Pools); i++ {
		if swag.IsZero(m.Pools[i]) { // not required
			continue
		}

		if m.Pools[i] != nil {
			if err := m.Pools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadBalancer) validatePrivateIps(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateIps) { // not required
		return nil
	}

	for i := 0; i < len(m.PrivateIps); i++ {
		if swag.IsZero(m.PrivateIps[i]) { // not required
			continue
		}

		if m.PrivateIps[i] != nil {
			if err := m.PrivateIps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("private_ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var loadBalancerTypeProvisioningStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","create_pending","update_pending","delete_pending","maintenance_pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loadBalancerTypeProvisioningStatusPropEnum = append(loadBalancerTypeProvisioningStatusPropEnum, v)
	}
}

const (

	// LoadBalancerProvisioningStatusActive captures enum value "active"
	LoadBalancerProvisioningStatusActive string = "active"

	// LoadBalancerProvisioningStatusCreatePending captures enum value "create_pending"
	LoadBalancerProvisioningStatusCreatePending string = "create_pending"

	// LoadBalancerProvisioningStatusUpdatePending captures enum value "update_pending"
	LoadBalancerProvisioningStatusUpdatePending string = "update_pending"

	// LoadBalancerProvisioningStatusDeletePending captures enum value "delete_pending"
	LoadBalancerProvisioningStatusDeletePending string = "delete_pending"

	// LoadBalancerProvisioningStatusMaintenancePending captures enum value "maintenance_pending"
	LoadBalancerProvisioningStatusMaintenancePending string = "maintenance_pending"
)

// prop value enum
func (m *LoadBalancer) validateProvisioningStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, loadBalancerTypeProvisioningStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LoadBalancer) validateProvisioningStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ProvisioningStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateProvisioningStatusEnum("provisioning_status", "body", m.ProvisioningStatus); err != nil {
		return err
	}

	return nil
}

func (m *LoadBalancer) validatePublicIps(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicIps) { // not required
		return nil
	}

	for i := 0; i < len(m.PublicIps); i++ {
		if swag.IsZero(m.PublicIps[i]) { // not required
			continue
		}

		if m.PublicIps[i] != nil {
			if err := m.PublicIps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("public_ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoadBalancer) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancer) validateSubnets(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnets) { // not required
		return nil
	}

	for i := 0; i < len(m.Subnets); i++ {
		if swag.IsZero(m.Subnets[i]) { // not required
			continue
		}

		if m.Subnets[i] != nil {
			if err := m.Subnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancer) UnmarshalBinary(b []byte) error {
	var res LoadBalancer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoadBalancerSubnetsItems0 load balancer subnets items0
// swagger:model LoadBalancerSubnetsItems0
type LoadBalancerSubnetsItems0 struct {

	// The CRN for this subnet
	Crn string `json:"crn,omitempty"`

	// The unique identifier for this subnet
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this subnet
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this load balancer subnets items0
func (m *LoadBalancerSubnetsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancerSubnetsItems0) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoadBalancerSubnetsItems0) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancerSubnetsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancerSubnetsItems0) UnmarshalBinary(b []byte) error {
	var res LoadBalancerSubnetsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
