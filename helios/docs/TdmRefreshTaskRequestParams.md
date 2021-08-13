# TdmRefreshTaskRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneId** | **NullableString** | Specifies the ID of the clone, which needs to be refreshed. | 
**SnapshotId** | **NullableString** | Specifies the snapshot ID, using which the clone is to be refreshed. | 
**Description** | Pointer to **NullableString** | Specifies the description of the clone refresh task. | [optional] 

## Methods

### NewTdmRefreshTaskRequestParams

`func NewTdmRefreshTaskRequestParams(cloneId NullableString, snapshotId NullableString, ) *TdmRefreshTaskRequestParams`

NewTdmRefreshTaskRequestParams instantiates a new TdmRefreshTaskRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmRefreshTaskRequestParamsWithDefaults

`func NewTdmRefreshTaskRequestParamsWithDefaults() *TdmRefreshTaskRequestParams`

NewTdmRefreshTaskRequestParamsWithDefaults instantiates a new TdmRefreshTaskRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneId

`func (o *TdmRefreshTaskRequestParams) GetCloneId() string`

GetCloneId returns the CloneId field if non-nil, zero value otherwise.

### GetCloneIdOk

`func (o *TdmRefreshTaskRequestParams) GetCloneIdOk() (*string, bool)`

GetCloneIdOk returns a tuple with the CloneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneId

`func (o *TdmRefreshTaskRequestParams) SetCloneId(v string)`

SetCloneId sets CloneId field to given value.


### SetCloneIdNil

`func (o *TdmRefreshTaskRequestParams) SetCloneIdNil(b bool)`

 SetCloneIdNil sets the value for CloneId to be an explicit nil

### UnsetCloneId
`func (o *TdmRefreshTaskRequestParams) UnsetCloneId()`

UnsetCloneId ensures that no value is present for CloneId, not even an explicit nil
### GetSnapshotId

`func (o *TdmRefreshTaskRequestParams) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *TdmRefreshTaskRequestParams) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *TdmRefreshTaskRequestParams) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### SetSnapshotIdNil

`func (o *TdmRefreshTaskRequestParams) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *TdmRefreshTaskRequestParams) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetDescription

`func (o *TdmRefreshTaskRequestParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TdmRefreshTaskRequestParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TdmRefreshTaskRequestParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TdmRefreshTaskRequestParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TdmRefreshTaskRequestParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TdmRefreshTaskRequestParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


