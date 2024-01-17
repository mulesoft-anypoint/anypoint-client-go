/*
Secret Group Keystore API

Secret Group Keystore API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretgroup_keystore

import (
	"encoding/json"
)

// checks if the CertificateDetailsBasicConstraints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateDetailsBasicConstraints{}

// CertificateDetailsBasicConstraints This identifies if the subject of the certificate is a CA
type CertificateDetailsBasicConstraints struct {
	// If set to true, indicates that this is a CA certificate.
	CertificateAuthority *bool `json:"certificateAuthority,omitempty"`
}

// NewCertificateDetailsBasicConstraints instantiates a new CertificateDetailsBasicConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateDetailsBasicConstraints() *CertificateDetailsBasicConstraints {
	this := CertificateDetailsBasicConstraints{}
	return &this
}

// NewCertificateDetailsBasicConstraintsWithDefaults instantiates a new CertificateDetailsBasicConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateDetailsBasicConstraintsWithDefaults() *CertificateDetailsBasicConstraints {
	this := CertificateDetailsBasicConstraints{}
	return &this
}

// GetCertificateAuthority returns the CertificateAuthority field value if set, zero value otherwise.
func (o *CertificateDetailsBasicConstraints) GetCertificateAuthority() bool {
	if o == nil || IsNil(o.CertificateAuthority) {
		var ret bool
		return ret
	}
	return *o.CertificateAuthority
}

// GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetailsBasicConstraints) GetCertificateAuthorityOk() (*bool, bool) {
	if o == nil || IsNil(o.CertificateAuthority) {
		return nil, false
	}
	return o.CertificateAuthority, true
}

// HasCertificateAuthority returns a boolean if a field has been set.
func (o *CertificateDetailsBasicConstraints) HasCertificateAuthority() bool {
	if o != nil && !IsNil(o.CertificateAuthority) {
		return true
	}

	return false
}

// SetCertificateAuthority gets a reference to the given bool and assigns it to the CertificateAuthority field.
func (o *CertificateDetailsBasicConstraints) SetCertificateAuthority(v bool) {
	o.CertificateAuthority = &v
}

func (o CertificateDetailsBasicConstraints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateDetailsBasicConstraints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificateAuthority) {
		toSerialize["certificateAuthority"] = o.CertificateAuthority
	}
	return toSerialize, nil
}

type NullableCertificateDetailsBasicConstraints struct {
	value *CertificateDetailsBasicConstraints
	isSet bool
}

func (v NullableCertificateDetailsBasicConstraints) Get() *CertificateDetailsBasicConstraints {
	return v.value
}

func (v *NullableCertificateDetailsBasicConstraints) Set(val *CertificateDetailsBasicConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateDetailsBasicConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateDetailsBasicConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateDetailsBasicConstraints(val *CertificateDetailsBasicConstraints) *NullableCertificateDetailsBasicConstraints {
	return &NullableCertificateDetailsBasicConstraints{value: val, isSet: true}
}

func (v NullableCertificateDetailsBasicConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateDetailsBasicConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


