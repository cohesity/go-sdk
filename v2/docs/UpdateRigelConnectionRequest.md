# UpdateRigelConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the connection. | 
**Scalable** | Pointer to **NullableBool** | Flag to specify if the connection is scalable. | [optional] 
**TenantId** | **NullableString** | Specifies the id of the tenant which the connection belongs to. | 
**ConnectorGroups** | Pointer to [**[]ConnectorGroup**](ConnectorGroup.md) | Specifies the connector groups in the connection. | [optional] 
**UngroupedConnectors** | Pointer to **[]int64** | Specifies the ids of the connectors which are not grouped in this connection | [optional] 

## Methods

### NewUpdateRigelConnectionRequest

`func NewUpdateRigelConnectionRequest(name NullableString, tenantId NullableString, ) *UpdateRigelConnectionRequest`

NewUpdateRigelConnectionRequest instantiates a new UpdateRigelConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRigelConnectionRequestWithDefaults

`func NewUpdateRigelConnectionRequestWithDefaults() *UpdateRigelConnectionRequest`

NewUpdateRigelConnectionRequestWithDefaults instantiates a new UpdateRigelConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRigelConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRigelConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRigelConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateRigelConnectionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateRigelConnectionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScalable

`func (o *UpdateRigelConnectionRequest) GetScalable() bool`

GetScalable returns the Scalable field if non-nil, zero value otherwise.

### GetScalableOk

`func (o *UpdateRigelConnectionRequest) GetScalableOk() (*bool, bool)`

GetScalableOk returns a tuple with the Scalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalable

`func (o *UpdateRigelConnectionRequest) SetScalable(v bool)`

SetScalable sets Scalable field to given value.

### HasScalable

`func (o *UpdateRigelConnectionRequest) HasScalable() bool`

HasScalable returns a boolean if a field has been set.

### SetScalableNil

`func (o *UpdateRigelConnectionRequest) SetScalableNil(b bool)`

 SetScalableNil sets the value for Scalable to be an explicit nil

### UnsetScalable
`func (o *UpdateRigelConnectionRequest) UnsetScalable()`

UnsetScalable ensures that no value is present for Scalable, not even an explicit nil
### GetTenantId

`func (o *UpdateRigelConnectionRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateRigelConnectionRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateRigelConnectionRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *UpdateRigelConnectionRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UpdateRigelConnectionRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetConnectorGroups

`func (o *UpdateRigelConnectionRequest) GetConnectorGroups() []ConnectorGroup`

GetConnectorGroups returns the ConnectorGroups field if non-nil, zero value otherwise.

### GetConnectorGroupsOk

`func (o *UpdateRigelConnectionRequest) GetConnectorGroupsOk() (*[]ConnectorGroup, bool)`

GetConnectorGroupsOk returns a tuple with the ConnectorGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGroups

`func (o *UpdateRigelConnectionRequest) SetConnectorGroups(v []ConnectorGroup)`

SetConnectorGroups sets ConnectorGroups field to given value.

### HasConnectorGroups

`func (o *UpdateRigelConnectionRequest) HasConnectorGroups() bool`

HasConnectorGroups returns a boolean if a field has been set.

### GetUngroupedConnectors

`func (o *UpdateRigelConnectionRequest) GetUngroupedConnectors() []int64`

GetUngroupedConnectors returns the UngroupedConnectors field if non-nil, zero value otherwise.

### GetUngroupedConnectorsOk

`func (o *UpdateRigelConnectionRequest) GetUngroupedConnectorsOk() (*[]int64, bool)`

GetUngroupedConnectorsOk returns a tuple with the UngroupedConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUngroupedConnectors

`func (o *UpdateRigelConnectionRequest) SetUngroupedConnectors(v []int64)`

SetUngroupedConnectors sets UngroupedConnectors field to given value.

### HasUngroupedConnectors

`func (o *UpdateRigelConnectionRequest) HasUngroupedConnectors() bool`

HasUngroupedConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


