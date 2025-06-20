/*
Private Space API

Description of the Private Space API

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package private_space

import (
	"encoding/json"
)

// checks if the IngressConfigurationDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngressConfigurationDeployment{}

// IngressConfigurationDeployment The deployment configuration of the ingress configuration.
type IngressConfigurationDeployment struct {
	// The status of the deployment configuration.
	Status *string `json:"status,omitempty"`
	// The last seen timestamp of the deployment configuration.
	LastSeenTimestamp *int64 `json:"lastSeenTimestamp,omitempty"`
}

// NewIngressConfigurationDeployment instantiates a new IngressConfigurationDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngressConfigurationDeployment() *IngressConfigurationDeployment {
	this := IngressConfigurationDeployment{}
	return &this
}

// NewIngressConfigurationDeploymentWithDefaults instantiates a new IngressConfigurationDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngressConfigurationDeploymentWithDefaults() *IngressConfigurationDeployment {
	this := IngressConfigurationDeployment{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IngressConfigurationDeployment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngressConfigurationDeployment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IngressConfigurationDeployment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IngressConfigurationDeployment) SetStatus(v string) {
	o.Status = &v
}

// GetLastSeenTimestamp returns the LastSeenTimestamp field value if set, zero value otherwise.
func (o *IngressConfigurationDeployment) GetLastSeenTimestamp() int64 {
	if o == nil || IsNil(o.LastSeenTimestamp) {
		var ret int64
		return ret
	}
	return *o.LastSeenTimestamp
}

// GetLastSeenTimestampOk returns a tuple with the LastSeenTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngressConfigurationDeployment) GetLastSeenTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.LastSeenTimestamp) {
		return nil, false
	}
	return o.LastSeenTimestamp, true
}

// HasLastSeenTimestamp returns a boolean if a field has been set.
func (o *IngressConfigurationDeployment) HasLastSeenTimestamp() bool {
	if o != nil && !IsNil(o.LastSeenTimestamp) {
		return true
	}

	return false
}

// SetLastSeenTimestamp gets a reference to the given int64 and assigns it to the LastSeenTimestamp field.
func (o *IngressConfigurationDeployment) SetLastSeenTimestamp(v int64) {
	o.LastSeenTimestamp = &v
}

func (o IngressConfigurationDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngressConfigurationDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.LastSeenTimestamp) {
		toSerialize["lastSeenTimestamp"] = o.LastSeenTimestamp
	}
	return toSerialize, nil
}

type NullableIngressConfigurationDeployment struct {
	value *IngressConfigurationDeployment
	isSet bool
}

func (v NullableIngressConfigurationDeployment) Get() *IngressConfigurationDeployment {
	return v.value
}

func (v *NullableIngressConfigurationDeployment) Set(val *IngressConfigurationDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableIngressConfigurationDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableIngressConfigurationDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngressConfigurationDeployment(val *IngressConfigurationDeployment) *NullableIngressConfigurationDeployment {
	return &NullableIngressConfigurationDeployment{value: val, isSet: true}
}

func (v NullableIngressConfigurationDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngressConfigurationDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


