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

// checks if the RuntimeCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuntimeCollection{}

// RuntimeCollection struct for RuntimeCollection
type RuntimeCollection struct {
	Data []Runtime `json:"data,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewRuntimeCollection instantiates a new RuntimeCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuntimeCollection() *RuntimeCollection {
	this := RuntimeCollection{}
	return &this
}

// NewRuntimeCollectionWithDefaults instantiates a new RuntimeCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuntimeCollectionWithDefaults() *RuntimeCollection {
	this := RuntimeCollection{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RuntimeCollection) GetData() []Runtime {
	if o == nil || IsNil(o.Data) {
		var ret []Runtime
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeCollection) GetDataOk() ([]Runtime, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RuntimeCollection) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Runtime and assigns it to the Data field.
func (o *RuntimeCollection) SetData(v []Runtime) {
	o.Data = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RuntimeCollection) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeCollection) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RuntimeCollection) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *RuntimeCollection) SetTotal(v int32) {
	o.Total = &v
}

func (o RuntimeCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuntimeCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableRuntimeCollection struct {
	value *RuntimeCollection
	isSet bool
}

func (v NullableRuntimeCollection) Get() *RuntimeCollection {
	return v.value
}

func (v *NullableRuntimeCollection) Set(val *RuntimeCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeCollection(val *RuntimeCollection) *NullableRuntimeCollection {
	return &NullableRuntimeCollection{value: val, isSet: true}
}

func (v NullableRuntimeCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


