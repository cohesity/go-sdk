# SQLServerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the unique id of the SQL server instance. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the SQL server instance. | [optional] 
**IsOnline** | Pointer to **NullableString** | Specifies the wehther the SQL server instance is online or not. | [optional] 
**Endpoints** | Pointer to [**[]ResourceEndpoint**](ResourceEndpoint.md) | Specifies the information about endpoint associated with this SQL server instance. | [optional] 
**IsPartofFCI** | Pointer to **NullableBool** | Specifies whether this SQL server instance is a part of Failover cluster or not. | [optional] 

## Methods

### NewSQLServerInstance

`func NewSQLServerInstance() *SQLServerInstance`

NewSQLServerInstance instantiates a new SQLServerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSQLServerInstanceWithDefaults

`func NewSQLServerInstanceWithDefaults() *SQLServerInstance`

NewSQLServerInstanceWithDefaults instantiates a new SQLServerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SQLServerInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SQLServerInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SQLServerInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SQLServerInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SQLServerInstance) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SQLServerInstance) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SQLServerInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SQLServerInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SQLServerInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SQLServerInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SQLServerInstance) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SQLServerInstance) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsOnline

`func (o *SQLServerInstance) GetIsOnline() string`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *SQLServerInstance) GetIsOnlineOk() (*string, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *SQLServerInstance) SetIsOnline(v string)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *SQLServerInstance) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### SetIsOnlineNil

`func (o *SQLServerInstance) SetIsOnlineNil(b bool)`

 SetIsOnlineNil sets the value for IsOnline to be an explicit nil

### UnsetIsOnline
`func (o *SQLServerInstance) UnsetIsOnline()`

UnsetIsOnline ensures that no value is present for IsOnline, not even an explicit nil
### GetEndpoints

`func (o *SQLServerInstance) GetEndpoints() []ResourceEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *SQLServerInstance) GetEndpointsOk() (*[]ResourceEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *SQLServerInstance) SetEndpoints(v []ResourceEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *SQLServerInstance) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpointsNil

`func (o *SQLServerInstance) SetEndpointsNil(b bool)`

 SetEndpointsNil sets the value for Endpoints to be an explicit nil

### UnsetEndpoints
`func (o *SQLServerInstance) UnsetEndpoints()`

UnsetEndpoints ensures that no value is present for Endpoints, not even an explicit nil
### GetIsPartofFCI

`func (o *SQLServerInstance) GetIsPartofFCI() bool`

GetIsPartofFCI returns the IsPartofFCI field if non-nil, zero value otherwise.

### GetIsPartofFCIOk

`func (o *SQLServerInstance) GetIsPartofFCIOk() (*bool, bool)`

GetIsPartofFCIOk returns a tuple with the IsPartofFCI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartofFCI

`func (o *SQLServerInstance) SetIsPartofFCI(v bool)`

SetIsPartofFCI sets IsPartofFCI field to given value.

### HasIsPartofFCI

`func (o *SQLServerInstance) HasIsPartofFCI() bool`

HasIsPartofFCI returns a boolean if a field has been set.

### SetIsPartofFCINil

`func (o *SQLServerInstance) SetIsPartofFCINil(b bool)`

 SetIsPartofFCINil sets the value for IsPartofFCI to be an explicit nil

### UnsetIsPartofFCI
`func (o *SQLServerInstance) UnsetIsPartofFCI()`

UnsetIsPartofFCI ensures that no value is present for IsPartofFCI, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


