/*
Deployment v1

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Cloudhub 1.0 targets only. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v1

import (
	"encoding/json"
)

// checks if the Workers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Workers{}

// Workers Worker configuration for the application deployment.
type Workers struct {
	Type *WorkerType `json:"type,omitempty"`
	// The number of worker instances allocated.
	Amount *int32 `json:"amount,omitempty"`
	// Remaining worker count available to the organization.
	RemainingOrgWorkers *float32 `json:"remainingOrgWorkers,omitempty"`
	// Total number of workers allocated to the organization.
	TotalOrgWorkers *float32 `json:"totalOrgWorkers,omitempty"`
	RecentStatistics *WorkersRecentStatistics `json:"recentStatistics,omitempty"`
}

// NewWorkers instantiates a new Workers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkers() *Workers {
	this := Workers{}
	return &this
}

// NewWorkersWithDefaults instantiates a new Workers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkersWithDefaults() *Workers {
	this := Workers{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Workers) GetType() WorkerType {
	if o == nil || IsNil(o.Type) {
		var ret WorkerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workers) GetTypeOk() (*WorkerType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Workers) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given WorkerType and assigns it to the Type field.
func (o *Workers) SetType(v WorkerType) {
	o.Type = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Workers) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workers) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Workers) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *Workers) SetAmount(v int32) {
	o.Amount = &v
}

// GetRemainingOrgWorkers returns the RemainingOrgWorkers field value if set, zero value otherwise.
func (o *Workers) GetRemainingOrgWorkers() float32 {
	if o == nil || IsNil(o.RemainingOrgWorkers) {
		var ret float32
		return ret
	}
	return *o.RemainingOrgWorkers
}

// GetRemainingOrgWorkersOk returns a tuple with the RemainingOrgWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workers) GetRemainingOrgWorkersOk() (*float32, bool) {
	if o == nil || IsNil(o.RemainingOrgWorkers) {
		return nil, false
	}
	return o.RemainingOrgWorkers, true
}

// HasRemainingOrgWorkers returns a boolean if a field has been set.
func (o *Workers) HasRemainingOrgWorkers() bool {
	if o != nil && !IsNil(o.RemainingOrgWorkers) {
		return true
	}

	return false
}

// SetRemainingOrgWorkers gets a reference to the given float32 and assigns it to the RemainingOrgWorkers field.
func (o *Workers) SetRemainingOrgWorkers(v float32) {
	o.RemainingOrgWorkers = &v
}

// GetTotalOrgWorkers returns the TotalOrgWorkers field value if set, zero value otherwise.
func (o *Workers) GetTotalOrgWorkers() float32 {
	if o == nil || IsNil(o.TotalOrgWorkers) {
		var ret float32
		return ret
	}
	return *o.TotalOrgWorkers
}

// GetTotalOrgWorkersOk returns a tuple with the TotalOrgWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workers) GetTotalOrgWorkersOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalOrgWorkers) {
		return nil, false
	}
	return o.TotalOrgWorkers, true
}

// HasTotalOrgWorkers returns a boolean if a field has been set.
func (o *Workers) HasTotalOrgWorkers() bool {
	if o != nil && !IsNil(o.TotalOrgWorkers) {
		return true
	}

	return false
}

// SetTotalOrgWorkers gets a reference to the given float32 and assigns it to the TotalOrgWorkers field.
func (o *Workers) SetTotalOrgWorkers(v float32) {
	o.TotalOrgWorkers = &v
}

// GetRecentStatistics returns the RecentStatistics field value if set, zero value otherwise.
func (o *Workers) GetRecentStatistics() WorkersRecentStatistics {
	if o == nil || IsNil(o.RecentStatistics) {
		var ret WorkersRecentStatistics
		return ret
	}
	return *o.RecentStatistics
}

// GetRecentStatisticsOk returns a tuple with the RecentStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workers) GetRecentStatisticsOk() (*WorkersRecentStatistics, bool) {
	if o == nil || IsNil(o.RecentStatistics) {
		return nil, false
	}
	return o.RecentStatistics, true
}

// HasRecentStatistics returns a boolean if a field has been set.
func (o *Workers) HasRecentStatistics() bool {
	if o != nil && !IsNil(o.RecentStatistics) {
		return true
	}

	return false
}

// SetRecentStatistics gets a reference to the given WorkersRecentStatistics and assigns it to the RecentStatistics field.
func (o *Workers) SetRecentStatistics(v WorkersRecentStatistics) {
	o.RecentStatistics = &v
}

func (o Workers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Workers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.RemainingOrgWorkers) {
		toSerialize["remainingOrgWorkers"] = o.RemainingOrgWorkers
	}
	if !IsNil(o.TotalOrgWorkers) {
		toSerialize["totalOrgWorkers"] = o.TotalOrgWorkers
	}
	if !IsNil(o.RecentStatistics) {
		toSerialize["recentStatistics"] = o.RecentStatistics
	}
	return toSerialize, nil
}

type NullableWorkers struct {
	value *Workers
	isSet bool
}

func (v NullableWorkers) Get() *Workers {
	return v.value
}

func (v *NullableWorkers) Set(val *Workers) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkers) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkers(val *Workers) *NullableWorkers {
	return &NullableWorkers{value: val, isSet: true}
}

func (v NullableWorkers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


