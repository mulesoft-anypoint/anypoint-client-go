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

// checks if the IdpSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpSummary{}

// IdpSummary struct for IdpSummary
type IdpSummary struct {
	ProviderId string `json:"provider_id"`
	OrgId string `json:"org_id"`
	Name string `json:"name"`
	Type IdpSummaryType `json:"type"`
}

// NewIdpSummary instantiates a new IdpSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpSummary(providerId string, orgId string, name string, type_ IdpSummaryType) *IdpSummary {
	this := IdpSummary{}
	this.ProviderId = providerId
	this.OrgId = orgId
	this.Name = name
	this.Type = type_
	return &this
}

// NewIdpSummaryWithDefaults instantiates a new IdpSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpSummaryWithDefaults() *IdpSummary {
	this := IdpSummary{}
	return &this
}

// GetProviderId returns the ProviderId field value
func (o *IdpSummary) GetProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *IdpSummary) GetProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value
func (o *IdpSummary) SetProviderId(v string) {
	o.ProviderId = v
}

// GetOrgId returns the OrgId field value
func (o *IdpSummary) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *IdpSummary) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *IdpSummary) SetOrgId(v string) {
	o.OrgId = v
}

// GetName returns the Name field value
func (o *IdpSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdpSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdpSummary) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *IdpSummary) GetType() IdpSummaryType {
	if o == nil {
		var ret IdpSummaryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IdpSummary) GetTypeOk() (*IdpSummaryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IdpSummary) SetType(v IdpSummaryType) {
	o.Type = v
}

func (o IdpSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider_id"] = o.ProviderId
	toSerialize["org_id"] = o.OrgId
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableIdpSummary struct {
	value *IdpSummary
	isSet bool
}

func (v NullableIdpSummary) Get() *IdpSummary {
	return v.value
}

func (v *NullableIdpSummary) Set(val *IdpSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpSummary(val *IdpSummary) *NullableIdpSummary {
	return &NullableIdpSummary{value: val, isSet: true}
}

func (v NullableIdpSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


