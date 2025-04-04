/*
Team Roles API

Description of the Team Roles API

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package team_roles

import (
	"encoding/json"
)

// checks if the TeamRolePostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamRolePostBody{}

// TeamRolePostBody struct for TeamRolePostBody
type TeamRolePostBody struct {
	RoleId *string `json:"role_id,omitempty"`
	ContextParams *ContextParams `json:"context_params,omitempty"`
}

// NewTeamRolePostBody instantiates a new TeamRolePostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamRolePostBody() *TeamRolePostBody {
	this := TeamRolePostBody{}
	return &this
}

// NewTeamRolePostBodyWithDefaults instantiates a new TeamRolePostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamRolePostBodyWithDefaults() *TeamRolePostBody {
	this := TeamRolePostBody{}
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *TeamRolePostBody) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRolePostBody) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *TeamRolePostBody) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *TeamRolePostBody) SetRoleId(v string) {
	o.RoleId = &v
}

// GetContextParams returns the ContextParams field value if set, zero value otherwise.
func (o *TeamRolePostBody) GetContextParams() ContextParams {
	if o == nil || IsNil(o.ContextParams) {
		var ret ContextParams
		return ret
	}
	return *o.ContextParams
}

// GetContextParamsOk returns a tuple with the ContextParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRolePostBody) GetContextParamsOk() (*ContextParams, bool) {
	if o == nil || IsNil(o.ContextParams) {
		return nil, false
	}
	return o.ContextParams, true
}

// HasContextParams returns a boolean if a field has been set.
func (o *TeamRolePostBody) HasContextParams() bool {
	if o != nil && !IsNil(o.ContextParams) {
		return true
	}

	return false
}

// SetContextParams gets a reference to the given ContextParams and assigns it to the ContextParams field.
func (o *TeamRolePostBody) SetContextParams(v ContextParams) {
	o.ContextParams = &v
}

func (o TeamRolePostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamRolePostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleId) {
		toSerialize["role_id"] = o.RoleId
	}
	if !IsNil(o.ContextParams) {
		toSerialize["context_params"] = o.ContextParams
	}
	return toSerialize, nil
}

type NullableTeamRolePostBody struct {
	value *TeamRolePostBody
	isSet bool
}

func (v NullableTeamRolePostBody) Get() *TeamRolePostBody {
	return v.value
}

func (v *NullableTeamRolePostBody) Set(val *TeamRolePostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamRolePostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamRolePostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamRolePostBody(val *TeamRolePostBody) *NullableTeamRolePostBody {
	return &NullableTeamRolePostBody{value: val, isSet: true}
}

func (v NullableTeamRolePostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamRolePostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


