/*
Deployment v1

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Cloudhub 1.0 targets only. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v1

import (
	"encoding/json"
)

// checks if the DeploymentInfoWorkers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentInfoWorkers{}

// DeploymentInfoWorkers struct for DeploymentInfoWorkers
type DeploymentInfoWorkers struct {
	// The number or replicas
	Amount *int32 `json:"amount,omitempty"`
	Type *DeploymentInfoWorkersType `json:"type,omitempty"`
}

// NewDeploymentInfoWorkers instantiates a new DeploymentInfoWorkers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentInfoWorkers() *DeploymentInfoWorkers {
	this := DeploymentInfoWorkers{}
	return &this
}

// NewDeploymentInfoWorkersWithDefaults instantiates a new DeploymentInfoWorkers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentInfoWorkersWithDefaults() *DeploymentInfoWorkers {
	this := DeploymentInfoWorkers{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DeploymentInfoWorkers) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentInfoWorkers) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DeploymentInfoWorkers) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *DeploymentInfoWorkers) SetAmount(v int32) {
	o.Amount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploymentInfoWorkers) GetType() DeploymentInfoWorkersType {
	if o == nil || IsNil(o.Type) {
		var ret DeploymentInfoWorkersType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentInfoWorkers) GetTypeOk() (*DeploymentInfoWorkersType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploymentInfoWorkers) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DeploymentInfoWorkersType and assigns it to the Type field.
func (o *DeploymentInfoWorkers) SetType(v DeploymentInfoWorkersType) {
	o.Type = &v
}

func (o DeploymentInfoWorkers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentInfoWorkers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDeploymentInfoWorkers struct {
	value *DeploymentInfoWorkers
	isSet bool
}

func (v NullableDeploymentInfoWorkers) Get() *DeploymentInfoWorkers {
	return v.value
}

func (v *NullableDeploymentInfoWorkers) Set(val *DeploymentInfoWorkers) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentInfoWorkers) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentInfoWorkers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentInfoWorkers(val *DeploymentInfoWorkers) *NullableDeploymentInfoWorkers {
	return &NullableDeploymentInfoWorkers{value: val, isSet: true}
}

func (v NullableDeploymentInfoWorkers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentInfoWorkers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


