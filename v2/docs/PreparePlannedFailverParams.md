# PreparePlannedFailverParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReverseReplication** | Pointer to **NullableBool** | Specifies whether a reverse replication needs to be set for the view on target cluster after failover. | [optional] 

## Methods

### NewPreparePlannedFailverParams

`func NewPreparePlannedFailverParams() *PreparePlannedFailverParams`

NewPreparePlannedFailverParams instantiates a new PreparePlannedFailverParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreparePlannedFailverParamsWithDefaults

`func NewPreparePlannedFailverParamsWithDefaults() *PreparePlannedFailverParams`

NewPreparePlannedFailverParamsWithDefaults instantiates a new PreparePlannedFailverParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReverseReplication

`func (o *PreparePlannedFailverParams) GetReverseReplication() bool`

GetReverseReplication returns the ReverseReplication field if non-nil, zero value otherwise.

### GetReverseReplicationOk

`func (o *PreparePlannedFailverParams) GetReverseReplicationOk() (*bool, bool)`

GetReverseReplicationOk returns a tuple with the ReverseReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseReplication

`func (o *PreparePlannedFailverParams) SetReverseReplication(v bool)`

SetReverseReplication sets ReverseReplication field to given value.

### HasReverseReplication

`func (o *PreparePlannedFailverParams) HasReverseReplication() bool`

HasReverseReplication returns a boolean if a field has been set.

### SetReverseReplicationNil

`func (o *PreparePlannedFailverParams) SetReverseReplicationNil(b bool)`

 SetReverseReplicationNil sets the value for ReverseReplication to be an explicit nil

### UnsetReverseReplication
`func (o *PreparePlannedFailverParams) UnsetReverseReplication()`

UnsetReverseReplication ensures that no value is present for ReverseReplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


