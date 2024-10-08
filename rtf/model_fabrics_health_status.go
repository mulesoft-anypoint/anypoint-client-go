/*
Runtime Fabrics API

Runtime Fabrics API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtf

import (
	"encoding/json"
)

// checks if the FabricsHealthStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricsHealthStatus{}

// FabricsHealthStatus struct for FabricsHealthStatus
type FabricsHealthStatus struct {
	// True is the component is healthy.
	Healthy *bool `json:"healthy,omitempty"`
	UpdatedAt *int64 `json:"updatedAt,omitempty"`
	// Probes collected for this health check. Only applicable for Appliance probes.
	Probes *string `json:"probes,omitempty"`
	// Probe failures attributing to the result of this health check.
	FailedProbes []FabricsHealthStatusFailedProbesInner `json:"failedProbes,omitempty"`
}

// NewFabricsHealthStatus instantiates a new FabricsHealthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricsHealthStatus() *FabricsHealthStatus {
	this := FabricsHealthStatus{}
	return &this
}

// NewFabricsHealthStatusWithDefaults instantiates a new FabricsHealthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricsHealthStatusWithDefaults() *FabricsHealthStatus {
	this := FabricsHealthStatus{}
	return &this
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *FabricsHealthStatus) GetHealthy() bool {
	if o == nil || IsNil(o.Healthy) {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealthStatus) GetHealthyOk() (*bool, bool) {
	if o == nil || IsNil(o.Healthy) {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *FabricsHealthStatus) HasHealthy() bool {
	if o != nil && !IsNil(o.Healthy) {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *FabricsHealthStatus) SetHealthy(v bool) {
	o.Healthy = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FabricsHealthStatus) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealthStatus) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FabricsHealthStatus) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *FabricsHealthStatus) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetProbes returns the Probes field value if set, zero value otherwise.
func (o *FabricsHealthStatus) GetProbes() string {
	if o == nil || IsNil(o.Probes) {
		var ret string
		return ret
	}
	return *o.Probes
}

// GetProbesOk returns a tuple with the Probes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealthStatus) GetProbesOk() (*string, bool) {
	if o == nil || IsNil(o.Probes) {
		return nil, false
	}
	return o.Probes, true
}

// HasProbes returns a boolean if a field has been set.
func (o *FabricsHealthStatus) HasProbes() bool {
	if o != nil && !IsNil(o.Probes) {
		return true
	}

	return false
}

// SetProbes gets a reference to the given string and assigns it to the Probes field.
func (o *FabricsHealthStatus) SetProbes(v string) {
	o.Probes = &v
}

// GetFailedProbes returns the FailedProbes field value if set, zero value otherwise.
func (o *FabricsHealthStatus) GetFailedProbes() []FabricsHealthStatusFailedProbesInner {
	if o == nil || IsNil(o.FailedProbes) {
		var ret []FabricsHealthStatusFailedProbesInner
		return ret
	}
	return o.FailedProbes
}

// GetFailedProbesOk returns a tuple with the FailedProbes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHealthStatus) GetFailedProbesOk() ([]FabricsHealthStatusFailedProbesInner, bool) {
	if o == nil || IsNil(o.FailedProbes) {
		return nil, false
	}
	return o.FailedProbes, true
}

// HasFailedProbes returns a boolean if a field has been set.
func (o *FabricsHealthStatus) HasFailedProbes() bool {
	if o != nil && !IsNil(o.FailedProbes) {
		return true
	}

	return false
}

// SetFailedProbes gets a reference to the given []FabricsHealthStatusFailedProbesInner and assigns it to the FailedProbes field.
func (o *FabricsHealthStatus) SetFailedProbes(v []FabricsHealthStatusFailedProbesInner) {
	o.FailedProbes = v
}

func (o FabricsHealthStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricsHealthStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Healthy) {
		toSerialize["healthy"] = o.Healthy
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Probes) {
		toSerialize["probes"] = o.Probes
	}
	if !IsNil(o.FailedProbes) {
		toSerialize["failedProbes"] = o.FailedProbes
	}
	return toSerialize, nil
}

type NullableFabricsHealthStatus struct {
	value *FabricsHealthStatus
	isSet bool
}

func (v NullableFabricsHealthStatus) Get() *FabricsHealthStatus {
	return v.value
}

func (v *NullableFabricsHealthStatus) Set(val *FabricsHealthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricsHealthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricsHealthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricsHealthStatus(val *FabricsHealthStatus) *NullableFabricsHealthStatus {
	return &NullableFabricsHealthStatus{value: val, isSet: true}
}

func (v NullableFabricsHealthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricsHealthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


