# ExchangeProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DagInfo** | Pointer to [**DagInfo**](DagInfo.md) |  | [optional] 
**ApplicationServerInfo** | Pointer to [**ApplicationServerInfo**](ApplicationServerInfo.md) |  | [optional] 
**DagDatabaseCopyInfo** | Pointer to [**ExchangeDatabaseCopyInfo**](ExchangeDatabaseCopyInfo.md) |  | [optional] 
**DagDatabaseInfo** | Pointer to [**ExchangeDAGDatabase**](ExchangeDAGDatabase.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OwnerId** | Pointer to **NullableInt64** | Specifies the entity id of the owner of the Exchange Protection Source. | [optional] 
**StandaloneDatabaseCopyInfo** | Pointer to [**ExchangeDatabaseInfo**](ExchangeDatabaseInfo.md) |  | [optional] 
**Type** | Pointer to **NullableInt32** | Specifies the type of the Exchange Protection Source. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the Exchange entity. | [optional] 

## Methods

### NewExchangeProtectionSource

`func NewExchangeProtectionSource() *ExchangeProtectionSource`

NewExchangeProtectionSource instantiates a new ExchangeProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeProtectionSourceWithDefaults

`func NewExchangeProtectionSourceWithDefaults() *ExchangeProtectionSource`

NewExchangeProtectionSourceWithDefaults instantiates a new ExchangeProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDagInfo

`func (o *ExchangeProtectionSource) GetDagInfo() DagInfo`

GetDagInfo returns the DagInfo field if non-nil, zero value otherwise.

### GetDagInfoOk

`func (o *ExchangeProtectionSource) GetDagInfoOk() (*DagInfo, bool)`

GetDagInfoOk returns a tuple with the DagInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagInfo

`func (o *ExchangeProtectionSource) SetDagInfo(v DagInfo)`

SetDagInfo sets DagInfo field to given value.

### HasDagInfo

`func (o *ExchangeProtectionSource) HasDagInfo() bool`

HasDagInfo returns a boolean if a field has been set.

### GetApplicationServerInfo

`func (o *ExchangeProtectionSource) GetApplicationServerInfo() ApplicationServerInfo`

GetApplicationServerInfo returns the ApplicationServerInfo field if non-nil, zero value otherwise.

### GetApplicationServerInfoOk

`func (o *ExchangeProtectionSource) GetApplicationServerInfoOk() (*ApplicationServerInfo, bool)`

GetApplicationServerInfoOk returns a tuple with the ApplicationServerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationServerInfo

`func (o *ExchangeProtectionSource) SetApplicationServerInfo(v ApplicationServerInfo)`

SetApplicationServerInfo sets ApplicationServerInfo field to given value.

### HasApplicationServerInfo

`func (o *ExchangeProtectionSource) HasApplicationServerInfo() bool`

HasApplicationServerInfo returns a boolean if a field has been set.

### GetDagDatabaseCopyInfo

`func (o *ExchangeProtectionSource) GetDagDatabaseCopyInfo() ExchangeDatabaseCopyInfo`

GetDagDatabaseCopyInfo returns the DagDatabaseCopyInfo field if non-nil, zero value otherwise.

### GetDagDatabaseCopyInfoOk

`func (o *ExchangeProtectionSource) GetDagDatabaseCopyInfoOk() (*ExchangeDatabaseCopyInfo, bool)`

GetDagDatabaseCopyInfoOk returns a tuple with the DagDatabaseCopyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagDatabaseCopyInfo

`func (o *ExchangeProtectionSource) SetDagDatabaseCopyInfo(v ExchangeDatabaseCopyInfo)`

SetDagDatabaseCopyInfo sets DagDatabaseCopyInfo field to given value.

### HasDagDatabaseCopyInfo

`func (o *ExchangeProtectionSource) HasDagDatabaseCopyInfo() bool`

HasDagDatabaseCopyInfo returns a boolean if a field has been set.

### GetDagDatabaseInfo

`func (o *ExchangeProtectionSource) GetDagDatabaseInfo() ExchangeDAGDatabase`

GetDagDatabaseInfo returns the DagDatabaseInfo field if non-nil, zero value otherwise.

### GetDagDatabaseInfoOk

`func (o *ExchangeProtectionSource) GetDagDatabaseInfoOk() (*ExchangeDAGDatabase, bool)`

GetDagDatabaseInfoOk returns a tuple with the DagDatabaseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDagDatabaseInfo

`func (o *ExchangeProtectionSource) SetDagDatabaseInfo(v ExchangeDAGDatabase)`

SetDagDatabaseInfo sets DagDatabaseInfo field to given value.

### HasDagDatabaseInfo

`func (o *ExchangeProtectionSource) HasDagDatabaseInfo() bool`

HasDagDatabaseInfo returns a boolean if a field has been set.

### GetName

`func (o *ExchangeProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExchangeProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExchangeProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *ExchangeProtectionSource) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ExchangeProtectionSource) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ExchangeProtectionSource) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ExchangeProtectionSource) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *ExchangeProtectionSource) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *ExchangeProtectionSource) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetStandaloneDatabaseCopyInfo

`func (o *ExchangeProtectionSource) GetStandaloneDatabaseCopyInfo() ExchangeDatabaseInfo`

GetStandaloneDatabaseCopyInfo returns the StandaloneDatabaseCopyInfo field if non-nil, zero value otherwise.

### GetStandaloneDatabaseCopyInfoOk

`func (o *ExchangeProtectionSource) GetStandaloneDatabaseCopyInfoOk() (*ExchangeDatabaseInfo, bool)`

GetStandaloneDatabaseCopyInfoOk returns a tuple with the StandaloneDatabaseCopyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneDatabaseCopyInfo

`func (o *ExchangeProtectionSource) SetStandaloneDatabaseCopyInfo(v ExchangeDatabaseInfo)`

SetStandaloneDatabaseCopyInfo sets StandaloneDatabaseCopyInfo field to given value.

### HasStandaloneDatabaseCopyInfo

`func (o *ExchangeProtectionSource) HasStandaloneDatabaseCopyInfo() bool`

HasStandaloneDatabaseCopyInfo returns a boolean if a field has been set.

### GetType

`func (o *ExchangeProtectionSource) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangeProtectionSource) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangeProtectionSource) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ExchangeProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ExchangeProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExchangeProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *ExchangeProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExchangeProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExchangeProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExchangeProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ExchangeProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ExchangeProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


