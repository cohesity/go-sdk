# RunArchivalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTargetType** | **NullableString** | Specifies the snapshot&#39;s archival target type from which recovery has been performed. | 
**CopyOnlyFullySuccessful** | Pointer to **NullableBool** | Specifies if Snapshots are copied from a fully successful Protection Group Run or a partially successful Protection Group Run. If false, Snapshots are copied the Protection Group Run, even if the Run was not fully successful i.e. Snapshots were not captured for all Objects in the Protection Group. If true, Snapshots are copied only when the run is fully successful. | [optional] 
**Id** | **NullableInt64** | Specifies the Archival target to copy the Snapshots to. | 
**OnLegalHold** | Pointer to **NullableBool** | Specifies if the Run is on legal hold. | [optional] 
**Retention** | Pointer to [**Retention**](Retention.md) |  | [optional] 

## Methods

### NewRunArchivalConfig

`func NewRunArchivalConfig(archivalTargetType NullableString, id NullableInt64, ) *RunArchivalConfig`

NewRunArchivalConfig instantiates a new RunArchivalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunArchivalConfigWithDefaults

`func NewRunArchivalConfigWithDefaults() *RunArchivalConfig`

NewRunArchivalConfigWithDefaults instantiates a new RunArchivalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTargetType

`func (o *RunArchivalConfig) GetArchivalTargetType() string`

GetArchivalTargetType returns the ArchivalTargetType field if non-nil, zero value otherwise.

### GetArchivalTargetTypeOk

`func (o *RunArchivalConfig) GetArchivalTargetTypeOk() (*string, bool)`

GetArchivalTargetTypeOk returns a tuple with the ArchivalTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetType

`func (o *RunArchivalConfig) SetArchivalTargetType(v string)`

SetArchivalTargetType sets ArchivalTargetType field to given value.


### SetArchivalTargetTypeNil

`func (o *RunArchivalConfig) SetArchivalTargetTypeNil(b bool)`

 SetArchivalTargetTypeNil sets the value for ArchivalTargetType to be an explicit nil

### UnsetArchivalTargetType
`func (o *RunArchivalConfig) UnsetArchivalTargetType()`

UnsetArchivalTargetType ensures that no value is present for ArchivalTargetType, not even an explicit nil
### GetCopyOnlyFullySuccessful

`func (o *RunArchivalConfig) GetCopyOnlyFullySuccessful() bool`

GetCopyOnlyFullySuccessful returns the CopyOnlyFullySuccessful field if non-nil, zero value otherwise.

### GetCopyOnlyFullySuccessfulOk

`func (o *RunArchivalConfig) GetCopyOnlyFullySuccessfulOk() (*bool, bool)`

GetCopyOnlyFullySuccessfulOk returns a tuple with the CopyOnlyFullySuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOnlyFullySuccessful

`func (o *RunArchivalConfig) SetCopyOnlyFullySuccessful(v bool)`

SetCopyOnlyFullySuccessful sets CopyOnlyFullySuccessful field to given value.

### HasCopyOnlyFullySuccessful

`func (o *RunArchivalConfig) HasCopyOnlyFullySuccessful() bool`

HasCopyOnlyFullySuccessful returns a boolean if a field has been set.

### SetCopyOnlyFullySuccessfulNil

`func (o *RunArchivalConfig) SetCopyOnlyFullySuccessfulNil(b bool)`

 SetCopyOnlyFullySuccessfulNil sets the value for CopyOnlyFullySuccessful to be an explicit nil

### UnsetCopyOnlyFullySuccessful
`func (o *RunArchivalConfig) UnsetCopyOnlyFullySuccessful()`

UnsetCopyOnlyFullySuccessful ensures that no value is present for CopyOnlyFullySuccessful, not even an explicit nil
### GetId

`func (o *RunArchivalConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunArchivalConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunArchivalConfig) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *RunArchivalConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RunArchivalConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOnLegalHold

`func (o *RunArchivalConfig) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *RunArchivalConfig) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *RunArchivalConfig) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *RunArchivalConfig) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *RunArchivalConfig) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *RunArchivalConfig) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil
### GetRetention

`func (o *RunArchivalConfig) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *RunArchivalConfig) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *RunArchivalConfig) SetRetention(v Retention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *RunArchivalConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


