# TenantViewBoxUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 
**ViewBoxIds** | Pointer to **[]int64** | Specifies the ViewBoxIds for respective tenant. | [optional] 

## Methods

### NewTenantViewBoxUpdateParameters

`func NewTenantViewBoxUpdateParameters() *TenantViewBoxUpdateParameters`

NewTenantViewBoxUpdateParameters instantiates a new TenantViewBoxUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantViewBoxUpdateParametersWithDefaults

`func NewTenantViewBoxUpdateParametersWithDefaults() *TenantViewBoxUpdateParameters`

NewTenantViewBoxUpdateParametersWithDefaults instantiates a new TenantViewBoxUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantViewBoxUpdateParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantViewBoxUpdateParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantViewBoxUpdateParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantViewBoxUpdateParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantViewBoxUpdateParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantViewBoxUpdateParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetViewBoxIds

`func (o *TenantViewBoxUpdateParameters) GetViewBoxIds() []int64`

GetViewBoxIds returns the ViewBoxIds field if non-nil, zero value otherwise.

### GetViewBoxIdsOk

`func (o *TenantViewBoxUpdateParameters) GetViewBoxIdsOk() (*[]int64, bool)`

GetViewBoxIdsOk returns a tuple with the ViewBoxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxIds

`func (o *TenantViewBoxUpdateParameters) SetViewBoxIds(v []int64)`

SetViewBoxIds sets ViewBoxIds field to given value.

### HasViewBoxIds

`func (o *TenantViewBoxUpdateParameters) HasViewBoxIds() bool`

HasViewBoxIds returns a boolean if a field has been set.

### SetViewBoxIdsNil

`func (o *TenantViewBoxUpdateParameters) SetViewBoxIdsNil(b bool)`

 SetViewBoxIdsNil sets the value for ViewBoxIds to be an explicit nil

### UnsetViewBoxIds
`func (o *TenantViewBoxUpdateParameters) UnsetViewBoxIds()`

UnsetViewBoxIds ensures that no value is present for ViewBoxIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


