/*
 * Auth API
 *
 * Use the Auth API to manage tokens for secure access to IONOS Cloud APIs (Auth API, Cloud API, Reseller API, Activity Log API, and others).
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Jwt struct for Jwt
type Jwt struct {
	// JSON Web Token (JWT) Base64url strings separated by dots
	Token *string `json:"token,omitempty"`
}

// NewJwt instantiates a new Jwt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJwt() *Jwt {
	this := Jwt{}

	return &this
}

// NewJwtWithDefaults instantiates a new Jwt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJwtWithDefaults() *Jwt {
	this := Jwt{}
	return &this
}

// GetToken returns the Token field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Jwt) GetToken() *string {
	if o == nil {
		return nil
	}

	return o.Token

}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Jwt) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Token, true
}

// SetToken sets field value
func (o *Jwt) SetToken(v string) {

	o.Token = &v

}

// HasToken returns a boolean if a field has been set.
func (o *Jwt) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

func (o Jwt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	return json.Marshal(toSerialize)
}

type NullableJwt struct {
	value *Jwt
	isSet bool
}

func (v NullableJwt) Get() *Jwt {
	return v.value
}

func (v *NullableJwt) Set(val *Jwt) {
	v.value = val
	v.isSet = true
}

func (v NullableJwt) IsSet() bool {
	return v.isSet
}

func (v *NullableJwt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJwt(val *Jwt) *NullableJwt {
	return &NullableJwt{value: val, isSet: true}
}

func (v NullableJwt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJwt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
