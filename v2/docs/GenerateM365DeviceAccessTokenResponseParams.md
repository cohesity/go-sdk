# GenerateM365DeviceAccessTokenResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **NullableString** | Specifies the access token for Microsoft365 Azure PowerShell. | [optional] 
**ExpiresInSecs** | Pointer to **NullableInt64** | Specifies the number of seconds before the included access token is valid for. | [optional] 

## Methods

### NewGenerateM365DeviceAccessTokenResponseParams

`func NewGenerateM365DeviceAccessTokenResponseParams() *GenerateM365DeviceAccessTokenResponseParams`

NewGenerateM365DeviceAccessTokenResponseParams instantiates a new GenerateM365DeviceAccessTokenResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateM365DeviceAccessTokenResponseParamsWithDefaults

`func NewGenerateM365DeviceAccessTokenResponseParamsWithDefaults() *GenerateM365DeviceAccessTokenResponseParams`

NewGenerateM365DeviceAccessTokenResponseParamsWithDefaults instantiates a new GenerateM365DeviceAccessTokenResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GenerateM365DeviceAccessTokenResponseParams) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GenerateM365DeviceAccessTokenResponseParams) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GenerateM365DeviceAccessTokenResponseParams) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *GenerateM365DeviceAccessTokenResponseParams) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *GenerateM365DeviceAccessTokenResponseParams) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *GenerateM365DeviceAccessTokenResponseParams) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetExpiresInSecs

`func (o *GenerateM365DeviceAccessTokenResponseParams) GetExpiresInSecs() int64`

GetExpiresInSecs returns the ExpiresInSecs field if non-nil, zero value otherwise.

### GetExpiresInSecsOk

`func (o *GenerateM365DeviceAccessTokenResponseParams) GetExpiresInSecsOk() (*int64, bool)`

GetExpiresInSecsOk returns a tuple with the ExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSecs

`func (o *GenerateM365DeviceAccessTokenResponseParams) SetExpiresInSecs(v int64)`

SetExpiresInSecs sets ExpiresInSecs field to given value.

### HasExpiresInSecs

`func (o *GenerateM365DeviceAccessTokenResponseParams) HasExpiresInSecs() bool`

HasExpiresInSecs returns a boolean if a field has been set.

### SetExpiresInSecsNil

`func (o *GenerateM365DeviceAccessTokenResponseParams) SetExpiresInSecsNil(b bool)`

 SetExpiresInSecsNil sets the value for ExpiresInSecs to be an explicit nil

### UnsetExpiresInSecs
`func (o *GenerateM365DeviceAccessTokenResponseParams) UnsetExpiresInSecs()`

UnsetExpiresInSecs ensures that no value is present for ExpiresInSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


