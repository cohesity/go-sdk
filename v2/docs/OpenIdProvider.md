# OpenIdProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceIds** | **[]string** | Specifies the audience IDs of the configuration. This is used for validation. We will check this against the &#39;aud&#39; field sent in the JWT at authorization time and if they do not match against at least one of the elements in this list, then authentication will fail. | 
**PollingFrequencyMins** | Pointer to **NullableInt64** | Specifies the number of minutes the cluster should wait before polling for a new public key. Default value is 1440 (24 hours). | [optional] [default to 1440]
**PublicKeyUrl** | **NullableString** | Specifies the URL to poll for the public key. | 

## Methods

### NewOpenIdProvider

`func NewOpenIdProvider(audienceIds []string, publicKeyUrl NullableString, ) *OpenIdProvider`

NewOpenIdProvider instantiates a new OpenIdProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdProviderWithDefaults

`func NewOpenIdProviderWithDefaults() *OpenIdProvider`

NewOpenIdProviderWithDefaults instantiates a new OpenIdProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceIds

`func (o *OpenIdProvider) GetAudienceIds() []string`

GetAudienceIds returns the AudienceIds field if non-nil, zero value otherwise.

### GetAudienceIdsOk

`func (o *OpenIdProvider) GetAudienceIdsOk() (*[]string, bool)`

GetAudienceIdsOk returns a tuple with the AudienceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceIds

`func (o *OpenIdProvider) SetAudienceIds(v []string)`

SetAudienceIds sets AudienceIds field to given value.


### SetAudienceIdsNil

`func (o *OpenIdProvider) SetAudienceIdsNil(b bool)`

 SetAudienceIdsNil sets the value for AudienceIds to be an explicit nil

### UnsetAudienceIds
`func (o *OpenIdProvider) UnsetAudienceIds()`

UnsetAudienceIds ensures that no value is present for AudienceIds, not even an explicit nil
### GetPollingFrequencyMins

`func (o *OpenIdProvider) GetPollingFrequencyMins() int64`

GetPollingFrequencyMins returns the PollingFrequencyMins field if non-nil, zero value otherwise.

### GetPollingFrequencyMinsOk

`func (o *OpenIdProvider) GetPollingFrequencyMinsOk() (*int64, bool)`

GetPollingFrequencyMinsOk returns a tuple with the PollingFrequencyMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingFrequencyMins

`func (o *OpenIdProvider) SetPollingFrequencyMins(v int64)`

SetPollingFrequencyMins sets PollingFrequencyMins field to given value.

### HasPollingFrequencyMins

`func (o *OpenIdProvider) HasPollingFrequencyMins() bool`

HasPollingFrequencyMins returns a boolean if a field has been set.

### SetPollingFrequencyMinsNil

`func (o *OpenIdProvider) SetPollingFrequencyMinsNil(b bool)`

 SetPollingFrequencyMinsNil sets the value for PollingFrequencyMins to be an explicit nil

### UnsetPollingFrequencyMins
`func (o *OpenIdProvider) UnsetPollingFrequencyMins()`

UnsetPollingFrequencyMins ensures that no value is present for PollingFrequencyMins, not even an explicit nil
### GetPublicKeyUrl

`func (o *OpenIdProvider) GetPublicKeyUrl() string`

GetPublicKeyUrl returns the PublicKeyUrl field if non-nil, zero value otherwise.

### GetPublicKeyUrlOk

`func (o *OpenIdProvider) GetPublicKeyUrlOk() (*string, bool)`

GetPublicKeyUrlOk returns a tuple with the PublicKeyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyUrl

`func (o *OpenIdProvider) SetPublicKeyUrl(v string)`

SetPublicKeyUrl sets PublicKeyUrl field to given value.


### SetPublicKeyUrlNil

`func (o *OpenIdProvider) SetPublicKeyUrlNil(b bool)`

 SetPublicKeyUrlNil sets the value for PublicKeyUrl to be an explicit nil

### UnsetPublicKeyUrl
`func (o *OpenIdProvider) UnsetPublicKeyUrl()`

UnsetPublicKeyUrl ensures that no value is present for PublicKeyUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


