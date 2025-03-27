/*
Deployment v1

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:   - Deploying a Mule Application or API to a Mule Runtime   - Scaling up or down a running application   - Managing application settings (ie: properties)   - Deleting a Mule Application Deployment   - Changing artifact version or configurations of a deployment   - Starting, Stopping or Restarting applications   - etc. This API currently supports deployments to Cloudhub 1.0 targets only. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package application_manager_v1

import (
	"encoding/json"
)

// checks if the RuntimeLatestUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuntimeLatestUpdate{}

// RuntimeLatestUpdate struct for RuntimeLatestUpdate
type RuntimeLatestUpdate struct {
	// The unique identifier of the update.
	Id *string `json:"id,omitempty"`
	// The name of the update.
	Name *string `json:"name,omitempty"`
	// The release date as a timestamp (in milliseconds).
	ReleaseDate *float32 `json:"releaseDate,omitempty"`
	// URL or text of the release notes.
	ReleaseNotes *string `json:"releaseNotes,omitempty"`
	Flags *RuntimeLatestUpdateFlags `json:"flags,omitempty"`
}

// NewRuntimeLatestUpdate instantiates a new RuntimeLatestUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuntimeLatestUpdate() *RuntimeLatestUpdate {
	this := RuntimeLatestUpdate{}
	return &this
}

// NewRuntimeLatestUpdateWithDefaults instantiates a new RuntimeLatestUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuntimeLatestUpdateWithDefaults() *RuntimeLatestUpdate {
	this := RuntimeLatestUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuntimeLatestUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeLatestUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuntimeLatestUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuntimeLatestUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuntimeLatestUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeLatestUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuntimeLatestUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuntimeLatestUpdate) SetName(v string) {
	o.Name = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *RuntimeLatestUpdate) GetReleaseDate() float32 {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret float32
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeLatestUpdate) GetReleaseDateOk() (*float32, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *RuntimeLatestUpdate) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given float32 and assigns it to the ReleaseDate field.
func (o *RuntimeLatestUpdate) SetReleaseDate(v float32) {
	o.ReleaseDate = &v
}

// GetReleaseNotes returns the ReleaseNotes field value if set, zero value otherwise.
func (o *RuntimeLatestUpdate) GetReleaseNotes() string {
	if o == nil || IsNil(o.ReleaseNotes) {
		var ret string
		return ret
	}
	return *o.ReleaseNotes
}

// GetReleaseNotesOk returns a tuple with the ReleaseNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeLatestUpdate) GetReleaseNotesOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseNotes) {
		return nil, false
	}
	return o.ReleaseNotes, true
}

// HasReleaseNotes returns a boolean if a field has been set.
func (o *RuntimeLatestUpdate) HasReleaseNotes() bool {
	if o != nil && !IsNil(o.ReleaseNotes) {
		return true
	}

	return false
}

// SetReleaseNotes gets a reference to the given string and assigns it to the ReleaseNotes field.
func (o *RuntimeLatestUpdate) SetReleaseNotes(v string) {
	o.ReleaseNotes = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *RuntimeLatestUpdate) GetFlags() RuntimeLatestUpdateFlags {
	if o == nil || IsNil(o.Flags) {
		var ret RuntimeLatestUpdateFlags
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeLatestUpdate) GetFlagsOk() (*RuntimeLatestUpdateFlags, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *RuntimeLatestUpdate) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given RuntimeLatestUpdateFlags and assigns it to the Flags field.
func (o *RuntimeLatestUpdate) SetFlags(v RuntimeLatestUpdateFlags) {
	o.Flags = &v
}

func (o RuntimeLatestUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuntimeLatestUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	if !IsNil(o.ReleaseNotes) {
		toSerialize["releaseNotes"] = o.ReleaseNotes
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

type NullableRuntimeLatestUpdate struct {
	value *RuntimeLatestUpdate
	isSet bool
}

func (v NullableRuntimeLatestUpdate) Get() *RuntimeLatestUpdate {
	return v.value
}

func (v *NullableRuntimeLatestUpdate) Set(val *RuntimeLatestUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeLatestUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeLatestUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeLatestUpdate(val *RuntimeLatestUpdate) *NullableRuntimeLatestUpdate {
	return &NullableRuntimeLatestUpdate{value: val, isSet: true}
}

func (v NullableRuntimeLatestUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeLatestUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


