# SqlAagHostAndDatabases

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AagDatabases** | Pointer to [**[]AagAndDatabases**](AagAndDatabases.md) | Specifies a list of AAGs and database members in each AAG. | [optional] 
**ApplicationNode** | Pointer to [**ProtectionSourceNode**](ProtectionSourceNode.md) |  | [optional] 
**Databases** | Pointer to [**[]ProtectionSource**](ProtectionSource.md) | Specifies all database entities found on the server. Database may or may not be in an AAG. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies an error message when the host is not registered as an SQL host. | [optional] 
**UnknownHostName** | Pointer to **NullableString** | Specifies the name of the host that is not registered as an SQL server on Cohesity Cluser. | [optional] 

## Methods

### NewSqlAagHostAndDatabases

`func NewSqlAagHostAndDatabases() *SqlAagHostAndDatabases`

NewSqlAagHostAndDatabases instantiates a new SqlAagHostAndDatabases object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAagHostAndDatabasesWithDefaults

`func NewSqlAagHostAndDatabasesWithDefaults() *SqlAagHostAndDatabases`

NewSqlAagHostAndDatabasesWithDefaults instantiates a new SqlAagHostAndDatabases object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAagDatabases

`func (o *SqlAagHostAndDatabases) GetAagDatabases() []AagAndDatabases`

GetAagDatabases returns the AagDatabases field if non-nil, zero value otherwise.

### GetAagDatabasesOk

`func (o *SqlAagHostAndDatabases) GetAagDatabasesOk() (*[]AagAndDatabases, bool)`

GetAagDatabasesOk returns a tuple with the AagDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAagDatabases

`func (o *SqlAagHostAndDatabases) SetAagDatabases(v []AagAndDatabases)`

SetAagDatabases sets AagDatabases field to given value.

### HasAagDatabases

`func (o *SqlAagHostAndDatabases) HasAagDatabases() bool`

HasAagDatabases returns a boolean if a field has been set.

### SetAagDatabasesNil

`func (o *SqlAagHostAndDatabases) SetAagDatabasesNil(b bool)`

 SetAagDatabasesNil sets the value for AagDatabases to be an explicit nil

### UnsetAagDatabases
`func (o *SqlAagHostAndDatabases) UnsetAagDatabases()`

UnsetAagDatabases ensures that no value is present for AagDatabases, not even an explicit nil
### GetApplicationNode

`func (o *SqlAagHostAndDatabases) GetApplicationNode() ProtectionSourceNode`

GetApplicationNode returns the ApplicationNode field if non-nil, zero value otherwise.

### GetApplicationNodeOk

`func (o *SqlAagHostAndDatabases) GetApplicationNodeOk() (*ProtectionSourceNode, bool)`

GetApplicationNodeOk returns a tuple with the ApplicationNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNode

`func (o *SqlAagHostAndDatabases) SetApplicationNode(v ProtectionSourceNode)`

SetApplicationNode sets ApplicationNode field to given value.

### HasApplicationNode

`func (o *SqlAagHostAndDatabases) HasApplicationNode() bool`

HasApplicationNode returns a boolean if a field has been set.

### GetDatabases

`func (o *SqlAagHostAndDatabases) GetDatabases() []ProtectionSource`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *SqlAagHostAndDatabases) GetDatabasesOk() (*[]ProtectionSource, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *SqlAagHostAndDatabases) SetDatabases(v []ProtectionSource)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *SqlAagHostAndDatabases) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### SetDatabasesNil

`func (o *SqlAagHostAndDatabases) SetDatabasesNil(b bool)`

 SetDatabasesNil sets the value for Databases to be an explicit nil

### UnsetDatabases
`func (o *SqlAagHostAndDatabases) UnsetDatabases()`

UnsetDatabases ensures that no value is present for Databases, not even an explicit nil
### GetErrorMessage

`func (o *SqlAagHostAndDatabases) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SqlAagHostAndDatabases) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SqlAagHostAndDatabases) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SqlAagHostAndDatabases) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *SqlAagHostAndDatabases) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *SqlAagHostAndDatabases) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetUnknownHostName

`func (o *SqlAagHostAndDatabases) GetUnknownHostName() string`

GetUnknownHostName returns the UnknownHostName field if non-nil, zero value otherwise.

### GetUnknownHostNameOk

`func (o *SqlAagHostAndDatabases) GetUnknownHostNameOk() (*string, bool)`

GetUnknownHostNameOk returns a tuple with the UnknownHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownHostName

`func (o *SqlAagHostAndDatabases) SetUnknownHostName(v string)`

SetUnknownHostName sets UnknownHostName field to given value.

### HasUnknownHostName

`func (o *SqlAagHostAndDatabases) HasUnknownHostName() bool`

HasUnknownHostName returns a boolean if a field has been set.

### SetUnknownHostNameNil

`func (o *SqlAagHostAndDatabases) SetUnknownHostNameNil(b bool)`

 SetUnknownHostNameNil sets the value for UnknownHostName to be an explicit nil

### UnsetUnknownHostName
`func (o *SqlAagHostAndDatabases) UnsetUnknownHostName()`

UnsetUnknownHostName ensures that no value is present for UnknownHostName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


