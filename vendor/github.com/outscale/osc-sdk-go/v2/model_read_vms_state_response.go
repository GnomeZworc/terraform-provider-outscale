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

// ReadVmsStateResponse struct for ReadVmsStateResponse
type ReadVmsStateResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more VM states.
	VmStates *[]VmStates `json:"VmStates,omitempty"`
}

// NewReadVmsStateResponse instantiates a new ReadVmsStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVmsStateResponse() *ReadVmsStateResponse {
	this := ReadVmsStateResponse{}
	return &this
}

// NewReadVmsStateResponseWithDefaults instantiates a new ReadVmsStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVmsStateResponseWithDefaults() *ReadVmsStateResponse {
	this := ReadVmsStateResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVmsStateResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsStateResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVmsStateResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVmsStateResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVmStates returns the VmStates field value if set, zero value otherwise.
func (o *ReadVmsStateResponse) GetVmStates() []VmStates {
	if o == nil || o.VmStates == nil {
		var ret []VmStates
		return ret
	}
	return *o.VmStates
}

// GetVmStatesOk returns a tuple with the VmStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmsStateResponse) GetVmStatesOk() (*[]VmStates, bool) {
	if o == nil || o.VmStates == nil {
		return nil, false
	}
	return o.VmStates, true
}

// HasVmStates returns a boolean if a field has been set.
func (o *ReadVmsStateResponse) HasVmStates() bool {
	if o != nil && o.VmStates != nil {
		return true
	}

	return false
}

// SetVmStates gets a reference to the given []VmStates and assigns it to the VmStates field.
func (o *ReadVmsStateResponse) SetVmStates(v []VmStates) {
	o.VmStates = &v
}

func (o ReadVmsStateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.VmStates != nil {
		toSerialize["VmStates"] = o.VmStates
	}
	return json.Marshal(toSerialize)
}

type NullableReadVmsStateResponse struct {
	value *ReadVmsStateResponse
	isSet bool
}

func (v NullableReadVmsStateResponse) Get() *ReadVmsStateResponse {
	return v.value
}

func (v *NullableReadVmsStateResponse) Set(val *ReadVmsStateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVmsStateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVmsStateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVmsStateResponse(val *ReadVmsStateResponse) *NullableReadVmsStateResponse {
	return &NullableReadVmsStateResponse{value: val, isSet: true}
}

func (v NullableReadVmsStateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVmsStateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
