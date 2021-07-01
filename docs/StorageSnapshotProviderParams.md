# StorageSnapshotProviderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorParams** | Pointer to [**ConnectorParams**](ConnectorParams.md) |  | [optional] 
**Entity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewStorageSnapshotProviderParams

`func NewStorageSnapshotProviderParams() *StorageSnapshotProviderParams`

NewStorageSnapshotProviderParams instantiates a new StorageSnapshotProviderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSnapshotProviderParamsWithDefaults

`func NewStorageSnapshotProviderParamsWithDefaults() *StorageSnapshotProviderParams`

NewStorageSnapshotProviderParamsWithDefaults instantiates a new StorageSnapshotProviderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorParams

`func (o *StorageSnapshotProviderParams) GetConnectorParams() ConnectorParams`

GetConnectorParams returns the ConnectorParams field if non-nil, zero value otherwise.

### GetConnectorParamsOk

`func (o *StorageSnapshotProviderParams) GetConnectorParamsOk() (*ConnectorParams, bool)`

GetConnectorParamsOk returns a tuple with the ConnectorParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorParams

`func (o *StorageSnapshotProviderParams) SetConnectorParams(v ConnectorParams)`

SetConnectorParams sets ConnectorParams field to given value.

### HasConnectorParams

`func (o *StorageSnapshotProviderParams) HasConnectorParams() bool`

HasConnectorParams returns a boolean if a field has been set.

### GetEntity

`func (o *StorageSnapshotProviderParams) GetEntity() EntityProto`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *StorageSnapshotProviderParams) GetEntityOk() (*EntityProto, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *StorageSnapshotProviderParams) SetEntity(v EntityProto)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *StorageSnapshotProviderParams) HasEntity() bool`

HasEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


