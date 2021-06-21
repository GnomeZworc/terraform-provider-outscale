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

// CreateSnapshotResponse struct for CreateSnapshotResponse
type CreateSnapshotResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	Snapshot        *Snapshot        `json:"Snapshot,omitempty"`
}

// NewCreateSnapshotResponse instantiates a new CreateSnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSnapshotResponse() *CreateSnapshotResponse {
	this := CreateSnapshotResponse{}
	return &this
}

// NewCreateSnapshotResponseWithDefaults instantiates a new CreateSnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSnapshotResponseWithDefaults() *CreateSnapshotResponse {
	this := CreateSnapshotResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateSnapshotResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateSnapshotResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateSnapshotResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *CreateSnapshotResponse) GetSnapshot() Snapshot {
	if o == nil || o.Snapshot == nil {
		var ret Snapshot
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotResponse) GetSnapshotOk() (*Snapshot, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *CreateSnapshotResponse) HasSnapshot() bool {
	if o != nil && o.Snapshot != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given Snapshot and assigns it to the Snapshot field.
func (o *CreateSnapshotResponse) SetSnapshot(v Snapshot) {
	o.Snapshot = &v
}

func (o CreateSnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Snapshot != nil {
		toSerialize["Snapshot"] = o.Snapshot
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSnapshotResponse struct {
	value *CreateSnapshotResponse
	isSet bool
}

func (v NullableCreateSnapshotResponse) Get() *CreateSnapshotResponse {
	return v.value
}

func (v *NullableCreateSnapshotResponse) Set(val *CreateSnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSnapshotResponse(val *CreateSnapshotResponse) *NullableCreateSnapshotResponse {
	return &NullableCreateSnapshotResponse{value: val, isSet: true}
}

func (v NullableCreateSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
