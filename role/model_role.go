/*
 * Role API
 *
 * Description of the Role API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package role

import (
	"encoding/json"
)

// Role struct for Role
type Role struct {
	RoleId *string `json:"role_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Internal *bool `json:"internal,omitempty"`
	OrgId NullableString `json:"org_id,omitempty"`
	Namespaces *[]string `json:"namespaces,omitempty"`
	Shareable *bool `json:"shareable,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole() *Role {
	this := Role{}
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *Role) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *Role) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *Role) SetRoleId(v string) {
	o.RoleId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Role) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Role) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Role) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Role) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Role) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Role) SetDescription(v string) {
	o.Description = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *Role) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *Role) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *Role) SetInternal(v bool) {
	o.Internal = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Role) GetOrgId() string {
	if o == nil || o.OrgId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrgId.Get()
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Role) GetOrgIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OrgId.Get(), o.OrgId.IsSet()
}

// HasOrgId returns a boolean if a field has been set.
func (o *Role) HasOrgId() bool {
	if o != nil && o.OrgId.IsSet() {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given NullableString and assigns it to the OrgId field.
func (o *Role) SetOrgId(v string) {
	o.OrgId.Set(&v)
}
// SetOrgIdNil sets the value for OrgId to be an explicit nil
func (o *Role) SetOrgIdNil() {
	o.OrgId.Set(nil)
}

// UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
func (o *Role) UnsetOrgId() {
	o.OrgId.Unset()
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *Role) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return *o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetNamespacesOk() (*[]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *Role) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *Role) SetNamespaces(v []string) {
	o.Namespaces = &v
}

// GetShareable returns the Shareable field value if set, zero value otherwise.
func (o *Role) GetShareable() bool {
	if o == nil || o.Shareable == nil {
		var ret bool
		return ret
	}
	return *o.Shareable
}

// GetShareableOk returns a tuple with the Shareable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetShareableOk() (*bool, bool) {
	if o == nil || o.Shareable == nil {
		return nil, false
	}
	return o.Shareable, true
}

// HasShareable returns a boolean if a field has been set.
func (o *Role) HasShareable() bool {
	if o != nil && o.Shareable != nil {
		return true
	}

	return false
}

// SetShareable gets a reference to the given bool and assigns it to the Shareable field.
func (o *Role) SetShareable(v bool) {
	o.Shareable = &v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleId != nil {
		toSerialize["role_id"] = o.RoleId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}
	if o.OrgId.IsSet() {
		toSerialize["org_id"] = o.OrgId.Get()
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Shareable != nil {
		toSerialize["shareable"] = o.Shareable
	}
	return json.Marshal(toSerialize)
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

