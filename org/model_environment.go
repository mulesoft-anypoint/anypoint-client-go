/*
Organization API

Description of the Organization API

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// checks if the Environment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Environment{}

// Environment an instance of an environment.
type Environment struct {
	// The client id of the environment.
	ClientId *string `json:"clientId,omitempty"`
	// The id of the environment.
	Id *string `json:"id,omitempty"`
	// Whether the environment is a production environment.
	IsProduction *bool `json:"isProduction,omitempty"`
	// The name of the environment.
	Name *string `json:"name,omitempty"`
	// The id of the organization.
	OrganizationId *string `json:"organizationId,omitempty"`
	// The type of the environment.
	Type *string `json:"type,omitempty"`
}

// NewEnvironment instantiates a new Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironment() *Environment {
	this := Environment{}
	var clientId string = ""
	this.ClientId = &clientId
	var id string = ""
	this.Id = &id
	var isProduction bool = false
	this.IsProduction = &isProduction
	var name string = ""
	this.Name = &name
	var organizationId string = ""
	this.OrganizationId = &organizationId
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	var clientId string = ""
	this.ClientId = &clientId
	var id string = ""
	this.Id = &id
	var isProduction bool = false
	this.IsProduction = &isProduction
	var name string = ""
	this.Name = &name
	var organizationId string = ""
	this.OrganizationId = &organizationId
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Environment) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Environment) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Environment) SetClientId(v string) {
	o.ClientId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Environment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Environment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Environment) SetId(v string) {
	o.Id = &v
}

// GetIsProduction returns the IsProduction field value if set, zero value otherwise.
func (o *Environment) GetIsProduction() bool {
	if o == nil || IsNil(o.IsProduction) {
		var ret bool
		return ret
	}
	return *o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetIsProductionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProduction) {
		return nil, false
	}
	return o.IsProduction, true
}

// HasIsProduction returns a boolean if a field has been set.
func (o *Environment) HasIsProduction() bool {
	if o != nil && !IsNil(o.IsProduction) {
		return true
	}

	return false
}

// SetIsProduction gets a reference to the given bool and assigns it to the IsProduction field.
func (o *Environment) SetIsProduction(v bool) {
	o.IsProduction = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Environment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Environment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Environment) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *Environment) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Environment) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *Environment) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Environment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Environment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Environment) SetType(v string) {
	o.Type = &v
}

func (o Environment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Environment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsProduction) {
		toSerialize["isProduction"] = o.IsProduction
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableEnvironment struct {
	value *Environment
	isSet bool
}

func (v NullableEnvironment) Get() *Environment {
	return v.value
}

func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


