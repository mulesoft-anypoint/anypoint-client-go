/*
API Manager Policy API

API Manager Policy API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apim_policy

import (
	"encoding/json"
	"fmt"
)

// ApimPolicyCollection - struct for ApimPolicyCollection
type ApimPolicyCollection struct {
	ApimPolicyFullData *ApimPolicyFullData
	ArrayOfApimPolicy *[]ApimPolicy
}

// ApimPolicyFullDataAsApimPolicyCollection is a convenience function that returns ApimPolicyFullData wrapped in ApimPolicyCollection
func ApimPolicyFullDataAsApimPolicyCollection(v *ApimPolicyFullData) ApimPolicyCollection {
	return ApimPolicyCollection{
		ApimPolicyFullData: v,
	}
}

// []ApimPolicyAsApimPolicyCollection is a convenience function that returns []ApimPolicy wrapped in ApimPolicyCollection
func ArrayOfApimPolicyAsApimPolicyCollection(v *[]ApimPolicy) ApimPolicyCollection {
	return ApimPolicyCollection{
		ArrayOfApimPolicy: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApimPolicyCollection) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApimPolicyFullData
	err = newStrictDecoder(data).Decode(&dst.ApimPolicyFullData)
	if err == nil {
		jsonApimPolicyFullData, _ := json.Marshal(dst.ApimPolicyFullData)
		if string(jsonApimPolicyFullData) == "{}" { // empty struct
			dst.ApimPolicyFullData = nil
		} else {
			match++
		}
	} else {
		dst.ApimPolicyFullData = nil
	}

	// try to unmarshal data into ArrayOfApimPolicy
	err = newStrictDecoder(data).Decode(&dst.ArrayOfApimPolicy)
	if err == nil {
		jsonArrayOfApimPolicy, _ := json.Marshal(dst.ArrayOfApimPolicy)
		if string(jsonArrayOfApimPolicy) == "{}" { // empty struct
			dst.ArrayOfApimPolicy = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfApimPolicy = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApimPolicyFullData = nil
		dst.ArrayOfApimPolicy = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApimPolicyCollection)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApimPolicyCollection)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApimPolicyCollection) MarshalJSON() ([]byte, error) {
	if src.ApimPolicyFullData != nil {
		return json.Marshal(&src.ApimPolicyFullData)
	}

	if src.ArrayOfApimPolicy != nil {
		return json.Marshal(&src.ArrayOfApimPolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApimPolicyCollection) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApimPolicyFullData != nil {
		return obj.ApimPolicyFullData
	}

	if obj.ArrayOfApimPolicy != nil {
		return obj.ArrayOfApimPolicy
	}

	// all schemas are nil
	return nil
}

type NullableApimPolicyCollection struct {
	value *ApimPolicyCollection
	isSet bool
}

func (v NullableApimPolicyCollection) Get() *ApimPolicyCollection {
	return v.value
}

func (v *NullableApimPolicyCollection) Set(val *ApimPolicyCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableApimPolicyCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableApimPolicyCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApimPolicyCollection(val *ApimPolicyCollection) *NullableApimPolicyCollection {
	return &NullableApimPolicyCollection{value: val, isSet: true}
}

func (v NullableApimPolicyCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApimPolicyCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


