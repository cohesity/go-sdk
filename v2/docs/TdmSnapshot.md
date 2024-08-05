# TdmSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **NullableString** | Specifies the label for the snapshot. | [optional] 
**CreatedAt** | Pointer to **NullableInt64** | Specifies the time (in usecs from epoch) when the snapshot was created. | [optional] 
**CreatedByUser** | Pointer to **map[string]interface{}** | Specifies the details of the user, who created the snapshot. This will be null for snapshots, that are taken by system, such as a scheduler. | [optional] 
**Id** | **NullableString** | Specifies the ID of the snapshot. | 
**IsAutomated** | Pointer to **NullableBool** | Specifies whether the snapshot was taken automatically by the scheduler. | [optional] 
**UpdatedAt** | Pointer to **NullableInt64** | Specifies the time (in usecs from epoch) when the snapshot was last updated. | [optional] 
**UpdatedByUser** | Pointer to **map[string]interface{}** | Specifies the details of the user, who last updated the snapshot. | [optional] 

## Methods

### NewTdmSnapshot

`func NewTdmSnapshot(id NullableString, ) *TdmSnapshot`

NewTdmSnapshot instantiates a new TdmSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmSnapshotWithDefaults

`func NewTdmSnapshotWithDefaults() *TdmSnapshot`

NewTdmSnapshotWithDefaults instantiates a new TdmSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *TdmSnapshot) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TdmSnapshot) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TdmSnapshot) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TdmSnapshot) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *TdmSnapshot) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *TdmSnapshot) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetCreatedAt

`func (o *TdmSnapshot) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TdmSnapshot) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TdmSnapshot) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TdmSnapshot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TdmSnapshot) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TdmSnapshot) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedByUser

`func (o *TdmSnapshot) GetCreatedByUser() map[string]interface{}`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *TdmSnapshot) GetCreatedByUserOk() (*map[string]interface{}, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *TdmSnapshot) SetCreatedByUser(v map[string]interface{})`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *TdmSnapshot) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUserNil

`func (o *TdmSnapshot) SetCreatedByUserNil(b bool)`

 SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil

### UnsetCreatedByUser
`func (o *TdmSnapshot) UnsetCreatedByUser()`

UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
### GetId

`func (o *TdmSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TdmSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TdmSnapshot) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TdmSnapshot) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TdmSnapshot) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsAutomated

`func (o *TdmSnapshot) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TdmSnapshot) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TdmSnapshot) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TdmSnapshot) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *TdmSnapshot) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *TdmSnapshot) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetUpdatedAt

`func (o *TdmSnapshot) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TdmSnapshot) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TdmSnapshot) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TdmSnapshot) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TdmSnapshot) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TdmSnapshot) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedByUser

`func (o *TdmSnapshot) GetUpdatedByUser() map[string]interface{}`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *TdmSnapshot) GetUpdatedByUserOk() (*map[string]interface{}, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *TdmSnapshot) SetUpdatedByUser(v map[string]interface{})`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *TdmSnapshot) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### SetUpdatedByUserNil

`func (o *TdmSnapshot) SetUpdatedByUserNil(b bool)`

 SetUpdatedByUserNil sets the value for UpdatedByUser to be an explicit nil

### UnsetUpdatedByUser
`func (o *TdmSnapshot) UnsetUpdatedByUser()`

UnsetUpdatedByUser ensures that no value is present for UpdatedByUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


