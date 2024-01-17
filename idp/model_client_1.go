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

// checks if the Client1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Client1{}

// Client1 either contains urls or credentials (mutually exclusive)
type Client1 struct {
	Urls *Urls1 `json:"urls,omitempty"`
	Credentials *Credentials1 `json:"credentials,omitempty"`
}

// NewClient1 instantiates a new Client1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClient1() *Client1 {
	this := Client1{}
	return &this
}

// NewClient1WithDefaults instantiates a new Client1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClient1WithDefaults() *Client1 {
	this := Client1{}
	return &this
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *Client1) GetUrls() Urls1 {
	if o == nil || IsNil(o.Urls) {
		var ret Urls1
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client1) GetUrlsOk() (*Urls1, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *Client1) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given Urls1 and assigns it to the Urls field.
func (o *Client1) SetUrls(v Urls1) {
	o.Urls = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *Client1) GetCredentials() Credentials1 {
	if o == nil || IsNil(o.Credentials) {
		var ret Credentials1
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client1) GetCredentialsOk() (*Credentials1, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *Client1) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given Credentials1 and assigns it to the Credentials field.
func (o *Client1) SetCredentials(v Credentials1) {
	o.Credentials = &v
}

func (o Client1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Client1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	return toSerialize, nil
}

type NullableClient1 struct {
	value *Client1
	isSet bool
}

func (v NullableClient1) Get() *Client1 {
	return v.value
}

func (v *NullableClient1) Set(val *Client1) {
	v.value = val
	v.isSet = true
}

func (v NullableClient1) IsSet() bool {
	return v.isSet
}

func (v *NullableClient1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClient1(val *Client1) *NullableClient1 {
	return &NullableClient1{value: val, isSet: true}
}

func (v NullableClient1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClient1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


