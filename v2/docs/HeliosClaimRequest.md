# HeliosClaimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationToken** | **NullableString** | Specifies the Helios registration token. | 
**RigelGuid** | Pointer to **NullableInt64** | Specifies the rigel guid to be used for registration. | [optional] 

## Methods

### NewHeliosClaimRequest

`func NewHeliosClaimRequest(registrationToken NullableString, ) *HeliosClaimRequest`

NewHeliosClaimRequest instantiates a new HeliosClaimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosClaimRequestWithDefaults

`func NewHeliosClaimRequestWithDefaults() *HeliosClaimRequest`

NewHeliosClaimRequestWithDefaults instantiates a new HeliosClaimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationToken

`func (o *HeliosClaimRequest) GetRegistrationToken() string`

GetRegistrationToken returns the RegistrationToken field if non-nil, zero value otherwise.

### GetRegistrationTokenOk

`func (o *HeliosClaimRequest) GetRegistrationTokenOk() (*string, bool)`

GetRegistrationTokenOk returns a tuple with the RegistrationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationToken

`func (o *HeliosClaimRequest) SetRegistrationToken(v string)`

SetRegistrationToken sets RegistrationToken field to given value.


### SetRegistrationTokenNil

`func (o *HeliosClaimRequest) SetRegistrationTokenNil(b bool)`

 SetRegistrationTokenNil sets the value for RegistrationToken to be an explicit nil

### UnsetRegistrationToken
`func (o *HeliosClaimRequest) UnsetRegistrationToken()`

UnsetRegistrationToken ensures that no value is present for RegistrationToken, not even an explicit nil
### GetRigelGuid

`func (o *HeliosClaimRequest) GetRigelGuid() int64`

GetRigelGuid returns the RigelGuid field if non-nil, zero value otherwise.

### GetRigelGuidOk

`func (o *HeliosClaimRequest) GetRigelGuidOk() (*int64, bool)`

GetRigelGuidOk returns a tuple with the RigelGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelGuid

`func (o *HeliosClaimRequest) SetRigelGuid(v int64)`

SetRigelGuid sets RigelGuid field to given value.

### HasRigelGuid

`func (o *HeliosClaimRequest) HasRigelGuid() bool`

HasRigelGuid returns a boolean if a field has been set.

### SetRigelGuidNil

`func (o *HeliosClaimRequest) SetRigelGuidNil(b bool)`

 SetRigelGuidNil sets the value for RigelGuid to be an explicit nil

### UnsetRigelGuid
`func (o *HeliosClaimRequest) UnsetRigelGuid()`

UnsetRigelGuid ensures that no value is present for RigelGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


