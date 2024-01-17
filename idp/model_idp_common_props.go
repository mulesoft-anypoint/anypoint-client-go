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

// checks if the IdpCommonProps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpCommonProps{}

// IdpCommonProps struct for IdpCommonProps
type IdpCommonProps struct {
	ServiceProvider *ServiceProvider `json:"service_provider,omitempty"`
}

// NewIdpCommonProps instantiates a new IdpCommonProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpCommonProps() *IdpCommonProps {
	this := IdpCommonProps{}
	return &this
}

// NewIdpCommonPropsWithDefaults instantiates a new IdpCommonProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpCommonPropsWithDefaults() *IdpCommonProps {
	this := IdpCommonProps{}
	return &this
}

// GetServiceProvider returns the ServiceProvider field value if set, zero value otherwise.
func (o *IdpCommonProps) GetServiceProvider() ServiceProvider {
	if o == nil || IsNil(o.ServiceProvider) {
		var ret ServiceProvider
		return ret
	}
	return *o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpCommonProps) GetServiceProviderOk() (*ServiceProvider, bool) {
	if o == nil || IsNil(o.ServiceProvider) {
		return nil, false
	}
	return o.ServiceProvider, true
}

// HasServiceProvider returns a boolean if a field has been set.
func (o *IdpCommonProps) HasServiceProvider() bool {
	if o != nil && !IsNil(o.ServiceProvider) {
		return true
	}

	return false
}

// SetServiceProvider gets a reference to the given ServiceProvider and assigns it to the ServiceProvider field.
func (o *IdpCommonProps) SetServiceProvider(v ServiceProvider) {
	o.ServiceProvider = &v
}

func (o IdpCommonProps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpCommonProps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceProvider) {
		toSerialize["service_provider"] = o.ServiceProvider
	}
	return toSerialize, nil
}

type NullableIdpCommonProps struct {
	value *IdpCommonProps
	isSet bool
}

func (v NullableIdpCommonProps) Get() *IdpCommonProps {
	return v.value
}

func (v *NullableIdpCommonProps) Set(val *IdpCommonProps) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpCommonProps) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpCommonProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpCommonProps(val *IdpCommonProps) *NullableIdpCommonProps {
	return &NullableIdpCommonProps{value: val, isSet: true}
}

func (v NullableIdpCommonProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpCommonProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


