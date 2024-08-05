# CreateRigelConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the connection. | 
**Scalable** | Pointer to **NullableBool** | Flag to specify if the connection is scalable. | [optional] 
**TenantId** | **NullableString** | Specifies the id of the tenant which the connection belongs to. | 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection. Required if environment type is &#39;kHelios&#39; | [optional] 
**EnvType** | Pointer to **NullableString** | Specifies the type of the connection. | [optional] 

## Methods

### NewCreateRigelConnectionRequest

`func NewCreateRigelConnectionRequest(name NullableString, tenantId NullableString, ) *CreateRigelConnectionRequest`

NewCreateRigelConnectionRequest instantiates a new CreateRigelConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRigelConnectionRequestWithDefaults

`func NewCreateRigelConnectionRequestWithDefaults() *CreateRigelConnectionRequest`

NewCreateRigelConnectionRequestWithDefaults instantiates a new CreateRigelConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRigelConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRigelConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRigelConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateRigelConnectionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateRigelConnectionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScalable

`func (o *CreateRigelConnectionRequest) GetScalable() bool`

GetScalable returns the Scalable field if non-nil, zero value otherwise.

### GetScalableOk

`func (o *CreateRigelConnectionRequest) GetScalableOk() (*bool, bool)`

GetScalableOk returns a tuple with the Scalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalable

`func (o *CreateRigelConnectionRequest) SetScalable(v bool)`

SetScalable sets Scalable field to given value.

### HasScalable

`func (o *CreateRigelConnectionRequest) HasScalable() bool`

HasScalable returns a boolean if a field has been set.

### SetScalableNil

`func (o *CreateRigelConnectionRequest) SetScalableNil(b bool)`

 SetScalableNil sets the value for Scalable to be an explicit nil

### UnsetScalable
`func (o *CreateRigelConnectionRequest) UnsetScalable()`

UnsetScalable ensures that no value is present for Scalable, not even an explicit nil
### GetTenantId

`func (o *CreateRigelConnectionRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateRigelConnectionRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateRigelConnectionRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *CreateRigelConnectionRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateRigelConnectionRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetConnectionId

`func (o *CreateRigelConnectionRequest) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CreateRigelConnectionRequest) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CreateRigelConnectionRequest) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CreateRigelConnectionRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *CreateRigelConnectionRequest) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *CreateRigelConnectionRequest) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetEnvType

`func (o *CreateRigelConnectionRequest) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *CreateRigelConnectionRequest) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *CreateRigelConnectionRequest) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *CreateRigelConnectionRequest) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *CreateRigelConnectionRequest) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *CreateRigelConnectionRequest) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


