# TdmSnapshotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | Specifies the ID of the snapshot. | 
**IsAutomated** | Pointer to **NullableBool** | Specifies whether the snapshot was taken automatically by the scheduler. | [optional] 
**CreatedAt** | Pointer to **NullableInt64** | Specifies the time (in usecs from epoch) when the snapshot was created. | [optional] 
**UpdatedAt** | Pointer to **NullableInt64** | Specifies the time (in usecs from epoch) when the snapshot was last updated. | [optional] 
**CreatedByUser** | Pointer to [**NullableUser**](User.md) | Specifies the details of the user, who created the snapshot. This will be null for snapshots, that are taken by system, such as a scheduler. | [optional] 
**UpdatedByUser** | Pointer to [**NullableUser**](User.md) | Specifies the details of the user, who last updated the snapshot. | [optional] 

## Methods

### NewTdmSnapshotAllOf

`func NewTdmSnapshotAllOf(id NullableString, ) *TdmSnapshotAllOf`

NewTdmSnapshotAllOf instantiates a new TdmSnapshotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmSnapshotAllOfWithDefaults

`func NewTdmSnapshotAllOfWithDefaults() *TdmSnapshotAllOf`

NewTdmSnapshotAllOfWithDefaults instantiates a new TdmSnapshotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TdmSnapshotAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TdmSnapshotAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TdmSnapshotAllOf) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TdmSnapshotAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TdmSnapshotAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsAutomated

`func (o *TdmSnapshotAllOf) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TdmSnapshotAllOf) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TdmSnapshotAllOf) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TdmSnapshotAllOf) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *TdmSnapshotAllOf) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *TdmSnapshotAllOf) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetCreatedAt

`func (o *TdmSnapshotAllOf) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TdmSnapshotAllOf) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TdmSnapshotAllOf) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TdmSnapshotAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TdmSnapshotAllOf) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TdmSnapshotAllOf) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *TdmSnapshotAllOf) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TdmSnapshotAllOf) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TdmSnapshotAllOf) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TdmSnapshotAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TdmSnapshotAllOf) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TdmSnapshotAllOf) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCreatedByUser

`func (o *TdmSnapshotAllOf) GetCreatedByUser() User`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *TdmSnapshotAllOf) GetCreatedByUserOk() (*User, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *TdmSnapshotAllOf) SetCreatedByUser(v User)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *TdmSnapshotAllOf) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUserNil

`func (o *TdmSnapshotAllOf) SetCreatedByUserNil(b bool)`

 SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil

### UnsetCreatedByUser
`func (o *TdmSnapshotAllOf) UnsetCreatedByUser()`

UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
### GetUpdatedByUser

`func (o *TdmSnapshotAllOf) GetUpdatedByUser() User`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *TdmSnapshotAllOf) GetUpdatedByUserOk() (*User, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *TdmSnapshotAllOf) SetUpdatedByUser(v User)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *TdmSnapshotAllOf) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### SetUpdatedByUserNil

`func (o *TdmSnapshotAllOf) SetUpdatedByUserNil(b bool)`

 SetUpdatedByUserNil sets the value for UpdatedByUser to be an explicit nil

### UnsetUpdatedByUser
`func (o *TdmSnapshotAllOf) UnsetUpdatedByUser()`

UnsetUpdatedByUser ensures that no value is present for UpdatedByUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


