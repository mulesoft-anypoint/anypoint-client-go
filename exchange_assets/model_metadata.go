/*
Exchange Assets

Description of the Exchange Assets API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package exchange_assets

import (
	"encoding/json"
)

// checks if the Metadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metadata{}

// Metadata struct for Metadata
type Metadata struct {
	Classifier *string `json:"classifier,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata() *Metadata {
	this := Metadata{}
	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetClassifier returns the Classifier field value if set, zero value otherwise.
func (o *Metadata) GetClassifier() string {
	if o == nil || IsNil(o.Classifier) {
		var ret string
		return ret
	}
	return *o.Classifier
}

// GetClassifierOk returns a tuple with the Classifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetClassifierOk() (*string, bool) {
	if o == nil || IsNil(o.Classifier) {
		return nil, false
	}
	return o.Classifier, true
}

// HasClassifier returns a boolean if a field has been set.
func (o *Metadata) HasClassifier() bool {
	if o != nil && !IsNil(o.Classifier) {
		return true
	}

	return false
}

// SetClassifier gets a reference to the given string and assigns it to the Classifier field.
func (o *Metadata) SetClassifier(v string) {
	o.Classifier = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Metadata) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Metadata) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Metadata) SetTags(v []string) {
	o.Tags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Metadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Metadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Metadata) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Metadata) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Metadata) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Metadata) SetType(v string) {
	o.Type = &v
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Classifier) {
		toSerialize["classifier"] = o.Classifier
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


