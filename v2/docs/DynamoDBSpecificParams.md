# DynamoDBSpecificParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3URI** | Pointer to **NullableString** | Specifies the s3 bucket URI which is used for import and export during backup and recovery. | [optional] 

## Methods

### NewDynamoDBSpecificParams

`func NewDynamoDBSpecificParams() *DynamoDBSpecificParams`

NewDynamoDBSpecificParams instantiates a new DynamoDBSpecificParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamoDBSpecificParamsWithDefaults

`func NewDynamoDBSpecificParamsWithDefaults() *DynamoDBSpecificParams`

NewDynamoDBSpecificParamsWithDefaults instantiates a new DynamoDBSpecificParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3URI

`func (o *DynamoDBSpecificParams) GetS3URI() string`

GetS3URI returns the S3URI field if non-nil, zero value otherwise.

### GetS3URIOk

`func (o *DynamoDBSpecificParams) GetS3URIOk() (*string, bool)`

GetS3URIOk returns a tuple with the S3URI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3URI

`func (o *DynamoDBSpecificParams) SetS3URI(v string)`

SetS3URI sets S3URI field to given value.

### HasS3URI

`func (o *DynamoDBSpecificParams) HasS3URI() bool`

HasS3URI returns a boolean if a field has been set.

### SetS3URINil

`func (o *DynamoDBSpecificParams) SetS3URINil(b bool)`

 SetS3URINil sets the value for S3URI to be an explicit nil

### UnsetS3URI
`func (o *DynamoDBSpecificParams) UnsetS3URI()`

UnsetS3URI ensures that no value is present for S3URI, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


