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

// checks if the TeamRoleCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamRoleCollection{}

// TeamRoleCollection struct for TeamRoleCollection
type TeamRoleCollection struct {
	Data []TeamRole `json:"data,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewTeamRoleCollection instantiates a new TeamRoleCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamRoleCollection() *TeamRoleCollection {
	this := TeamRoleCollection{}
	return &this
}

// NewTeamRoleCollectionWithDefaults instantiates a new TeamRoleCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamRoleCollectionWithDefaults() *TeamRoleCollection {
	this := TeamRoleCollection{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TeamRoleCollection) GetData() []TeamRole {
	if o == nil || IsNil(o.Data) {
		var ret []TeamRole
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoleCollection) GetDataOk() ([]TeamRole, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TeamRoleCollection) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TeamRole and assigns it to the Data field.
func (o *TeamRoleCollection) SetData(v []TeamRole) {
	o.Data = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TeamRoleCollection) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamRoleCollection) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TeamRoleCollection) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *TeamRoleCollection) SetTotal(v int32) {
	o.Total = &v
}

func (o TeamRoleCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamRoleCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableTeamRoleCollection struct {
	value *TeamRoleCollection
	isSet bool
}

func (v NullableTeamRoleCollection) Get() *TeamRoleCollection {
	return v.value
}

func (v *NullableTeamRoleCollection) Set(val *TeamRoleCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamRoleCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamRoleCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamRoleCollection(val *TeamRoleCollection) *NullableTeamRoleCollection {
	return &NullableTeamRoleCollection{value: val, isSet: true}
}

func (v NullableTeamRoleCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamRoleCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


