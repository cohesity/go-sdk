# KubernetesProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]KubernetesProtectionGroupObjectParams**](KubernetesProtectionGroupObjectParams.md) | Specifies the objects included in the Protection Group. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**LabelIds** | Pointer to **[][]int64** | Array of array of label IDs that specify labels to protect. Optionally specify a list of labels to protect by listing protection source ids of labels in this two dimensional array. Using this two dimensional array of label IDs, the cluster generates a list of namespaces to protect, which are derived from intersections of the inner arrays and union of the outer array. | [optional] 
**ExcludeLabelIds** | Pointer to **[][]int64** | Array of arrays of label IDs that specify labels to exclude. Optionally specify a list of labels to exclude from protecting by listing protection source ids of labels in this two dimensional array. Using this two dimensional array of label IDs, the Cluster generates a list of namespaces to exclude from protecting, which are derived from intersections of the inner arrays and union of the outer array. | [optional] 
**IncludeParams** | Pointer to [**NullableKubernetesFilterParams**](KubernetesFilterParams.md) |  | [optional] 
**ExcludeParams** | Pointer to [**NullableKubernetesFilterParams**](KubernetesFilterParams.md) |  | [optional] 

## Methods

### NewKubernetesProtectionGroupParams

`func NewKubernetesProtectionGroupParams() *KubernetesProtectionGroupParams`

NewKubernetesProtectionGroupParams instantiates a new KubernetesProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesProtectionGroupParamsWithDefaults

`func NewKubernetesProtectionGroupParamsWithDefaults() *KubernetesProtectionGroupParams`

NewKubernetesProtectionGroupParamsWithDefaults instantiates a new KubernetesProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *KubernetesProtectionGroupParams) GetObjects() []KubernetesProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *KubernetesProtectionGroupParams) GetObjectsOk() (*[]KubernetesProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *KubernetesProtectionGroupParams) SetObjects(v []KubernetesProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *KubernetesProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *KubernetesProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *KubernetesProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *KubernetesProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *KubernetesProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetSourceId

`func (o *KubernetesProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *KubernetesProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *KubernetesProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *KubernetesProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *KubernetesProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *KubernetesProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *KubernetesProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *KubernetesProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *KubernetesProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *KubernetesProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *KubernetesProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *KubernetesProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetLabelIds

`func (o *KubernetesProtectionGroupParams) GetLabelIds() [][]int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *KubernetesProtectionGroupParams) GetLabelIdsOk() (*[][]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *KubernetesProtectionGroupParams) SetLabelIds(v [][]int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *KubernetesProtectionGroupParams) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### SetLabelIdsNil

`func (o *KubernetesProtectionGroupParams) SetLabelIdsNil(b bool)`

 SetLabelIdsNil sets the value for LabelIds to be an explicit nil

### UnsetLabelIds
`func (o *KubernetesProtectionGroupParams) UnsetLabelIds()`

UnsetLabelIds ensures that no value is present for LabelIds, not even an explicit nil
### GetExcludeLabelIds

`func (o *KubernetesProtectionGroupParams) GetExcludeLabelIds() [][]int64`

GetExcludeLabelIds returns the ExcludeLabelIds field if non-nil, zero value otherwise.

### GetExcludeLabelIdsOk

`func (o *KubernetesProtectionGroupParams) GetExcludeLabelIdsOk() (*[][]int64, bool)`

GetExcludeLabelIdsOk returns a tuple with the ExcludeLabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLabelIds

`func (o *KubernetesProtectionGroupParams) SetExcludeLabelIds(v [][]int64)`

SetExcludeLabelIds sets ExcludeLabelIds field to given value.

### HasExcludeLabelIds

`func (o *KubernetesProtectionGroupParams) HasExcludeLabelIds() bool`

HasExcludeLabelIds returns a boolean if a field has been set.

### GetIncludeParams

`func (o *KubernetesProtectionGroupParams) GetIncludeParams() KubernetesFilterParams`

GetIncludeParams returns the IncludeParams field if non-nil, zero value otherwise.

### GetIncludeParamsOk

`func (o *KubernetesProtectionGroupParams) GetIncludeParamsOk() (*KubernetesFilterParams, bool)`

GetIncludeParamsOk returns a tuple with the IncludeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeParams

`func (o *KubernetesProtectionGroupParams) SetIncludeParams(v KubernetesFilterParams)`

SetIncludeParams sets IncludeParams field to given value.

### HasIncludeParams

`func (o *KubernetesProtectionGroupParams) HasIncludeParams() bool`

HasIncludeParams returns a boolean if a field has been set.

### SetIncludeParamsNil

`func (o *KubernetesProtectionGroupParams) SetIncludeParamsNil(b bool)`

 SetIncludeParamsNil sets the value for IncludeParams to be an explicit nil

### UnsetIncludeParams
`func (o *KubernetesProtectionGroupParams) UnsetIncludeParams()`

UnsetIncludeParams ensures that no value is present for IncludeParams, not even an explicit nil
### GetExcludeParams

`func (o *KubernetesProtectionGroupParams) GetExcludeParams() KubernetesFilterParams`

GetExcludeParams returns the ExcludeParams field if non-nil, zero value otherwise.

### GetExcludeParamsOk

`func (o *KubernetesProtectionGroupParams) GetExcludeParamsOk() (*KubernetesFilterParams, bool)`

GetExcludeParamsOk returns a tuple with the ExcludeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeParams

`func (o *KubernetesProtectionGroupParams) SetExcludeParams(v KubernetesFilterParams)`

SetExcludeParams sets ExcludeParams field to given value.

### HasExcludeParams

`func (o *KubernetesProtectionGroupParams) HasExcludeParams() bool`

HasExcludeParams returns a boolean if a field has been set.

### SetExcludeParamsNil

`func (o *KubernetesProtectionGroupParams) SetExcludeParamsNil(b bool)`

 SetExcludeParamsNil sets the value for ExcludeParams to be an explicit nil

### UnsetExcludeParams
`func (o *KubernetesProtectionGroupParams) UnsetExcludeParams()`

UnsetExcludeParams ensures that no value is present for ExcludeParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


