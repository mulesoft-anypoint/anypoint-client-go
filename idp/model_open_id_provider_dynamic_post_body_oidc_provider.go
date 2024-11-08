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

// checks if the OpenIDProviderDynamicPostBodyOidcProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIDProviderDynamicPostBodyOidcProvider{}

// OpenIDProviderDynamicPostBodyOidcProvider struct for OpenIDProviderDynamicPostBodyOidcProvider
type OpenIDProviderDynamicPostBodyOidcProvider struct {
	Client OpenIDProviderDynamicPostBodyOidcProviderClient `json:"client"`
	Issuer *string `json:"issuer,omitempty"`
	Urls OpenIDProviderManualPostBodyOidcProviderUrls `json:"urls"`
	GroupScope *string `json:"group_scope,omitempty"`
	ClaimsMapping *OpenIDProviderManualPostBodyOidcProviderClaimsMapping `json:"claims_mapping,omitempty"`
}

// NewOpenIDProviderDynamicPostBodyOidcProvider instantiates a new OpenIDProviderDynamicPostBodyOidcProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIDProviderDynamicPostBodyOidcProvider(client OpenIDProviderDynamicPostBodyOidcProviderClient, urls OpenIDProviderManualPostBodyOidcProviderUrls) *OpenIDProviderDynamicPostBodyOidcProvider {
	this := OpenIDProviderDynamicPostBodyOidcProvider{}
	this.Client = client
	this.Urls = urls
	return &this
}

// NewOpenIDProviderDynamicPostBodyOidcProviderWithDefaults instantiates a new OpenIDProviderDynamicPostBodyOidcProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIDProviderDynamicPostBodyOidcProviderWithDefaults() *OpenIDProviderDynamicPostBodyOidcProvider {
	this := OpenIDProviderDynamicPostBodyOidcProvider{}
	return &this
}

// GetClient returns the Client field value
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetClient() OpenIDProviderDynamicPostBodyOidcProviderClient {
	if o == nil {
		var ret OpenIDProviderDynamicPostBodyOidcProviderClient
		return ret
	}

	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetClientOk() (*OpenIDProviderDynamicPostBodyOidcProviderClient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value
func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetClient(v OpenIDProviderDynamicPostBodyOidcProviderClient) {
	o.Client = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetIssuer(v string) {
	o.Issuer = &v
}

// GetUrls returns the Urls field value
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetUrls() OpenIDProviderManualPostBodyOidcProviderUrls {
	if o == nil {
		var ret OpenIDProviderManualPostBodyOidcProviderUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetUrlsOk() (*OpenIDProviderManualPostBodyOidcProviderUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetUrls(v OpenIDProviderManualPostBodyOidcProviderUrls) {
	o.Urls = v
}

// GetGroupScope returns the GroupScope field value if set, zero value otherwise.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetGroupScope() string {
	if o == nil || IsNil(o.GroupScope) {
		var ret string
		return ret
	}
	return *o.GroupScope
}

// GetGroupScopeOk returns a tuple with the GroupScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetGroupScopeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupScope) {
		return nil, false
	}
	return o.GroupScope, true
}

// HasGroupScope returns a boolean if a field has been set.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) HasGroupScope() bool {
	if o != nil && !IsNil(o.GroupScope) {
		return true
	}

	return false
}

// SetGroupScope gets a reference to the given string and assigns it to the GroupScope field.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetGroupScope(v string) {
	o.GroupScope = &v
}

// GetClaimsMapping returns the ClaimsMapping field value if set, zero value otherwise.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetClaimsMapping() OpenIDProviderManualPostBodyOidcProviderClaimsMapping {
	if o == nil || IsNil(o.ClaimsMapping) {
		var ret OpenIDProviderManualPostBodyOidcProviderClaimsMapping
		return ret
	}
	return *o.ClaimsMapping
}

// GetClaimsMappingOk returns a tuple with the ClaimsMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) GetClaimsMappingOk() (*OpenIDProviderManualPostBodyOidcProviderClaimsMapping, bool) {
	if o == nil || IsNil(o.ClaimsMapping) {
		return nil, false
	}
	return o.ClaimsMapping, true
}

// HasClaimsMapping returns a boolean if a field has been set.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) HasClaimsMapping() bool {
	if o != nil && !IsNil(o.ClaimsMapping) {
		return true
	}

	return false
}

// SetClaimsMapping gets a reference to the given OpenIDProviderManualPostBodyOidcProviderClaimsMapping and assigns it to the ClaimsMapping field.
func (o *OpenIDProviderDynamicPostBodyOidcProvider) SetClaimsMapping(v OpenIDProviderManualPostBodyOidcProviderClaimsMapping) {
	o.ClaimsMapping = &v
}

func (o OpenIDProviderDynamicPostBodyOidcProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIDProviderDynamicPostBodyOidcProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client"] = o.Client
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	toSerialize["urls"] = o.Urls
	if !IsNil(o.GroupScope) {
		toSerialize["group_scope"] = o.GroupScope
	}
	if !IsNil(o.ClaimsMapping) {
		toSerialize["claims_mapping"] = o.ClaimsMapping
	}
	return toSerialize, nil
}

type NullableOpenIDProviderDynamicPostBodyOidcProvider struct {
	value *OpenIDProviderDynamicPostBodyOidcProvider
	isSet bool
}

func (v NullableOpenIDProviderDynamicPostBodyOidcProvider) Get() *OpenIDProviderDynamicPostBodyOidcProvider {
	return v.value
}

func (v *NullableOpenIDProviderDynamicPostBodyOidcProvider) Set(val *OpenIDProviderDynamicPostBodyOidcProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIDProviderDynamicPostBodyOidcProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIDProviderDynamicPostBodyOidcProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIDProviderDynamicPostBodyOidcProvider(val *OpenIDProviderDynamicPostBodyOidcProvider) *NullableOpenIDProviderDynamicPostBodyOidcProvider {
	return &NullableOpenIDProviderDynamicPostBodyOidcProvider{value: val, isSet: true}
}

func (v NullableOpenIDProviderDynamicPostBodyOidcProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIDProviderDynamicPostBodyOidcProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


