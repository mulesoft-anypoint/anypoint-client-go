/*
API Manager Contract API

API Manager Contract API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_contract

import (
	"encoding/json"
)

// checks if the Limit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Limit{}

// Limit struct for Limit
type Limit struct {
	MaximumRequests *int64 `json:"maximumRequests,omitempty"`
	TimePeriodInMilliseconds *int64 `json:"timePeriodInMilliseconds,omitempty"`
	Visible *bool `json:"visible,omitempty"`
}

// NewLimit instantiates a new Limit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimit() *Limit {
	this := Limit{}
	return &this
}

// NewLimitWithDefaults instantiates a new Limit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitWithDefaults() *Limit {
	this := Limit{}
	return &this
}

// GetMaximumRequests returns the MaximumRequests field value if set, zero value otherwise.
func (o *Limit) GetMaximumRequests() int64 {
	if o == nil || IsNil(o.MaximumRequests) {
		var ret int64
		return ret
	}
	return *o.MaximumRequests
}

// GetMaximumRequestsOk returns a tuple with the MaximumRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Limit) GetMaximumRequestsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumRequests) {
		return nil, false
	}
	return o.MaximumRequests, true
}

// HasMaximumRequests returns a boolean if a field has been set.
func (o *Limit) HasMaximumRequests() bool {
	if o != nil && !IsNil(o.MaximumRequests) {
		return true
	}

	return false
}

// SetMaximumRequests gets a reference to the given int64 and assigns it to the MaximumRequests field.
func (o *Limit) SetMaximumRequests(v int64) {
	o.MaximumRequests = &v
}

// GetTimePeriodInMilliseconds returns the TimePeriodInMilliseconds field value if set, zero value otherwise.
func (o *Limit) GetTimePeriodInMilliseconds() int64 {
	if o == nil || IsNil(o.TimePeriodInMilliseconds) {
		var ret int64
		return ret
	}
	return *o.TimePeriodInMilliseconds
}

// GetTimePeriodInMillisecondsOk returns a tuple with the TimePeriodInMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Limit) GetTimePeriodInMillisecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TimePeriodInMilliseconds) {
		return nil, false
	}
	return o.TimePeriodInMilliseconds, true
}

// HasTimePeriodInMilliseconds returns a boolean if a field has been set.
func (o *Limit) HasTimePeriodInMilliseconds() bool {
	if o != nil && !IsNil(o.TimePeriodInMilliseconds) {
		return true
	}

	return false
}

// SetTimePeriodInMilliseconds gets a reference to the given int64 and assigns it to the TimePeriodInMilliseconds field.
func (o *Limit) SetTimePeriodInMilliseconds(v int64) {
	o.TimePeriodInMilliseconds = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *Limit) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Limit) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *Limit) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *Limit) SetVisible(v bool) {
	o.Visible = &v
}

func (o Limit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Limit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaximumRequests) {
		toSerialize["maximumRequests"] = o.MaximumRequests
	}
	if !IsNil(o.TimePeriodInMilliseconds) {
		toSerialize["timePeriodInMilliseconds"] = o.TimePeriodInMilliseconds
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableLimit struct {
	value *Limit
	isSet bool
}

func (v NullableLimit) Get() *Limit {
	return v.value
}

func (v *NullableLimit) Set(val *Limit) {
	v.value = val
	v.isSet = true
}

func (v NullableLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimit(val *Limit) *NullableLimit {
	return &NullableLimit{value: val, isSet: true}
}

func (v NullableLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


