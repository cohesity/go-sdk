# ApplicationRestoreObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdRestoreParameters** | Pointer to [**AdRestoreParameters**](AdRestoreParameters.md) |  | [optional] 
**ApplicationServerId** | Pointer to **NullableInt64** | Specifies the Application Server to restore (for example, kSQL). | [optional] 
**ExchangeRestoreParameters** | Pointer to [**ExchangeRestoreParameters**](ExchangeRestoreParameters.md) |  | [optional] 
**SqlRestoreParameters** | Pointer to [**SqlRestoreParameters**](SqlRestoreParameters.md) |  | [optional] 
**TargetHostId** | Pointer to **NullableInt64** | Specifies the target host if the application is to be restored to a different host. If this is empty, then the application is restored to the original host, which is the hosting Protection Source. | [optional] 
**TargetRootNodeId** | Pointer to **NullableInt64** | Specifies the registered root node, like vCenter, of targetHost. If this is empty, then it is assumed the root node of the target host is the same as the host Protection Source of the application. | [optional] 

## Methods

### NewApplicationRestoreObject

`func NewApplicationRestoreObject() *ApplicationRestoreObject`

NewApplicationRestoreObject instantiates a new ApplicationRestoreObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRestoreObjectWithDefaults

`func NewApplicationRestoreObjectWithDefaults() *ApplicationRestoreObject`

NewApplicationRestoreObjectWithDefaults instantiates a new ApplicationRestoreObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdRestoreParameters

`func (o *ApplicationRestoreObject) GetAdRestoreParameters() AdRestoreParameters`

GetAdRestoreParameters returns the AdRestoreParameters field if non-nil, zero value otherwise.

### GetAdRestoreParametersOk

`func (o *ApplicationRestoreObject) GetAdRestoreParametersOk() (*AdRestoreParameters, bool)`

GetAdRestoreParametersOk returns a tuple with the AdRestoreParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRestoreParameters

`func (o *ApplicationRestoreObject) SetAdRestoreParameters(v AdRestoreParameters)`

SetAdRestoreParameters sets AdRestoreParameters field to given value.

### HasAdRestoreParameters

`func (o *ApplicationRestoreObject) HasAdRestoreParameters() bool`

HasAdRestoreParameters returns a boolean if a field has been set.

### GetApplicationServerId

`func (o *ApplicationRestoreObject) GetApplicationServerId() int64`

GetApplicationServerId returns the ApplicationServerId field if non-nil, zero value otherwise.

### GetApplicationServerIdOk

`func (o *ApplicationRestoreObject) GetApplicationServerIdOk() (*int64, bool)`

GetApplicationServerIdOk returns a tuple with the ApplicationServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationServerId

`func (o *ApplicationRestoreObject) SetApplicationServerId(v int64)`

SetApplicationServerId sets ApplicationServerId field to given value.

### HasApplicationServerId

`func (o *ApplicationRestoreObject) HasApplicationServerId() bool`

HasApplicationServerId returns a boolean if a field has been set.

### SetApplicationServerIdNil

`func (o *ApplicationRestoreObject) SetApplicationServerIdNil(b bool)`

 SetApplicationServerIdNil sets the value for ApplicationServerId to be an explicit nil

### UnsetApplicationServerId
`func (o *ApplicationRestoreObject) UnsetApplicationServerId()`

UnsetApplicationServerId ensures that no value is present for ApplicationServerId, not even an explicit nil
### GetExchangeRestoreParameters

`func (o *ApplicationRestoreObject) GetExchangeRestoreParameters() ExchangeRestoreParameters`

GetExchangeRestoreParameters returns the ExchangeRestoreParameters field if non-nil, zero value otherwise.

### GetExchangeRestoreParametersOk

`func (o *ApplicationRestoreObject) GetExchangeRestoreParametersOk() (*ExchangeRestoreParameters, bool)`

GetExchangeRestoreParametersOk returns a tuple with the ExchangeRestoreParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRestoreParameters

`func (o *ApplicationRestoreObject) SetExchangeRestoreParameters(v ExchangeRestoreParameters)`

SetExchangeRestoreParameters sets ExchangeRestoreParameters field to given value.

### HasExchangeRestoreParameters

`func (o *ApplicationRestoreObject) HasExchangeRestoreParameters() bool`

HasExchangeRestoreParameters returns a boolean if a field has been set.

### GetSqlRestoreParameters

`func (o *ApplicationRestoreObject) GetSqlRestoreParameters() SqlRestoreParameters`

GetSqlRestoreParameters returns the SqlRestoreParameters field if non-nil, zero value otherwise.

### GetSqlRestoreParametersOk

`func (o *ApplicationRestoreObject) GetSqlRestoreParametersOk() (*SqlRestoreParameters, bool)`

GetSqlRestoreParametersOk returns a tuple with the SqlRestoreParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlRestoreParameters

`func (o *ApplicationRestoreObject) SetSqlRestoreParameters(v SqlRestoreParameters)`

SetSqlRestoreParameters sets SqlRestoreParameters field to given value.

### HasSqlRestoreParameters

`func (o *ApplicationRestoreObject) HasSqlRestoreParameters() bool`

HasSqlRestoreParameters returns a boolean if a field has been set.

### GetTargetHostId

`func (o *ApplicationRestoreObject) GetTargetHostId() int64`

GetTargetHostId returns the TargetHostId field if non-nil, zero value otherwise.

### GetTargetHostIdOk

`func (o *ApplicationRestoreObject) GetTargetHostIdOk() (*int64, bool)`

GetTargetHostIdOk returns a tuple with the TargetHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHostId

`func (o *ApplicationRestoreObject) SetTargetHostId(v int64)`

SetTargetHostId sets TargetHostId field to given value.

### HasTargetHostId

`func (o *ApplicationRestoreObject) HasTargetHostId() bool`

HasTargetHostId returns a boolean if a field has been set.

### SetTargetHostIdNil

`func (o *ApplicationRestoreObject) SetTargetHostIdNil(b bool)`

 SetTargetHostIdNil sets the value for TargetHostId to be an explicit nil

### UnsetTargetHostId
`func (o *ApplicationRestoreObject) UnsetTargetHostId()`

UnsetTargetHostId ensures that no value is present for TargetHostId, not even an explicit nil
### GetTargetRootNodeId

`func (o *ApplicationRestoreObject) GetTargetRootNodeId() int64`

GetTargetRootNodeId returns the TargetRootNodeId field if non-nil, zero value otherwise.

### GetTargetRootNodeIdOk

`func (o *ApplicationRestoreObject) GetTargetRootNodeIdOk() (*int64, bool)`

GetTargetRootNodeIdOk returns a tuple with the TargetRootNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRootNodeId

`func (o *ApplicationRestoreObject) SetTargetRootNodeId(v int64)`

SetTargetRootNodeId sets TargetRootNodeId field to given value.

### HasTargetRootNodeId

`func (o *ApplicationRestoreObject) HasTargetRootNodeId() bool`

HasTargetRootNodeId returns a boolean if a field has been set.

### SetTargetRootNodeIdNil

`func (o *ApplicationRestoreObject) SetTargetRootNodeIdNil(b bool)`

 SetTargetRootNodeIdNil sets the value for TargetRootNodeId to be an explicit nil

### UnsetTargetRootNodeId
`func (o *ApplicationRestoreObject) UnsetTargetRootNodeId()`

UnsetTargetRootNodeId ensures that no value is present for TargetRootNodeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


