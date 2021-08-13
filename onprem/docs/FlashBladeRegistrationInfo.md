# FlashBladeRegistrationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **NullableString** | Specifies management ip of pure flashblade server. | [optional] 
**ApiToken** | **NullableString** | Specifies the api token of the pure flashblade. | 

## Methods

### NewFlashBladeRegistrationInfo

`func NewFlashBladeRegistrationInfo(apiToken NullableString, ) *FlashBladeRegistrationInfo`

NewFlashBladeRegistrationInfo instantiates a new FlashBladeRegistrationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashBladeRegistrationInfoWithDefaults

`func NewFlashBladeRegistrationInfoWithDefaults() *FlashBladeRegistrationInfo`

NewFlashBladeRegistrationInfoWithDefaults instantiates a new FlashBladeRegistrationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *FlashBladeRegistrationInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FlashBladeRegistrationInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FlashBladeRegistrationInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FlashBladeRegistrationInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *FlashBladeRegistrationInfo) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *FlashBladeRegistrationInfo) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetApiToken

`func (o *FlashBladeRegistrationInfo) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *FlashBladeRegistrationInfo) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *FlashBladeRegistrationInfo) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### SetApiTokenNil

`func (o *FlashBladeRegistrationInfo) SetApiTokenNil(b bool)`

 SetApiTokenNil sets the value for ApiToken to be an explicit nil

### UnsetApiToken
`func (o *FlashBladeRegistrationInfo) UnsetApiToken()`

UnsetApiToken ensures that no value is present for ApiToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


