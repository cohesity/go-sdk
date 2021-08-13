# UpdateRigelConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the id of the tenant which the connection belongs to. | 
**Name** | **NullableString** | Specifies the name of the connection. | 

## Methods

### NewUpdateRigelConnectionRequest

`func NewUpdateRigelConnectionRequest(tenantId NullableString, name NullableString, ) *UpdateRigelConnectionRequest`

NewUpdateRigelConnectionRequest instantiates a new UpdateRigelConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRigelConnectionRequestWithDefaults

`func NewUpdateRigelConnectionRequestWithDefaults() *UpdateRigelConnectionRequest`

NewUpdateRigelConnectionRequestWithDefaults instantiates a new UpdateRigelConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


