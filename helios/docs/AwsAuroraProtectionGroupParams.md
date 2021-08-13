# AwsAuroraProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]AwsAuroraProtectionGroupObjectParams**](AwsAuroraProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**AuroraTagIds** | Pointer to **[][]int64** | Array of arrays of Aurora Tag Ids that specify aurora clusters to protect. | [optional] 
**ExcludeAuroraTagIds** | Pointer to **[][]int64** | Array of arrays of RDS Tag Ids that specify aurora clusters to exclude. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 

## Methods

### NewAwsAuroraProtectionGroupParams

`func NewAwsAuroraProtectionGroupParams() *AwsAuroraProtectionGroupParams`

NewAwsAuroraProtectionGroupParams instantiates a new AwsAuroraProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAuroraProtectionGroupParamsWithDefaults

`func NewAwsAuroraProtectionGroupParamsWithDefaults() *AwsAuroraProtectionGroupParams`

NewAwsAuroraProtectionGroupParamsWithDefaults instantiates a new AwsAuroraProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *AwsAuroraProtectionGroupParams) GetObjects() []AwsAuroraProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AwsAuroraProtectionGroupParams) GetObjectsOk() (*[]AwsAuroraProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AwsAuroraProtectionGroupParams) SetObjects(v []AwsAuroraProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AwsAuroraProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *AwsAuroraProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *AwsAuroraProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *AwsAuroraProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *AwsAuroraProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *AwsAuroraProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *AwsAuroraProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetAuroraTagIds

`func (o *AwsAuroraProtectionGroupParams) GetAuroraTagIds() [][]int64`

GetAuroraTagIds returns the AuroraTagIds field if non-nil, zero value otherwise.

### GetAuroraTagIdsOk

`func (o *AwsAuroraProtectionGroupParams) GetAuroraTagIdsOk() (*[][]int64, bool)`

GetAuroraTagIdsOk returns a tuple with the AuroraTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuroraTagIds

`func (o *AwsAuroraProtectionGroupParams) SetAuroraTagIds(v [][]int64)`

SetAuroraTagIds sets AuroraTagIds field to given value.

### HasAuroraTagIds

`func (o *AwsAuroraProtectionGroupParams) HasAuroraTagIds() bool`

HasAuroraTagIds returns a boolean if a field has been set.

### SetAuroraTagIdsNil

`func (o *AwsAuroraProtectionGroupParams) SetAuroraTagIdsNil(b bool)`

 SetAuroraTagIdsNil sets the value for AuroraTagIds to be an explicit nil

### UnsetAuroraTagIds
`func (o *AwsAuroraProtectionGroupParams) UnsetAuroraTagIds()`

UnsetAuroraTagIds ensures that no value is present for AuroraTagIds, not even an explicit nil
### GetExcludeAuroraTagIds

`func (o *AwsAuroraProtectionGroupParams) GetExcludeAuroraTagIds() [][]int64`

GetExcludeAuroraTagIds returns the ExcludeAuroraTagIds field if non-nil, zero value otherwise.

### GetExcludeAuroraTagIdsOk

`func (o *AwsAuroraProtectionGroupParams) GetExcludeAuroraTagIdsOk() (*[][]int64, bool)`

GetExcludeAuroraTagIdsOk returns a tuple with the ExcludeAuroraTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAuroraTagIds

`func (o *AwsAuroraProtectionGroupParams) SetExcludeAuroraTagIds(v [][]int64)`

SetExcludeAuroraTagIds sets ExcludeAuroraTagIds field to given value.

### HasExcludeAuroraTagIds

`func (o *AwsAuroraProtectionGroupParams) HasExcludeAuroraTagIds() bool`

HasExcludeAuroraTagIds returns a boolean if a field has been set.

### SetExcludeAuroraTagIdsNil

`func (o *AwsAuroraProtectionGroupParams) SetExcludeAuroraTagIdsNil(b bool)`

 SetExcludeAuroraTagIdsNil sets the value for ExcludeAuroraTagIds to be an explicit nil

### UnsetExcludeAuroraTagIds
`func (o *AwsAuroraProtectionGroupParams) UnsetExcludeAuroraTagIds()`

UnsetExcludeAuroraTagIds ensures that no value is present for ExcludeAuroraTagIds, not even an explicit nil
### GetSourceId

`func (o *AwsAuroraProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AwsAuroraProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AwsAuroraProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AwsAuroraProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *AwsAuroraProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AwsAuroraProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *AwsAuroraProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AwsAuroraProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AwsAuroraProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AwsAuroraProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *AwsAuroraProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *AwsAuroraProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


