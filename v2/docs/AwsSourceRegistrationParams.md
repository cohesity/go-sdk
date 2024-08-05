# AwsSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DynamoDBParams** | Pointer to [**DynamoDBSpecificParams**](DynamoDBSpecificParams.md) |  | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the dataplane region of cluster. | [optional] 
**S3Params** | Pointer to [**S3SpecificParams**](S3SpecificParams.md) |  | [optional] 
**StandardParams** | Pointer to [**StandardParams**](StandardParams.md) |  | [optional] 
**SubscriptionType** | **NullableString** | Specifies the AWS Subscription type (Commercial/Gov). | 
**UseCases** | Pointer to **[]string** | The use cases for which the source is to be registered. | [optional] 

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

### GetDynamoDBParams

`func (o *AwsSourceRegistrationParams) GetDynamoDBParams() DynamoDBSpecificParams`

GetDynamoDBParams returns the DynamoDBParams field if non-nil, zero value otherwise.

### GetDynamoDBParamsOk

`func (o *AwsSourceRegistrationParams) GetDynamoDBParamsOk() (*DynamoDBSpecificParams, bool)`

GetDynamoDBParamsOk returns a tuple with the DynamoDBParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamoDBParams

`func (o *AwsSourceRegistrationParams) SetDynamoDBParams(v DynamoDBSpecificParams)`

SetDynamoDBParams sets DynamoDBParams field to given value.

### HasDynamoDBParams

`func (o *AwsSourceRegistrationParams) HasDynamoDBParams() bool`

HasDynamoDBParams returns a boolean if a field has been set.

### GetRegionId

`func (o *AwsSourceRegistrationParams) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AwsSourceRegistrationParams) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AwsSourceRegistrationParams) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *AwsSourceRegistrationParams) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *AwsSourceRegistrationParams) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *AwsSourceRegistrationParams) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetS3Params

`func (o *AwsSourceRegistrationParams) GetS3Params() S3SpecificParams`

GetS3Params returns the S3Params field if non-nil, zero value otherwise.

### GetS3ParamsOk

`func (o *AwsSourceRegistrationParams) GetS3ParamsOk() (*S3SpecificParams, bool)`

GetS3ParamsOk returns a tuple with the S3Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Params

`func (o *AwsSourceRegistrationParams) SetS3Params(v S3SpecificParams)`

SetS3Params sets S3Params field to given value.

### HasS3Params

`func (o *AwsSourceRegistrationParams) HasS3Params() bool`

HasS3Params returns a boolean if a field has been set.

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
### GetUseCases

`func (o *AwsSourceRegistrationParams) GetUseCases() []string`

GetUseCases returns the UseCases field if non-nil, zero value otherwise.

### GetUseCasesOk

`func (o *AwsSourceRegistrationParams) GetUseCasesOk() (*[]string, bool)`

GetUseCasesOk returns a tuple with the UseCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCases

`func (o *AwsSourceRegistrationParams) SetUseCases(v []string)`

SetUseCases sets UseCases field to given value.

### HasUseCases

`func (o *AwsSourceRegistrationParams) HasUseCases() bool`

HasUseCases returns a boolean if a field has been set.

### SetUseCasesNil

`func (o *AwsSourceRegistrationParams) SetUseCasesNil(b bool)`

 SetUseCasesNil sets the value for UseCases to be an explicit nil

### UnsetUseCases
`func (o *AwsSourceRegistrationParams) UnsetUseCases()`

UnsetUseCases ensures that no value is present for UseCases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


