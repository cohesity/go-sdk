# BackupSourceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEntityIdVec** | Pointer to **[]int64** | If we are backing up an application (such as SQL), this contains the entity ids of the app entities (such as SQL instances and databases) that will be protected on the backup source.  If this vector is empty, it implies that we are protecting all app entities on the source. | [optional] 
**OracleParams** | Pointer to [**OracleSourceParams**](OracleSourceParams.md) |  | [optional] 
**PhysicalParams** | Pointer to [**PhysicalBackupSourceParams**](PhysicalBackupSourceParams.md) |  | [optional] 
**SkipIndexing** | Pointer to **NullableBool** | Set to true, if indexing is not required for given source. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Source entity id. NOTE: This is expected to point to a leaf-level entity. | [optional] 
**VmwareParams** | Pointer to [**VMwareBackupSourceParams**](VMwareBackupSourceParams.md) |  | [optional] 

## Methods

### NewBackupSourceParams

`func NewBackupSourceParams() *BackupSourceParams`

NewBackupSourceParams instantiates a new BackupSourceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSourceParamsWithDefaults

`func NewBackupSourceParamsWithDefaults() *BackupSourceParams`

NewBackupSourceParamsWithDefaults instantiates a new BackupSourceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEntityIdVec

`func (o *BackupSourceParams) GetAppEntityIdVec() []int64`

GetAppEntityIdVec returns the AppEntityIdVec field if non-nil, zero value otherwise.

### GetAppEntityIdVecOk

`func (o *BackupSourceParams) GetAppEntityIdVecOk() (*[]int64, bool)`

GetAppEntityIdVecOk returns a tuple with the AppEntityIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEntityIdVec

`func (o *BackupSourceParams) SetAppEntityIdVec(v []int64)`

SetAppEntityIdVec sets AppEntityIdVec field to given value.

### HasAppEntityIdVec

`func (o *BackupSourceParams) HasAppEntityIdVec() bool`

HasAppEntityIdVec returns a boolean if a field has been set.

### SetAppEntityIdVecNil

`func (o *BackupSourceParams) SetAppEntityIdVecNil(b bool)`

 SetAppEntityIdVecNil sets the value for AppEntityIdVec to be an explicit nil

### UnsetAppEntityIdVec
`func (o *BackupSourceParams) UnsetAppEntityIdVec()`

UnsetAppEntityIdVec ensures that no value is present for AppEntityIdVec, not even an explicit nil
### GetOracleParams

`func (o *BackupSourceParams) GetOracleParams() OracleSourceParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *BackupSourceParams) GetOracleParamsOk() (*OracleSourceParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *BackupSourceParams) SetOracleParams(v OracleSourceParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *BackupSourceParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetPhysicalParams

`func (o *BackupSourceParams) GetPhysicalParams() PhysicalBackupSourceParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *BackupSourceParams) GetPhysicalParamsOk() (*PhysicalBackupSourceParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *BackupSourceParams) SetPhysicalParams(v PhysicalBackupSourceParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *BackupSourceParams) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSkipIndexing

`func (o *BackupSourceParams) GetSkipIndexing() bool`

GetSkipIndexing returns the SkipIndexing field if non-nil, zero value otherwise.

### GetSkipIndexingOk

`func (o *BackupSourceParams) GetSkipIndexingOk() (*bool, bool)`

GetSkipIndexingOk returns a tuple with the SkipIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIndexing

`func (o *BackupSourceParams) SetSkipIndexing(v bool)`

SetSkipIndexing sets SkipIndexing field to given value.

### HasSkipIndexing

`func (o *BackupSourceParams) HasSkipIndexing() bool`

HasSkipIndexing returns a boolean if a field has been set.

### SetSkipIndexingNil

`func (o *BackupSourceParams) SetSkipIndexingNil(b bool)`

 SetSkipIndexingNil sets the value for SkipIndexing to be an explicit nil

### UnsetSkipIndexing
`func (o *BackupSourceParams) UnsetSkipIndexing()`

UnsetSkipIndexing ensures that no value is present for SkipIndexing, not even an explicit nil
### GetSourceId

`func (o *BackupSourceParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BackupSourceParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BackupSourceParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BackupSourceParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *BackupSourceParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *BackupSourceParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetVmwareParams

`func (o *BackupSourceParams) GetVmwareParams() VMwareBackupSourceParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *BackupSourceParams) GetVmwareParamsOk() (*VMwareBackupSourceParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *BackupSourceParams) SetVmwareParams(v VMwareBackupSourceParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *BackupSourceParams) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


