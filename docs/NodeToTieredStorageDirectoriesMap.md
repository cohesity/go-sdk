# NodeToTieredStorageDirectoriesMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraNodeName** | Pointer to **NullableString** | Name of the Cassandra node. | [optional] 
**TieredStorageDirectoriesVec** | Pointer to **[]string** | Array of tiered storage directories. | [optional] 

## Methods

### NewNodeToTieredStorageDirectoriesMap

`func NewNodeToTieredStorageDirectoriesMap() *NodeToTieredStorageDirectoriesMap`

NewNodeToTieredStorageDirectoriesMap instantiates a new NodeToTieredStorageDirectoriesMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeToTieredStorageDirectoriesMapWithDefaults

`func NewNodeToTieredStorageDirectoriesMapWithDefaults() *NodeToTieredStorageDirectoriesMap`

NewNodeToTieredStorageDirectoriesMapWithDefaults instantiates a new NodeToTieredStorageDirectoriesMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraNodeName

`func (o *NodeToTieredStorageDirectoriesMap) GetCassandraNodeName() string`

GetCassandraNodeName returns the CassandraNodeName field if non-nil, zero value otherwise.

### GetCassandraNodeNameOk

`func (o *NodeToTieredStorageDirectoriesMap) GetCassandraNodeNameOk() (*string, bool)`

GetCassandraNodeNameOk returns a tuple with the CassandraNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraNodeName

`func (o *NodeToTieredStorageDirectoriesMap) SetCassandraNodeName(v string)`

SetCassandraNodeName sets CassandraNodeName field to given value.

### HasCassandraNodeName

`func (o *NodeToTieredStorageDirectoriesMap) HasCassandraNodeName() bool`

HasCassandraNodeName returns a boolean if a field has been set.

### SetCassandraNodeNameNil

`func (o *NodeToTieredStorageDirectoriesMap) SetCassandraNodeNameNil(b bool)`

 SetCassandraNodeNameNil sets the value for CassandraNodeName to be an explicit nil

### UnsetCassandraNodeName
`func (o *NodeToTieredStorageDirectoriesMap) UnsetCassandraNodeName()`

UnsetCassandraNodeName ensures that no value is present for CassandraNodeName, not even an explicit nil
### GetTieredStorageDirectoriesVec

`func (o *NodeToTieredStorageDirectoriesMap) GetTieredStorageDirectoriesVec() []string`

GetTieredStorageDirectoriesVec returns the TieredStorageDirectoriesVec field if non-nil, zero value otherwise.

### GetTieredStorageDirectoriesVecOk

`func (o *NodeToTieredStorageDirectoriesMap) GetTieredStorageDirectoriesVecOk() (*[]string, bool)`

GetTieredStorageDirectoriesVecOk returns a tuple with the TieredStorageDirectoriesVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredStorageDirectoriesVec

`func (o *NodeToTieredStorageDirectoriesMap) SetTieredStorageDirectoriesVec(v []string)`

SetTieredStorageDirectoriesVec sets TieredStorageDirectoriesVec field to given value.

### HasTieredStorageDirectoriesVec

`func (o *NodeToTieredStorageDirectoriesMap) HasTieredStorageDirectoriesVec() bool`

HasTieredStorageDirectoriesVec returns a boolean if a field has been set.

### SetTieredStorageDirectoriesVecNil

`func (o *NodeToTieredStorageDirectoriesMap) SetTieredStorageDirectoriesVecNil(b bool)`

 SetTieredStorageDirectoriesVecNil sets the value for TieredStorageDirectoriesVec to be an explicit nil

### UnsetTieredStorageDirectoriesVec
`func (o *NodeToTieredStorageDirectoriesMap) UnsetTieredStorageDirectoriesVec()`

UnsetTieredStorageDirectoriesVec ensures that no value is present for TieredStorageDirectoriesVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


