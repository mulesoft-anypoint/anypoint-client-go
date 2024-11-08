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

// checks if the OpenIDProviderGetServiceProviderUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIDProviderGetServiceProviderUrls{}

// OpenIDProviderGetServiceProviderUrls struct for OpenIDProviderGetServiceProviderUrls
type OpenIDProviderGetServiceProviderUrls struct {
	SignOn string `json:"sign_on"`
}

// NewOpenIDProviderGetServiceProviderUrls instantiates a new OpenIDProviderGetServiceProviderUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIDProviderGetServiceProviderUrls(signOn string) *OpenIDProviderGetServiceProviderUrls {
	this := OpenIDProviderGetServiceProviderUrls{}
	this.SignOn = signOn
	return &this
}

// NewOpenIDProviderGetServiceProviderUrlsWithDefaults instantiates a new OpenIDProviderGetServiceProviderUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIDProviderGetServiceProviderUrlsWithDefaults() *OpenIDProviderGetServiceProviderUrls {
	this := OpenIDProviderGetServiceProviderUrls{}
	return &this
}

// GetSignOn returns the SignOn field value
func (o *OpenIDProviderGetServiceProviderUrls) GetSignOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignOn
}

// GetSignOnOk returns a tuple with the SignOn field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderGetServiceProviderUrls) GetSignOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignOn, true
}

// SetSignOn sets field value
func (o *OpenIDProviderGetServiceProviderUrls) SetSignOn(v string) {
	o.SignOn = v
}

func (o OpenIDProviderGetServiceProviderUrls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIDProviderGetServiceProviderUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sign_on"] = o.SignOn
	return toSerialize, nil
}

type NullableOpenIDProviderGetServiceProviderUrls struct {
	value *OpenIDProviderGetServiceProviderUrls
	isSet bool
}

func (v NullableOpenIDProviderGetServiceProviderUrls) Get() *OpenIDProviderGetServiceProviderUrls {
	return v.value
}

func (v *NullableOpenIDProviderGetServiceProviderUrls) Set(val *OpenIDProviderGetServiceProviderUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIDProviderGetServiceProviderUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIDProviderGetServiceProviderUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIDProviderGetServiceProviderUrls(val *OpenIDProviderGetServiceProviderUrls) *NullableOpenIDProviderGetServiceProviderUrls {
	return &NullableOpenIDProviderGetServiceProviderUrls{value: val, isSet: true}
}

func (v NullableOpenIDProviderGetServiceProviderUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIDProviderGetServiceProviderUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


