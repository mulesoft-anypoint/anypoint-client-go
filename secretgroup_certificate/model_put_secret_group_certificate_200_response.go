/*
Secret Group Certificate API

Secret Group Certificate API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_certificate

import (
	"encoding/json"
)

// checks if the PutSecretGroupCertificate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutSecretGroupCertificate200Response{}

// PutSecretGroupCertificate200Response struct for PutSecretGroupCertificate200Response
type PutSecretGroupCertificate200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewPutSecretGroupCertificate200Response instantiates a new PutSecretGroupCertificate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutSecretGroupCertificate200Response() *PutSecretGroupCertificate200Response {
	this := PutSecretGroupCertificate200Response{}
	return &this
}

// NewPutSecretGroupCertificate200ResponseWithDefaults instantiates a new PutSecretGroupCertificate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutSecretGroupCertificate200ResponseWithDefaults() *PutSecretGroupCertificate200Response {
	this := PutSecretGroupCertificate200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PutSecretGroupCertificate200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutSecretGroupCertificate200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PutSecretGroupCertificate200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PutSecretGroupCertificate200Response) SetMessage(v string) {
	o.Message = &v
}

func (o PutSecretGroupCertificate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutSecretGroupCertificate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePutSecretGroupCertificate200Response struct {
	value *PutSecretGroupCertificate200Response
	isSet bool
}

func (v NullablePutSecretGroupCertificate200Response) Get() *PutSecretGroupCertificate200Response {
	return v.value
}

func (v *NullablePutSecretGroupCertificate200Response) Set(val *PutSecretGroupCertificate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePutSecretGroupCertificate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePutSecretGroupCertificate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutSecretGroupCertificate200Response(val *PutSecretGroupCertificate200Response) *NullablePutSecretGroupCertificate200Response {
	return &NullablePutSecretGroupCertificate200Response{value: val, isSet: true}
}

func (v NullablePutSecretGroupCertificate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutSecretGroupCertificate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

