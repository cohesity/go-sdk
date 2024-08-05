# DeleteRigelConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the id of the tenant which the connector belongs to. | 

## Methods

### NewDeleteRigelConnectorRequest

`func NewDeleteRigelConnectorRequest(tenantId NullableString, ) *DeleteRigelConnectorRequest`

NewDeleteRigelConnectorRequest instantiates a new DeleteRigelConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRigelConnectorRequestWithDefaults

`func NewDeleteRigelConnectorRequestWithDefaults() *DeleteRigelConnectorRequest`

NewDeleteRigelConnectorRequestWithDefaults instantiates a new DeleteRigelConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DeleteRigelConnectorRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeleteRigelConnectorRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeleteRigelConnectorRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *DeleteRigelConnectorRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DeleteRigelConnectorRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


