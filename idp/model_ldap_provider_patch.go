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

// checks if the LdapProviderPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapProviderPatch{}

// LdapProviderPatch struct for LdapProviderPatch
type LdapProviderPatch struct {
	Name *string `json:"name,omitempty"`
	Type *SamlProviderPatchType `json:"type,omitempty"`
	Connection *LdapProviderPatchConnection `json:"connection,omitempty"`
	SearchBases *LdapProviderPatchSearchBases `json:"search_bases,omitempty"`
	Dns *LdapProviderPatchSearchBases `json:"dns,omitempty"`
	Filters *LdapProviderPatchFilters `json:"filters,omitempty"`
	UserMapping *LdapProviderPatchUserMapping `json:"user_mapping,omitempty"`
	GroupMapping *LdapProviderPatchGroupMapping `json:"group_mapping,omitempty"`
}

// NewLdapProviderPatch instantiates a new LdapProviderPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapProviderPatch() *LdapProviderPatch {
	this := LdapProviderPatch{}
	return &this
}

// NewLdapProviderPatchWithDefaults instantiates a new LdapProviderPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapProviderPatchWithDefaults() *LdapProviderPatch {
	this := LdapProviderPatch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LdapProviderPatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LdapProviderPatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LdapProviderPatch) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LdapProviderPatch) GetType() SamlProviderPatchType {
	if o == nil || IsNil(o.Type) {
		var ret SamlProviderPatchType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatch) GetTypeOk() (*SamlProviderPatchType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LdapProviderPatch) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SamlProviderPatchType and assigns it to the Type field.
func (o *LdapProviderPatch) SetType(v SamlProviderPatchType) {
	o.Type = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *LdapProviderPatch) GetConnection() LdapProviderPatchConnection {
	if o == nil || IsNil(o.Connection) {
		var ret LdapProviderPatchConnection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatch) GetConnectionOk() (*LdapProviderPatchConnection, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *LdapProviderPatch) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given LdapProviderPatchConnection and assigns it to the Connection field.
func (o *LdapProviderPatch) SetConnection(v LdapProviderPatchConnection) {
	o.Connection = &v
}

// GetSearchBases returns the SearchBases field value if set, zero value otherwise.
func (o *LdapProviderPatch) GetSearchBases() LdapProviderPatchSearchBases {
	if o == nil || IsNil(o.SearchBases) {
		var ret LdapProviderPatchSearchBases
		return ret
	}
	return *o.SearchBases
}

// GetSearchBasesOk returns a tuple with the SearchBases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatch) GetSearchBasesOk() (*LdapProviderPatchSearchBases, bool) {
	if o == nil || IsNil(o.SearchBases) {
		return nil, false
	}
	return o.SearchBases, true
}

// HasSearchBases returns a boolean if a field has been set.
func (o *LdapProviderPatch) HasSearchBases() bool {
	if o != nil && !IsNil(o.SearchBases) {
		return true
	}

	return false
}

// SetSearchBases gets a reference to the given LdapProviderPatchSearchBases and assigns it to the SearchBases field.
func (o *LdapProviderPatch) SetSearchBases(v LdapProviderPatchSearchBases) {
	o.SearchBases = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *LdapProviderPatch) GetDns() LdapProviderPatchSearchBases {
	if o == nil || IsNil(o.Dns) {
		var ret LdapProviderPatchSearchBases
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatch) GetDnsOk() (*LdapProviderPatchSearchBases, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *LdapProviderPatch) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given LdapProviderPatchSearchBases and assigns it to the Dns field.
func (o *LdapProviderPatch) SetDns(v LdapProviderPatchSearchBases) {
	o.Dns = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *LdapProviderPatch) GetFilters() LdapProviderPatchFilters {
	if o == nil || IsNil(o.Filters) {
		var ret LdapProviderPatchFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatch) GetFiltersOk() (*LdapProviderPatchFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *LdapProviderPatch) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given LdapProviderPatchFilters and assigns it to the Filters field.
func (o *LdapProviderPatch) SetFilters(v LdapProviderPatchFilters) {
	o.Filters = &v
}

// GetUserMapping returns the UserMapping field value if set, zero value otherwise.
func (o *LdapProviderPatch) GetUserMapping() LdapProviderPatchUserMapping {
	if o == nil || IsNil(o.UserMapping) {
		var ret LdapProviderPatchUserMapping
		return ret
	}
	return *o.UserMapping
}

// GetUserMappingOk returns a tuple with the UserMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatch) GetUserMappingOk() (*LdapProviderPatchUserMapping, bool) {
	if o == nil || IsNil(o.UserMapping) {
		return nil, false
	}
	return o.UserMapping, true
}

// HasUserMapping returns a boolean if a field has been set.
func (o *LdapProviderPatch) HasUserMapping() bool {
	if o != nil && !IsNil(o.UserMapping) {
		return true
	}

	return false
}

// SetUserMapping gets a reference to the given LdapProviderPatchUserMapping and assigns it to the UserMapping field.
func (o *LdapProviderPatch) SetUserMapping(v LdapProviderPatchUserMapping) {
	o.UserMapping = &v
}

// GetGroupMapping returns the GroupMapping field value if set, zero value otherwise.
func (o *LdapProviderPatch) GetGroupMapping() LdapProviderPatchGroupMapping {
	if o == nil || IsNil(o.GroupMapping) {
		var ret LdapProviderPatchGroupMapping
		return ret
	}
	return *o.GroupMapping
}

// GetGroupMappingOk returns a tuple with the GroupMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapProviderPatch) GetGroupMappingOk() (*LdapProviderPatchGroupMapping, bool) {
	if o == nil || IsNil(o.GroupMapping) {
		return nil, false
	}
	return o.GroupMapping, true
}

// HasGroupMapping returns a boolean if a field has been set.
func (o *LdapProviderPatch) HasGroupMapping() bool {
	if o != nil && !IsNil(o.GroupMapping) {
		return true
	}

	return false
}

// SetGroupMapping gets a reference to the given LdapProviderPatchGroupMapping and assigns it to the GroupMapping field.
func (o *LdapProviderPatch) SetGroupMapping(v LdapProviderPatchGroupMapping) {
	o.GroupMapping = &v
}

func (o LdapProviderPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapProviderPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	if !IsNil(o.SearchBases) {
		toSerialize["search_bases"] = o.SearchBases
	}
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.UserMapping) {
		toSerialize["user_mapping"] = o.UserMapping
	}
	if !IsNil(o.GroupMapping) {
		toSerialize["group_mapping"] = o.GroupMapping
	}
	return toSerialize, nil
}

type NullableLdapProviderPatch struct {
	value *LdapProviderPatch
	isSet bool
}

func (v NullableLdapProviderPatch) Get() *LdapProviderPatch {
	return v.value
}

func (v *NullableLdapProviderPatch) Set(val *LdapProviderPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapProviderPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapProviderPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapProviderPatch(val *LdapProviderPatch) *NullableLdapProviderPatch {
	return &NullableLdapProviderPatch{value: val, isSet: true}
}

func (v NullableLdapProviderPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapProviderPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


