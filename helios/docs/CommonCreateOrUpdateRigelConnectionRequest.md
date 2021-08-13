# CommonCreateOrUpdateRigelConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the id of the tenant which the connection belongs to. | 
**Name** | **NullableString** | Specifies the name of the connection. | 

## Methods

### NewCommonCreateOrUpdateRigelConnectionRequest

`func NewCommonCreateOrUpdateRigelConnectionRequest(tenantId NullableString, name NullableString, ) *CommonCreateOrUpdateRigelConnectionRequest`

NewCommonCreateOrUpdateRigelConnectionRequest instantiates a new CommonCreateOrUpdateRigelConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonCreateOrUpdateRigelConnectionRequestWithDefaults

`func NewCommonCreateOrUpdateRigelConnectionRequestWithDefaults() *CommonCreateOrUpdateRigelConnectionRequest`

NewCommonCreateOrUpdateRigelConnectionRequestWithDefaults instantiates a new CommonCreateOrUpdateRigelConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


