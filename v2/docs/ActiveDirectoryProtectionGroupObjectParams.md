# ActiveDirectoryProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppParams** | Pointer to [**[]ActiveDirectoryAppParams**](ActiveDirectoryAppParams.md) | Specifies the specific parameters required for active directory app configuration. | [optional] 
**EnableSystemBackup** | Pointer to **NullableBool** | Specifies whether to take bmr backup. If this is not specified, the bmr backup won&#39;t be enabled. | [optional] 
**SourceId** | **NullableInt64** | Specifies the id of the registered active directory source. | 
**SourceName** | Pointer to **NullableString** | Specifies the name of the registered active directory source. | [optional] [readonly] 

## Methods

### NewActiveDirectoryProtectionGroupObjectParams

`func NewActiveDirectoryProtectionGroupObjectParams(sourceId NullableInt64, ) *ActiveDirectoryProtectionGroupObjectParams`

NewActiveDirectoryProtectionGroupObjectParams instantiates a new ActiveDirectoryProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryProtectionGroupObjectParamsWithDefaults

`func NewActiveDirectoryProtectionGroupObjectParamsWithDefaults() *ActiveDirectoryProtectionGroupObjectParams`

NewActiveDirectoryProtectionGroupObjectParamsWithDefaults instantiates a new ActiveDirectoryProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppParams

`func (o *ActiveDirectoryProtectionGroupObjectParams) GetAppParams() []ActiveDirectoryAppParams`

GetAppParams returns the AppParams field if non-nil, zero value otherwise.

### GetAppParamsOk

`func (o *ActiveDirectoryProtectionGroupObjectParams) GetAppParamsOk() (*[]ActiveDirectoryAppParams, bool)`

GetAppParamsOk returns a tuple with the AppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppParams

`func (o *ActiveDirectoryProtectionGroupObjectParams) SetAppParams(v []ActiveDirectoryAppParams)`

SetAppParams sets AppParams field to given value.

### HasAppParams

`func (o *ActiveDirectoryProtectionGroupObjectParams) HasAppParams() bool`

HasAppParams returns a boolean if a field has been set.

### SetAppParamsNil

`func (o *ActiveDirectoryProtectionGroupObjectParams) SetAppParamsNil(b bool)`

 SetAppParamsNil sets the value for AppParams to be an explicit nil

### UnsetAppParams
`func (o *ActiveDirectoryProtectionGroupObjectParams) UnsetAppParams()`

UnsetAppParams ensures that no value is present for AppParams, not even an explicit nil
### GetEnableSystemBackup

`func (o *ActiveDirectoryProtectionGroupObjectParams) GetEnableSystemBackup() bool`

GetEnableSystemBackup returns the EnableSystemBackup field if non-nil, zero value otherwise.

### GetEnableSystemBackupOk

`func (o *ActiveDirectoryProtectionGroupObjectParams) GetEnableSystemBackupOk() (*bool, bool)`

GetEnableSystemBackupOk returns a tuple with the EnableSystemBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSystemBackup

`func (o *ActiveDirectoryProtectionGroupObjectParams) SetEnableSystemBackup(v bool)`

SetEnableSystemBackup sets EnableSystemBackup field to given value.

### HasEnableSystemBackup

`func (o *ActiveDirectoryProtectionGroupObjectParams) HasEnableSystemBackup() bool`

HasEnableSystemBackup returns a boolean if a field has been set.

### SetEnableSystemBackupNil

`func (o *ActiveDirectoryProtectionGroupObjectParams) SetEnableSystemBackupNil(b bool)`

 SetEnableSystemBackupNil sets the value for EnableSystemBackup to be an explicit nil

### UnsetEnableSystemBackup
`func (o *ActiveDirectoryProtectionGroupObjectParams) UnsetEnableSystemBackup()`

UnsetEnableSystemBackup ensures that no value is present for EnableSystemBackup, not even an explicit nil
### GetSourceId

`func (o *ActiveDirectoryProtectionGroupObjectParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ActiveDirectoryProtectionGroupObjectParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ActiveDirectoryProtectionGroupObjectParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *ActiveDirectoryProtectionGroupObjectParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ActiveDirectoryProtectionGroupObjectParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *ActiveDirectoryProtectionGroupObjectParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ActiveDirectoryProtectionGroupObjectParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ActiveDirectoryProtectionGroupObjectParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ActiveDirectoryProtectionGroupObjectParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *ActiveDirectoryProtectionGroupObjectParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *ActiveDirectoryProtectionGroupObjectParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


