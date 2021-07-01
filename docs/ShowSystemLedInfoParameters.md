# ShowSystemLedInfoParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIp** | Pointer to **NullableString** | Optional field. If Node IP is not specified, LED info for the current node is displayed. | [optional] 

## Methods

### NewShowSystemLedInfoParameters

`func NewShowSystemLedInfoParameters() *ShowSystemLedInfoParameters`

NewShowSystemLedInfoParameters instantiates a new ShowSystemLedInfoParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowSystemLedInfoParametersWithDefaults

`func NewShowSystemLedInfoParametersWithDefaults() *ShowSystemLedInfoParameters`

NewShowSystemLedInfoParametersWithDefaults instantiates a new ShowSystemLedInfoParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIp

`func (o *ShowSystemLedInfoParameters) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *ShowSystemLedInfoParameters) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *ShowSystemLedInfoParameters) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *ShowSystemLedInfoParameters) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *ShowSystemLedInfoParameters) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *ShowSystemLedInfoParameters) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


