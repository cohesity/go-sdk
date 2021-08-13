# VmwareObjectEnableAppProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppType** | **NullableString** | Specifies the app from which protection must be enabled. | 
**EnableSqlProtectionParams** | Pointer to [**VmwareObjectEnableSqlProtectionParams**](VmwareObjectEnableSqlProtectionParams.md) |  | [optional] 

## Methods

### NewVmwareObjectEnableAppProtectionParams

`func NewVmwareObjectEnableAppProtectionParams(appType NullableString, ) *VmwareObjectEnableAppProtectionParams`

NewVmwareObjectEnableAppProtectionParams instantiates a new VmwareObjectEnableAppProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectEnableAppProtectionParamsWithDefaults

`func NewVmwareObjectEnableAppProtectionParamsWithDefaults() *VmwareObjectEnableAppProtectionParams`

NewVmwareObjectEnableAppProtectionParamsWithDefaults instantiates a new VmwareObjectEnableAppProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppType

`func (o *VmwareObjectEnableAppProtectionParams) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *VmwareObjectEnableAppProtectionParams) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *VmwareObjectEnableAppProtectionParams) SetAppType(v string)`

SetAppType sets AppType field to given value.


### SetAppTypeNil

`func (o *VmwareObjectEnableAppProtectionParams) SetAppTypeNil(b bool)`

 SetAppTypeNil sets the value for AppType to be an explicit nil

### UnsetAppType
`func (o *VmwareObjectEnableAppProtectionParams) UnsetAppType()`

UnsetAppType ensures that no value is present for AppType, not even an explicit nil
### GetEnableSqlProtectionParams

`func (o *VmwareObjectEnableAppProtectionParams) GetEnableSqlProtectionParams() VmwareObjectEnableSqlProtectionParams`

GetEnableSqlProtectionParams returns the EnableSqlProtectionParams field if non-nil, zero value otherwise.

### GetEnableSqlProtectionParamsOk

`func (o *VmwareObjectEnableAppProtectionParams) GetEnableSqlProtectionParamsOk() (*VmwareObjectEnableSqlProtectionParams, bool)`

GetEnableSqlProtectionParamsOk returns a tuple with the EnableSqlProtectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSqlProtectionParams

`func (o *VmwareObjectEnableAppProtectionParams) SetEnableSqlProtectionParams(v VmwareObjectEnableSqlProtectionParams)`

SetEnableSqlProtectionParams sets EnableSqlProtectionParams field to given value.

### HasEnableSqlProtectionParams

`func (o *VmwareObjectEnableAppProtectionParams) HasEnableSqlProtectionParams() bool`

HasEnableSqlProtectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


