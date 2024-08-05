# BifrostConnector

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

### NewBifrostConnector

`func NewBifrostConnector(connectionId NullableInt64, id NullableString, ) *BifrostConnector`

NewBifrostConnector instantiates a new BifrostConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBifrostConnectorWithDefaults

`func NewBifrostConnectorWithDefaults() *BifrostConnector`

NewBifrostConnectorWithDefaults instantiates a new BifrostConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCohesitySideIp

`func (o *BifrostConnector) GetCohesitySideIp() string`

GetCohesitySideIp returns the CohesitySideIp field if non-nil, zero value otherwise.

### GetCohesitySideIpOk

`func (o *BifrostConnector) GetCohesitySideIpOk() (*string, bool)`

GetCohesitySideIpOk returns a tuple with the CohesitySideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesitySideIp

`func (o *BifrostConnector) SetCohesitySideIp(v string)`

SetCohesitySideIp sets CohesitySideIp field to given value.

### HasCohesitySideIp

`func (o *BifrostConnector) HasCohesitySideIp() bool`

HasCohesitySideIp returns a boolean if a field has been set.

### SetCohesitySideIpNil

`func (o *BifrostConnector) SetCohesitySideIpNil(b bool)`

 SetCohesitySideIpNil sets the value for CohesitySideIp to be an explicit nil

### UnsetCohesitySideIp
`func (o *BifrostConnector) UnsetCohesitySideIp()`

UnsetCohesitySideIp ensures that no value is present for CohesitySideIp, not even an explicit nil
### GetConnectionId

`func (o *BifrostConnector) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BifrostConnector) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BifrostConnector) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.


### SetConnectionIdNil

`func (o *BifrostConnector) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *BifrostConnector) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnectionStatus

`func (o *BifrostConnector) GetConnectionStatus() ConnectorConnectionInfo`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *BifrostConnector) GetConnectionStatusOk() (*ConnectorConnectionInfo, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *BifrostConnector) SetConnectionStatus(v ConnectorConnectionInfo)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *BifrostConnector) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetHyxVersion

`func (o *BifrostConnector) GetHyxVersion() string`

GetHyxVersion returns the HyxVersion field if non-nil, zero value otherwise.

### GetHyxVersionOk

`func (o *BifrostConnector) GetHyxVersionOk() (*string, bool)`

GetHyxVersionOk returns a tuple with the HyxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyxVersion

`func (o *BifrostConnector) SetHyxVersion(v string)`

SetHyxVersion sets HyxVersion field to given value.

### HasHyxVersion

`func (o *BifrostConnector) HasHyxVersion() bool`

HasHyxVersion returns a boolean if a field has been set.

### SetHyxVersionNil

`func (o *BifrostConnector) SetHyxVersionNil(b bool)`

 SetHyxVersionNil sets the value for HyxVersion to be an explicit nil

### UnsetHyxVersion
`func (o *BifrostConnector) UnsetHyxVersion()`

UnsetHyxVersion ensures that no value is present for HyxVersion, not even an explicit nil
### GetId

`func (o *BifrostConnector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BifrostConnector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BifrostConnector) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *BifrostConnector) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BifrostConnector) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BifrostConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BifrostConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BifrostConnector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BifrostConnector) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BifrostConnector) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BifrostConnector) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantSourceSideIp

`func (o *BifrostConnector) GetTenantSourceSideIp() string`

GetTenantSourceSideIp returns the TenantSourceSideIp field if non-nil, zero value otherwise.

### GetTenantSourceSideIpOk

`func (o *BifrostConnector) GetTenantSourceSideIpOk() (*string, bool)`

GetTenantSourceSideIpOk returns a tuple with the TenantSourceSideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantSourceSideIp

`func (o *BifrostConnector) SetTenantSourceSideIp(v string)`

SetTenantSourceSideIp sets TenantSourceSideIp field to given value.

### HasTenantSourceSideIp

`func (o *BifrostConnector) HasTenantSourceSideIp() bool`

HasTenantSourceSideIp returns a boolean if a field has been set.

### SetTenantSourceSideIpNil

`func (o *BifrostConnector) SetTenantSourceSideIpNil(b bool)`

 SetTenantSourceSideIpNil sets the value for TenantSourceSideIp to be an explicit nil

### UnsetTenantSourceSideIp
`func (o *BifrostConnector) UnsetTenantSourceSideIp()`

UnsetTenantSourceSideIp ensures that no value is present for TenantSourceSideIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


