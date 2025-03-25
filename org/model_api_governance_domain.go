/*
Organization API

Description of the Organization API

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// checks if the ApiGovernanceDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiGovernanceDomain{}

// ApiGovernanceDomain struct for ApiGovernanceDomain
type ApiGovernanceDomain struct {
	// API Governance domain
	Domain *string `json:"domain,omitempty"`
}

// NewApiGovernanceDomain instantiates a new ApiGovernanceDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiGovernanceDomain() *ApiGovernanceDomain {
	this := ApiGovernanceDomain{}
	var domain string = "ASSET"
	this.Domain = &domain
	return &this
}

// NewApiGovernanceDomainWithDefaults instantiates a new ApiGovernanceDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiGovernanceDomainWithDefaults() *ApiGovernanceDomain {
	this := ApiGovernanceDomain{}
	var domain string = "ASSET"
	this.Domain = &domain
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiGovernanceDomain) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiGovernanceDomain) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiGovernanceDomain) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiGovernanceDomain) SetDomain(v string) {
	o.Domain = &v
}

func (o ApiGovernanceDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiGovernanceDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	return toSerialize, nil
}

type NullableApiGovernanceDomain struct {
	value *ApiGovernanceDomain
	isSet bool
}

func (v NullableApiGovernanceDomain) Get() *ApiGovernanceDomain {
	return v.value
}

func (v *NullableApiGovernanceDomain) Set(val *ApiGovernanceDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableApiGovernanceDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableApiGovernanceDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiGovernanceDomain(val *ApiGovernanceDomain) *NullableApiGovernanceDomain {
	return &NullableApiGovernanceDomain{value: val, isSet: true}
}

func (v NullableApiGovernanceDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiGovernanceDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


