/*
Secret Group Truststore API

Secret Group Truststore API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_truststore

import (
	"encoding/json"
)

// checks if the TruststoreDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TruststoreDetails{}

// TruststoreDetails Details about each of the trusted certificate from the truststore
type TruststoreDetails struct {
	CertificateEntries []CertificateEntry `json:"certificateEntries,omitempty"`
}

// NewTruststoreDetails instantiates a new TruststoreDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTruststoreDetails() *TruststoreDetails {
	this := TruststoreDetails{}
	return &this
}

// NewTruststoreDetailsWithDefaults instantiates a new TruststoreDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTruststoreDetailsWithDefaults() *TruststoreDetails {
	this := TruststoreDetails{}
	return &this
}

// GetCertificateEntries returns the CertificateEntries field value if set, zero value otherwise.
func (o *TruststoreDetails) GetCertificateEntries() []CertificateEntry {
	if o == nil || IsNil(o.CertificateEntries) {
		var ret []CertificateEntry
		return ret
	}
	return o.CertificateEntries
}

// GetCertificateEntriesOk returns a tuple with the CertificateEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TruststoreDetails) GetCertificateEntriesOk() ([]CertificateEntry, bool) {
	if o == nil || IsNil(o.CertificateEntries) {
		return nil, false
	}
	return o.CertificateEntries, true
}

// HasCertificateEntries returns a boolean if a field has been set.
func (o *TruststoreDetails) HasCertificateEntries() bool {
	if o != nil && !IsNil(o.CertificateEntries) {
		return true
	}

	return false
}

// SetCertificateEntries gets a reference to the given []CertificateEntry and assigns it to the CertificateEntries field.
func (o *TruststoreDetails) SetCertificateEntries(v []CertificateEntry) {
	o.CertificateEntries = v
}

func (o TruststoreDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TruststoreDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificateEntries) {
		toSerialize["certificateEntries"] = o.CertificateEntries
	}
	return toSerialize, nil
}

type NullableTruststoreDetails struct {
	value *TruststoreDetails
	isSet bool
}

func (v NullableTruststoreDetails) Get() *TruststoreDetails {
	return v.value
}

func (v *NullableTruststoreDetails) Set(val *TruststoreDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTruststoreDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTruststoreDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTruststoreDetails(val *TruststoreDetails) *NullableTruststoreDetails {
	return &NullableTruststoreDetails{value: val, isSet: true}
}

func (v NullableTruststoreDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTruststoreDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


