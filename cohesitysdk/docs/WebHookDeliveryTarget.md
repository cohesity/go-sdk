# WebHookDeliveryTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurlOptions** | Pointer to **NullableString** | Specifies curl options used to invoke external api url defined above. | [optional] 
**ExternalApiUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebHookDeliveryTarget

`func NewWebHookDeliveryTarget() *WebHookDeliveryTarget`

NewWebHookDeliveryTarget instantiates a new WebHookDeliveryTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookDeliveryTargetWithDefaults

`func NewWebHookDeliveryTargetWithDefaults() *WebHookDeliveryTarget`

NewWebHookDeliveryTargetWithDefaults instantiates a new WebHookDeliveryTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurlOptions

`func (o *WebHookDeliveryTarget) GetCurlOptions() string`

GetCurlOptions returns the CurlOptions field if non-nil, zero value otherwise.

### GetCurlOptionsOk

`func (o *WebHookDeliveryTarget) GetCurlOptionsOk() (*string, bool)`

GetCurlOptionsOk returns a tuple with the CurlOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurlOptions

`func (o *WebHookDeliveryTarget) SetCurlOptions(v string)`

SetCurlOptions sets CurlOptions field to given value.

### HasCurlOptions

`func (o *WebHookDeliveryTarget) HasCurlOptions() bool`

HasCurlOptions returns a boolean if a field has been set.

### SetCurlOptionsNil

`func (o *WebHookDeliveryTarget) SetCurlOptionsNil(b bool)`

 SetCurlOptionsNil sets the value for CurlOptions to be an explicit nil

### UnsetCurlOptions
`func (o *WebHookDeliveryTarget) UnsetCurlOptions()`

UnsetCurlOptions ensures that no value is present for CurlOptions, not even an explicit nil
### GetExternalApiUrl

`func (o *WebHookDeliveryTarget) GetExternalApiUrl() string`

GetExternalApiUrl returns the ExternalApiUrl field if non-nil, zero value otherwise.

### GetExternalApiUrlOk

`func (o *WebHookDeliveryTarget) GetExternalApiUrlOk() (*string, bool)`

GetExternalApiUrlOk returns a tuple with the ExternalApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApiUrl

`func (o *WebHookDeliveryTarget) SetExternalApiUrl(v string)`

SetExternalApiUrl sets ExternalApiUrl field to given value.

### HasExternalApiUrl

`func (o *WebHookDeliveryTarget) HasExternalApiUrl() bool`

HasExternalApiUrl returns a boolean if a field has been set.

### SetExternalApiUrlNil

`func (o *WebHookDeliveryTarget) SetExternalApiUrlNil(b bool)`

 SetExternalApiUrlNil sets the value for ExternalApiUrl to be an explicit nil

### UnsetExternalApiUrl
`func (o *WebHookDeliveryTarget) UnsetExternalApiUrl()`

UnsetExternalApiUrl ensures that no value is present for ExternalApiUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


