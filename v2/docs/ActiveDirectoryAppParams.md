# ActiveDirectoryAppParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableInt64** | Specifies the application id of the Active Directory instance. | [optional] 
**AppName** | Pointer to **NullableString** | Specifies the application name of the Active Directory instance. | [optional] [readonly] 

## Methods

### NewActiveDirectoryAppParams

`func NewActiveDirectoryAppParams() *ActiveDirectoryAppParams`

NewActiveDirectoryAppParams instantiates a new ActiveDirectoryAppParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryAppParamsWithDefaults

`func NewActiveDirectoryAppParamsWithDefaults() *ActiveDirectoryAppParams`

NewActiveDirectoryAppParamsWithDefaults instantiates a new ActiveDirectoryAppParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ActiveDirectoryAppParams) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ActiveDirectoryAppParams) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ActiveDirectoryAppParams) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ActiveDirectoryAppParams) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *ActiveDirectoryAppParams) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *ActiveDirectoryAppParams) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAppName

`func (o *ActiveDirectoryAppParams) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ActiveDirectoryAppParams) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ActiveDirectoryAppParams) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ActiveDirectoryAppParams) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *ActiveDirectoryAppParams) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *ActiveDirectoryAppParams) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


