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

// DeleteSecurityGroupRuleRequest struct for DeleteSecurityGroupRuleRequest
type DeleteSecurityGroupRuleRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The direction of the flow: `Inbound` or `Outbound`. You can specify `Outbound` for Nets only.
	Flow string `json:"Flow"`
	// The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
	FromPortRange *int32 `json:"FromPortRange,omitempty"`
	// The IP protocol name (`tcp`, `udp`, `icmp`) or protocol number. By default, `-1`, which means all protocols.
	IpProtocol *string `json:"IpProtocol,omitempty"`
	// The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16).
	IpRange *string `json:"IpRange,omitempty"`
	// One or more rules you want to delete from the security group.
	Rules *[]SecurityGroupRule `json:"Rules,omitempty"`
	// The account ID of the owner of the security group you want to delete a rule from.
	SecurityGroupAccountIdToUnlink *string `json:"SecurityGroupAccountIdToUnlink,omitempty"`
	// The ID of the security group you want to delete a rule from.
	SecurityGroupId string `json:"SecurityGroupId"`
	// The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group.
	SecurityGroupNameToUnlink *string `json:"SecurityGroupNameToUnlink,omitempty"`
	// The end of the port range for the TCP and UDP protocols, or an ICMP type number.
	ToPortRange *int32 `json:"ToPortRange,omitempty"`
}

// NewDeleteSecurityGroupRuleRequest instantiates a new DeleteSecurityGroupRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSecurityGroupRuleRequest(flow string, securityGroupId string) *DeleteSecurityGroupRuleRequest {
	this := DeleteSecurityGroupRuleRequest{}
	this.Flow = flow
	this.SecurityGroupId = securityGroupId
	return &this
}

// NewDeleteSecurityGroupRuleRequestWithDefaults instantiates a new DeleteSecurityGroupRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSecurityGroupRuleRequestWithDefaults() *DeleteSecurityGroupRuleRequest {
	this := DeleteSecurityGroupRuleRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteSecurityGroupRuleRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFlow returns the Flow field value
func (o *DeleteSecurityGroupRuleRequest) GetFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Flow
}

// GetFlowOk returns a tuple with the Flow field value
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flow, true
}

// SetFlow sets field value
func (o *DeleteSecurityGroupRuleRequest) SetFlow(v string) {
	o.Flow = v
}

// GetFromPortRange returns the FromPortRange field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleRequest) GetFromPortRange() int32 {
	if o == nil || o.FromPortRange == nil {
		var ret int32
		return ret
	}
	return *o.FromPortRange
}

// GetFromPortRangeOk returns a tuple with the FromPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetFromPortRangeOk() (*int32, bool) {
	if o == nil || o.FromPortRange == nil {
		return nil, false
	}
	return o.FromPortRange, true
}

// HasFromPortRange returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleRequest) HasFromPortRange() bool {
	if o != nil && o.FromPortRange != nil {
		return true
	}

	return false
}

// SetFromPortRange gets a reference to the given int32 and assigns it to the FromPortRange field.
func (o *DeleteSecurityGroupRuleRequest) SetFromPortRange(v int32) {
	o.FromPortRange = &v
}

// GetIpProtocol returns the IpProtocol field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleRequest) GetIpProtocol() string {
	if o == nil || o.IpProtocol == nil {
		var ret string
		return ret
	}
	return *o.IpProtocol
}

// GetIpProtocolOk returns a tuple with the IpProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetIpProtocolOk() (*string, bool) {
	if o == nil || o.IpProtocol == nil {
		return nil, false
	}
	return o.IpProtocol, true
}

// HasIpProtocol returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleRequest) HasIpProtocol() bool {
	if o != nil && o.IpProtocol != nil {
		return true
	}

	return false
}

// SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.
func (o *DeleteSecurityGroupRuleRequest) SetIpProtocol(v string) {
	o.IpProtocol = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleRequest) GetIpRange() string {
	if o == nil || o.IpRange == nil {
		var ret string
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetIpRangeOk() (*string, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleRequest) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *DeleteSecurityGroupRuleRequest) SetIpRange(v string) {
	o.IpRange = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleRequest) GetRules() []SecurityGroupRule {
	if o == nil || o.Rules == nil {
		var ret []SecurityGroupRule
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetRulesOk() (*[]SecurityGroupRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleRequest) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []SecurityGroupRule and assigns it to the Rules field.
func (o *DeleteSecurityGroupRuleRequest) SetRules(v []SecurityGroupRule) {
	o.Rules = &v
}

// GetSecurityGroupAccountIdToUnlink returns the SecurityGroupAccountIdToUnlink field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupAccountIdToUnlink() string {
	if o == nil || o.SecurityGroupAccountIdToUnlink == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupAccountIdToUnlink
}

// GetSecurityGroupAccountIdToUnlinkOk returns a tuple with the SecurityGroupAccountIdToUnlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupAccountIdToUnlinkOk() (*string, bool) {
	if o == nil || o.SecurityGroupAccountIdToUnlink == nil {
		return nil, false
	}
	return o.SecurityGroupAccountIdToUnlink, true
}

// HasSecurityGroupAccountIdToUnlink returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleRequest) HasSecurityGroupAccountIdToUnlink() bool {
	if o != nil && o.SecurityGroupAccountIdToUnlink != nil {
		return true
	}

	return false
}

// SetSecurityGroupAccountIdToUnlink gets a reference to the given string and assigns it to the SecurityGroupAccountIdToUnlink field.
func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupAccountIdToUnlink(v string) {
	o.SecurityGroupAccountIdToUnlink = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value
func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityGroupId, true
}

// SetSecurityGroupId sets field value
func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupId(v string) {
	o.SecurityGroupId = v
}

// GetSecurityGroupNameToUnlink returns the SecurityGroupNameToUnlink field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupNameToUnlink() string {
	if o == nil || o.SecurityGroupNameToUnlink == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupNameToUnlink
}

// GetSecurityGroupNameToUnlinkOk returns a tuple with the SecurityGroupNameToUnlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupNameToUnlinkOk() (*string, bool) {
	if o == nil || o.SecurityGroupNameToUnlink == nil {
		return nil, false
	}
	return o.SecurityGroupNameToUnlink, true
}

// HasSecurityGroupNameToUnlink returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleRequest) HasSecurityGroupNameToUnlink() bool {
	if o != nil && o.SecurityGroupNameToUnlink != nil {
		return true
	}

	return false
}

// SetSecurityGroupNameToUnlink gets a reference to the given string and assigns it to the SecurityGroupNameToUnlink field.
func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupNameToUnlink(v string) {
	o.SecurityGroupNameToUnlink = &v
}

// GetToPortRange returns the ToPortRange field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRuleRequest) GetToPortRange() int32 {
	if o == nil || o.ToPortRange == nil {
		var ret int32
		return ret
	}
	return *o.ToPortRange
}

// GetToPortRangeOk returns a tuple with the ToPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRuleRequest) GetToPortRangeOk() (*int32, bool) {
	if o == nil || o.ToPortRange == nil {
		return nil, false
	}
	return o.ToPortRange, true
}

// HasToPortRange returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRuleRequest) HasToPortRange() bool {
	if o != nil && o.ToPortRange != nil {
		return true
	}

	return false
}

// SetToPortRange gets a reference to the given int32 and assigns it to the ToPortRange field.
func (o *DeleteSecurityGroupRuleRequest) SetToPortRange(v int32) {
	o.ToPortRange = &v
}

func (o DeleteSecurityGroupRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Flow"] = o.Flow
	}
	if o.FromPortRange != nil {
		toSerialize["FromPortRange"] = o.FromPortRange
	}
	if o.IpProtocol != nil {
		toSerialize["IpProtocol"] = o.IpProtocol
	}
	if o.IpRange != nil {
		toSerialize["IpRange"] = o.IpRange
	}
	if o.Rules != nil {
		toSerialize["Rules"] = o.Rules
	}
	if o.SecurityGroupAccountIdToUnlink != nil {
		toSerialize["SecurityGroupAccountIdToUnlink"] = o.SecurityGroupAccountIdToUnlink
	}
	if true {
		toSerialize["SecurityGroupId"] = o.SecurityGroupId
	}
	if o.SecurityGroupNameToUnlink != nil {
		toSerialize["SecurityGroupNameToUnlink"] = o.SecurityGroupNameToUnlink
	}
	if o.ToPortRange != nil {
		toSerialize["ToPortRange"] = o.ToPortRange
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteSecurityGroupRuleRequest struct {
	value *DeleteSecurityGroupRuleRequest
	isSet bool
}

func (v NullableDeleteSecurityGroupRuleRequest) Get() *DeleteSecurityGroupRuleRequest {
	return v.value
}

func (v *NullableDeleteSecurityGroupRuleRequest) Set(val *DeleteSecurityGroupRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSecurityGroupRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSecurityGroupRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSecurityGroupRuleRequest(val *DeleteSecurityGroupRuleRequest) *NullableDeleteSecurityGroupRuleRequest {
	return &NullableDeleteSecurityGroupRuleRequest{value: val, isSet: true}
}

func (v NullableDeleteSecurityGroupRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSecurityGroupRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
