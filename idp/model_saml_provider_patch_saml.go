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

// checks if the SamlProviderPatchSaml type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlProviderPatchSaml{}

// SamlProviderPatchSaml struct for SamlProviderPatchSaml
type SamlProviderPatchSaml struct {
	Audience *string `json:"audience,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	PublicKey []string `json:"public_key,omitempty"`
	ClaimsMapping *SamlProviderPostBodySamlClaimsMapping `json:"claims_mapping,omitempty"`
	IdpInitiatedSsoEnabled *bool `json:"idp_initiated_sso_enabled,omitempty"`
	SpInitiatedSsoEnabled *bool `json:"sp_initiated_sso_enabled,omitempty"`
	UseComposerAcsUrl *bool `json:"use_composer_acs_url,omitempty"`
	RequireEncryptedSamlAssertions *bool `json:"require_encrypted_saml_assertions,omitempty"`
}

// NewSamlProviderPatchSaml instantiates a new SamlProviderPatchSaml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlProviderPatchSaml() *SamlProviderPatchSaml {
	this := SamlProviderPatchSaml{}
	return &this
}

// NewSamlProviderPatchSamlWithDefaults instantiates a new SamlProviderPatchSaml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlProviderPatchSamlWithDefaults() *SamlProviderPatchSaml {
	this := SamlProviderPatchSaml{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *SamlProviderPatchSaml) GetAudience() string {
	if o == nil || IsNil(o.Audience) {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchSaml) GetAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.Audience) {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *SamlProviderPatchSaml) HasAudience() bool {
	if o != nil && !IsNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *SamlProviderPatchSaml) SetAudience(v string) {
	o.Audience = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SamlProviderPatchSaml) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchSaml) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SamlProviderPatchSaml) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SamlProviderPatchSaml) SetIssuer(v string) {
	o.Issuer = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *SamlProviderPatchSaml) GetPublicKey() []string {
	if o == nil || IsNil(o.PublicKey) {
		var ret []string
		return ret
	}
	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchSaml) GetPublicKeyOk() ([]string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *SamlProviderPatchSaml) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given []string and assigns it to the PublicKey field.
func (o *SamlProviderPatchSaml) SetPublicKey(v []string) {
	o.PublicKey = v
}

// GetClaimsMapping returns the ClaimsMapping field value if set, zero value otherwise.
func (o *SamlProviderPatchSaml) GetClaimsMapping() SamlProviderPostBodySamlClaimsMapping {
	if o == nil || IsNil(o.ClaimsMapping) {
		var ret SamlProviderPostBodySamlClaimsMapping
		return ret
	}
	return *o.ClaimsMapping
}

// GetClaimsMappingOk returns a tuple with the ClaimsMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchSaml) GetClaimsMappingOk() (*SamlProviderPostBodySamlClaimsMapping, bool) {
	if o == nil || IsNil(o.ClaimsMapping) {
		return nil, false
	}
	return o.ClaimsMapping, true
}

// HasClaimsMapping returns a boolean if a field has been set.
func (o *SamlProviderPatchSaml) HasClaimsMapping() bool {
	if o != nil && !IsNil(o.ClaimsMapping) {
		return true
	}

	return false
}

// SetClaimsMapping gets a reference to the given SamlProviderPostBodySamlClaimsMapping and assigns it to the ClaimsMapping field.
func (o *SamlProviderPatchSaml) SetClaimsMapping(v SamlProviderPostBodySamlClaimsMapping) {
	o.ClaimsMapping = &v
}

// GetIdpInitiatedSsoEnabled returns the IdpInitiatedSsoEnabled field value if set, zero value otherwise.
func (o *SamlProviderPatchSaml) GetIdpInitiatedSsoEnabled() bool {
	if o == nil || IsNil(o.IdpInitiatedSsoEnabled) {
		var ret bool
		return ret
	}
	return *o.IdpInitiatedSsoEnabled
}

// GetIdpInitiatedSsoEnabledOk returns a tuple with the IdpInitiatedSsoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchSaml) GetIdpInitiatedSsoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IdpInitiatedSsoEnabled) {
		return nil, false
	}
	return o.IdpInitiatedSsoEnabled, true
}

// HasIdpInitiatedSsoEnabled returns a boolean if a field has been set.
func (o *SamlProviderPatchSaml) HasIdpInitiatedSsoEnabled() bool {
	if o != nil && !IsNil(o.IdpInitiatedSsoEnabled) {
		return true
	}

	return false
}

// SetIdpInitiatedSsoEnabled gets a reference to the given bool and assigns it to the IdpInitiatedSsoEnabled field.
func (o *SamlProviderPatchSaml) SetIdpInitiatedSsoEnabled(v bool) {
	o.IdpInitiatedSsoEnabled = &v
}

// GetSpInitiatedSsoEnabled returns the SpInitiatedSsoEnabled field value if set, zero value otherwise.
func (o *SamlProviderPatchSaml) GetSpInitiatedSsoEnabled() bool {
	if o == nil || IsNil(o.SpInitiatedSsoEnabled) {
		var ret bool
		return ret
	}
	return *o.SpInitiatedSsoEnabled
}

// GetSpInitiatedSsoEnabledOk returns a tuple with the SpInitiatedSsoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchSaml) GetSpInitiatedSsoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SpInitiatedSsoEnabled) {
		return nil, false
	}
	return o.SpInitiatedSsoEnabled, true
}

// HasSpInitiatedSsoEnabled returns a boolean if a field has been set.
func (o *SamlProviderPatchSaml) HasSpInitiatedSsoEnabled() bool {
	if o != nil && !IsNil(o.SpInitiatedSsoEnabled) {
		return true
	}

	return false
}

// SetSpInitiatedSsoEnabled gets a reference to the given bool and assigns it to the SpInitiatedSsoEnabled field.
func (o *SamlProviderPatchSaml) SetSpInitiatedSsoEnabled(v bool) {
	o.SpInitiatedSsoEnabled = &v
}

// GetUseComposerAcsUrl returns the UseComposerAcsUrl field value if set, zero value otherwise.
func (o *SamlProviderPatchSaml) GetUseComposerAcsUrl() bool {
	if o == nil || IsNil(o.UseComposerAcsUrl) {
		var ret bool
		return ret
	}
	return *o.UseComposerAcsUrl
}

// GetUseComposerAcsUrlOk returns a tuple with the UseComposerAcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchSaml) GetUseComposerAcsUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseComposerAcsUrl) {
		return nil, false
	}
	return o.UseComposerAcsUrl, true
}

// HasUseComposerAcsUrl returns a boolean if a field has been set.
func (o *SamlProviderPatchSaml) HasUseComposerAcsUrl() bool {
	if o != nil && !IsNil(o.UseComposerAcsUrl) {
		return true
	}

	return false
}

// SetUseComposerAcsUrl gets a reference to the given bool and assigns it to the UseComposerAcsUrl field.
func (o *SamlProviderPatchSaml) SetUseComposerAcsUrl(v bool) {
	o.UseComposerAcsUrl = &v
}

// GetRequireEncryptedSamlAssertions returns the RequireEncryptedSamlAssertions field value if set, zero value otherwise.
func (o *SamlProviderPatchSaml) GetRequireEncryptedSamlAssertions() bool {
	if o == nil || IsNil(o.RequireEncryptedSamlAssertions) {
		var ret bool
		return ret
	}
	return *o.RequireEncryptedSamlAssertions
}

// GetRequireEncryptedSamlAssertionsOk returns a tuple with the RequireEncryptedSamlAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlProviderPatchSaml) GetRequireEncryptedSamlAssertionsOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireEncryptedSamlAssertions) {
		return nil, false
	}
	return o.RequireEncryptedSamlAssertions, true
}

// HasRequireEncryptedSamlAssertions returns a boolean if a field has been set.
func (o *SamlProviderPatchSaml) HasRequireEncryptedSamlAssertions() bool {
	if o != nil && !IsNil(o.RequireEncryptedSamlAssertions) {
		return true
	}

	return false
}

// SetRequireEncryptedSamlAssertions gets a reference to the given bool and assigns it to the RequireEncryptedSamlAssertions field.
func (o *SamlProviderPatchSaml) SetRequireEncryptedSamlAssertions(v bool) {
	o.RequireEncryptedSamlAssertions = &v
}

func (o SamlProviderPatchSaml) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlProviderPatchSaml) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audience) {
		toSerialize["audience"] = o.Audience
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.ClaimsMapping) {
		toSerialize["claims_mapping"] = o.ClaimsMapping
	}
	if !IsNil(o.IdpInitiatedSsoEnabled) {
		toSerialize["idp_initiated_sso_enabled"] = o.IdpInitiatedSsoEnabled
	}
	if !IsNil(o.SpInitiatedSsoEnabled) {
		toSerialize["sp_initiated_sso_enabled"] = o.SpInitiatedSsoEnabled
	}
	if !IsNil(o.UseComposerAcsUrl) {
		toSerialize["use_composer_acs_url"] = o.UseComposerAcsUrl
	}
	if !IsNil(o.RequireEncryptedSamlAssertions) {
		toSerialize["require_encrypted_saml_assertions"] = o.RequireEncryptedSamlAssertions
	}
	return toSerialize, nil
}

type NullableSamlProviderPatchSaml struct {
	value *SamlProviderPatchSaml
	isSet bool
}

func (v NullableSamlProviderPatchSaml) Get() *SamlProviderPatchSaml {
	return v.value
}

func (v *NullableSamlProviderPatchSaml) Set(val *SamlProviderPatchSaml) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlProviderPatchSaml) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlProviderPatchSaml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlProviderPatchSaml(val *SamlProviderPatchSaml) *NullableSamlProviderPatchSaml {
	return &NullableSamlProviderPatchSaml{value: val, isSet: true}
}

func (v NullableSamlProviderPatchSaml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlProviderPatchSaml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


