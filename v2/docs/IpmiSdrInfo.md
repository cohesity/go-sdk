# IpmiSdrInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdrEntries** | Pointer to [**[]IpmiSdrInfoEntry**](IpmiSdrInfoEntry.md) | Specifies the list of sdr entries for the given node. | [optional] 

## Methods

### NewIpmiSdrInfo

`func NewIpmiSdrInfo() *IpmiSdrInfo`

NewIpmiSdrInfo instantiates a new IpmiSdrInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiSdrInfoWithDefaults

`func NewIpmiSdrInfoWithDefaults() *IpmiSdrInfo`

NewIpmiSdrInfoWithDefaults instantiates a new IpmiSdrInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdrEntries

`func (o *IpmiSdrInfo) GetSdrEntries() []IpmiSdrInfoEntry`

GetSdrEntries returns the SdrEntries field if non-nil, zero value otherwise.

### GetSdrEntriesOk

`func (o *IpmiSdrInfo) GetSdrEntriesOk() (*[]IpmiSdrInfoEntry, bool)`

GetSdrEntriesOk returns a tuple with the SdrEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdrEntries

`func (o *IpmiSdrInfo) SetSdrEntries(v []IpmiSdrInfoEntry)`

SetSdrEntries sets SdrEntries field to given value.

### HasSdrEntries

`func (o *IpmiSdrInfo) HasSdrEntries() bool`

HasSdrEntries returns a boolean if a field has been set.

### SetSdrEntriesNil

`func (o *IpmiSdrInfo) SetSdrEntriesNil(b bool)`

 SetSdrEntriesNil sets the value for SdrEntries to be an explicit nil

### UnsetSdrEntries
`func (o *IpmiSdrInfo) UnsetSdrEntries()`

UnsetSdrEntries ensures that no value is present for SdrEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


