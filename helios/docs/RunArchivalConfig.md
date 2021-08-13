# RunArchivalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the Archival target to copy the Snapshots to. | 
**ArchivalTargetType** | **NullableString** | Specifies the snapshot&#39;s archival target type from which recovery has been performed. | 
**Retention** | Pointer to [**Retention**](Retention.md) |  | [optional] 

## Methods

### NewRunArchivalConfig

`func NewRunArchivalConfig(id NullableInt64, archivalTargetType NullableString, ) *RunArchivalConfig`

NewRunArchivalConfig instantiates a new RunArchivalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunArchivalConfigWithDefaults

`func NewRunArchivalConfigWithDefaults() *RunArchivalConfig`

NewRunArchivalConfigWithDefaults instantiates a new RunArchivalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


