/*
Secret Group TLS Context API

Secret Group TLS Context API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_tlscontext

import (
	"encoding/json"
)

// checks if the GetSecretGroupTlsContextDetails404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSecretGroupTlsContextDetails404Response{}

// GetSecretGroupTlsContextDetails404Response struct for GetSecretGroupTlsContextDetails404Response
type GetSecretGroupTlsContextDetails404Response struct {
	Name *string `json:"name,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGetSecretGroupTlsContextDetails404Response instantiates a new GetSecretGroupTlsContextDetails404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSecretGroupTlsContextDetails404Response() *GetSecretGroupTlsContextDetails404Response {
	this := GetSecretGroupTlsContextDetails404Response{}
	return &this
}

// NewGetSecretGroupTlsContextDetails404ResponseWithDefaults instantiates a new GetSecretGroupTlsContextDetails404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSecretGroupTlsContextDetails404ResponseWithDefaults() *GetSecretGroupTlsContextDetails404Response {
	this := GetSecretGroupTlsContextDetails404Response{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetSecretGroupTlsContextDetails404Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecretGroupTlsContextDetails404Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetSecretGroupTlsContextDetails404Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetSecretGroupTlsContextDetails404Response) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetSecretGroupTlsContextDetails404Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecretGroupTlsContextDetails404Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetSecretGroupTlsContextDetails404Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetSecretGroupTlsContextDetails404Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetSecretGroupTlsContextDetails404Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSecretGroupTlsContextDetails404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetSecretGroupTlsContextDetails404Response struct {
	value *GetSecretGroupTlsContextDetails404Response
	isSet bool
}

func (v NullableGetSecretGroupTlsContextDetails404Response) Get() *GetSecretGroupTlsContextDetails404Response {
	return v.value
}

func (v *NullableGetSecretGroupTlsContextDetails404Response) Set(val *GetSecretGroupTlsContextDetails404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSecretGroupTlsContextDetails404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSecretGroupTlsContextDetails404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSecretGroupTlsContextDetails404Response(val *GetSecretGroupTlsContextDetails404Response) *NullableGetSecretGroupTlsContextDetails404Response {
	return &NullableGetSecretGroupTlsContextDetails404Response{value: val, isSet: true}
}

func (v NullableGetSecretGroupTlsContextDetails404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSecretGroupTlsContextDetails404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


