# AwsSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionType** | **NullableString** | Specifies the AWS Subscription type (Commercial/Gov). | 
**StandardParams** | Pointer to [**StandardParams**](StandardParams.md) |  | [optional] 

## Methods

### NewAwsSourceRegistrationParams

`func NewAwsSourceRegistrationParams(subscriptionType NullableString, ) *AwsSourceRegistrationParams`

NewAwsSourceRegistrationParams instantiates a new AwsSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSourceRegistrationParamsWithDefaults

`func NewAwsSourceRegistrationParamsWithDefaults() *AwsSourceRegistrationParams`

NewAwsSourceRegistrationParamsWithDefaults instantiates a new AwsSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionType

`func (o *AwsSourceRegistrationParams) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *AwsSourceRegistrationParams) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *AwsSourceRegistrationParams) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.


### SetSubscriptionTypeNil

`func (o *AwsSourceRegistrationParams) SetSubscriptionTypeNil(b bool)`

 SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil

### UnsetSubscriptionType
`func (o *AwsSourceRegistrationParams) UnsetSubscriptionType()`

UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
### GetStandardParams

`func (o *AwsSourceRegistrationParams) GetStandardParams() StandardParams`

GetStandardParams returns the StandardParams field if non-nil, zero value otherwise.

### GetStandardParamsOk

`func (o *AwsSourceRegistrationParams) GetStandardParamsOk() (*StandardParams, bool)`

GetStandardParamsOk returns a tuple with the StandardParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardParams

`func (o *AwsSourceRegistrationParams) SetStandardParams(v StandardParams)`

SetStandardParams sets StandardParams field to given value.

### HasStandardParams

`func (o *AwsSourceRegistrationParams) HasStandardParams() bool`

HasStandardParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


