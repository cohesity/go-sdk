# UpdateDirQuotaArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableDirQuota** | Pointer to **NullableBool** | Specifies directory quota to be disabled on the view. | [optional] 
**Quota** | Pointer to [**DirQuotaPolicy**](DirQuotaPolicy.md) |  | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the view. | [optional] 

## Methods

### NewUpdateDirQuotaArgs

`func NewUpdateDirQuotaArgs() *UpdateDirQuotaArgs`

NewUpdateDirQuotaArgs instantiates a new UpdateDirQuotaArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDirQuotaArgsWithDefaults

`func NewUpdateDirQuotaArgsWithDefaults() *UpdateDirQuotaArgs`

NewUpdateDirQuotaArgsWithDefaults instantiates a new UpdateDirQuotaArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableDirQuota

`func (o *UpdateDirQuotaArgs) GetDisableDirQuota() bool`

GetDisableDirQuota returns the DisableDirQuota field if non-nil, zero value otherwise.

### GetDisableDirQuotaOk

`func (o *UpdateDirQuotaArgs) GetDisableDirQuotaOk() (*bool, bool)`

GetDisableDirQuotaOk returns a tuple with the DisableDirQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDirQuota

`func (o *UpdateDirQuotaArgs) SetDisableDirQuota(v bool)`

SetDisableDirQuota sets DisableDirQuota field to given value.

### HasDisableDirQuota

`func (o *UpdateDirQuotaArgs) HasDisableDirQuota() bool`

HasDisableDirQuota returns a boolean if a field has been set.

### SetDisableDirQuotaNil

`func (o *UpdateDirQuotaArgs) SetDisableDirQuotaNil(b bool)`

 SetDisableDirQuotaNil sets the value for DisableDirQuota to be an explicit nil

### UnsetDisableDirQuota
`func (o *UpdateDirQuotaArgs) UnsetDisableDirQuota()`

UnsetDisableDirQuota ensures that no value is present for DisableDirQuota, not even an explicit nil
### GetQuota

`func (o *UpdateDirQuotaArgs) GetQuota() DirQuotaPolicy`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateDirQuotaArgs) GetQuotaOk() (*DirQuotaPolicy, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateDirQuotaArgs) SetQuota(v DirQuotaPolicy)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateDirQuotaArgs) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetViewName

`func (o *UpdateDirQuotaArgs) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *UpdateDirQuotaArgs) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *UpdateDirQuotaArgs) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *UpdateDirQuotaArgs) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *UpdateDirQuotaArgs) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *UpdateDirQuotaArgs) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


