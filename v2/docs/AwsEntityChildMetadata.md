# AwsEntityChildMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuroraMetadata** | Pointer to [**AwsAuroraEntityMetadata**](AwsAuroraEntityMetadata.md) |  | [optional] 
**RdsMetadata** | Pointer to [**AwsRdsEntityMetadata**](AwsRdsEntityMetadata.md) |  | [optional] 

## Methods

### NewAwsEntityChildMetadata

`func NewAwsEntityChildMetadata() *AwsEntityChildMetadata`

NewAwsEntityChildMetadata instantiates a new AwsEntityChildMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEntityChildMetadataWithDefaults

`func NewAwsEntityChildMetadataWithDefaults() *AwsEntityChildMetadata`

NewAwsEntityChildMetadataWithDefaults instantiates a new AwsEntityChildMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuroraMetadata

`func (o *AwsEntityChildMetadata) GetAuroraMetadata() AwsAuroraEntityMetadata`

GetAuroraMetadata returns the AuroraMetadata field if non-nil, zero value otherwise.

### GetAuroraMetadataOk

`func (o *AwsEntityChildMetadata) GetAuroraMetadataOk() (*AwsAuroraEntityMetadata, bool)`

GetAuroraMetadataOk returns a tuple with the AuroraMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuroraMetadata

`func (o *AwsEntityChildMetadata) SetAuroraMetadata(v AwsAuroraEntityMetadata)`

SetAuroraMetadata sets AuroraMetadata field to given value.

### HasAuroraMetadata

`func (o *AwsEntityChildMetadata) HasAuroraMetadata() bool`

HasAuroraMetadata returns a boolean if a field has been set.

### GetRdsMetadata

`func (o *AwsEntityChildMetadata) GetRdsMetadata() AwsRdsEntityMetadata`

GetRdsMetadata returns the RdsMetadata field if non-nil, zero value otherwise.

### GetRdsMetadataOk

`func (o *AwsEntityChildMetadata) GetRdsMetadataOk() (*AwsRdsEntityMetadata, bool)`

GetRdsMetadataOk returns a tuple with the RdsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsMetadata

`func (o *AwsEntityChildMetadata) SetRdsMetadata(v AwsRdsEntityMetadata)`

SetRdsMetadata sets RdsMetadata field to given value.

### HasRdsMetadata

`func (o *AwsEntityChildMetadata) HasRdsMetadata() bool`

HasRdsMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


