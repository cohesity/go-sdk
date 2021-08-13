# CreateOrUpdateBifrostConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **NullableInt64** | Specifies the Id of the connection which this connector belongs to. | 
**TenantId** | **NullableString** | Specifies the id of the tenant which the connector belongs to. | 
**Name** | **NullableString** | Specifies the name of the connector. | 

## Methods

### NewCreateOrUpdateBifrostConnectorRequest

`func NewCreateOrUpdateBifrostConnectorRequest(connectionId NullableInt64, tenantId NullableString, name NullableString, ) *CreateOrUpdateBifrostConnectorRequest`

NewCreateOrUpdateBifrostConnectorRequest instantiates a new CreateOrUpdateBifrostConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateBifrostConnectorRequestWithDefaults

`func NewCreateOrUpdateBifrostConnectorRequestWithDefaults() *CreateOrUpdateBifrostConnectorRequest`

NewCreateOrUpdateBifrostConnectorRequestWithDefaults instantiates a new CreateOrUpdateBifrostConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *CreateOrUpdateBifrostConnectorRequest) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CreateOrUpdateBifrostConnectorRequest) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CreateOrUpdateBifrostConnectorRequest) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.


### SetConnectionIdNil

`func (o *CreateOrUpdateBifrostConnectorRequest) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *CreateOrUpdateBifrostConnectorRequest) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetTenantId

`func (o *CreateOrUpdateBifrostConnectorRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateOrUpdateBifrostConnectorRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateOrUpdateBifrostConnectorRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *CreateOrUpdateBifrostConnectorRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateOrUpdateBifrostConnectorRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetName

`func (o *CreateOrUpdateBifrostConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateBifrostConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateBifrostConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateOrUpdateBifrostConnectorRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOrUpdateBifrostConnectorRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


