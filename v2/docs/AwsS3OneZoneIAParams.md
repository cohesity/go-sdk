# AwsS3OneZoneIAParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **NullableString** | Specifies the AWS External Target type. | 
**AwsCloudGovParams** | Pointer to [**AwsCloudGovParams**](AwsCloudGovParams.md) |  | [optional] 
**AwsCloudStandardParams** | Pointer to [**AwsCloudStandardParams**](AwsCloudStandardParams.md) |  | [optional] 

## Methods

### NewAwsS3OneZoneIAParams

`func NewAwsS3OneZoneIAParams(cloudType NullableString, ) *AwsS3OneZoneIAParams`

NewAwsS3OneZoneIAParams instantiates a new AwsS3OneZoneIAParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsS3OneZoneIAParamsWithDefaults

`func NewAwsS3OneZoneIAParamsWithDefaults() *AwsS3OneZoneIAParams`

NewAwsS3OneZoneIAParamsWithDefaults instantiates a new AwsS3OneZoneIAParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *AwsS3OneZoneIAParams) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *AwsS3OneZoneIAParams) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *AwsS3OneZoneIAParams) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.


### SetCloudTypeNil

`func (o *AwsS3OneZoneIAParams) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *AwsS3OneZoneIAParams) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
### GetAwsCloudGovParams

`func (o *AwsS3OneZoneIAParams) GetAwsCloudGovParams() AwsCloudGovParams`

GetAwsCloudGovParams returns the AwsCloudGovParams field if non-nil, zero value otherwise.

### GetAwsCloudGovParamsOk

`func (o *AwsS3OneZoneIAParams) GetAwsCloudGovParamsOk() (*AwsCloudGovParams, bool)`

GetAwsCloudGovParamsOk returns a tuple with the AwsCloudGovParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudGovParams

`func (o *AwsS3OneZoneIAParams) SetAwsCloudGovParams(v AwsCloudGovParams)`

SetAwsCloudGovParams sets AwsCloudGovParams field to given value.

### HasAwsCloudGovParams

`func (o *AwsS3OneZoneIAParams) HasAwsCloudGovParams() bool`

HasAwsCloudGovParams returns a boolean if a field has been set.

### GetAwsCloudStandardParams

`func (o *AwsS3OneZoneIAParams) GetAwsCloudStandardParams() AwsCloudStandardParams`

GetAwsCloudStandardParams returns the AwsCloudStandardParams field if non-nil, zero value otherwise.

### GetAwsCloudStandardParamsOk

`func (o *AwsS3OneZoneIAParams) GetAwsCloudStandardParamsOk() (*AwsCloudStandardParams, bool)`

GetAwsCloudStandardParamsOk returns a tuple with the AwsCloudStandardParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudStandardParams

`func (o *AwsS3OneZoneIAParams) SetAwsCloudStandardParams(v AwsCloudStandardParams)`

SetAwsCloudStandardParams sets AwsCloudStandardParams field to given value.

### HasAwsCloudStandardParams

`func (o *AwsS3OneZoneIAParams) HasAwsCloudStandardParams() bool`

HasAwsCloudStandardParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


