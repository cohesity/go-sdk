# RunReplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies id of Remote Cluster to copy the Snapshots to. | 
**OnLegalHold** | Pointer to **NullableBool** | Specifies if the Run is on legal hold. | [optional] 
**Retention** | Pointer to [**Retention**](Retention.md) |  | [optional] 

## Methods

### NewRunReplicationConfig

`func NewRunReplicationConfig(id NullableInt64, ) *RunReplicationConfig`

NewRunReplicationConfig instantiates a new RunReplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunReplicationConfigWithDefaults

`func NewRunReplicationConfigWithDefaults() *RunReplicationConfig`

NewRunReplicationConfigWithDefaults instantiates a new RunReplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunReplicationConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunReplicationConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunReplicationConfig) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *RunReplicationConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RunReplicationConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOnLegalHold

`func (o *RunReplicationConfig) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *RunReplicationConfig) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *RunReplicationConfig) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *RunReplicationConfig) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *RunReplicationConfig) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *RunReplicationConfig) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil
### GetRetention

`func (o *RunReplicationConfig) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *RunReplicationConfig) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *RunReplicationConfig) SetRetention(v Retention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *RunReplicationConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


