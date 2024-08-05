# GenerateM365DeviceCodeResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceCode** | Pointer to **NullableString** | Specifies the string used to verify the session between the client and the authorization server. The client uses this parameter to request the access token from the authorization server. | [optional] 
**ExpiresInSecs** | Pointer to **NullableInt64** | The number of seconds before the deviceCode and userCode expire. | [optional] 
**IntervalSecs** | Pointer to **NullableInt64** | The number of seconds the client should wait between polling requests to check for token. | [optional] 
**UserCode** | Pointer to **NullableString** | A short string shown to the user that&#39;s used to identify the session on a secondary device. | [optional] 
**VerificationUri** | Pointer to **NullableString** | The URI the user should go to with the userCode in order to sign in. | [optional] 

## Methods

### NewGenerateM365DeviceCodeResponseParams

`func NewGenerateM365DeviceCodeResponseParams() *GenerateM365DeviceCodeResponseParams`

NewGenerateM365DeviceCodeResponseParams instantiates a new GenerateM365DeviceCodeResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateM365DeviceCodeResponseParamsWithDefaults

`func NewGenerateM365DeviceCodeResponseParamsWithDefaults() *GenerateM365DeviceCodeResponseParams`

NewGenerateM365DeviceCodeResponseParamsWithDefaults instantiates a new GenerateM365DeviceCodeResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceCode

`func (o *GenerateM365DeviceCodeResponseParams) GetDeviceCode() string`

GetDeviceCode returns the DeviceCode field if non-nil, zero value otherwise.

### GetDeviceCodeOk

`func (o *GenerateM365DeviceCodeResponseParams) GetDeviceCodeOk() (*string, bool)`

GetDeviceCodeOk returns a tuple with the DeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCode

`func (o *GenerateM365DeviceCodeResponseParams) SetDeviceCode(v string)`

SetDeviceCode sets DeviceCode field to given value.

### HasDeviceCode

`func (o *GenerateM365DeviceCodeResponseParams) HasDeviceCode() bool`

HasDeviceCode returns a boolean if a field has been set.

### SetDeviceCodeNil

`func (o *GenerateM365DeviceCodeResponseParams) SetDeviceCodeNil(b bool)`

 SetDeviceCodeNil sets the value for DeviceCode to be an explicit nil

### UnsetDeviceCode
`func (o *GenerateM365DeviceCodeResponseParams) UnsetDeviceCode()`

UnsetDeviceCode ensures that no value is present for DeviceCode, not even an explicit nil
### GetExpiresInSecs

`func (o *GenerateM365DeviceCodeResponseParams) GetExpiresInSecs() int64`

GetExpiresInSecs returns the ExpiresInSecs field if non-nil, zero value otherwise.

### GetExpiresInSecsOk

`func (o *GenerateM365DeviceCodeResponseParams) GetExpiresInSecsOk() (*int64, bool)`

GetExpiresInSecsOk returns a tuple with the ExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSecs

`func (o *GenerateM365DeviceCodeResponseParams) SetExpiresInSecs(v int64)`

SetExpiresInSecs sets ExpiresInSecs field to given value.

### HasExpiresInSecs

`func (o *GenerateM365DeviceCodeResponseParams) HasExpiresInSecs() bool`

HasExpiresInSecs returns a boolean if a field has been set.

### SetExpiresInSecsNil

`func (o *GenerateM365DeviceCodeResponseParams) SetExpiresInSecsNil(b bool)`

 SetExpiresInSecsNil sets the value for ExpiresInSecs to be an explicit nil

### UnsetExpiresInSecs
`func (o *GenerateM365DeviceCodeResponseParams) UnsetExpiresInSecs()`

UnsetExpiresInSecs ensures that no value is present for ExpiresInSecs, not even an explicit nil
### GetIntervalSecs

`func (o *GenerateM365DeviceCodeResponseParams) GetIntervalSecs() int64`

GetIntervalSecs returns the IntervalSecs field if non-nil, zero value otherwise.

### GetIntervalSecsOk

`func (o *GenerateM365DeviceCodeResponseParams) GetIntervalSecsOk() (*int64, bool)`

GetIntervalSecsOk returns a tuple with the IntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSecs

`func (o *GenerateM365DeviceCodeResponseParams) SetIntervalSecs(v int64)`

SetIntervalSecs sets IntervalSecs field to given value.

### HasIntervalSecs

`func (o *GenerateM365DeviceCodeResponseParams) HasIntervalSecs() bool`

HasIntervalSecs returns a boolean if a field has been set.

### SetIntervalSecsNil

`func (o *GenerateM365DeviceCodeResponseParams) SetIntervalSecsNil(b bool)`

 SetIntervalSecsNil sets the value for IntervalSecs to be an explicit nil

### UnsetIntervalSecs
`func (o *GenerateM365DeviceCodeResponseParams) UnsetIntervalSecs()`

UnsetIntervalSecs ensures that no value is present for IntervalSecs, not even an explicit nil
### GetUserCode

`func (o *GenerateM365DeviceCodeResponseParams) GetUserCode() string`

GetUserCode returns the UserCode field if non-nil, zero value otherwise.

### GetUserCodeOk

`func (o *GenerateM365DeviceCodeResponseParams) GetUserCodeOk() (*string, bool)`

GetUserCodeOk returns a tuple with the UserCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCode

`func (o *GenerateM365DeviceCodeResponseParams) SetUserCode(v string)`

SetUserCode sets UserCode field to given value.

### HasUserCode

`func (o *GenerateM365DeviceCodeResponseParams) HasUserCode() bool`

HasUserCode returns a boolean if a field has been set.

### SetUserCodeNil

`func (o *GenerateM365DeviceCodeResponseParams) SetUserCodeNil(b bool)`

 SetUserCodeNil sets the value for UserCode to be an explicit nil

### UnsetUserCode
`func (o *GenerateM365DeviceCodeResponseParams) UnsetUserCode()`

UnsetUserCode ensures that no value is present for UserCode, not even an explicit nil
### GetVerificationUri

`func (o *GenerateM365DeviceCodeResponseParams) GetVerificationUri() string`

GetVerificationUri returns the VerificationUri field if non-nil, zero value otherwise.

### GetVerificationUriOk

`func (o *GenerateM365DeviceCodeResponseParams) GetVerificationUriOk() (*string, bool)`

GetVerificationUriOk returns a tuple with the VerificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationUri

`func (o *GenerateM365DeviceCodeResponseParams) SetVerificationUri(v string)`

SetVerificationUri sets VerificationUri field to given value.

### HasVerificationUri

`func (o *GenerateM365DeviceCodeResponseParams) HasVerificationUri() bool`

HasVerificationUri returns a boolean if a field has been set.

### SetVerificationUriNil

`func (o *GenerateM365DeviceCodeResponseParams) SetVerificationUriNil(b bool)`

 SetVerificationUriNil sets the value for VerificationUri to be an explicit nil

### UnsetVerificationUri
`func (o *GenerateM365DeviceCodeResponseParams) UnsetVerificationUri()`

UnsetVerificationUri ensures that no value is present for VerificationUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


