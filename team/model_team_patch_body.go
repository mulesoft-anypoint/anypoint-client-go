/*
Team API

Description of the Team API

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package team

import (
	"encoding/json"
)

// checks if the TeamPatchBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamPatchBody{}

// TeamPatchBody struct for TeamPatchBody
type TeamPatchBody struct {
	TeamName *string `json:"team_name,omitempty"`
	TeamType *string `json:"team_type,omitempty"`
}

// NewTeamPatchBody instantiates a new TeamPatchBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamPatchBody() *TeamPatchBody {
	this := TeamPatchBody{}
	var teamType string = "internal"
	this.TeamType = &teamType
	return &this
}

// NewTeamPatchBodyWithDefaults instantiates a new TeamPatchBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamPatchBodyWithDefaults() *TeamPatchBody {
	this := TeamPatchBody{}
	var teamType string = "internal"
	this.TeamType = &teamType
	return &this
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *TeamPatchBody) GetTeamName() string {
	if o == nil || IsNil(o.TeamName) {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPatchBody) GetTeamNameOk() (*string, bool) {
	if o == nil || IsNil(o.TeamName) {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *TeamPatchBody) HasTeamName() bool {
	if o != nil && !IsNil(o.TeamName) {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *TeamPatchBody) SetTeamName(v string) {
	o.TeamName = &v
}

// GetTeamType returns the TeamType field value if set, zero value otherwise.
func (o *TeamPatchBody) GetTeamType() string {
	if o == nil || IsNil(o.TeamType) {
		var ret string
		return ret
	}
	return *o.TeamType
}

// GetTeamTypeOk returns a tuple with the TeamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamPatchBody) GetTeamTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TeamType) {
		return nil, false
	}
	return o.TeamType, true
}

// HasTeamType returns a boolean if a field has been set.
func (o *TeamPatchBody) HasTeamType() bool {
	if o != nil && !IsNil(o.TeamType) {
		return true
	}

	return false
}

// SetTeamType gets a reference to the given string and assigns it to the TeamType field.
func (o *TeamPatchBody) SetTeamType(v string) {
	o.TeamType = &v
}

func (o TeamPatchBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamPatchBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TeamName) {
		toSerialize["team_name"] = o.TeamName
	}
	if !IsNil(o.TeamType) {
		toSerialize["team_type"] = o.TeamType
	}
	return toSerialize, nil
}

type NullableTeamPatchBody struct {
	value *TeamPatchBody
	isSet bool
}

func (v NullableTeamPatchBody) Get() *TeamPatchBody {
	return v.value
}

func (v *NullableTeamPatchBody) Set(val *TeamPatchBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamPatchBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamPatchBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamPatchBody(val *TeamPatchBody) *NullableTeamPatchBody {
	return &NullableTeamPatchBody{value: val, isSet: true}
}

func (v NullableTeamPatchBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamPatchBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


