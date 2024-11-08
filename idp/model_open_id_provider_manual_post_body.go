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

// checks if the OpenIDProviderManualPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIDProviderManualPostBody{}

// OpenIDProviderManualPostBody struct for OpenIDProviderManualPostBody
type OpenIDProviderManualPostBody struct {
	Name string `json:"name"`
	ArcNamespace *string `json:"arc_namespace,omitempty"`
	Type OpenIDProviderManualPostBodyType `json:"type"`
	OidcProvider OpenIDProviderManualPostBodyOidcProvider `json:"oidc_provider"`
	AllowUntrustedCertificates *bool `json:"allow_untrusted_certificates,omitempty"`
	LoginDisabled *bool `json:"login_disabled,omitempty"`
}

// NewOpenIDProviderManualPostBody instantiates a new OpenIDProviderManualPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIDProviderManualPostBody(name string, type_ OpenIDProviderManualPostBodyType, oidcProvider OpenIDProviderManualPostBodyOidcProvider) *OpenIDProviderManualPostBody {
	this := OpenIDProviderManualPostBody{}
	this.Name = name
	this.Type = type_
	this.OidcProvider = oidcProvider
	return &this
}

// NewOpenIDProviderManualPostBodyWithDefaults instantiates a new OpenIDProviderManualPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIDProviderManualPostBodyWithDefaults() *OpenIDProviderManualPostBody {
	this := OpenIDProviderManualPostBody{}
	return &this
}

// GetName returns the Name field value
func (o *OpenIDProviderManualPostBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderManualPostBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpenIDProviderManualPostBody) SetName(v string) {
	o.Name = v
}

// GetArcNamespace returns the ArcNamespace field value if set, zero value otherwise.
func (o *OpenIDProviderManualPostBody) GetArcNamespace() string {
	if o == nil || IsNil(o.ArcNamespace) {
		var ret string
		return ret
	}
	return *o.ArcNamespace
}

// GetArcNamespaceOk returns a tuple with the ArcNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderManualPostBody) GetArcNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.ArcNamespace) {
		return nil, false
	}
	return o.ArcNamespace, true
}

// HasArcNamespace returns a boolean if a field has been set.
func (o *OpenIDProviderManualPostBody) HasArcNamespace() bool {
	if o != nil && !IsNil(o.ArcNamespace) {
		return true
	}

	return false
}

// SetArcNamespace gets a reference to the given string and assigns it to the ArcNamespace field.
func (o *OpenIDProviderManualPostBody) SetArcNamespace(v string) {
	o.ArcNamespace = &v
}

// GetType returns the Type field value
func (o *OpenIDProviderManualPostBody) GetType() OpenIDProviderManualPostBodyType {
	if o == nil {
		var ret OpenIDProviderManualPostBodyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderManualPostBody) GetTypeOk() (*OpenIDProviderManualPostBodyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OpenIDProviderManualPostBody) SetType(v OpenIDProviderManualPostBodyType) {
	o.Type = v
}

// GetOidcProvider returns the OidcProvider field value
func (o *OpenIDProviderManualPostBody) GetOidcProvider() OpenIDProviderManualPostBodyOidcProvider {
	if o == nil {
		var ret OpenIDProviderManualPostBodyOidcProvider
		return ret
	}

	return o.OidcProvider
}

// GetOidcProviderOk returns a tuple with the OidcProvider field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderManualPostBody) GetOidcProviderOk() (*OpenIDProviderManualPostBodyOidcProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcProvider, true
}

// SetOidcProvider sets field value
func (o *OpenIDProviderManualPostBody) SetOidcProvider(v OpenIDProviderManualPostBodyOidcProvider) {
	o.OidcProvider = v
}

// GetAllowUntrustedCertificates returns the AllowUntrustedCertificates field value if set, zero value otherwise.
func (o *OpenIDProviderManualPostBody) GetAllowUntrustedCertificates() bool {
	if o == nil || IsNil(o.AllowUntrustedCertificates) {
		var ret bool
		return ret
	}
	return *o.AllowUntrustedCertificates
}

// GetAllowUntrustedCertificatesOk returns a tuple with the AllowUntrustedCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderManualPostBody) GetAllowUntrustedCertificatesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUntrustedCertificates) {
		return nil, false
	}
	return o.AllowUntrustedCertificates, true
}

// HasAllowUntrustedCertificates returns a boolean if a field has been set.
func (o *OpenIDProviderManualPostBody) HasAllowUntrustedCertificates() bool {
	if o != nil && !IsNil(o.AllowUntrustedCertificates) {
		return true
	}

	return false
}

// SetAllowUntrustedCertificates gets a reference to the given bool and assigns it to the AllowUntrustedCertificates field.
func (o *OpenIDProviderManualPostBody) SetAllowUntrustedCertificates(v bool) {
	o.AllowUntrustedCertificates = &v
}

// GetLoginDisabled returns the LoginDisabled field value if set, zero value otherwise.
func (o *OpenIDProviderManualPostBody) GetLoginDisabled() bool {
	if o == nil || IsNil(o.LoginDisabled) {
		var ret bool
		return ret
	}
	return *o.LoginDisabled
}

// GetLoginDisabledOk returns a tuple with the LoginDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderManualPostBody) GetLoginDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LoginDisabled) {
		return nil, false
	}
	return o.LoginDisabled, true
}

// HasLoginDisabled returns a boolean if a field has been set.
func (o *OpenIDProviderManualPostBody) HasLoginDisabled() bool {
	if o != nil && !IsNil(o.LoginDisabled) {
		return true
	}

	return false
}

// SetLoginDisabled gets a reference to the given bool and assigns it to the LoginDisabled field.
func (o *OpenIDProviderManualPostBody) SetLoginDisabled(v bool) {
	o.LoginDisabled = &v
}

func (o OpenIDProviderManualPostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIDProviderManualPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.ArcNamespace) {
		toSerialize["arc_namespace"] = o.ArcNamespace
	}
	toSerialize["type"] = o.Type
	toSerialize["oidc_provider"] = o.OidcProvider
	if !IsNil(o.AllowUntrustedCertificates) {
		toSerialize["allow_untrusted_certificates"] = o.AllowUntrustedCertificates
	}
	if !IsNil(o.LoginDisabled) {
		toSerialize["login_disabled"] = o.LoginDisabled
	}
	return toSerialize, nil
}

type NullableOpenIDProviderManualPostBody struct {
	value *OpenIDProviderManualPostBody
	isSet bool
}

func (v NullableOpenIDProviderManualPostBody) Get() *OpenIDProviderManualPostBody {
	return v.value
}

func (v *NullableOpenIDProviderManualPostBody) Set(val *OpenIDProviderManualPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIDProviderManualPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIDProviderManualPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIDProviderManualPostBody(val *OpenIDProviderManualPostBody) *NullableOpenIDProviderManualPostBody {
	return &NullableOpenIDProviderManualPostBody{value: val, isSet: true}
}

func (v NullableOpenIDProviderManualPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIDProviderManualPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


