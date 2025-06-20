/*
Application Manager v1

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Cloudhub 1.0 targets only. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v1

import (
	"encoding/json"
)

// checks if the WorkersRecentStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkersRecentStatistics{}

// WorkersRecentStatistics Recent performance statistics for the workers.
type WorkersRecentStatistics struct {
	// The number of transactions processed.
	Transactions *int32 `json:"transactions,omitempty"`
	// The current CPU usage percentage.
	Cpu *int32 `json:"cpu,omitempty"`
}

// NewWorkersRecentStatistics instantiates a new WorkersRecentStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkersRecentStatistics() *WorkersRecentStatistics {
	this := WorkersRecentStatistics{}
	return &this
}

// NewWorkersRecentStatisticsWithDefaults instantiates a new WorkersRecentStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkersRecentStatisticsWithDefaults() *WorkersRecentStatistics {
	this := WorkersRecentStatistics{}
	return &this
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *WorkersRecentStatistics) GetTransactions() int32 {
	if o == nil || IsNil(o.Transactions) {
		var ret int32
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkersRecentStatistics) GetTransactionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Transactions) {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *WorkersRecentStatistics) HasTransactions() bool {
	if o != nil && !IsNil(o.Transactions) {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given int32 and assigns it to the Transactions field.
func (o *WorkersRecentStatistics) SetTransactions(v int32) {
	o.Transactions = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *WorkersRecentStatistics) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkersRecentStatistics) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *WorkersRecentStatistics) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *WorkersRecentStatistics) SetCpu(v int32) {
	o.Cpu = &v
}

func (o WorkersRecentStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkersRecentStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transactions) {
		toSerialize["transactions"] = o.Transactions
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	return toSerialize, nil
}

type NullableWorkersRecentStatistics struct {
	value *WorkersRecentStatistics
	isSet bool
}

func (v NullableWorkersRecentStatistics) Get() *WorkersRecentStatistics {
	return v.value
}

func (v *NullableWorkersRecentStatistics) Set(val *WorkersRecentStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkersRecentStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkersRecentStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkersRecentStatistics(val *WorkersRecentStatistics) *NullableWorkersRecentStatistics {
	return &NullableWorkersRecentStatistics{value: val, isSet: true}
}

func (v NullableWorkersRecentStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkersRecentStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


