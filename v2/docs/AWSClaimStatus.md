# AWSClaimStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **NullableString** | Specifies the error, if any, that occurred during appliance registration to AWS. | [optional] 
**IsRegistered** | Pointer to **NullableBool** | Specifies whether the appliance is registered to AWS | [optional] 

## Methods

### NewAWSClaimStatus

`func NewAWSClaimStatus() *AWSClaimStatus`

NewAWSClaimStatus instantiates a new AWSClaimStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSClaimStatusWithDefaults

`func NewAWSClaimStatusWithDefaults() *AWSClaimStatus`

NewAWSClaimStatusWithDefaults instantiates a new AWSClaimStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AWSClaimStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AWSClaimStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AWSClaimStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AWSClaimStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AWSClaimStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AWSClaimStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetIsRegistered

`func (o *AWSClaimStatus) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *AWSClaimStatus) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *AWSClaimStatus) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.

### HasIsRegistered

`func (o *AWSClaimStatus) HasIsRegistered() bool`

HasIsRegistered returns a boolean if a field has been set.

### SetIsRegisteredNil

`func (o *AWSClaimStatus) SetIsRegisteredNil(b bool)`

 SetIsRegisteredNil sets the value for IsRegistered to be an explicit nil

### UnsetIsRegistered
`func (o *AWSClaimStatus) UnsetIsRegistered()`

UnsetIsRegistered ensures that no value is present for IsRegistered, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


