# ElastifileContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** | Specifies the creation date of the container. | [optional] 
**Id** | Pointer to **NullableInt32** | Specifies id of a Elastifile Container in a Cluster. | [optional] 
**IsNfsInterface** | Pointer to **NullableBool** | Specifies if the container has NFS volumes or not. | [optional] 
**IsSmbInterface** | Pointer to **NullableBool** | Specifies if the container has SMB volumes or not. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the container. | [optional] 
**Protocols** | Pointer to **[]string** | Specifies Elastifile supported Protocol information enabled on Elastifile container. &#39;kNfs&#39; indicates NFS protocol in an elastifile container. &#39;kSmb&#39; indicates SMB protocol in an elastifile container. | [optional] 
**UsedBytes** | Pointer to **NullableInt64** | Specifies the bytes used by the container. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID of the container. | [optional] 

## Methods

### NewElastifileContainer

`func NewElastifileContainer() *ElastifileContainer`

NewElastifileContainer instantiates a new ElastifileContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElastifileContainerWithDefaults

`func NewElastifileContainerWithDefaults() *ElastifileContainer`

NewElastifileContainerWithDefaults instantiates a new ElastifileContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ElastifileContainer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ElastifileContainer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ElastifileContainer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ElastifileContainer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ElastifileContainer) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ElastifileContainer) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetId

`func (o *ElastifileContainer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ElastifileContainer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ElastifileContainer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ElastifileContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ElastifileContainer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ElastifileContainer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsNfsInterface

`func (o *ElastifileContainer) GetIsNfsInterface() bool`

GetIsNfsInterface returns the IsNfsInterface field if non-nil, zero value otherwise.

### GetIsNfsInterfaceOk

`func (o *ElastifileContainer) GetIsNfsInterfaceOk() (*bool, bool)`

GetIsNfsInterfaceOk returns a tuple with the IsNfsInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNfsInterface

`func (o *ElastifileContainer) SetIsNfsInterface(v bool)`

SetIsNfsInterface sets IsNfsInterface field to given value.

### HasIsNfsInterface

`func (o *ElastifileContainer) HasIsNfsInterface() bool`

HasIsNfsInterface returns a boolean if a field has been set.

### SetIsNfsInterfaceNil

`func (o *ElastifileContainer) SetIsNfsInterfaceNil(b bool)`

 SetIsNfsInterfaceNil sets the value for IsNfsInterface to be an explicit nil

### UnsetIsNfsInterface
`func (o *ElastifileContainer) UnsetIsNfsInterface()`

UnsetIsNfsInterface ensures that no value is present for IsNfsInterface, not even an explicit nil
### GetIsSmbInterface

`func (o *ElastifileContainer) GetIsSmbInterface() bool`

GetIsSmbInterface returns the IsSmbInterface field if non-nil, zero value otherwise.

### GetIsSmbInterfaceOk

`func (o *ElastifileContainer) GetIsSmbInterfaceOk() (*bool, bool)`

GetIsSmbInterfaceOk returns a tuple with the IsSmbInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSmbInterface

`func (o *ElastifileContainer) SetIsSmbInterface(v bool)`

SetIsSmbInterface sets IsSmbInterface field to given value.

### HasIsSmbInterface

`func (o *ElastifileContainer) HasIsSmbInterface() bool`

HasIsSmbInterface returns a boolean if a field has been set.

### SetIsSmbInterfaceNil

`func (o *ElastifileContainer) SetIsSmbInterfaceNil(b bool)`

 SetIsSmbInterfaceNil sets the value for IsSmbInterface to be an explicit nil

### UnsetIsSmbInterface
`func (o *ElastifileContainer) UnsetIsSmbInterface()`

UnsetIsSmbInterface ensures that no value is present for IsSmbInterface, not even an explicit nil
### GetName

`func (o *ElastifileContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ElastifileContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ElastifileContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ElastifileContainer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ElastifileContainer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ElastifileContainer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtocols

`func (o *ElastifileContainer) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ElastifileContainer) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ElastifileContainer) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *ElastifileContainer) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### SetProtocolsNil

`func (o *ElastifileContainer) SetProtocolsNil(b bool)`

 SetProtocolsNil sets the value for Protocols to be an explicit nil

### UnsetProtocols
`func (o *ElastifileContainer) UnsetProtocols()`

UnsetProtocols ensures that no value is present for Protocols, not even an explicit nil
### GetUsedBytes

`func (o *ElastifileContainer) GetUsedBytes() int64`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *ElastifileContainer) GetUsedBytesOk() (*int64, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *ElastifileContainer) SetUsedBytes(v int64)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *ElastifileContainer) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### SetUsedBytesNil

`func (o *ElastifileContainer) SetUsedBytesNil(b bool)`

 SetUsedBytesNil sets the value for UsedBytes to be an explicit nil

### UnsetUsedBytes
`func (o *ElastifileContainer) UnsetUsedBytes()`

UnsetUsedBytes ensures that no value is present for UsedBytes, not even an explicit nil
### GetUuid

`func (o *ElastifileContainer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ElastifileContainer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ElastifileContainer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ElastifileContainer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ElastifileContainer) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ElastifileContainer) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


