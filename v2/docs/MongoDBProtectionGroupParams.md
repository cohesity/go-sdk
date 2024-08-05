# MongoDBProtectionGroupParams

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
**CdpInfo** | Pointer to **map[string]interface{}** | Specifies the CDP related information for a given protection group. This field will only be populated when protection group is configured with a CDP policy. | [optional] 

## Methods

### NewMongoDBProtectionGroupParams

`func NewMongoDBProtectionGroupParams() *MongoDBProtectionGroupParams`

NewMongoDBProtectionGroupParams instantiates a new MongoDBProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBProtectionGroupParamsWithDefaults

`func NewMongoDBProtectionGroupParamsWithDefaults() *MongoDBProtectionGroupParams`

NewMongoDBProtectionGroupParamsWithDefaults instantiates a new MongoDBProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthMBPS

`func (o *MongoDBProtectionGroupParams) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *MongoDBProtectionGroupParams) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *MongoDBProtectionGroupParams) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *MongoDBProtectionGroupParams) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *MongoDBProtectionGroupParams) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *MongoDBProtectionGroupParams) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetConcurrency

`func (o *MongoDBProtectionGroupParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *MongoDBProtectionGroupParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *MongoDBProtectionGroupParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *MongoDBProtectionGroupParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *MongoDBProtectionGroupParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *MongoDBProtectionGroupParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetCustomSourceName

`func (o *MongoDBProtectionGroupParams) GetCustomSourceName() string`

GetCustomSourceName returns the CustomSourceName field if non-nil, zero value otherwise.

### GetCustomSourceNameOk

`func (o *MongoDBProtectionGroupParams) GetCustomSourceNameOk() (*string, bool)`

GetCustomSourceNameOk returns a tuple with the CustomSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSourceName

`func (o *MongoDBProtectionGroupParams) SetCustomSourceName(v string)`

SetCustomSourceName sets CustomSourceName field to given value.

### HasCustomSourceName

`func (o *MongoDBProtectionGroupParams) HasCustomSourceName() bool`

HasCustomSourceName returns a boolean if a field has been set.

### SetCustomSourceNameNil

`func (o *MongoDBProtectionGroupParams) SetCustomSourceNameNil(b bool)`

 SetCustomSourceNameNil sets the value for CustomSourceName to be an explicit nil

### UnsetCustomSourceName
`func (o *MongoDBProtectionGroupParams) UnsetCustomSourceName()`

UnsetCustomSourceName ensures that no value is present for CustomSourceName, not even an explicit nil
### GetExcludeObjectIds

`func (o *MongoDBProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *MongoDBProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *MongoDBProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *MongoDBProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *MongoDBProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *MongoDBProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetExcludeObjectlist

`func (o *MongoDBProtectionGroupParams) GetExcludeObjectlist() []string`

GetExcludeObjectlist returns the ExcludeObjectlist field if non-nil, zero value otherwise.

### GetExcludeObjectlistOk

`func (o *MongoDBProtectionGroupParams) GetExcludeObjectlistOk() (*[]string, bool)`

GetExcludeObjectlistOk returns a tuple with the ExcludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectlist

`func (o *MongoDBProtectionGroupParams) SetExcludeObjectlist(v []string)`

SetExcludeObjectlist sets ExcludeObjectlist field to given value.

### HasExcludeObjectlist

`func (o *MongoDBProtectionGroupParams) HasExcludeObjectlist() bool`

HasExcludeObjectlist returns a boolean if a field has been set.

### SetExcludeObjectlistNil

`func (o *MongoDBProtectionGroupParams) SetExcludeObjectlistNil(b bool)`

 SetExcludeObjectlistNil sets the value for ExcludeObjectlist to be an explicit nil

### UnsetExcludeObjectlist
`func (o *MongoDBProtectionGroupParams) UnsetExcludeObjectlist()`

UnsetExcludeObjectlist ensures that no value is present for ExcludeObjectlist, not even an explicit nil
### GetIncludeObjectlist

`func (o *MongoDBProtectionGroupParams) GetIncludeObjectlist() []string`

GetIncludeObjectlist returns the IncludeObjectlist field if non-nil, zero value otherwise.

### GetIncludeObjectlistOk

`func (o *MongoDBProtectionGroupParams) GetIncludeObjectlistOk() (*[]string, bool)`

GetIncludeObjectlistOk returns a tuple with the IncludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectlist

`func (o *MongoDBProtectionGroupParams) SetIncludeObjectlist(v []string)`

SetIncludeObjectlist sets IncludeObjectlist field to given value.

### HasIncludeObjectlist

`func (o *MongoDBProtectionGroupParams) HasIncludeObjectlist() bool`

HasIncludeObjectlist returns a boolean if a field has been set.

### SetIncludeObjectlistNil

`func (o *MongoDBProtectionGroupParams) SetIncludeObjectlistNil(b bool)`

 SetIncludeObjectlistNil sets the value for IncludeObjectlist to be an explicit nil

### UnsetIncludeObjectlist
`func (o *MongoDBProtectionGroupParams) UnsetIncludeObjectlist()`

UnsetIncludeObjectlist ensures that no value is present for IncludeObjectlist, not even an explicit nil
### GetObjects

`func (o *MongoDBProtectionGroupParams) GetObjects() []NoSqlProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MongoDBProtectionGroupParams) GetObjectsOk() (*[]NoSqlProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MongoDBProtectionGroupParams) SetObjects(v []NoSqlProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *MongoDBProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetOverwriteExcludeObjectlist

`func (o *MongoDBProtectionGroupParams) GetOverwriteExcludeObjectlist() bool`

GetOverwriteExcludeObjectlist returns the OverwriteExcludeObjectlist field if non-nil, zero value otherwise.

### GetOverwriteExcludeObjectlistOk

`func (o *MongoDBProtectionGroupParams) GetOverwriteExcludeObjectlistOk() (*bool, bool)`

GetOverwriteExcludeObjectlistOk returns a tuple with the OverwriteExcludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExcludeObjectlist

`func (o *MongoDBProtectionGroupParams) SetOverwriteExcludeObjectlist(v bool)`

SetOverwriteExcludeObjectlist sets OverwriteExcludeObjectlist field to given value.

### HasOverwriteExcludeObjectlist

`func (o *MongoDBProtectionGroupParams) HasOverwriteExcludeObjectlist() bool`

HasOverwriteExcludeObjectlist returns a boolean if a field has been set.

### SetOverwriteExcludeObjectlistNil

`func (o *MongoDBProtectionGroupParams) SetOverwriteExcludeObjectlistNil(b bool)`

 SetOverwriteExcludeObjectlistNil sets the value for OverwriteExcludeObjectlist to be an explicit nil

### UnsetOverwriteExcludeObjectlist
`func (o *MongoDBProtectionGroupParams) UnsetOverwriteExcludeObjectlist()`

UnsetOverwriteExcludeObjectlist ensures that no value is present for OverwriteExcludeObjectlist, not even an explicit nil
### GetOverwriteIncludeObjectlist

`func (o *MongoDBProtectionGroupParams) GetOverwriteIncludeObjectlist() bool`

GetOverwriteIncludeObjectlist returns the OverwriteIncludeObjectlist field if non-nil, zero value otherwise.

### GetOverwriteIncludeObjectlistOk

`func (o *MongoDBProtectionGroupParams) GetOverwriteIncludeObjectlistOk() (*bool, bool)`

GetOverwriteIncludeObjectlistOk returns a tuple with the OverwriteIncludeObjectlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteIncludeObjectlist

`func (o *MongoDBProtectionGroupParams) SetOverwriteIncludeObjectlist(v bool)`

SetOverwriteIncludeObjectlist sets OverwriteIncludeObjectlist field to given value.

### HasOverwriteIncludeObjectlist

`func (o *MongoDBProtectionGroupParams) HasOverwriteIncludeObjectlist() bool`

HasOverwriteIncludeObjectlist returns a boolean if a field has been set.

### SetOverwriteIncludeObjectlistNil

`func (o *MongoDBProtectionGroupParams) SetOverwriteIncludeObjectlistNil(b bool)`

 SetOverwriteIncludeObjectlistNil sets the value for OverwriteIncludeObjectlist to be an explicit nil

### UnsetOverwriteIncludeObjectlist
`func (o *MongoDBProtectionGroupParams) UnsetOverwriteIncludeObjectlist()`

UnsetOverwriteIncludeObjectlist ensures that no value is present for OverwriteIncludeObjectlist, not even an explicit nil
### GetSourceId

`func (o *MongoDBProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *MongoDBProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *MongoDBProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *MongoDBProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *MongoDBProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *MongoDBProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *MongoDBProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *MongoDBProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *MongoDBProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *MongoDBProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *MongoDBProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *MongoDBProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetCdpInfo

`func (o *MongoDBProtectionGroupParams) GetCdpInfo() map[string]interface{}`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *MongoDBProtectionGroupParams) GetCdpInfoOk() (*map[string]interface{}, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *MongoDBProtectionGroupParams) SetCdpInfo(v map[string]interface{})`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *MongoDBProtectionGroupParams) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


