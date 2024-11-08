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

// checks if the OpenIDProviderGetOidcProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIDProviderGetOidcProvider{}

// OpenIDProviderGetOidcProvider struct for OpenIDProviderGetOidcProvider
type OpenIDProviderGetOidcProvider struct {
	Client OpenIDProviderGetOidcProviderClient `json:"client"`
	Issuer string `json:"issuer"`
	Urls OpenIDProviderGetOidcProviderUrls `json:"urls"`
	GroupScope *string `json:"group_scope,omitempty"`
	ClaimsMapping *OpenIDProviderManualPostBodyOidcProviderClaimsMapping `json:"claims_mapping,omitempty"`
}

// NewOpenIDProviderGetOidcProvider instantiates a new OpenIDProviderGetOidcProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIDProviderGetOidcProvider(client OpenIDProviderGetOidcProviderClient, issuer string, urls OpenIDProviderGetOidcProviderUrls) *OpenIDProviderGetOidcProvider {
	this := OpenIDProviderGetOidcProvider{}
	this.Client = client
	this.Issuer = issuer
	this.Urls = urls
	return &this
}

// NewOpenIDProviderGetOidcProviderWithDefaults instantiates a new OpenIDProviderGetOidcProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIDProviderGetOidcProviderWithDefaults() *OpenIDProviderGetOidcProvider {
	this := OpenIDProviderGetOidcProvider{}
	return &this
}

// GetClient returns the Client field value
func (o *OpenIDProviderGetOidcProvider) GetClient() OpenIDProviderGetOidcProviderClient {
	if o == nil {
		var ret OpenIDProviderGetOidcProviderClient
		return ret
	}

	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderGetOidcProvider) GetClientOk() (*OpenIDProviderGetOidcProviderClient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value
func (o *OpenIDProviderGetOidcProvider) SetClient(v OpenIDProviderGetOidcProviderClient) {
	o.Client = v
}

// GetIssuer returns the Issuer field value
func (o *OpenIDProviderGetOidcProvider) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderGetOidcProvider) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *OpenIDProviderGetOidcProvider) SetIssuer(v string) {
	o.Issuer = v
}

// GetUrls returns the Urls field value
func (o *OpenIDProviderGetOidcProvider) GetUrls() OpenIDProviderGetOidcProviderUrls {
	if o == nil {
		var ret OpenIDProviderGetOidcProviderUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderGetOidcProvider) GetUrlsOk() (*OpenIDProviderGetOidcProviderUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *OpenIDProviderGetOidcProvider) SetUrls(v OpenIDProviderGetOidcProviderUrls) {
	o.Urls = v
}

// GetGroupScope returns the GroupScope field value if set, zero value otherwise.
func (o *OpenIDProviderGetOidcProvider) GetGroupScope() string {
	if o == nil || IsNil(o.GroupScope) {
		var ret string
		return ret
	}
	return *o.GroupScope
}

// GetGroupScopeOk returns a tuple with the GroupScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderGetOidcProvider) GetGroupScopeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupScope) {
		return nil, false
	}
	return o.GroupScope, true
}

// HasGroupScope returns a boolean if a field has been set.
func (o *OpenIDProviderGetOidcProvider) HasGroupScope() bool {
	if o != nil && !IsNil(o.GroupScope) {
		return true
	}

	return false
}

// SetGroupScope gets a reference to the given string and assigns it to the GroupScope field.
func (o *OpenIDProviderGetOidcProvider) SetGroupScope(v string) {
	o.GroupScope = &v
}

// GetClaimsMapping returns the ClaimsMapping field value if set, zero value otherwise.
func (o *OpenIDProviderGetOidcProvider) GetClaimsMapping() OpenIDProviderManualPostBodyOidcProviderClaimsMapping {
	if o == nil || IsNil(o.ClaimsMapping) {
		var ret OpenIDProviderManualPostBodyOidcProviderClaimsMapping
		return ret
	}
	return *o.ClaimsMapping
}

// GetClaimsMappingOk returns a tuple with the ClaimsMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderGetOidcProvider) GetClaimsMappingOk() (*OpenIDProviderManualPostBodyOidcProviderClaimsMapping, bool) {
	if o == nil || IsNil(o.ClaimsMapping) {
		return nil, false
	}
	return o.ClaimsMapping, true
}

// HasClaimsMapping returns a boolean if a field has been set.
func (o *OpenIDProviderGetOidcProvider) HasClaimsMapping() bool {
	if o != nil && !IsNil(o.ClaimsMapping) {
		return true
	}

	return false
}

// SetClaimsMapping gets a reference to the given OpenIDProviderManualPostBodyOidcProviderClaimsMapping and assigns it to the ClaimsMapping field.
func (o *OpenIDProviderGetOidcProvider) SetClaimsMapping(v OpenIDProviderManualPostBodyOidcProviderClaimsMapping) {
	o.ClaimsMapping = &v
}

func (o OpenIDProviderGetOidcProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIDProviderGetOidcProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client"] = o.Client
	toSerialize["issuer"] = o.Issuer
	toSerialize["urls"] = o.Urls
	if !IsNil(o.GroupScope) {
		toSerialize["group_scope"] = o.GroupScope
	}
	if !IsNil(o.ClaimsMapping) {
		toSerialize["claims_mapping"] = o.ClaimsMapping
	}
	return toSerialize, nil
}

type NullableOpenIDProviderGetOidcProvider struct {
	value *OpenIDProviderGetOidcProvider
	isSet bool
}

func (v NullableOpenIDProviderGetOidcProvider) Get() *OpenIDProviderGetOidcProvider {
	return v.value
}

func (v *NullableOpenIDProviderGetOidcProvider) Set(val *OpenIDProviderGetOidcProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIDProviderGetOidcProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIDProviderGetOidcProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIDProviderGetOidcProvider(val *OpenIDProviderGetOidcProvider) *NullableOpenIDProviderGetOidcProvider {
	return &NullableOpenIDProviderGetOidcProvider{value: val, isSet: true}
}

func (v NullableOpenIDProviderGetOidcProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIDProviderGetOidcProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


