# CassandraAdditionalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraClasspathSuffix** | Pointer to **NullableString** | Cassandra classpath suffix. | [optional] 
**CassandraPartitioner** | Pointer to **NullableString** | Required in compaction. | [optional] 
**CassandraVersion** | Pointer to **NullableString** | Cassandra and DSE Versions. Discovery code will attempt to discover the versions. | [optional] 
**DataCenterVec** | Pointer to **[]string** | Data center information is required for backup and recovery. | [optional] 
**DseSolrInfo** | Pointer to [**DSESolrInfo**](DSESolrInfo.md) |  | [optional] 
**DseVersion** | Pointer to **NullableString** |  | [optional] 
**TieredStorageDirsMap** | Pointer to [**[]NodeToTieredStorageDirectoriesMap**](NodeToTieredStorageDirectoriesMap.md) | Map of nodes to tiered storage directories | [optional] 

## Methods

### NewCassandraAdditionalParams

`func NewCassandraAdditionalParams() *CassandraAdditionalParams`

NewCassandraAdditionalParams instantiates a new CassandraAdditionalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraAdditionalParamsWithDefaults

`func NewCassandraAdditionalParamsWithDefaults() *CassandraAdditionalParams`

NewCassandraAdditionalParamsWithDefaults instantiates a new CassandraAdditionalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraClasspathSuffix

`func (o *CassandraAdditionalParams) GetCassandraClasspathSuffix() string`

GetCassandraClasspathSuffix returns the CassandraClasspathSuffix field if non-nil, zero value otherwise.

### GetCassandraClasspathSuffixOk

`func (o *CassandraAdditionalParams) GetCassandraClasspathSuffixOk() (*string, bool)`

GetCassandraClasspathSuffixOk returns a tuple with the CassandraClasspathSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraClasspathSuffix

`func (o *CassandraAdditionalParams) SetCassandraClasspathSuffix(v string)`

SetCassandraClasspathSuffix sets CassandraClasspathSuffix field to given value.

### HasCassandraClasspathSuffix

`func (o *CassandraAdditionalParams) HasCassandraClasspathSuffix() bool`

HasCassandraClasspathSuffix returns a boolean if a field has been set.

### SetCassandraClasspathSuffixNil

`func (o *CassandraAdditionalParams) SetCassandraClasspathSuffixNil(b bool)`

 SetCassandraClasspathSuffixNil sets the value for CassandraClasspathSuffix to be an explicit nil

### UnsetCassandraClasspathSuffix
`func (o *CassandraAdditionalParams) UnsetCassandraClasspathSuffix()`

UnsetCassandraClasspathSuffix ensures that no value is present for CassandraClasspathSuffix, not even an explicit nil
### GetCassandraPartitioner

`func (o *CassandraAdditionalParams) GetCassandraPartitioner() string`

GetCassandraPartitioner returns the CassandraPartitioner field if non-nil, zero value otherwise.

### GetCassandraPartitionerOk

`func (o *CassandraAdditionalParams) GetCassandraPartitionerOk() (*string, bool)`

GetCassandraPartitionerOk returns a tuple with the CassandraPartitioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraPartitioner

`func (o *CassandraAdditionalParams) SetCassandraPartitioner(v string)`

SetCassandraPartitioner sets CassandraPartitioner field to given value.

### HasCassandraPartitioner

`func (o *CassandraAdditionalParams) HasCassandraPartitioner() bool`

HasCassandraPartitioner returns a boolean if a field has been set.

### SetCassandraPartitionerNil

`func (o *CassandraAdditionalParams) SetCassandraPartitionerNil(b bool)`

 SetCassandraPartitionerNil sets the value for CassandraPartitioner to be an explicit nil

### UnsetCassandraPartitioner
`func (o *CassandraAdditionalParams) UnsetCassandraPartitioner()`

UnsetCassandraPartitioner ensures that no value is present for CassandraPartitioner, not even an explicit nil
### GetCassandraVersion

`func (o *CassandraAdditionalParams) GetCassandraVersion() string`

GetCassandraVersion returns the CassandraVersion field if non-nil, zero value otherwise.

### GetCassandraVersionOk

`func (o *CassandraAdditionalParams) GetCassandraVersionOk() (*string, bool)`

GetCassandraVersionOk returns a tuple with the CassandraVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraVersion

`func (o *CassandraAdditionalParams) SetCassandraVersion(v string)`

SetCassandraVersion sets CassandraVersion field to given value.

### HasCassandraVersion

`func (o *CassandraAdditionalParams) HasCassandraVersion() bool`

HasCassandraVersion returns a boolean if a field has been set.

### SetCassandraVersionNil

`func (o *CassandraAdditionalParams) SetCassandraVersionNil(b bool)`

 SetCassandraVersionNil sets the value for CassandraVersion to be an explicit nil

### UnsetCassandraVersion
`func (o *CassandraAdditionalParams) UnsetCassandraVersion()`

UnsetCassandraVersion ensures that no value is present for CassandraVersion, not even an explicit nil
### GetDataCenterVec

`func (o *CassandraAdditionalParams) GetDataCenterVec() []string`

GetDataCenterVec returns the DataCenterVec field if non-nil, zero value otherwise.

### GetDataCenterVecOk

`func (o *CassandraAdditionalParams) GetDataCenterVecOk() (*[]string, bool)`

GetDataCenterVecOk returns a tuple with the DataCenterVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterVec

`func (o *CassandraAdditionalParams) SetDataCenterVec(v []string)`

SetDataCenterVec sets DataCenterVec field to given value.

### HasDataCenterVec

`func (o *CassandraAdditionalParams) HasDataCenterVec() bool`

HasDataCenterVec returns a boolean if a field has been set.

### SetDataCenterVecNil

`func (o *CassandraAdditionalParams) SetDataCenterVecNil(b bool)`

 SetDataCenterVecNil sets the value for DataCenterVec to be an explicit nil

### UnsetDataCenterVec
`func (o *CassandraAdditionalParams) UnsetDataCenterVec()`

UnsetDataCenterVec ensures that no value is present for DataCenterVec, not even an explicit nil
### GetDseSolrInfo

`func (o *CassandraAdditionalParams) GetDseSolrInfo() DSESolrInfo`

GetDseSolrInfo returns the DseSolrInfo field if non-nil, zero value otherwise.

### GetDseSolrInfoOk

`func (o *CassandraAdditionalParams) GetDseSolrInfoOk() (*DSESolrInfo, bool)`

GetDseSolrInfoOk returns a tuple with the DseSolrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseSolrInfo

`func (o *CassandraAdditionalParams) SetDseSolrInfo(v DSESolrInfo)`

SetDseSolrInfo sets DseSolrInfo field to given value.

### HasDseSolrInfo

`func (o *CassandraAdditionalParams) HasDseSolrInfo() bool`

HasDseSolrInfo returns a boolean if a field has been set.

### GetDseVersion

`func (o *CassandraAdditionalParams) GetDseVersion() string`

GetDseVersion returns the DseVersion field if non-nil, zero value otherwise.

### GetDseVersionOk

`func (o *CassandraAdditionalParams) GetDseVersionOk() (*string, bool)`

GetDseVersionOk returns a tuple with the DseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseVersion

`func (o *CassandraAdditionalParams) SetDseVersion(v string)`

SetDseVersion sets DseVersion field to given value.

### HasDseVersion

`func (o *CassandraAdditionalParams) HasDseVersion() bool`

HasDseVersion returns a boolean if a field has been set.

### SetDseVersionNil

`func (o *CassandraAdditionalParams) SetDseVersionNil(b bool)`

 SetDseVersionNil sets the value for DseVersion to be an explicit nil

### UnsetDseVersion
`func (o *CassandraAdditionalParams) UnsetDseVersion()`

UnsetDseVersion ensures that no value is present for DseVersion, not even an explicit nil
### GetTieredStorageDirsMap

`func (o *CassandraAdditionalParams) GetTieredStorageDirsMap() []NodeToTieredStorageDirectoriesMap`

GetTieredStorageDirsMap returns the TieredStorageDirsMap field if non-nil, zero value otherwise.

### GetTieredStorageDirsMapOk

`func (o *CassandraAdditionalParams) GetTieredStorageDirsMapOk() (*[]NodeToTieredStorageDirectoriesMap, bool)`

GetTieredStorageDirsMapOk returns a tuple with the TieredStorageDirsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredStorageDirsMap

`func (o *CassandraAdditionalParams) SetTieredStorageDirsMap(v []NodeToTieredStorageDirectoriesMap)`

SetTieredStorageDirsMap sets TieredStorageDirsMap field to given value.

### HasTieredStorageDirsMap

`func (o *CassandraAdditionalParams) HasTieredStorageDirsMap() bool`

HasTieredStorageDirsMap returns a boolean if a field has been set.

### SetTieredStorageDirsMapNil

`func (o *CassandraAdditionalParams) SetTieredStorageDirsMapNil(b bool)`

 SetTieredStorageDirsMapNil sets the value for TieredStorageDirsMap to be an explicit nil

### UnsetTieredStorageDirsMap
`func (o *CassandraAdditionalParams) UnsetTieredStorageDirsMap()`

UnsetTieredStorageDirsMap ensures that no value is present for TieredStorageDirsMap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


