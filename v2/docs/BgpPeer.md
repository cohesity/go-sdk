# BgpPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressOrTag** | Pointer to **NullableString** | IP Address. | [optional] 
**Description** | Pointer to **NullableString** | Peer&#39;s description. | [optional] 
**RemoteAS** | Pointer to **NullableInt32** | Remote AS. | [optional] 
**Timers** | Pointer to [**BgpTimers**](BgpTimers.md) |  | [optional] 

## Methods

### NewBgpPeer

`func NewBgpPeer() *BgpPeer`

NewBgpPeer instantiates a new BgpPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpPeerWithDefaults

`func NewBgpPeerWithDefaults() *BgpPeer`

NewBgpPeerWithDefaults instantiates a new BgpPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressOrTag

`func (o *BgpPeer) GetAddressOrTag() string`

GetAddressOrTag returns the AddressOrTag field if non-nil, zero value otherwise.

### GetAddressOrTagOk

`func (o *BgpPeer) GetAddressOrTagOk() (*string, bool)`

GetAddressOrTagOk returns a tuple with the AddressOrTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOrTag

`func (o *BgpPeer) SetAddressOrTag(v string)`

SetAddressOrTag sets AddressOrTag field to given value.

### HasAddressOrTag

`func (o *BgpPeer) HasAddressOrTag() bool`

HasAddressOrTag returns a boolean if a field has been set.

### SetAddressOrTagNil

`func (o *BgpPeer) SetAddressOrTagNil(b bool)`

 SetAddressOrTagNil sets the value for AddressOrTag to be an explicit nil

### UnsetAddressOrTag
`func (o *BgpPeer) UnsetAddressOrTag()`

UnsetAddressOrTag ensures that no value is present for AddressOrTag, not even an explicit nil
### GetDescription

`func (o *BgpPeer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BgpPeer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BgpPeer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BgpPeer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BgpPeer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BgpPeer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRemoteAS

`func (o *BgpPeer) GetRemoteAS() int32`

GetRemoteAS returns the RemoteAS field if non-nil, zero value otherwise.

### GetRemoteASOk

`func (o *BgpPeer) GetRemoteASOk() (*int32, bool)`

GetRemoteASOk returns a tuple with the RemoteAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAS

`func (o *BgpPeer) SetRemoteAS(v int32)`

SetRemoteAS sets RemoteAS field to given value.

### HasRemoteAS

`func (o *BgpPeer) HasRemoteAS() bool`

HasRemoteAS returns a boolean if a field has been set.

### SetRemoteASNil

`func (o *BgpPeer) SetRemoteASNil(b bool)`

 SetRemoteASNil sets the value for RemoteAS to be an explicit nil

### UnsetRemoteAS
`func (o *BgpPeer) UnsetRemoteAS()`

UnsetRemoteAS ensures that no value is present for RemoteAS, not even an explicit nil
### GetTimers

`func (o *BgpPeer) GetTimers() BgpTimers`

GetTimers returns the Timers field if non-nil, zero value otherwise.

### GetTimersOk

`func (o *BgpPeer) GetTimersOk() (*BgpTimers, bool)`

GetTimersOk returns a tuple with the Timers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimers

`func (o *BgpPeer) SetTimers(v BgpTimers)`

SetTimers sets Timers field to given value.

### HasTimers

`func (o *BgpPeer) HasTimers() bool`

HasTimers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


