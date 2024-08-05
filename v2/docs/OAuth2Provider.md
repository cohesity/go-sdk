# OAuth2Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | [**[]OAuthAudience**](OAuthAudience.md) | Specifies the audiences of the configuration. This is used for validation. We will check this against the &#39;aud&#39; field sent in the JWT at authorization time and if they do not match against at least one of the elements in this list, then authentication will fail. We will also check the &#39;clientIds&#39; under the specified audience to make sure it matches the &#39;appid&#39; in the token. | 
**PollingFrequencyMins** | Pointer to **NullableInt64** | Specifies the number of minutes the cluster should wait before polling for a new public key. Default value is 1440 (24 hours). | [optional] [default to 1440]
**PublicKeyUrl** | **NullableString** | Specifies the URL to poll for the public key. | 

## Methods

### NewOAuth2Provider

`func NewOAuth2Provider(audiences []OAuthAudience, publicKeyUrl NullableString, ) *OAuth2Provider`

NewOAuth2Provider instantiates a new OAuth2Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ProviderWithDefaults

`func NewOAuth2ProviderWithDefaults() *OAuth2Provider`

NewOAuth2ProviderWithDefaults instantiates a new OAuth2Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *OAuth2Provider) GetAudiences() []OAuthAudience`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *OAuth2Provider) GetAudiencesOk() (*[]OAuthAudience, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *OAuth2Provider) SetAudiences(v []OAuthAudience)`

SetAudiences sets Audiences field to given value.


### SetAudiencesNil

`func (o *OAuth2Provider) SetAudiencesNil(b bool)`

 SetAudiencesNil sets the value for Audiences to be an explicit nil

### UnsetAudiences
`func (o *OAuth2Provider) UnsetAudiences()`

UnsetAudiences ensures that no value is present for Audiences, not even an explicit nil
### GetPollingFrequencyMins

`func (o *OAuth2Provider) GetPollingFrequencyMins() int64`

GetPollingFrequencyMins returns the PollingFrequencyMins field if non-nil, zero value otherwise.

### GetPollingFrequencyMinsOk

`func (o *OAuth2Provider) GetPollingFrequencyMinsOk() (*int64, bool)`

GetPollingFrequencyMinsOk returns a tuple with the PollingFrequencyMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingFrequencyMins

`func (o *OAuth2Provider) SetPollingFrequencyMins(v int64)`

SetPollingFrequencyMins sets PollingFrequencyMins field to given value.

### HasPollingFrequencyMins

`func (o *OAuth2Provider) HasPollingFrequencyMins() bool`

HasPollingFrequencyMins returns a boolean if a field has been set.

### SetPollingFrequencyMinsNil

`func (o *OAuth2Provider) SetPollingFrequencyMinsNil(b bool)`

 SetPollingFrequencyMinsNil sets the value for PollingFrequencyMins to be an explicit nil

### UnsetPollingFrequencyMins
`func (o *OAuth2Provider) UnsetPollingFrequencyMins()`

UnsetPollingFrequencyMins ensures that no value is present for PollingFrequencyMins, not even an explicit nil
### GetPublicKeyUrl

`func (o *OAuth2Provider) GetPublicKeyUrl() string`

GetPublicKeyUrl returns the PublicKeyUrl field if non-nil, zero value otherwise.

### GetPublicKeyUrlOk

`func (o *OAuth2Provider) GetPublicKeyUrlOk() (*string, bool)`

GetPublicKeyUrlOk returns a tuple with the PublicKeyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyUrl

`func (o *OAuth2Provider) SetPublicKeyUrl(v string)`

SetPublicKeyUrl sets PublicKeyUrl field to given value.


### SetPublicKeyUrlNil

`func (o *OAuth2Provider) SetPublicKeyUrlNil(b bool)`

 SetPublicKeyUrlNil sets the value for PublicKeyUrl to be an explicit nil

### UnsetPublicKeyUrl
`func (o *OAuth2Provider) UnsetPublicKeyUrl()`

UnsetPublicKeyUrl ensures that no value is present for PublicKeyUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


