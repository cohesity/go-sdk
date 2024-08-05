# S3SpecificParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryReportBucket** | Pointer to **NullableString** | Specifies the ARN for S3 bucket where inventory reports are to be stored. | [optional] 
**InventoryReportPrefix** | Pointer to **NullableString** | The inventory bucket prefix where inventory reports are to be stored. | [optional] 

## Methods

### NewS3SpecificParams

`func NewS3SpecificParams() *S3SpecificParams`

NewS3SpecificParams instantiates a new S3SpecificParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3SpecificParamsWithDefaults

`func NewS3SpecificParamsWithDefaults() *S3SpecificParams`

NewS3SpecificParamsWithDefaults instantiates a new S3SpecificParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventoryReportBucket

`func (o *S3SpecificParams) GetInventoryReportBucket() string`

GetInventoryReportBucket returns the InventoryReportBucket field if non-nil, zero value otherwise.

### GetInventoryReportBucketOk

`func (o *S3SpecificParams) GetInventoryReportBucketOk() (*string, bool)`

GetInventoryReportBucketOk returns a tuple with the InventoryReportBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReportBucket

`func (o *S3SpecificParams) SetInventoryReportBucket(v string)`

SetInventoryReportBucket sets InventoryReportBucket field to given value.

### HasInventoryReportBucket

`func (o *S3SpecificParams) HasInventoryReportBucket() bool`

HasInventoryReportBucket returns a boolean if a field has been set.

### SetInventoryReportBucketNil

`func (o *S3SpecificParams) SetInventoryReportBucketNil(b bool)`

 SetInventoryReportBucketNil sets the value for InventoryReportBucket to be an explicit nil

### UnsetInventoryReportBucket
`func (o *S3SpecificParams) UnsetInventoryReportBucket()`

UnsetInventoryReportBucket ensures that no value is present for InventoryReportBucket, not even an explicit nil
### GetInventoryReportPrefix

`func (o *S3SpecificParams) GetInventoryReportPrefix() string`

GetInventoryReportPrefix returns the InventoryReportPrefix field if non-nil, zero value otherwise.

### GetInventoryReportPrefixOk

`func (o *S3SpecificParams) GetInventoryReportPrefixOk() (*string, bool)`

GetInventoryReportPrefixOk returns a tuple with the InventoryReportPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryReportPrefix

`func (o *S3SpecificParams) SetInventoryReportPrefix(v string)`

SetInventoryReportPrefix sets InventoryReportPrefix field to given value.

### HasInventoryReportPrefix

`func (o *S3SpecificParams) HasInventoryReportPrefix() bool`

HasInventoryReportPrefix returns a boolean if a field has been set.

### SetInventoryReportPrefixNil

`func (o *S3SpecificParams) SetInventoryReportPrefixNil(b bool)`

 SetInventoryReportPrefixNil sets the value for InventoryReportPrefix to be an explicit nil

### UnsetInventoryReportPrefix
`func (o *S3SpecificParams) UnsetInventoryReportPrefix()`

UnsetInventoryReportPrefix ensures that no value is present for InventoryReportPrefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


