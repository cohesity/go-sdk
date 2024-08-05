# CreateM365SelfServiceConfigRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailboxParams** | Pointer to [**M365SelfServiceWorkloadParams**](M365SelfServiceWorkloadParams.md) |  | [optional] 
**OneDriveParams** | Pointer to [**M365SelfServiceWorkloadParams**](M365SelfServiceWorkloadParams.md) |  | [optional] 
**TenantId** | **NullableString** | Specifies the Cohesity Tenant ID for the Microsoft365 source owner. | 
**Uuid** | **NullableString** | Specifies the UUID of the Microsoft365 Source. | 

## Methods

### NewCreateM365SelfServiceConfigRequestParams

`func NewCreateM365SelfServiceConfigRequestParams(tenantId NullableString, uuid NullableString, ) *CreateM365SelfServiceConfigRequestParams`

NewCreateM365SelfServiceConfigRequestParams instantiates a new CreateM365SelfServiceConfigRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateM365SelfServiceConfigRequestParamsWithDefaults

`func NewCreateM365SelfServiceConfigRequestParamsWithDefaults() *CreateM365SelfServiceConfigRequestParams`

NewCreateM365SelfServiceConfigRequestParamsWithDefaults instantiates a new CreateM365SelfServiceConfigRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailboxParams

`func (o *CreateM365SelfServiceConfigRequestParams) GetMailboxParams() M365SelfServiceWorkloadParams`

GetMailboxParams returns the MailboxParams field if non-nil, zero value otherwise.

### GetMailboxParamsOk

`func (o *CreateM365SelfServiceConfigRequestParams) GetMailboxParamsOk() (*M365SelfServiceWorkloadParams, bool)`

GetMailboxParamsOk returns a tuple with the MailboxParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxParams

`func (o *CreateM365SelfServiceConfigRequestParams) SetMailboxParams(v M365SelfServiceWorkloadParams)`

SetMailboxParams sets MailboxParams field to given value.

### HasMailboxParams

`func (o *CreateM365SelfServiceConfigRequestParams) HasMailboxParams() bool`

HasMailboxParams returns a boolean if a field has been set.

### GetOneDriveParams

`func (o *CreateM365SelfServiceConfigRequestParams) GetOneDriveParams() M365SelfServiceWorkloadParams`

GetOneDriveParams returns the OneDriveParams field if non-nil, zero value otherwise.

### GetOneDriveParamsOk

`func (o *CreateM365SelfServiceConfigRequestParams) GetOneDriveParamsOk() (*M365SelfServiceWorkloadParams, bool)`

GetOneDriveParamsOk returns a tuple with the OneDriveParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneDriveParams

`func (o *CreateM365SelfServiceConfigRequestParams) SetOneDriveParams(v M365SelfServiceWorkloadParams)`

SetOneDriveParams sets OneDriveParams field to given value.

### HasOneDriveParams

`func (o *CreateM365SelfServiceConfigRequestParams) HasOneDriveParams() bool`

HasOneDriveParams returns a boolean if a field has been set.

### GetTenantId

`func (o *CreateM365SelfServiceConfigRequestParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateM365SelfServiceConfigRequestParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateM365SelfServiceConfigRequestParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *CreateM365SelfServiceConfigRequestParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateM365SelfServiceConfigRequestParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUuid

`func (o *CreateM365SelfServiceConfigRequestParams) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateM365SelfServiceConfigRequestParams) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateM365SelfServiceConfigRequestParams) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### SetUuidNil

`func (o *CreateM365SelfServiceConfigRequestParams) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *CreateM365SelfServiceConfigRequestParams) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


