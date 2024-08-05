# AwsGlacierParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **NullableString** | Specifies the AWS External Target type. | 
**AwsCloudC2SParams** | Pointer to [**AwsCloudC2SParams**](AwsCloudC2SParams.md) |  | [optional] 
**AwsCloudGovParams** | Pointer to [**AwsCloudGovParams**](AwsCloudGovParams.md) |  | [optional] 
**AwsCloudStandardParams** | Pointer to [**AwsCloudStandardParams**](AwsCloudStandardParams.md) |  | [optional] 

## Methods

### NewAwsGlacierParams

`func NewAwsGlacierParams(cloudType NullableString, ) *AwsGlacierParams`

NewAwsGlacierParams instantiates a new AwsGlacierParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsGlacierParamsWithDefaults

`func NewAwsGlacierParamsWithDefaults() *AwsGlacierParams`

NewAwsGlacierParamsWithDefaults instantiates a new AwsGlacierParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *AwsGlacierParams) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *AwsGlacierParams) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *AwsGlacierParams) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.


### SetCloudTypeNil

`func (o *AwsGlacierParams) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *AwsGlacierParams) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
### GetAwsCloudC2SParams

`func (o *AwsGlacierParams) GetAwsCloudC2SParams() AwsCloudC2SParams`

GetAwsCloudC2SParams returns the AwsCloudC2SParams field if non-nil, zero value otherwise.

### GetAwsCloudC2SParamsOk

`func (o *AwsGlacierParams) GetAwsCloudC2SParamsOk() (*AwsCloudC2SParams, bool)`

GetAwsCloudC2SParamsOk returns a tuple with the AwsCloudC2SParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudC2SParams

`func (o *AwsGlacierParams) SetAwsCloudC2SParams(v AwsCloudC2SParams)`

SetAwsCloudC2SParams sets AwsCloudC2SParams field to given value.

### HasAwsCloudC2SParams

`func (o *AwsGlacierParams) HasAwsCloudC2SParams() bool`

HasAwsCloudC2SParams returns a boolean if a field has been set.

### GetAwsCloudGovParams

`func (o *AwsGlacierParams) GetAwsCloudGovParams() AwsCloudGovParams`

GetAwsCloudGovParams returns the AwsCloudGovParams field if non-nil, zero value otherwise.

### GetAwsCloudGovParamsOk

`func (o *AwsGlacierParams) GetAwsCloudGovParamsOk() (*AwsCloudGovParams, bool)`

GetAwsCloudGovParamsOk returns a tuple with the AwsCloudGovParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudGovParams

`func (o *AwsGlacierParams) SetAwsCloudGovParams(v AwsCloudGovParams)`

SetAwsCloudGovParams sets AwsCloudGovParams field to given value.

### HasAwsCloudGovParams

`func (o *AwsGlacierParams) HasAwsCloudGovParams() bool`

HasAwsCloudGovParams returns a boolean if a field has been set.

### GetAwsCloudStandardParams

`func (o *AwsGlacierParams) GetAwsCloudStandardParams() AwsCloudStandardParams`

GetAwsCloudStandardParams returns the AwsCloudStandardParams field if non-nil, zero value otherwise.

### GetAwsCloudStandardParamsOk

`func (o *AwsGlacierParams) GetAwsCloudStandardParamsOk() (*AwsCloudStandardParams, bool)`

GetAwsCloudStandardParamsOk returns a tuple with the AwsCloudStandardParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudStandardParams

`func (o *AwsGlacierParams) SetAwsCloudStandardParams(v AwsCloudStandardParams)`

SetAwsCloudStandardParams sets AwsCloudStandardParams field to given value.

### HasAwsCloudStandardParams

`func (o *AwsGlacierParams) HasAwsCloudStandardParams() bool`

HasAwsCloudStandardParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


