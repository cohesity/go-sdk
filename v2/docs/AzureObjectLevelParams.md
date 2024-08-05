# AzureObjectLevelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskExclusionParams** | Pointer to [**AzureDiskExclusionParams**](AzureDiskExclusionParams.md) |  | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects not to be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag. This can be used to ignore specific objects (can include tags) under a parent object which has been included for protection. | [optional] 
**Id** | **NullableInt64** | Specifies the id of the object being protected. This can be a leaf level or non leaf level object. | 

## Methods

### NewAzureObjectLevelParams

`func NewAzureObjectLevelParams(id NullableInt64, ) *AzureObjectLevelParams`

NewAzureObjectLevelParams instantiates a new AzureObjectLevelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureObjectLevelParamsWithDefaults

`func NewAzureObjectLevelParamsWithDefaults() *AzureObjectLevelParams`

NewAzureObjectLevelParamsWithDefaults instantiates a new AzureObjectLevelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskExclusionParams

`func (o *AzureObjectLevelParams) GetDiskExclusionParams() AzureDiskExclusionParams`

GetDiskExclusionParams returns the DiskExclusionParams field if non-nil, zero value otherwise.

### GetDiskExclusionParamsOk

`func (o *AzureObjectLevelParams) GetDiskExclusionParamsOk() (*AzureDiskExclusionParams, bool)`

GetDiskExclusionParamsOk returns a tuple with the DiskExclusionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskExclusionParams

`func (o *AzureObjectLevelParams) SetDiskExclusionParams(v AzureDiskExclusionParams)`

SetDiskExclusionParams sets DiskExclusionParams field to given value.

### HasDiskExclusionParams

`func (o *AzureObjectLevelParams) HasDiskExclusionParams() bool`

HasDiskExclusionParams returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *AzureObjectLevelParams) GetExcludeObjectIds() []*int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *AzureObjectLevelParams) GetExcludeObjectIdsOk() (*[]*int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *AzureObjectLevelParams) SetExcludeObjectIds(v []*int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *AzureObjectLevelParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetId

`func (o *AzureObjectLevelParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureObjectLevelParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureObjectLevelParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AzureObjectLevelParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AzureObjectLevelParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


