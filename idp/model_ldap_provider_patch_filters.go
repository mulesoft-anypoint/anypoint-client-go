/*
Identity Provider Management API

Description of Identity Provider API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// checks if the LdapProviderPatchFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapProviderPatchFilters{}

// LdapProviderPatchFilters struct for LdapProviderPatchFilters
type LdapProviderPatchFilters struct {
	GroupsByUsername *string `json:"groupsByUsername,omitempty"`
	UserByUsername *string `json:"userByUsername,omitempty"`
}

// NewLdapProviderPatchFilters instantiates a new LdapProviderPatchFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProviderPatchFilters() *LdapProviderPatchFilters {
	this := LdapProviderPatchFilters{}
	return &this
}

// NewLdapProviderPatchFiltersWithDefaults instantiates a new LdapProviderPatchFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderPatchFiltersWithDefaults() *LdapProviderPatchFilters {
	this := LdapProviderPatchFilters{}
	return &this
}

// GetGroupsByUsername returns the GroupsByUsername field value if set, zero value otherwise.
func (o *LdapProviderPatchFilters) GetGroupsByUsername() string {
	if o == nil || IsNil(o.GroupsByUsername) {
		var ret string
		return ret
	}
	return *o.GroupsByUsername
}

// GetGroupsByUsernameOk returns a tuple with the GroupsByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchFilters) GetGroupsByUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupsByUsername) {
		return nil, false
	}
	return o.GroupsByUsername, true
}

// HasGroupsByUsername returns a boolean if a field has been set.
func (o *LdapProviderPatchFilters) HasGroupsByUsername() bool {
	if o != nil && !IsNil(o.GroupsByUsername) {
		return true
	}

	return false
}

// SetGroupsByUsername gets a reference to the given string and assigns it to the GroupsByUsername field.
func (o *LdapProviderPatchFilters) SetGroupsByUsername(v string) {
	o.GroupsByUsername = &v
}

// GetUserByUsername returns the UserByUsername field value if set, zero value otherwise.
func (o *LdapProviderPatchFilters) GetUserByUsername() string {
	if o == nil || IsNil(o.UserByUsername) {
		var ret string
		return ret
	}
	return *o.UserByUsername
}

// GetUserByUsernameOk returns a tuple with the UserByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchFilters) GetUserByUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.UserByUsername) {
		return nil, false
	}
	return o.UserByUsername, true
}

// HasUserByUsername returns a boolean if a field has been set.
func (o *LdapProviderPatchFilters) HasUserByUsername() bool {
	if o != nil && !IsNil(o.UserByUsername) {
		return true
	}

	return false
}

// SetUserByUsername gets a reference to the given string and assigns it to the UserByUsername field.
func (o *LdapProviderPatchFilters) SetUserByUsername(v string) {
	o.UserByUsername = &v
}

func (o LdapProviderPatchFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapProviderPatchFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupsByUsername) {
		toSerialize["groupsByUsername"] = o.GroupsByUsername
	}
	if !IsNil(o.UserByUsername) {
		toSerialize["userByUsername"] = o.UserByUsername
	}
	return toSerialize, nil
}

type NullableLdapProviderPatchFilters struct {
	value *LdapProviderPatchFilters
	isSet bool
}

func (v NullableLdapProviderPatchFilters) Get() *LdapProviderPatchFilters {
	return v.value
}

func (v *NullableLdapProviderPatchFilters) Set(val *LdapProviderPatchFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProviderPatchFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProviderPatchFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProviderPatchFilters(val *LdapProviderPatchFilters) *NullableLdapProviderPatchFilters {
	return &NullableLdapProviderPatchFilters{value: val, isSet: true}
}

func (v NullableLdapProviderPatchFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProviderPatchFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


