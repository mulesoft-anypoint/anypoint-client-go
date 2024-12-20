/*
Identity Provider Management API

Description of Identity Provider API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idp

import (
	"encoding/json"
)

// checks if the SamlProviderPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlProviderPatch{}

// SamlProviderPatch struct for SamlProviderPatch
type SamlProviderPatch struct {
	Name *string `json:"name,omitempty"`
	Type *SamlProviderPatchType `json:"type,omitempty"`
	ServiceProvider *SamlProviderPatchServiceProvider `json:"service_provider,omitempty"`
	Saml *SamlProviderPatchSaml `json:"saml,omitempty"`
	LoginDisabled *bool `json:"login_disabled,omitempty"`
}

// NewSamlProviderPatch instantiates a new SamlProviderPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlProviderPatch() *SamlProviderPatch {
	this := SamlProviderPatch{}
	return &this
}

// NewSamlProviderPatchWithDefaults instantiates a new SamlProviderPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlProviderPatchWithDefaults() *SamlProviderPatch {
	this := SamlProviderPatch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SamlProviderPatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SamlProviderPatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SamlProviderPatch) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SamlProviderPatch) GetType() SamlProviderPatchType {
	if o == nil || IsNil(o.Type) {
		var ret SamlProviderPatchType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatch) GetTypeOk() (*SamlProviderPatchType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SamlProviderPatch) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SamlProviderPatchType and assigns it to the Type field.
func (o *SamlProviderPatch) SetType(v SamlProviderPatchType) {
	o.Type = &v
}

// GetServiceProvider returns the ServiceProvider field value if set, zero value otherwise.
func (o *SamlProviderPatch) GetServiceProvider() SamlProviderPatchServiceProvider {
	if o == nil || IsNil(o.ServiceProvider) {
		var ret SamlProviderPatchServiceProvider
		return ret
	}
	return *o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatch) GetServiceProviderOk() (*SamlProviderPatchServiceProvider, bool) {
	if o == nil || IsNil(o.ServiceProvider) {
		return nil, false
	}
	return o.ServiceProvider, true
}

// HasServiceProvider returns a boolean if a field has been set.
func (o *SamlProviderPatch) HasServiceProvider() bool {
	if o != nil && !IsNil(o.ServiceProvider) {
		return true
	}

	return false
}

// SetServiceProvider gets a reference to the given SamlProviderPatchServiceProvider and assigns it to the ServiceProvider field.
func (o *SamlProviderPatch) SetServiceProvider(v SamlProviderPatchServiceProvider) {
	o.ServiceProvider = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *SamlProviderPatch) GetSaml() SamlProviderPatchSaml {
	if o == nil || IsNil(o.Saml) {
		var ret SamlProviderPatchSaml
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatch) GetSamlOk() (*SamlProviderPatchSaml, bool) {
	if o == nil || IsNil(o.Saml) {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *SamlProviderPatch) HasSaml() bool {
	if o != nil && !IsNil(o.Saml) {
		return true
	}

	return false
}

// SetSaml gets a reference to the given SamlProviderPatchSaml and assigns it to the Saml field.
func (o *SamlProviderPatch) SetSaml(v SamlProviderPatchSaml) {
	o.Saml = &v
}

// GetLoginDisabled returns the LoginDisabled field value if set, zero value otherwise.
func (o *SamlProviderPatch) GetLoginDisabled() bool {
	if o == nil || IsNil(o.LoginDisabled) {
		var ret bool
		return ret
	}
	return *o.LoginDisabled
}

// GetLoginDisabledOk returns a tuple with the LoginDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatch) GetLoginDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LoginDisabled) {
		return nil, false
	}
	return o.LoginDisabled, true
}

// HasLoginDisabled returns a boolean if a field has been set.
func (o *SamlProviderPatch) HasLoginDisabled() bool {
	if o != nil && !IsNil(o.LoginDisabled) {
		return true
	}

	return false
}

// SetLoginDisabled gets a reference to the given bool and assigns it to the LoginDisabled field.
func (o *SamlProviderPatch) SetLoginDisabled(v bool) {
	o.LoginDisabled = &v
}

func (o SamlProviderPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlProviderPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ServiceProvider) {
		toSerialize["service_provider"] = o.ServiceProvider
	}
	if !IsNil(o.Saml) {
		toSerialize["saml"] = o.Saml
	}
	if !IsNil(o.LoginDisabled) {
		toSerialize["login_disabled"] = o.LoginDisabled
	}
	return toSerialize, nil
}

type NullableSamlProviderPatch struct {
	value *SamlProviderPatch
	isSet bool
}

func (v NullableSamlProviderPatch) Get() *SamlProviderPatch {
	return v.value
}

func (v *NullableSamlProviderPatch) Set(val *SamlProviderPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlProviderPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlProviderPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlProviderPatch(val *SamlProviderPatch) *NullableSamlProviderPatch {
	return &NullableSamlProviderPatch{value: val, isSet: true}
}

func (v NullableSamlProviderPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlProviderPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


