# TenantViewUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 
**ViewNames** | Pointer to **[]string** | Specifies the PolicyIds for respective tenant. | [optional] 

## Methods

### NewTenantViewUpdate

`func NewTenantViewUpdate() *TenantViewUpdate`

NewTenantViewUpdate instantiates a new TenantViewUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantViewUpdateWithDefaults

`func NewTenantViewUpdateWithDefaults() *TenantViewUpdate`

NewTenantViewUpdateWithDefaults instantiates a new TenantViewUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantViewUpdate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantViewUpdate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantViewUpdate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantViewUpdate) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantViewUpdate) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantViewUpdate) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetViewNames

`func (o *TenantViewUpdate) GetViewNames() []string`

GetViewNames returns the ViewNames field if non-nil, zero value otherwise.

### GetViewNamesOk

`func (o *TenantViewUpdate) GetViewNamesOk() (*[]string, bool)`

GetViewNamesOk returns a tuple with the ViewNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNames

`func (o *TenantViewUpdate) SetViewNames(v []string)`

SetViewNames sets ViewNames field to given value.

### HasViewNames

`func (o *TenantViewUpdate) HasViewNames() bool`

HasViewNames returns a boolean if a field has been set.

### SetViewNamesNil

`func (o *TenantViewUpdate) SetViewNamesNil(b bool)`

 SetViewNamesNil sets the value for ViewNames to be an explicit nil

### UnsetViewNames
`func (o *TenantViewUpdate) UnsetViewNames()`

UnsetViewNames ensures that no value is present for ViewNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


