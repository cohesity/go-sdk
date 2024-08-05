# NoSqlProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthMBPS** | Pointer to **NullableInt64** | Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. | [optional] 
**CustomSourceName** | Pointer to **NullableString** | The user specified name for the Source on which this protection was run. | [optional] [readonly] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**ExcludeObjectlist** | Pointer to **[]string** | Specifies the list of fully qualified name of the entities to exclude for protection. | [optional] 
**IncludeObjectlist** | Pointer to **[]string** | Specifies the list of fully qualified name of the entities to include for protection. | [optional] 
**Objects** | Pointer to [**[]NoSqlProtectionGroupObjectParams**](NoSqlProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | [optional] 
**OverwriteExcludeObjectlist** | Pointer to **NullableBool** | If disabled - The excludeObjectlist is merged with the existing exclude_sources_vec, preserving any existing elements while incorporating new ones. | [optional] [default to true]
**OverwriteIncludeObjectlist** | Pointer to **NullableBool** | If disabled - The includeObjectlist is merged with the existing sources_vec, preserving any existing elements while incorporating new ones. | [optional] [default to true]
**SourceId** | Pointer to **NullableInt64** | Object ID of the Source on which this protection was run . | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the Source on which this protection was run. | [optional] [readonly] 

## Methods

### NewNoSqlProtectionGroupParams

`func NewNoSqlProtectionGroupParams() *NoSqlProtectionGroupParams`

NewNoSqlProtectionGroupParams instantiates a new NoSqlProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoSqlProtectionGroupParamsWithDefaults

`func NewNoSqlProtectionGroupParamsWithDefaults() *NoSqlProtectionGroupParams`

NewNoSqlProtectionGroupParamsWithDefaults instantiates a new NoSqlProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthMBPS

`func (o *NoSqlProtectionGroupParams) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *NoSqlProtectionGroupParams) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *NoSqlProtectionGroupParams) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *NoSqlProtectionGroupParams) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *NoSqlProtectionGroupParams) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *NoSqlProtectionGroupParams) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetConcurrency

`func (o *NoSqlProtectionGroupParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *NoSqlProtectionGroupParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *NoSqlProtectionGroupParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *NoSqlProtectionGroupParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *NoSqlProtectionGroupParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *NoSqlProtectionGroupParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetCustomSourceName

`func (o *NoSqlProtectionGroupParams) GetCustomSourceName() string`

GetCustomSourceName returns the CustomSourceName field if non-nil, zero value otherwise.

### GetCustomSourceNameOk

`func (o *NoSqlProtectionGroupParams) GetCustomSourceNameOk() (*string, bool)`

GetCustomSourceNameOk returns a tuple with the CustomSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSourceName

`func (o *NoSqlProtectionGroupParams) SetCustomSourceName(v string)`

SetCustomSourceName sets CustomSourceName field to given value.

### HasCustomSourceName

`func (o *NoSqlProtectionGroupParams) HasCustomSourceName() bool`

HasCustomSourceName returns a boolean if a field has been set.

### SetCustomSourceNameNil

`func (o *NoSqlProtectionGroupParams) SetCustomSourceNameNil(b bool)`

 SetCustomSourceNameNil sets the value for CustomSourceName to be an explicit nil

### UnsetCustomSourceName
`func (o *NoSqlProtectionGroupParams) UnsetCustomSourceName()`

UnsetCustomSourceName ensures that no value is present for CustomSourceName, not even an explicit nil
### GetExcludeObjectIds

`func (o *NoSqlProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *NoSqlProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *NoSqlProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *NoSqlProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *NoSqlProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *NoSqlProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetExcludeObjectlist

`func (o *NoSqlProtectionGroupParams) GetExcludeObjectlist() []string`

GetExcludeObjectlist returns the ExcludeObjectlist field if non-nil, zero value otherwise.

### GetExcludeObjectlistOk

`func (o *NoSqlProtectionGroupParams) GetExcludeObjectlistOk() (*[]string, bool)`

GetExcludeObjectlistOk returns a tuple with the ExcludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectlist

`func (o *NoSqlProtectionGroupParams) SetExcludeObjectlist(v []string)`

SetExcludeObjectlist sets ExcludeObjectlist field to given value.

### HasExcludeObjectlist

`func (o *NoSqlProtectionGroupParams) HasExcludeObjectlist() bool`

HasExcludeObjectlist returns a boolean if a field has been set.

### SetExcludeObjectlistNil

`func (o *NoSqlProtectionGroupParams) SetExcludeObjectlistNil(b bool)`

 SetExcludeObjectlistNil sets the value for ExcludeObjectlist to be an explicit nil

### UnsetExcludeObjectlist
`func (o *NoSqlProtectionGroupParams) UnsetExcludeObjectlist()`

UnsetExcludeObjectlist ensures that no value is present for ExcludeObjectlist, not even an explicit nil
### GetIncludeObjectlist

`func (o *NoSqlProtectionGroupParams) GetIncludeObjectlist() []string`

GetIncludeObjectlist returns the IncludeObjectlist field if non-nil, zero value otherwise.

### GetIncludeObjectlistOk

`func (o *NoSqlProtectionGroupParams) GetIncludeObjectlistOk() (*[]string, bool)`

GetIncludeObjectlistOk returns a tuple with the IncludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectlist

`func (o *NoSqlProtectionGroupParams) SetIncludeObjectlist(v []string)`

SetIncludeObjectlist sets IncludeObjectlist field to given value.

### HasIncludeObjectlist

`func (o *NoSqlProtectionGroupParams) HasIncludeObjectlist() bool`

HasIncludeObjectlist returns a boolean if a field has been set.

### SetIncludeObjectlistNil

`func (o *NoSqlProtectionGroupParams) SetIncludeObjectlistNil(b bool)`

 SetIncludeObjectlistNil sets the value for IncludeObjectlist to be an explicit nil

### UnsetIncludeObjectlist
`func (o *NoSqlProtectionGroupParams) UnsetIncludeObjectlist()`

UnsetIncludeObjectlist ensures that no value is present for IncludeObjectlist, not even an explicit nil
### GetObjects

`func (o *NoSqlProtectionGroupParams) GetObjects() []NoSqlProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *NoSqlProtectionGroupParams) GetObjectsOk() (*[]NoSqlProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *NoSqlProtectionGroupParams) SetObjects(v []NoSqlProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *NoSqlProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetOverwriteExcludeObjectlist

`func (o *NoSqlProtectionGroupParams) GetOverwriteExcludeObjectlist() bool`

GetOverwriteExcludeObjectlist returns the OverwriteExcludeObjectlist field if non-nil, zero value otherwise.

### GetOverwriteExcludeObjectlistOk

`func (o *NoSqlProtectionGroupParams) GetOverwriteExcludeObjectlistOk() (*bool, bool)`

GetOverwriteExcludeObjectlistOk returns a tuple with the OverwriteExcludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExcludeObjectlist

`func (o *NoSqlProtectionGroupParams) SetOverwriteExcludeObjectlist(v bool)`

SetOverwriteExcludeObjectlist sets OverwriteExcludeObjectlist field to given value.

### HasOverwriteExcludeObjectlist

`func (o *NoSqlProtectionGroupParams) HasOverwriteExcludeObjectlist() bool`

HasOverwriteExcludeObjectlist returns a boolean if a field has been set.

### SetOverwriteExcludeObjectlistNil

`func (o *NoSqlProtectionGroupParams) SetOverwriteExcludeObjectlistNil(b bool)`

 SetOverwriteExcludeObjectlistNil sets the value for OverwriteExcludeObjectlist to be an explicit nil

### UnsetOverwriteExcludeObjectlist
`func (o *NoSqlProtectionGroupParams) UnsetOverwriteExcludeObjectlist()`

UnsetOverwriteExcludeObjectlist ensures that no value is present for OverwriteExcludeObjectlist, not even an explicit nil
### GetOverwriteIncludeObjectlist

`func (o *NoSqlProtectionGroupParams) GetOverwriteIncludeObjectlist() bool`

GetOverwriteIncludeObjectlist returns the OverwriteIncludeObjectlist field if non-nil, zero value otherwise.

### GetOverwriteIncludeObjectlistOk

`func (o *NoSqlProtectionGroupParams) GetOverwriteIncludeObjectlistOk() (*bool, bool)`

GetOverwriteIncludeObjectlistOk returns a tuple with the OverwriteIncludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteIncludeObjectlist

`func (o *NoSqlProtectionGroupParams) SetOverwriteIncludeObjectlist(v bool)`

SetOverwriteIncludeObjectlist sets OverwriteIncludeObjectlist field to given value.

### HasOverwriteIncludeObjectlist

`func (o *NoSqlProtectionGroupParams) HasOverwriteIncludeObjectlist() bool`

HasOverwriteIncludeObjectlist returns a boolean if a field has been set.

### SetOverwriteIncludeObjectlistNil

`func (o *NoSqlProtectionGroupParams) SetOverwriteIncludeObjectlistNil(b bool)`

 SetOverwriteIncludeObjectlistNil sets the value for OverwriteIncludeObjectlist to be an explicit nil

### UnsetOverwriteIncludeObjectlist
`func (o *NoSqlProtectionGroupParams) UnsetOverwriteIncludeObjectlist()`

UnsetOverwriteIncludeObjectlist ensures that no value is present for OverwriteIncludeObjectlist, not even an explicit nil
### GetSourceId

`func (o *NoSqlProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *NoSqlProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *NoSqlProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *NoSqlProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *NoSqlProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *NoSqlProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *NoSqlProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *NoSqlProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *NoSqlProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *NoSqlProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *NoSqlProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *NoSqlProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


