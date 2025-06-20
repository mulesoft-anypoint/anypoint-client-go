/*
Private Space TLS Context API

Description of the Private Space TLS Context API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package private_space_tlscontext

import (
	"encoding/json"
)

// checks if the TlsContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TlsContext{}

// TlsContext struct for TlsContext
type TlsContext struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TrustStore *TrustStore `json:"trustStore,omitempty"`
	KeyStore *KeyStore `json:"keyStore,omitempty"`
	Ciphers *Ciphers `json:"ciphers,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewTlsContext instantiates a new TlsContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsContext() *TlsContext {
	this := TlsContext{}
	return &this
}

// NewTlsContextWithDefaults instantiates a new TlsContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsContextWithDefaults() *TlsContext {
	this := TlsContext{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TlsContext) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContext) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TlsContext) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TlsContext) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TlsContext) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContext) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TlsContext) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TlsContext) SetName(v string) {
	o.Name = &v
}

// GetTrustStore returns the TrustStore field value if set, zero value otherwise.
func (o *TlsContext) GetTrustStore() TrustStore {
	if o == nil || IsNil(o.TrustStore) {
		var ret TrustStore
		return ret
	}
	return *o.TrustStore
}

// GetTrustStoreOk returns a tuple with the TrustStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContext) GetTrustStoreOk() (*TrustStore, bool) {
	if o == nil || IsNil(o.TrustStore) {
		return nil, false
	}
	return o.TrustStore, true
}

// HasTrustStore returns a boolean if a field has been set.
func (o *TlsContext) HasTrustStore() bool {
	if o != nil && !IsNil(o.TrustStore) {
		return true
	}

	return false
}

// SetTrustStore gets a reference to the given TrustStore and assigns it to the TrustStore field.
func (o *TlsContext) SetTrustStore(v TrustStore) {
	o.TrustStore = &v
}

// GetKeyStore returns the KeyStore field value if set, zero value otherwise.
func (o *TlsContext) GetKeyStore() KeyStore {
	if o == nil || IsNil(o.KeyStore) {
		var ret KeyStore
		return ret
	}
	return *o.KeyStore
}

// GetKeyStoreOk returns a tuple with the KeyStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContext) GetKeyStoreOk() (*KeyStore, bool) {
	if o == nil || IsNil(o.KeyStore) {
		return nil, false
	}
	return o.KeyStore, true
}

// HasKeyStore returns a boolean if a field has been set.
func (o *TlsContext) HasKeyStore() bool {
	if o != nil && !IsNil(o.KeyStore) {
		return true
	}

	return false
}

// SetKeyStore gets a reference to the given KeyStore and assigns it to the KeyStore field.
func (o *TlsContext) SetKeyStore(v KeyStore) {
	o.KeyStore = &v
}

// GetCiphers returns the Ciphers field value if set, zero value otherwise.
func (o *TlsContext) GetCiphers() Ciphers {
	if o == nil || IsNil(o.Ciphers) {
		var ret Ciphers
		return ret
	}
	return *o.Ciphers
}

// GetCiphersOk returns a tuple with the Ciphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContext) GetCiphersOk() (*Ciphers, bool) {
	if o == nil || IsNil(o.Ciphers) {
		return nil, false
	}
	return o.Ciphers, true
}

// HasCiphers returns a boolean if a field has been set.
func (o *TlsContext) HasCiphers() bool {
	if o != nil && !IsNil(o.Ciphers) {
		return true
	}

	return false
}

// SetCiphers gets a reference to the given Ciphers and assigns it to the Ciphers field.
func (o *TlsContext) SetCiphers(v Ciphers) {
	o.Ciphers = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsContext) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsContext) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsContext) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TlsContext) SetType(v string) {
	o.Type = &v
}

func (o TlsContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TlsContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TrustStore) {
		toSerialize["trustStore"] = o.TrustStore
	}
	if !IsNil(o.KeyStore) {
		toSerialize["keyStore"] = o.KeyStore
	}
	if !IsNil(o.Ciphers) {
		toSerialize["ciphers"] = o.Ciphers
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTlsContext struct {
	value *TlsContext
	isSet bool
}

func (v NullableTlsContext) Get() *TlsContext {
	return v.value
}

func (v *NullableTlsContext) Set(val *TlsContext) {
	v.value = val
	v.isSet = true
}

func (v NullableTlsContext) IsSet() bool {
	return v.isSet
}

func (v *NullableTlsContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTlsContext(val *TlsContext) *NullableTlsContext {
	return &NullableTlsContext{value: val, isSet: true}
}

func (v NullableTlsContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTlsContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


