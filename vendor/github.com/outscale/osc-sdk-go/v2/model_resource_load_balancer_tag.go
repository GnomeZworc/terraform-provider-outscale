/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ResourceLoadBalancerTag Information about the tag.
type ResourceLoadBalancerTag struct {
	// The key of the tag, with a minimum of 1 character.
	Key *string `json:"Key,omitempty"`
}

// NewResourceLoadBalancerTag instantiates a new ResourceLoadBalancerTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceLoadBalancerTag() *ResourceLoadBalancerTag {
	this := ResourceLoadBalancerTag{}
	return &this
}

// NewResourceLoadBalancerTagWithDefaults instantiates a new ResourceLoadBalancerTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceLoadBalancerTagWithDefaults() *ResourceLoadBalancerTag {
	this := ResourceLoadBalancerTag{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ResourceLoadBalancerTag) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLoadBalancerTag) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ResourceLoadBalancerTag) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ResourceLoadBalancerTag) SetKey(v string) {
	o.Key = &v
}

func (o ResourceLoadBalancerTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableResourceLoadBalancerTag struct {
	value *ResourceLoadBalancerTag
	isSet bool
}

func (v NullableResourceLoadBalancerTag) Get() *ResourceLoadBalancerTag {
	return v.value
}

func (v *NullableResourceLoadBalancerTag) Set(val *ResourceLoadBalancerTag) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLoadBalancerTag) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLoadBalancerTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLoadBalancerTag(val *ResourceLoadBalancerTag) *NullableResourceLoadBalancerTag {
	return &NullableResourceLoadBalancerTag{value: val, isSet: true}
}

func (v NullableResourceLoadBalancerTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLoadBalancerTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
