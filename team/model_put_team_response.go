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

// checks if the PutTeamResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutTeamResponse{}

// PutTeamResponse struct for PutTeamResponse
type PutTeamResponse struct {
	OrgId *string `json:"org_id,omitempty"`
	TeamId *string `json:"team_id,omitempty"`
	TeamName *string `json:"team_name,omitempty"`
	TeamType *string `json:"team_type,omitempty"`
	AncestorTeamIds []string `json:"ancestor_team_ids,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	PreviousAncestorTeamIds []string `json:"previous_ancestor_team_ids,omitempty"`
}

// NewPutTeamResponse instantiates a new PutTeamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutTeamResponse() *PutTeamResponse {
	this := PutTeamResponse{}
	return &this
}

// NewPutTeamResponseWithDefaults instantiates a new PutTeamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutTeamResponseWithDefaults() *PutTeamResponse {
	this := PutTeamResponse{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *PutTeamResponse) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTeamResponse) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *PutTeamResponse) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *PutTeamResponse) SetOrgId(v string) {
	o.OrgId = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *PutTeamResponse) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTeamResponse) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *PutTeamResponse) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *PutTeamResponse) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *PutTeamResponse) GetTeamName() string {
	if o == nil || IsNil(o.TeamName) {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTeamResponse) GetTeamNameOk() (*string, bool) {
	if o == nil || IsNil(o.TeamName) {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *PutTeamResponse) HasTeamName() bool {
	if o != nil && !IsNil(o.TeamName) {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *PutTeamResponse) SetTeamName(v string) {
	o.TeamName = &v
}

// GetTeamType returns the TeamType field value if set, zero value otherwise.
func (o *PutTeamResponse) GetTeamType() string {
	if o == nil || IsNil(o.TeamType) {
		var ret string
		return ret
	}
	return *o.TeamType
}

// GetTeamTypeOk returns a tuple with the TeamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTeamResponse) GetTeamTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TeamType) {
		return nil, false
	}
	return o.TeamType, true
}

// HasTeamType returns a boolean if a field has been set.
func (o *PutTeamResponse) HasTeamType() bool {
	if o != nil && !IsNil(o.TeamType) {
		return true
	}

	return false
}

// SetTeamType gets a reference to the given string and assigns it to the TeamType field.
func (o *PutTeamResponse) SetTeamType(v string) {
	o.TeamType = &v
}

// GetAncestorTeamIds returns the AncestorTeamIds field value if set, zero value otherwise.
func (o *PutTeamResponse) GetAncestorTeamIds() []string {
	if o == nil || IsNil(o.AncestorTeamIds) {
		var ret []string
		return ret
	}
	return o.AncestorTeamIds
}

// GetAncestorTeamIdsOk returns a tuple with the AncestorTeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTeamResponse) GetAncestorTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AncestorTeamIds) {
		return nil, false
	}
	return o.AncestorTeamIds, true
}

// HasAncestorTeamIds returns a boolean if a field has been set.
func (o *PutTeamResponse) HasAncestorTeamIds() bool {
	if o != nil && !IsNil(o.AncestorTeamIds) {
		return true
	}

	return false
}

// SetAncestorTeamIds gets a reference to the given []string and assigns it to the AncestorTeamIds field.
func (o *PutTeamResponse) SetAncestorTeamIds(v []string) {
	o.AncestorTeamIds = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PutTeamResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTeamResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PutTeamResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PutTeamResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PutTeamResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTeamResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PutTeamResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PutTeamResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetPreviousAncestorTeamIds returns the PreviousAncestorTeamIds field value if set, zero value otherwise.
func (o *PutTeamResponse) GetPreviousAncestorTeamIds() []string {
	if o == nil || IsNil(o.PreviousAncestorTeamIds) {
		var ret []string
		return ret
	}
	return o.PreviousAncestorTeamIds
}

// GetPreviousAncestorTeamIdsOk returns a tuple with the PreviousAncestorTeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutTeamResponse) GetPreviousAncestorTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PreviousAncestorTeamIds) {
		return nil, false
	}
	return o.PreviousAncestorTeamIds, true
}

// HasPreviousAncestorTeamIds returns a boolean if a field has been set.
func (o *PutTeamResponse) HasPreviousAncestorTeamIds() bool {
	if o != nil && !IsNil(o.PreviousAncestorTeamIds) {
		return true
	}

	return false
}

// SetPreviousAncestorTeamIds gets a reference to the given []string and assigns it to the PreviousAncestorTeamIds field.
func (o *PutTeamResponse) SetPreviousAncestorTeamIds(v []string) {
	o.PreviousAncestorTeamIds = v
}

func (o PutTeamResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutTeamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	if !IsNil(o.TeamName) {
		toSerialize["team_name"] = o.TeamName
	}
	if !IsNil(o.TeamType) {
		toSerialize["team_type"] = o.TeamType
	}
	if !IsNil(o.AncestorTeamIds) {
		toSerialize["ancestor_team_ids"] = o.AncestorTeamIds
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.PreviousAncestorTeamIds) {
		toSerialize["previous_ancestor_team_ids"] = o.PreviousAncestorTeamIds
	}
	return toSerialize, nil
}

type NullablePutTeamResponse struct {
	value *PutTeamResponse
	isSet bool
}

func (v NullablePutTeamResponse) Get() *PutTeamResponse {
	return v.value
}

func (v *NullablePutTeamResponse) Set(val *PutTeamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePutTeamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePutTeamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutTeamResponse(val *PutTeamResponse) *NullablePutTeamResponse {
	return &NullablePutTeamResponse{value: val, isSet: true}
}

func (v NullablePutTeamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutTeamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


