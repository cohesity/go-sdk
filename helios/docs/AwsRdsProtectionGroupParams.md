# AwsRdsProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]AwsRdsProtectionGroupObjectParams**](AwsRdsProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**RdsTagIds** | Pointer to **[][]int64** | Array of arrays of RDS Tag Ids that Specify db instaces to Protect. | [optional] 
**ExcludeRdsTagIds** | Pointer to **[][]int64** | Array of arrays of RDS Tag Ids that Specify db instaces to Exclude. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 

## Methods

### NewAwsRdsProtectionGroupParams

`func NewAwsRdsProtectionGroupParams() *AwsRdsProtectionGroupParams`

NewAwsRdsProtectionGroupParams instantiates a new AwsRdsProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRdsProtectionGroupParamsWithDefaults

`func NewAwsRdsProtectionGroupParamsWithDefaults() *AwsRdsProtectionGroupParams`

NewAwsRdsProtectionGroupParamsWithDefaults instantiates a new AwsRdsProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *AwsRdsProtectionGroupParams) GetObjects() []AwsRdsProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AwsRdsProtectionGroupParams) GetObjectsOk() (*[]AwsRdsProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AwsRdsProtectionGroupParams) SetObjects(v []AwsRdsProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AwsRdsProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *AwsRdsProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *AwsRdsProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *AwsRdsProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *AwsRdsProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *AwsRdsProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *AwsRdsProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetRdsTagIds

`func (o *AwsRdsProtectionGroupParams) GetRdsTagIds() [][]int64`

GetRdsTagIds returns the RdsTagIds field if non-nil, zero value otherwise.

### GetRdsTagIdsOk

`func (o *AwsRdsProtectionGroupParams) GetRdsTagIdsOk() (*[][]int64, bool)`

GetRdsTagIdsOk returns a tuple with the RdsTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsTagIds

`func (o *AwsRdsProtectionGroupParams) SetRdsTagIds(v [][]int64)`

SetRdsTagIds sets RdsTagIds field to given value.

### HasRdsTagIds

`func (o *AwsRdsProtectionGroupParams) HasRdsTagIds() bool`

HasRdsTagIds returns a boolean if a field has been set.

### SetRdsTagIdsNil

`func (o *AwsRdsProtectionGroupParams) SetRdsTagIdsNil(b bool)`

 SetRdsTagIdsNil sets the value for RdsTagIds to be an explicit nil

### UnsetRdsTagIds
`func (o *AwsRdsProtectionGroupParams) UnsetRdsTagIds()`

UnsetRdsTagIds ensures that no value is present for RdsTagIds, not even an explicit nil
### GetExcludeRdsTagIds

`func (o *AwsRdsProtectionGroupParams) GetExcludeRdsTagIds() [][]int64`

GetExcludeRdsTagIds returns the ExcludeRdsTagIds field if non-nil, zero value otherwise.

### GetExcludeRdsTagIdsOk

`func (o *AwsRdsProtectionGroupParams) GetExcludeRdsTagIdsOk() (*[][]int64, bool)`

GetExcludeRdsTagIdsOk returns a tuple with the ExcludeRdsTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRdsTagIds

`func (o *AwsRdsProtectionGroupParams) SetExcludeRdsTagIds(v [][]int64)`

SetExcludeRdsTagIds sets ExcludeRdsTagIds field to given value.

### HasExcludeRdsTagIds

`func (o *AwsRdsProtectionGroupParams) HasExcludeRdsTagIds() bool`

HasExcludeRdsTagIds returns a boolean if a field has been set.

### SetExcludeRdsTagIdsNil

`func (o *AwsRdsProtectionGroupParams) SetExcludeRdsTagIdsNil(b bool)`

 SetExcludeRdsTagIdsNil sets the value for ExcludeRdsTagIds to be an explicit nil

### UnsetExcludeRdsTagIds
`func (o *AwsRdsProtectionGroupParams) UnsetExcludeRdsTagIds()`

UnsetExcludeRdsTagIds ensures that no value is present for ExcludeRdsTagIds, not even an explicit nil
### GetSourceId

`func (o *AwsRdsProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AwsRdsProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AwsRdsProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AwsRdsProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *AwsRdsProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *AwsRdsProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *AwsRdsProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AwsRdsProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AwsRdsProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AwsRdsProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *AwsRdsProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *AwsRdsProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


