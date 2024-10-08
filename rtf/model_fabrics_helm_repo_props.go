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

// checks if the FabricsHelmRepoProps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricsHelmRepoProps{}

// FabricsHelmRepoProps struct for FabricsHelmRepoProps
type FabricsHelmRepoProps struct {
	// the fabrics image registry endpoint
	RTF_IMAGE_REGISTRY_ENDPOINT *string `json:"RTF_IMAGE_REGISTRY_ENDPOINT,omitempty"`
	// the username to access the fabrics registry
	RTF_IMAGE_REGISTRY_USER *string `json:"RTF_IMAGE_REGISTRY_USER,omitempty"`
	// the password to access the fabrics registry
	RTF_IMAGE_REGISTRY_PASSWORD *string `json:"RTF_IMAGE_REGISTRY_PASSWORD,omitempty"`
}

// NewFabricsHelmRepoProps instantiates a new FabricsHelmRepoProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricsHelmRepoProps() *FabricsHelmRepoProps {
	this := FabricsHelmRepoProps{}
	return &this
}

// NewFabricsHelmRepoPropsWithDefaults instantiates a new FabricsHelmRepoProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricsHelmRepoPropsWithDefaults() *FabricsHelmRepoProps {
	this := FabricsHelmRepoProps{}
	return &this
}

// GetRTF_IMAGE_REGISTRY_ENDPOINT returns the RTF_IMAGE_REGISTRY_ENDPOINT field value if set, zero value otherwise.
func (o *FabricsHelmRepoProps) GetRTF_IMAGE_REGISTRY_ENDPOINT() string {
	if o == nil || IsNil(o.RTF_IMAGE_REGISTRY_ENDPOINT) {
		var ret string
		return ret
	}
	return *o.RTF_IMAGE_REGISTRY_ENDPOINT
}

// GetRTF_IMAGE_REGISTRY_ENDPOINTOk returns a tuple with the RTF_IMAGE_REGISTRY_ENDPOINT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHelmRepoProps) GetRTF_IMAGE_REGISTRY_ENDPOINTOk() (*string, bool) {
	if o == nil || IsNil(o.RTF_IMAGE_REGISTRY_ENDPOINT) {
		return nil, false
	}
	return o.RTF_IMAGE_REGISTRY_ENDPOINT, true
}

// HasRTF_IMAGE_REGISTRY_ENDPOINT returns a boolean if a field has been set.
func (o *FabricsHelmRepoProps) HasRTF_IMAGE_REGISTRY_ENDPOINT() bool {
	if o != nil && !IsNil(o.RTF_IMAGE_REGISTRY_ENDPOINT) {
		return true
	}

	return false
}

// SetRTF_IMAGE_REGISTRY_ENDPOINT gets a reference to the given string and assigns it to the RTF_IMAGE_REGISTRY_ENDPOINT field.
func (o *FabricsHelmRepoProps) SetRTF_IMAGE_REGISTRY_ENDPOINT(v string) {
	o.RTF_IMAGE_REGISTRY_ENDPOINT = &v
}

// GetRTF_IMAGE_REGISTRY_USER returns the RTF_IMAGE_REGISTRY_USER field value if set, zero value otherwise.
func (o *FabricsHelmRepoProps) GetRTF_IMAGE_REGISTRY_USER() string {
	if o == nil || IsNil(o.RTF_IMAGE_REGISTRY_USER) {
		var ret string
		return ret
	}
	return *o.RTF_IMAGE_REGISTRY_USER
}

// GetRTF_IMAGE_REGISTRY_USEROk returns a tuple with the RTF_IMAGE_REGISTRY_USER field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHelmRepoProps) GetRTF_IMAGE_REGISTRY_USEROk() (*string, bool) {
	if o == nil || IsNil(o.RTF_IMAGE_REGISTRY_USER) {
		return nil, false
	}
	return o.RTF_IMAGE_REGISTRY_USER, true
}

// HasRTF_IMAGE_REGISTRY_USER returns a boolean if a field has been set.
func (o *FabricsHelmRepoProps) HasRTF_IMAGE_REGISTRY_USER() bool {
	if o != nil && !IsNil(o.RTF_IMAGE_REGISTRY_USER) {
		return true
	}

	return false
}

// SetRTF_IMAGE_REGISTRY_USER gets a reference to the given string and assigns it to the RTF_IMAGE_REGISTRY_USER field.
func (o *FabricsHelmRepoProps) SetRTF_IMAGE_REGISTRY_USER(v string) {
	o.RTF_IMAGE_REGISTRY_USER = &v
}

// GetRTF_IMAGE_REGISTRY_PASSWORD returns the RTF_IMAGE_REGISTRY_PASSWORD field value if set, zero value otherwise.
func (o *FabricsHelmRepoProps) GetRTF_IMAGE_REGISTRY_PASSWORD() string {
	if o == nil || IsNil(o.RTF_IMAGE_REGISTRY_PASSWORD) {
		var ret string
		return ret
	}
	return *o.RTF_IMAGE_REGISTRY_PASSWORD
}

// GetRTF_IMAGE_REGISTRY_PASSWORDOk returns a tuple with the RTF_IMAGE_REGISTRY_PASSWORD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricsHelmRepoProps) GetRTF_IMAGE_REGISTRY_PASSWORDOk() (*string, bool) {
	if o == nil || IsNil(o.RTF_IMAGE_REGISTRY_PASSWORD) {
		return nil, false
	}
	return o.RTF_IMAGE_REGISTRY_PASSWORD, true
}

// HasRTF_IMAGE_REGISTRY_PASSWORD returns a boolean if a field has been set.
func (o *FabricsHelmRepoProps) HasRTF_IMAGE_REGISTRY_PASSWORD() bool {
	if o != nil && !IsNil(o.RTF_IMAGE_REGISTRY_PASSWORD) {
		return true
	}

	return false
}

// SetRTF_IMAGE_REGISTRY_PASSWORD gets a reference to the given string and assigns it to the RTF_IMAGE_REGISTRY_PASSWORD field.
func (o *FabricsHelmRepoProps) SetRTF_IMAGE_REGISTRY_PASSWORD(v string) {
	o.RTF_IMAGE_REGISTRY_PASSWORD = &v
}

func (o FabricsHelmRepoProps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricsHelmRepoProps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RTF_IMAGE_REGISTRY_ENDPOINT) {
		toSerialize["RTF_IMAGE_REGISTRY_ENDPOINT"] = o.RTF_IMAGE_REGISTRY_ENDPOINT
	}
	if !IsNil(o.RTF_IMAGE_REGISTRY_USER) {
		toSerialize["RTF_IMAGE_REGISTRY_USER"] = o.RTF_IMAGE_REGISTRY_USER
	}
	if !IsNil(o.RTF_IMAGE_REGISTRY_PASSWORD) {
		toSerialize["RTF_IMAGE_REGISTRY_PASSWORD"] = o.RTF_IMAGE_REGISTRY_PASSWORD
	}
	return toSerialize, nil
}

type NullableFabricsHelmRepoProps struct {
	value *FabricsHelmRepoProps
	isSet bool
}

func (v NullableFabricsHelmRepoProps) Get() *FabricsHelmRepoProps {
	return v.value
}

func (v *NullableFabricsHelmRepoProps) Set(val *FabricsHelmRepoProps) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricsHelmRepoProps) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricsHelmRepoProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricsHelmRepoProps(val *FabricsHelmRepoProps) *NullableFabricsHelmRepoProps {
	return &NullableFabricsHelmRepoProps{value: val, isSet: true}
}

func (v NullableFabricsHelmRepoProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricsHelmRepoProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


