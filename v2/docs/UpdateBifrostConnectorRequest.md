# UpdateBifrostConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CohesitySideIp** | Pointer to **NullableString** | Specifies the cohesity side ip of the connector | [optional] [readonly] 
**ConnectionId** | **NullableInt64** | Specifies the Id of the connection which this connector belongs to. | 
**ConnectionStatus** | Pointer to [**ConnectorConnectionInfo**](ConnectorConnectionInfo.md) |  | [optional] 
**HyxVersion** | Pointer to **NullableString** | Specifies the connector&#39;s software Version | [optional] [readonly] 
**Id** | **NullableString** | Specifies the id of the connector. | 
**Name** | Pointer to **NullableString** | Specifies the name of the connector. | [optional] 
**TenantSourceSideIp** | Pointer to **NullableString** | Specifies the tenant source side ip of the connector | [optional] [readonly] 

## Methods

### NewUpdateBifrostConnectorRequest

`func NewUpdateBifrostConnectorRequest(connectionId NullableInt64, id NullableString, ) *UpdateBifrostConnectorRequest`

NewUpdateBifrostConnectorRequest instantiates a new UpdateBifrostConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBifrostConnectorRequestWithDefaults

`func NewUpdateBifrostConnectorRequestWithDefaults() *UpdateBifrostConnectorRequest`

NewUpdateBifrostConnectorRequestWithDefaults instantiates a new UpdateBifrostConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCohesitySideIp

`func (o *UpdateBifrostConnectorRequest) GetCohesitySideIp() string`

GetCohesitySideIp returns the CohesitySideIp field if non-nil, zero value otherwise.

### GetCohesitySideIpOk

`func (o *UpdateBifrostConnectorRequest) GetCohesitySideIpOk() (*string, bool)`

GetCohesitySideIpOk returns a tuple with the CohesitySideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesitySideIp

`func (o *UpdateBifrostConnectorRequest) SetCohesitySideIp(v string)`

SetCohesitySideIp sets CohesitySideIp field to given value.

### HasCohesitySideIp

`func (o *UpdateBifrostConnectorRequest) HasCohesitySideIp() bool`

HasCohesitySideIp returns a boolean if a field has been set.

### SetCohesitySideIpNil

`func (o *UpdateBifrostConnectorRequest) SetCohesitySideIpNil(b bool)`

 SetCohesitySideIpNil sets the value for CohesitySideIp to be an explicit nil

### UnsetCohesitySideIp
`func (o *UpdateBifrostConnectorRequest) UnsetCohesitySideIp()`

UnsetCohesitySideIp ensures that no value is present for CohesitySideIp, not even an explicit nil
### GetConnectionId

`func (o *UpdateBifrostConnectorRequest) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *UpdateBifrostConnectorRequest) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *UpdateBifrostConnectorRequest) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.


### SetConnectionIdNil

`func (o *UpdateBifrostConnectorRequest) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *UpdateBifrostConnectorRequest) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnectionStatus

`func (o *UpdateBifrostConnectorRequest) GetConnectionStatus() ConnectorConnectionInfo`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *UpdateBifrostConnectorRequest) GetConnectionStatusOk() (*ConnectorConnectionInfo, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *UpdateBifrostConnectorRequest) SetConnectionStatus(v ConnectorConnectionInfo)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *UpdateBifrostConnectorRequest) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetHyxVersion

`func (o *UpdateBifrostConnectorRequest) GetHyxVersion() string`

GetHyxVersion returns the HyxVersion field if non-nil, zero value otherwise.

### GetHyxVersionOk

`func (o *UpdateBifrostConnectorRequest) GetHyxVersionOk() (*string, bool)`

GetHyxVersionOk returns a tuple with the HyxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyxVersion

`func (o *UpdateBifrostConnectorRequest) SetHyxVersion(v string)`

SetHyxVersion sets HyxVersion field to given value.

### HasHyxVersion

`func (o *UpdateBifrostConnectorRequest) HasHyxVersion() bool`

HasHyxVersion returns a boolean if a field has been set.

### SetHyxVersionNil

`func (o *UpdateBifrostConnectorRequest) SetHyxVersionNil(b bool)`

 SetHyxVersionNil sets the value for HyxVersion to be an explicit nil

### UnsetHyxVersion
`func (o *UpdateBifrostConnectorRequest) UnsetHyxVersion()`

UnsetHyxVersion ensures that no value is present for HyxVersion, not even an explicit nil
### GetId

`func (o *UpdateBifrostConnectorRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateBifrostConnectorRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateBifrostConnectorRequest) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *UpdateBifrostConnectorRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateBifrostConnectorRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *UpdateBifrostConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBifrostConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBifrostConnectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBifrostConnectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateBifrostConnectorRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateBifrostConnectorRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantSourceSideIp

`func (o *UpdateBifrostConnectorRequest) GetTenantSourceSideIp() string`

GetTenantSourceSideIp returns the TenantSourceSideIp field if non-nil, zero value otherwise.

### GetTenantSourceSideIpOk

`func (o *UpdateBifrostConnectorRequest) GetTenantSourceSideIpOk() (*string, bool)`

GetTenantSourceSideIpOk returns a tuple with the TenantSourceSideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantSourceSideIp

`func (o *UpdateBifrostConnectorRequest) SetTenantSourceSideIp(v string)`

SetTenantSourceSideIp sets TenantSourceSideIp field to given value.

### HasTenantSourceSideIp

`func (o *UpdateBifrostConnectorRequest) HasTenantSourceSideIp() bool`

HasTenantSourceSideIp returns a boolean if a field has been set.

### SetTenantSourceSideIpNil

`func (o *UpdateBifrostConnectorRequest) SetTenantSourceSideIpNil(b bool)`

 SetTenantSourceSideIpNil sets the value for TenantSourceSideIp to be an explicit nil

### UnsetTenantSourceSideIp
`func (o *UpdateBifrostConnectorRequest) UnsetTenantSourceSideIp()`

UnsetTenantSourceSideIp ensures that no value is present for TenantSourceSideIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


