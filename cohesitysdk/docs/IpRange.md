# IpRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndIp** | Pointer to **NullableString** | Optional End IP of the range If not specified, EndIp is same as StartIp | [optional] 
**StartIp** | Pointer to **NullableString** | Start IP of the range | [optional] 

## Methods

### NewIpRange

`func NewIpRange() *IpRange`

NewIpRange instantiates a new IpRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpRangeWithDefaults

`func NewIpRangeWithDefaults() *IpRange`

NewIpRangeWithDefaults instantiates a new IpRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndIp

`func (o *IpRange) GetEndIp() string`

GetEndIp returns the EndIp field if non-nil, zero value otherwise.

### GetEndIpOk

`func (o *IpRange) GetEndIpOk() (*string, bool)`

GetEndIpOk returns a tuple with the EndIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIp

`func (o *IpRange) SetEndIp(v string)`

SetEndIp sets EndIp field to given value.

### HasEndIp

`func (o *IpRange) HasEndIp() bool`

HasEndIp returns a boolean if a field has been set.

### SetEndIpNil

`func (o *IpRange) SetEndIpNil(b bool)`

 SetEndIpNil sets the value for EndIp to be an explicit nil

### UnsetEndIp
`func (o *IpRange) UnsetEndIp()`

UnsetEndIp ensures that no value is present for EndIp, not even an explicit nil
### GetStartIp

`func (o *IpRange) GetStartIp() string`

GetStartIp returns the StartIp field if non-nil, zero value otherwise.

### GetStartIpOk

`func (o *IpRange) GetStartIpOk() (*string, bool)`

GetStartIpOk returns a tuple with the StartIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIp

`func (o *IpRange) SetStartIp(v string)`

SetStartIp sets StartIp field to given value.

### HasStartIp

`func (o *IpRange) HasStartIp() bool`

HasStartIp returns a boolean if a field has been set.

### SetStartIpNil

`func (o *IpRange) SetStartIpNil(b bool)`

 SetStartIpNil sets the value for StartIp to be an explicit nil

### UnsetStartIp
`func (o *IpRange) UnsetStartIp()`

UnsetStartIp ensures that no value is present for StartIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


