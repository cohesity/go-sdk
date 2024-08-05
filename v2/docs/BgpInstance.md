# BgpInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalAS** | Pointer to **NullableInt32** | Local AS. | [optional] 
**Peers** | Pointer to [**[]BgpPeer**](BgpPeer.md) | List of bgp peer groups. | [optional] 
**Timers** | Pointer to [**BgpTimers**](BgpTimers.md) |  | [optional] 

## Methods

### NewBgpInstance

`func NewBgpInstance() *BgpInstance`

NewBgpInstance instantiates a new BgpInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpInstanceWithDefaults

`func NewBgpInstanceWithDefaults() *BgpInstance`

NewBgpInstanceWithDefaults instantiates a new BgpInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalAS

`func (o *BgpInstance) GetLocalAS() int32`

GetLocalAS returns the LocalAS field if non-nil, zero value otherwise.

### GetLocalASOk

`func (o *BgpInstance) GetLocalASOk() (*int32, bool)`

GetLocalASOk returns a tuple with the LocalAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAS

`func (o *BgpInstance) SetLocalAS(v int32)`

SetLocalAS sets LocalAS field to given value.

### HasLocalAS

`func (o *BgpInstance) HasLocalAS() bool`

HasLocalAS returns a boolean if a field has been set.

### SetLocalASNil

`func (o *BgpInstance) SetLocalASNil(b bool)`

 SetLocalASNil sets the value for LocalAS to be an explicit nil

### UnsetLocalAS
`func (o *BgpInstance) UnsetLocalAS()`

UnsetLocalAS ensures that no value is present for LocalAS, not even an explicit nil
### GetPeers

`func (o *BgpInstance) GetPeers() []BgpPeer`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *BgpInstance) GetPeersOk() (*[]BgpPeer, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *BgpInstance) SetPeers(v []BgpPeer)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *BgpInstance) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### SetPeersNil

`func (o *BgpInstance) SetPeersNil(b bool)`

 SetPeersNil sets the value for Peers to be an explicit nil

### UnsetPeers
`func (o *BgpInstance) UnsetPeers()`

UnsetPeers ensures that no value is present for Peers, not even an explicit nil
### GetTimers

`func (o *BgpInstance) GetTimers() BgpTimers`

GetTimers returns the Timers field if non-nil, zero value otherwise.

### GetTimersOk

`func (o *BgpInstance) GetTimersOk() (*BgpTimers, bool)`

GetTimersOk returns a tuple with the Timers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimers

`func (o *BgpInstance) SetTimers(v BgpTimers)`

SetTimers sets Timers field to given value.

### HasTimers

`func (o *BgpInstance) HasTimers() bool`

HasTimers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


