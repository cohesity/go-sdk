# ContainerDatabaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluggableDatabaseList** | Pointer to [**[]PluggableDatabaseInfo**](PluggableDatabaseInfo.md) | Specifies the list of Pluggable databases within a container database. | [optional] 

## Methods

### NewContainerDatabaseInfo

`func NewContainerDatabaseInfo() *ContainerDatabaseInfo`

NewContainerDatabaseInfo instantiates a new ContainerDatabaseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerDatabaseInfoWithDefaults

`func NewContainerDatabaseInfoWithDefaults() *ContainerDatabaseInfo`

NewContainerDatabaseInfoWithDefaults instantiates a new ContainerDatabaseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluggableDatabaseList

`func (o *ContainerDatabaseInfo) GetPluggableDatabaseList() []PluggableDatabaseInfo`

GetPluggableDatabaseList returns the PluggableDatabaseList field if non-nil, zero value otherwise.

### GetPluggableDatabaseListOk

`func (o *ContainerDatabaseInfo) GetPluggableDatabaseListOk() (*[]PluggableDatabaseInfo, bool)`

GetPluggableDatabaseListOk returns a tuple with the PluggableDatabaseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluggableDatabaseList

`func (o *ContainerDatabaseInfo) SetPluggableDatabaseList(v []PluggableDatabaseInfo)`

SetPluggableDatabaseList sets PluggableDatabaseList field to given value.

### HasPluggableDatabaseList

`func (o *ContainerDatabaseInfo) HasPluggableDatabaseList() bool`

HasPluggableDatabaseList returns a boolean if a field has been set.

### SetPluggableDatabaseListNil

`func (o *ContainerDatabaseInfo) SetPluggableDatabaseListNil(b bool)`

 SetPluggableDatabaseListNil sets the value for PluggableDatabaseList to be an explicit nil

### UnsetPluggableDatabaseList
`func (o *ContainerDatabaseInfo) UnsetPluggableDatabaseList()`

UnsetPluggableDatabaseList ensures that no value is present for PluggableDatabaseList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


