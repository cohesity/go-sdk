# DatabaseEntityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerDatabaseInfo** | Pointer to [**ContainerDatabaseInfo**](ContainerDatabaseInfo.md) |  | [optional] 
**DataGuardInfo** | Pointer to [**OracleDataGuardInfo**](OracleDataGuardInfo.md) |  | [optional] 

## Methods

### NewDatabaseEntityInfo

`func NewDatabaseEntityInfo() *DatabaseEntityInfo`

NewDatabaseEntityInfo instantiates a new DatabaseEntityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseEntityInfoWithDefaults

`func NewDatabaseEntityInfoWithDefaults() *DatabaseEntityInfo`

NewDatabaseEntityInfoWithDefaults instantiates a new DatabaseEntityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerDatabaseInfo

`func (o *DatabaseEntityInfo) GetContainerDatabaseInfo() ContainerDatabaseInfo`

GetContainerDatabaseInfo returns the ContainerDatabaseInfo field if non-nil, zero value otherwise.

### GetContainerDatabaseInfoOk

`func (o *DatabaseEntityInfo) GetContainerDatabaseInfoOk() (*ContainerDatabaseInfo, bool)`

GetContainerDatabaseInfoOk returns a tuple with the ContainerDatabaseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDatabaseInfo

`func (o *DatabaseEntityInfo) SetContainerDatabaseInfo(v ContainerDatabaseInfo)`

SetContainerDatabaseInfo sets ContainerDatabaseInfo field to given value.

### HasContainerDatabaseInfo

`func (o *DatabaseEntityInfo) HasContainerDatabaseInfo() bool`

HasContainerDatabaseInfo returns a boolean if a field has been set.

### GetDataGuardInfo

`func (o *DatabaseEntityInfo) GetDataGuardInfo() OracleDataGuardInfo`

GetDataGuardInfo returns the DataGuardInfo field if non-nil, zero value otherwise.

### GetDataGuardInfoOk

`func (o *DatabaseEntityInfo) GetDataGuardInfoOk() (*OracleDataGuardInfo, bool)`

GetDataGuardInfoOk returns a tuple with the DataGuardInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataGuardInfo

`func (o *DatabaseEntityInfo) SetDataGuardInfo(v OracleDataGuardInfo)`

SetDataGuardInfo sets DataGuardInfo field to given value.

### HasDataGuardInfo

`func (o *DatabaseEntityInfo) HasDataGuardInfo() bool`

HasDataGuardInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


