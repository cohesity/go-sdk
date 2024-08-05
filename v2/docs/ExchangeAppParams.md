# ExchangeAppParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableInt64** | Specifies the application id of the Exchange database which has to be protected. | [optional] 
**AppName** | Pointer to **NullableString** | Specifies the application name of the Exchange database which has to be protected. | [optional] [readonly] 

## Methods

### NewExchangeAppParams

`func NewExchangeAppParams() *ExchangeAppParams`

NewExchangeAppParams instantiates a new ExchangeAppParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeAppParamsWithDefaults

`func NewExchangeAppParamsWithDefaults() *ExchangeAppParams`

NewExchangeAppParamsWithDefaults instantiates a new ExchangeAppParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ExchangeAppParams) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ExchangeAppParams) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ExchangeAppParams) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ExchangeAppParams) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *ExchangeAppParams) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *ExchangeAppParams) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAppName

`func (o *ExchangeAppParams) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ExchangeAppParams) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ExchangeAppParams) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ExchangeAppParams) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *ExchangeAppParams) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *ExchangeAppParams) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


