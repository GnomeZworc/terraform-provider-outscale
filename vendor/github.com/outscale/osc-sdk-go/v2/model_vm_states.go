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

// VmStates Information about the states of the VMs.
type VmStates struct {
	// One or more scheduled events associated with the VM.
	MaintenanceEvents *[]MaintenanceEvent `json:"MaintenanceEvents,omitempty"`
	// The name of the Subregion of the VM.
	SubregionName *string `json:"SubregionName,omitempty"`
	// The ID of the VM.
	VmId *string `json:"VmId,omitempty"`
	// The state of the VM (`pending` \\| `running` \\| `stopping` \\| `stopped` \\| `shutting-down` \\| `terminated` \\| `quarantine`).
	VmState *string `json:"VmState,omitempty"`
}

// NewVmStates instantiates a new VmStates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmStates() *VmStates {
	this := VmStates{}
	return &this
}

// NewVmStatesWithDefaults instantiates a new VmStates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmStatesWithDefaults() *VmStates {
	this := VmStates{}
	return &this
}

// GetMaintenanceEvents returns the MaintenanceEvents field value if set, zero value otherwise.
func (o *VmStates) GetMaintenanceEvents() []MaintenanceEvent {
	if o == nil || o.MaintenanceEvents == nil {
		var ret []MaintenanceEvent
		return ret
	}
	return *o.MaintenanceEvents
}

// GetMaintenanceEventsOk returns a tuple with the MaintenanceEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmStates) GetMaintenanceEventsOk() (*[]MaintenanceEvent, bool) {
	if o == nil || o.MaintenanceEvents == nil {
		return nil, false
	}
	return o.MaintenanceEvents, true
}

// HasMaintenanceEvents returns a boolean if a field has been set.
func (o *VmStates) HasMaintenanceEvents() bool {
	if o != nil && o.MaintenanceEvents != nil {
		return true
	}

	return false
}

// SetMaintenanceEvents gets a reference to the given []MaintenanceEvent and assigns it to the MaintenanceEvents field.
func (o *VmStates) SetMaintenanceEvents(v []MaintenanceEvent) {
	o.MaintenanceEvents = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *VmStates) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmStates) GetSubregionNameOk() (*string, bool) {
	if o == nil || o.SubregionName == nil {
		return nil, false
	}
	return o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *VmStates) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *VmStates) SetSubregionName(v string) {
	o.SubregionName = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *VmStates) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmStates) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *VmStates) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *VmStates) SetVmId(v string) {
	o.VmId = &v
}

// GetVmState returns the VmState field value if set, zero value otherwise.
func (o *VmStates) GetVmState() string {
	if o == nil || o.VmState == nil {
		var ret string
		return ret
	}
	return *o.VmState
}

// GetVmStateOk returns a tuple with the VmState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmStates) GetVmStateOk() (*string, bool) {
	if o == nil || o.VmState == nil {
		return nil, false
	}
	return o.VmState, true
}

// HasVmState returns a boolean if a field has been set.
func (o *VmStates) HasVmState() bool {
	if o != nil && o.VmState != nil {
		return true
	}

	return false
}

// SetVmState gets a reference to the given string and assigns it to the VmState field.
func (o *VmStates) SetVmState(v string) {
	o.VmState = &v
}

func (o VmStates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaintenanceEvents != nil {
		toSerialize["MaintenanceEvents"] = o.MaintenanceEvents
	}
	if o.SubregionName != nil {
		toSerialize["SubregionName"] = o.SubregionName
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	if o.VmState != nil {
		toSerialize["VmState"] = o.VmState
	}
	return json.Marshal(toSerialize)
}

type NullableVmStates struct {
	value *VmStates
	isSet bool
}

func (v NullableVmStates) Get() *VmStates {
	return v.value
}

func (v *NullableVmStates) Set(val *VmStates) {
	v.value = val
	v.isSet = true
}

func (v NullableVmStates) IsSet() bool {
	return v.isSet
}

func (v *NullableVmStates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmStates(val *VmStates) *NullableVmStates {
	return &NullableVmStates{value: val, isSet: true}
}

func (v NullableVmStates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmStates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
