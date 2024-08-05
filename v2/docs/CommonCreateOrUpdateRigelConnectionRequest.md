# CommonCreateOrUpdateRigelConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the connection. | 
**Scalable** | Pointer to **NullableBool** | Flag to specify if the connection is scalable. | [optional] 
**TenantId** | **NullableString** | Specifies the id of the tenant which the connection belongs to. | 

## Methods

### NewCommonCreateOrUpdateRigelConnectionRequest

`func NewCommonCreateOrUpdateRigelConnectionRequest(name NullableString, tenantId NullableString, ) *CommonCreateOrUpdateRigelConnectionRequest`

NewCommonCreateOrUpdateRigelConnectionRequest instantiates a new CommonCreateOrUpdateRigelConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonCreateOrUpdateRigelConnectionRequestWithDefaults

`func NewCommonCreateOrUpdateRigelConnectionRequestWithDefaults() *CommonCreateOrUpdateRigelConnectionRequest`

NewCommonCreateOrUpdateRigelConnectionRequestWithDefaults instantiates a new CommonCreateOrUpdateRigelConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonCreateOrUpdateRigelConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonCreateOrUpdateRigelConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonCreateOrUpdateRigelConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonCreateOrUpdateRigelConnectionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonCreateOrUpdateRigelConnectionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScalable

`func (o *CommonCreateOrUpdateRigelConnectionRequest) GetScalable() bool`

GetScalable returns the Scalable field if non-nil, zero value otherwise.

### GetScalableOk

`func (o *CommonCreateOrUpdateRigelConnectionRequest) GetScalableOk() (*bool, bool)`

GetScalableOk returns a tuple with the Scalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalable

`func (o *CommonCreateOrUpdateRigelConnectionRequest) SetScalable(v bool)`

SetScalable sets Scalable field to given value.

### HasScalable

`func (o *CommonCreateOrUpdateRigelConnectionRequest) HasScalable() bool`

HasScalable returns a boolean if a field has been set.

### SetScalableNil

`func (o *CommonCreateOrUpdateRigelConnectionRequest) SetScalableNil(b bool)`

 SetScalableNil sets the value for Scalable to be an explicit nil

### UnsetScalable
`func (o *CommonCreateOrUpdateRigelConnectionRequest) UnsetScalable()`

UnsetScalable ensures that no value is present for Scalable, not even an explicit nil
### GetTenantId

`func (o *CommonCreateOrUpdateRigelConnectionRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CommonCreateOrUpdateRigelConnectionRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CommonCreateOrUpdateRigelConnectionRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *CommonCreateOrUpdateRigelConnectionRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CommonCreateOrUpdateRigelConnectionRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


