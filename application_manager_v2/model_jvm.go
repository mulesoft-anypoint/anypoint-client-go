/*
Deployment

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Runtime Fabric and CloudHub 2.0 targets only. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v2

import (
	"encoding/json"
)

// checks if the Jvm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Jvm{}

// Jvm struct for Jvm
type Jvm struct {
	Args *string `json:"args,omitempty"`
}

// NewJvm instantiates a new Jvm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJvm() *Jvm {
	this := Jvm{}
	return &this
}

// NewJvmWithDefaults instantiates a new Jvm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJvmWithDefaults() *Jvm {
	this := Jvm{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *Jvm) GetArgs() string {
	if o == nil || IsNil(o.Args) {
		var ret string
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Jvm) GetArgsOk() (*string, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *Jvm) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given string and assigns it to the Args field.
func (o *Jvm) SetArgs(v string) {
	o.Args = &v
}

func (o Jvm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Jvm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	return toSerialize, nil
}

type NullableJvm struct {
	value *Jvm
	isSet bool
}

func (v NullableJvm) Get() *Jvm {
	return v.value
}

func (v *NullableJvm) Set(val *Jvm) {
	v.value = val
	v.isSet = true
}

func (v NullableJvm) IsSet() bool {
	return v.isSet
}

func (v *NullableJvm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJvm(val *Jvm) *NullableJvm {
	return &NullableJvm{value: val, isSet: true}
}

func (v NullableJvm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJvm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

