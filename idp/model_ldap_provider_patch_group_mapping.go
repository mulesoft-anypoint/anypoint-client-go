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

// checks if the LdapProviderPatchGroupMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapProviderPatchGroupMapping{}

// LdapProviderPatchGroupMapping struct for LdapProviderPatchGroupMapping
type LdapProviderPatchGroupMapping struct {
	GroupName *string `json:"groupName,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewLdapProviderPatchGroupMapping instantiates a new LdapProviderPatchGroupMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProviderPatchGroupMapping() *LdapProviderPatchGroupMapping {
	this := LdapProviderPatchGroupMapping{}
	return &this
}

// NewLdapProviderPatchGroupMappingWithDefaults instantiates a new LdapProviderPatchGroupMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderPatchGroupMappingWithDefaults() *LdapProviderPatchGroupMapping {
	this := LdapProviderPatchGroupMapping{}
	return &this
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *LdapProviderPatchGroupMapping) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchGroupMapping) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *LdapProviderPatchGroupMapping) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *LdapProviderPatchGroupMapping) SetGroupName(v string) {
	o.GroupName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LdapProviderPatchGroupMapping) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatchGroupMapping) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LdapProviderPatchGroupMapping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LdapProviderPatchGroupMapping) SetId(v string) {
	o.Id = &v
}

func (o LdapProviderPatchGroupMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapProviderPatchGroupMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableLdapProviderPatchGroupMapping struct {
	value *LdapProviderPatchGroupMapping
	isSet bool
}

func (v NullableLdapProviderPatchGroupMapping) Get() *LdapProviderPatchGroupMapping {
	return v.value
}

func (v *NullableLdapProviderPatchGroupMapping) Set(val *LdapProviderPatchGroupMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProviderPatchGroupMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProviderPatchGroupMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProviderPatchGroupMapping(val *LdapProviderPatchGroupMapping) *NullableLdapProviderPatchGroupMapping {
	return &NullableLdapProviderPatchGroupMapping{value: val, isSet: true}
}

func (v NullableLdapProviderPatchGroupMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProviderPatchGroupMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


