# AWSClaimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationURL** | **string** | The registration URL including the registration token, as obtained from AWS. | 

## Methods

### NewAWSClaimRequest

`func NewAWSClaimRequest(registrationURL string, ) *AWSClaimRequest`

NewAWSClaimRequest instantiates a new AWSClaimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSClaimRequestWithDefaults

`func NewAWSClaimRequestWithDefaults() *AWSClaimRequest`

NewAWSClaimRequestWithDefaults instantiates a new AWSClaimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationURL

`func (o *AWSClaimRequest) GetRegistrationURL() string`

GetRegistrationURL returns the RegistrationURL field if non-nil, zero value otherwise.

### GetRegistrationURLOk

`func (o *AWSClaimRequest) GetRegistrationURLOk() (*string, bool)`

GetRegistrationURLOk returns a tuple with the RegistrationURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationURL

`func (o *AWSClaimRequest) SetRegistrationURL(v string)`

SetRegistrationURL sets RegistrationURL field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


