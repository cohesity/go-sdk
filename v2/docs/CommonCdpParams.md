# CommonCdpParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableCdpSyncReplication** | Pointer to **NullableBool** | Specifies whether synchronous replication is enabled for CDP Protection Group when replication target is specified in attached policy. | [optional] 

## Methods

### NewCommonCdpParams

`func NewCommonCdpParams() *CommonCdpParams`

NewCommonCdpParams instantiates a new CommonCdpParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonCdpParamsWithDefaults

`func NewCommonCdpParamsWithDefaults() *CommonCdpParams`

NewCommonCdpParamsWithDefaults instantiates a new CommonCdpParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableCdpSyncReplication

`func (o *CommonCdpParams) GetEnableCdpSyncReplication() bool`

GetEnableCdpSyncReplication returns the EnableCdpSyncReplication field if non-nil, zero value otherwise.

### GetEnableCdpSyncReplicationOk

`func (o *CommonCdpParams) GetEnableCdpSyncReplicationOk() (*bool, bool)`

GetEnableCdpSyncReplicationOk returns a tuple with the EnableCdpSyncReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCdpSyncReplication

`func (o *CommonCdpParams) SetEnableCdpSyncReplication(v bool)`

SetEnableCdpSyncReplication sets EnableCdpSyncReplication field to given value.

### HasEnableCdpSyncReplication

`func (o *CommonCdpParams) HasEnableCdpSyncReplication() bool`

HasEnableCdpSyncReplication returns a boolean if a field has been set.

### SetEnableCdpSyncReplicationNil

`func (o *CommonCdpParams) SetEnableCdpSyncReplicationNil(b bool)`

 SetEnableCdpSyncReplicationNil sets the value for EnableCdpSyncReplication to be an explicit nil

### UnsetEnableCdpSyncReplication
`func (o *CommonCdpParams) UnsetEnableCdpSyncReplication()`

UnsetEnableCdpSyncReplication ensures that no value is present for EnableCdpSyncReplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


