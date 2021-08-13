# AwsS3StandardIAParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **NullableString** | Specifies the AWS External Target type. | 
**AwsCloudStandardParams** | Pointer to [**AwsCloudStandardParams**](AwsCloudStandardParams.md) |  | [optional] 
**AwsCloudGovParams** | Pointer to [**AwsCloudGovParams**](AwsCloudGovParams.md) |  | [optional] 
**AwsCloudC2SParams** | Pointer to [**AwsCloudC2SParams**](AwsCloudC2SParams.md) |  | [optional] 

## Methods

### NewAwsS3StandardIAParams

`func NewAwsS3StandardIAParams(cloudType NullableString, ) *AwsS3StandardIAParams`

NewAwsS3StandardIAParams instantiates a new AwsS3StandardIAParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsS3StandardIAParamsWithDefaults

`func NewAwsS3StandardIAParamsWithDefaults() *AwsS3StandardIAParams`

NewAwsS3StandardIAParamsWithDefaults instantiates a new AwsS3StandardIAParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *AwsS3StandardIAParams) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *AwsS3StandardIAParams) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *AwsS3StandardIAParams) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.


### SetCloudTypeNil

`func (o *AwsS3StandardIAParams) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *AwsS3StandardIAParams) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
### GetAwsCloudStandardParams

`func (o *AwsS3StandardIAParams) GetAwsCloudStandardParams() AwsCloudStandardParams`

GetAwsCloudStandardParams returns the AwsCloudStandardParams field if non-nil, zero value otherwise.

### GetAwsCloudStandardParamsOk

`func (o *AwsS3StandardIAParams) GetAwsCloudStandardParamsOk() (*AwsCloudStandardParams, bool)`

GetAwsCloudStandardParamsOk returns a tuple with the AwsCloudStandardParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudStandardParams

`func (o *AwsS3StandardIAParams) SetAwsCloudStandardParams(v AwsCloudStandardParams)`

SetAwsCloudStandardParams sets AwsCloudStandardParams field to given value.

### HasAwsCloudStandardParams

`func (o *AwsS3StandardIAParams) HasAwsCloudStandardParams() bool`

HasAwsCloudStandardParams returns a boolean if a field has been set.

### GetAwsCloudGovParams

`func (o *AwsS3StandardIAParams) GetAwsCloudGovParams() AwsCloudGovParams`

GetAwsCloudGovParams returns the AwsCloudGovParams field if non-nil, zero value otherwise.

### GetAwsCloudGovParamsOk

`func (o *AwsS3StandardIAParams) GetAwsCloudGovParamsOk() (*AwsCloudGovParams, bool)`

GetAwsCloudGovParamsOk returns a tuple with the AwsCloudGovParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudGovParams

`func (o *AwsS3StandardIAParams) SetAwsCloudGovParams(v AwsCloudGovParams)`

SetAwsCloudGovParams sets AwsCloudGovParams field to given value.

### HasAwsCloudGovParams

`func (o *AwsS3StandardIAParams) HasAwsCloudGovParams() bool`

HasAwsCloudGovParams returns a boolean if a field has been set.

### GetAwsCloudC2SParams

`func (o *AwsS3StandardIAParams) GetAwsCloudC2SParams() AwsCloudC2SParams`

GetAwsCloudC2SParams returns the AwsCloudC2SParams field if non-nil, zero value otherwise.

### GetAwsCloudC2SParamsOk

`func (o *AwsS3StandardIAParams) GetAwsCloudC2SParamsOk() (*AwsCloudC2SParams, bool)`

GetAwsCloudC2SParamsOk returns a tuple with the AwsCloudC2SParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudC2SParams

`func (o *AwsS3StandardIAParams) SetAwsCloudC2SParams(v AwsCloudC2SParams)`

SetAwsCloudC2SParams sets AwsCloudC2SParams field to given value.

### HasAwsCloudC2SParams

`func (o *AwsS3StandardIAParams) HasAwsCloudC2SParams() bool`

HasAwsCloudC2SParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


