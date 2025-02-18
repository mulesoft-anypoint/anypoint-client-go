/*
Exchange Client Apps

Description of the Exchange Client Applications API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package exchange_client_apps

import (
	"encoding/json"
)

// checks if the GetExchangeClientApps400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExchangeClientApps400Response{}

// GetExchangeClientApps400Response struct for GetExchangeClientApps400Response
type GetExchangeClientApps400Response struct {
	Status *int32 `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGetExchangeClientApps400Response instantiates a new GetExchangeClientApps400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeClientApps400Response() *GetExchangeClientApps400Response {
	this := GetExchangeClientApps400Response{}
	var status int32 = 400
	this.Status = &status
	return &this
}

// NewGetExchangeClientApps400ResponseWithDefaults instantiates a new GetExchangeClientApps400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeClientApps400ResponseWithDefaults() *GetExchangeClientApps400Response {
	this := GetExchangeClientApps400Response{}
	var status int32 = 400
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetExchangeClientApps400Response) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeClientApps400Response) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetExchangeClientApps400Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *GetExchangeClientApps400Response) SetStatus(v int32) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetExchangeClientApps400Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeClientApps400Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetExchangeClientApps400Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetExchangeClientApps400Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetExchangeClientApps400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeClientApps400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetExchangeClientApps400Response struct {
	value *GetExchangeClientApps400Response
	isSet bool
}

func (v NullableGetExchangeClientApps400Response) Get() *GetExchangeClientApps400Response {
	return v.value
}

func (v *NullableGetExchangeClientApps400Response) Set(val *GetExchangeClientApps400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeClientApps400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeClientApps400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeClientApps400Response(val *GetExchangeClientApps400Response) *NullableGetExchangeClientApps400Response {
	return &NullableGetExchangeClientApps400Response{value: val, isSet: true}
}

func (v NullableGetExchangeClientApps400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeClientApps400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


