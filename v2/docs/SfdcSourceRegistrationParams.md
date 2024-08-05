# SfdcSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **NullableString** | Specifies the token that will be used for fetching oAuth tokens from salesforce. | 
**CallbackUrl** | Pointer to **NullableString** | Specifies the URL added in the connected apps Callback URL field. You can find this URL on the connected apps Manage Connected Apps page or from the connected apps definition. This value must be URL encoded. | [optional] 
**ConcurrentApiRequestsLimit** | **NullableInt64** | Specifies the maximum number of concurrent API requests allowed for salesforce. | 
**ConsumerKey** | **NullableString** | Specifies Consumer key from the connected app in SFDC. | 
**ConsumerSecret** | **NullableString** | Specifies Consumer secret from the connected app in SFDC. | 
**DailyApiLimit** | **NullableInt64** | Specifies the maximum number of daily API requests allowed for salesforce. | 
**Endpoint** | **NullableString** | Specifies the SFDC endpoint URL. | 
**EndpointType** | **NullableString** | SFDC Endpoint type. | 
**MetadataEndpointUrl** | Pointer to **NullableString** | Specifies the url to access salesforce metadata requests. | [optional] [readonly] 
**Password** | Pointer to **NullableString** | Specifies the password to access salesforce. | [optional] 
**SoapEndpointUrl** | Pointer to **NullableString** | Specifies the url to access salesforce soap requests. | [optional] [readonly] 
**Username** | Pointer to **NullableString** | Specifies the username to access salesforce. | [optional] 

## Methods

### NewSfdcSourceRegistrationParams

`func NewSfdcSourceRegistrationParams(authToken NullableString, concurrentApiRequestsLimit NullableInt64, consumerKey NullableString, consumerSecret NullableString, dailyApiLimit NullableInt64, endpoint NullableString, endpointType NullableString, ) *SfdcSourceRegistrationParams`

NewSfdcSourceRegistrationParams instantiates a new SfdcSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcSourceRegistrationParamsWithDefaults

`func NewSfdcSourceRegistrationParamsWithDefaults() *SfdcSourceRegistrationParams`

NewSfdcSourceRegistrationParamsWithDefaults instantiates a new SfdcSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *SfdcSourceRegistrationParams) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *SfdcSourceRegistrationParams) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *SfdcSourceRegistrationParams) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.


### SetAuthTokenNil

`func (o *SfdcSourceRegistrationParams) SetAuthTokenNil(b bool)`

 SetAuthTokenNil sets the value for AuthToken to be an explicit nil

### UnsetAuthToken
`func (o *SfdcSourceRegistrationParams) UnsetAuthToken()`

UnsetAuthToken ensures that no value is present for AuthToken, not even an explicit nil
### GetCallbackUrl

`func (o *SfdcSourceRegistrationParams) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *SfdcSourceRegistrationParams) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *SfdcSourceRegistrationParams) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *SfdcSourceRegistrationParams) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *SfdcSourceRegistrationParams) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *SfdcSourceRegistrationParams) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetConcurrentApiRequestsLimit

`func (o *SfdcSourceRegistrationParams) GetConcurrentApiRequestsLimit() int64`

GetConcurrentApiRequestsLimit returns the ConcurrentApiRequestsLimit field if non-nil, zero value otherwise.

### GetConcurrentApiRequestsLimitOk

`func (o *SfdcSourceRegistrationParams) GetConcurrentApiRequestsLimitOk() (*int64, bool)`

GetConcurrentApiRequestsLimitOk returns a tuple with the ConcurrentApiRequestsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentApiRequestsLimit

`func (o *SfdcSourceRegistrationParams) SetConcurrentApiRequestsLimit(v int64)`

SetConcurrentApiRequestsLimit sets ConcurrentApiRequestsLimit field to given value.


### SetConcurrentApiRequestsLimitNil

`func (o *SfdcSourceRegistrationParams) SetConcurrentApiRequestsLimitNil(b bool)`

 SetConcurrentApiRequestsLimitNil sets the value for ConcurrentApiRequestsLimit to be an explicit nil

### UnsetConcurrentApiRequestsLimit
`func (o *SfdcSourceRegistrationParams) UnsetConcurrentApiRequestsLimit()`

UnsetConcurrentApiRequestsLimit ensures that no value is present for ConcurrentApiRequestsLimit, not even an explicit nil
### GetConsumerKey

`func (o *SfdcSourceRegistrationParams) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *SfdcSourceRegistrationParams) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *SfdcSourceRegistrationParams) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### SetConsumerKeyNil

`func (o *SfdcSourceRegistrationParams) SetConsumerKeyNil(b bool)`

 SetConsumerKeyNil sets the value for ConsumerKey to be an explicit nil

### UnsetConsumerKey
`func (o *SfdcSourceRegistrationParams) UnsetConsumerKey()`

UnsetConsumerKey ensures that no value is present for ConsumerKey, not even an explicit nil
### GetConsumerSecret

`func (o *SfdcSourceRegistrationParams) GetConsumerSecret() string`

GetConsumerSecret returns the ConsumerSecret field if non-nil, zero value otherwise.

### GetConsumerSecretOk

`func (o *SfdcSourceRegistrationParams) GetConsumerSecretOk() (*string, bool)`

GetConsumerSecretOk returns a tuple with the ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSecret

`func (o *SfdcSourceRegistrationParams) SetConsumerSecret(v string)`

SetConsumerSecret sets ConsumerSecret field to given value.


### SetConsumerSecretNil

`func (o *SfdcSourceRegistrationParams) SetConsumerSecretNil(b bool)`

 SetConsumerSecretNil sets the value for ConsumerSecret to be an explicit nil

### UnsetConsumerSecret
`func (o *SfdcSourceRegistrationParams) UnsetConsumerSecret()`

UnsetConsumerSecret ensures that no value is present for ConsumerSecret, not even an explicit nil
### GetDailyApiLimit

`func (o *SfdcSourceRegistrationParams) GetDailyApiLimit() int64`

GetDailyApiLimit returns the DailyApiLimit field if non-nil, zero value otherwise.

### GetDailyApiLimitOk

`func (o *SfdcSourceRegistrationParams) GetDailyApiLimitOk() (*int64, bool)`

GetDailyApiLimitOk returns a tuple with the DailyApiLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyApiLimit

`func (o *SfdcSourceRegistrationParams) SetDailyApiLimit(v int64)`

SetDailyApiLimit sets DailyApiLimit field to given value.


### SetDailyApiLimitNil

`func (o *SfdcSourceRegistrationParams) SetDailyApiLimitNil(b bool)`

 SetDailyApiLimitNil sets the value for DailyApiLimit to be an explicit nil

### UnsetDailyApiLimit
`func (o *SfdcSourceRegistrationParams) UnsetDailyApiLimit()`

UnsetDailyApiLimit ensures that no value is present for DailyApiLimit, not even an explicit nil
### GetEndpoint

`func (o *SfdcSourceRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SfdcSourceRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SfdcSourceRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *SfdcSourceRegistrationParams) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *SfdcSourceRegistrationParams) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetEndpointType

`func (o *SfdcSourceRegistrationParams) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *SfdcSourceRegistrationParams) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *SfdcSourceRegistrationParams) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.


### SetEndpointTypeNil

`func (o *SfdcSourceRegistrationParams) SetEndpointTypeNil(b bool)`

 SetEndpointTypeNil sets the value for EndpointType to be an explicit nil

### UnsetEndpointType
`func (o *SfdcSourceRegistrationParams) UnsetEndpointType()`

UnsetEndpointType ensures that no value is present for EndpointType, not even an explicit nil
### GetMetadataEndpointUrl

`func (o *SfdcSourceRegistrationParams) GetMetadataEndpointUrl() string`

GetMetadataEndpointUrl returns the MetadataEndpointUrl field if non-nil, zero value otherwise.

### GetMetadataEndpointUrlOk

`func (o *SfdcSourceRegistrationParams) GetMetadataEndpointUrlOk() (*string, bool)`

GetMetadataEndpointUrlOk returns a tuple with the MetadataEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataEndpointUrl

`func (o *SfdcSourceRegistrationParams) SetMetadataEndpointUrl(v string)`

SetMetadataEndpointUrl sets MetadataEndpointUrl field to given value.

### HasMetadataEndpointUrl

`func (o *SfdcSourceRegistrationParams) HasMetadataEndpointUrl() bool`

HasMetadataEndpointUrl returns a boolean if a field has been set.

### SetMetadataEndpointUrlNil

`func (o *SfdcSourceRegistrationParams) SetMetadataEndpointUrlNil(b bool)`

 SetMetadataEndpointUrlNil sets the value for MetadataEndpointUrl to be an explicit nil

### UnsetMetadataEndpointUrl
`func (o *SfdcSourceRegistrationParams) UnsetMetadataEndpointUrl()`

UnsetMetadataEndpointUrl ensures that no value is present for MetadataEndpointUrl, not even an explicit nil
### GetPassword

`func (o *SfdcSourceRegistrationParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SfdcSourceRegistrationParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SfdcSourceRegistrationParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SfdcSourceRegistrationParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SfdcSourceRegistrationParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SfdcSourceRegistrationParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSoapEndpointUrl

`func (o *SfdcSourceRegistrationParams) GetSoapEndpointUrl() string`

GetSoapEndpointUrl returns the SoapEndpointUrl field if non-nil, zero value otherwise.

### GetSoapEndpointUrlOk

`func (o *SfdcSourceRegistrationParams) GetSoapEndpointUrlOk() (*string, bool)`

GetSoapEndpointUrlOk returns a tuple with the SoapEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoapEndpointUrl

`func (o *SfdcSourceRegistrationParams) SetSoapEndpointUrl(v string)`

SetSoapEndpointUrl sets SoapEndpointUrl field to given value.

### HasSoapEndpointUrl

`func (o *SfdcSourceRegistrationParams) HasSoapEndpointUrl() bool`

HasSoapEndpointUrl returns a boolean if a field has been set.

### SetSoapEndpointUrlNil

`func (o *SfdcSourceRegistrationParams) SetSoapEndpointUrlNil(b bool)`

 SetSoapEndpointUrlNil sets the value for SoapEndpointUrl to be an explicit nil

### UnsetSoapEndpointUrl
`func (o *SfdcSourceRegistrationParams) UnsetSoapEndpointUrl()`

UnsetSoapEndpointUrl ensures that no value is present for SoapEndpointUrl, not even an explicit nil
### GetUsername

`func (o *SfdcSourceRegistrationParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SfdcSourceRegistrationParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SfdcSourceRegistrationParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SfdcSourceRegistrationParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *SfdcSourceRegistrationParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SfdcSourceRegistrationParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


