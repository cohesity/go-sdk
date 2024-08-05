# AwsEntityMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuroraParams** | Pointer to [**AwsAuroraEntityMetadata**](AwsAuroraEntityMetadata.md) |  | [optional] 
**ChildMetadata** | Pointer to [**AwsEntityChildMetadata**](AwsEntityChildMetadata.md) |  | [optional] 
**PostgresParams** | Pointer to [**AwsPostgresEntityMetadata**](AwsPostgresEntityMetadata.md) |  | [optional] 
**RdsParams** | Pointer to [**AwsRdsEntityMetadata**](AwsRdsEntityMetadata.md) |  | [optional] 

## Methods

### NewAwsEntityMetadata

`func NewAwsEntityMetadata() *AwsEntityMetadata`

NewAwsEntityMetadata instantiates a new AwsEntityMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEntityMetadataWithDefaults

`func NewAwsEntityMetadataWithDefaults() *AwsEntityMetadata`

NewAwsEntityMetadataWithDefaults instantiates a new AwsEntityMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuroraParams

`func (o *AwsEntityMetadata) GetAuroraParams() AwsAuroraEntityMetadata`

GetAuroraParams returns the AuroraParams field if non-nil, zero value otherwise.

### GetAuroraParamsOk

`func (o *AwsEntityMetadata) GetAuroraParamsOk() (*AwsAuroraEntityMetadata, bool)`

GetAuroraParamsOk returns a tuple with the AuroraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuroraParams

`func (o *AwsEntityMetadata) SetAuroraParams(v AwsAuroraEntityMetadata)`

SetAuroraParams sets AuroraParams field to given value.

### HasAuroraParams

`func (o *AwsEntityMetadata) HasAuroraParams() bool`

HasAuroraParams returns a boolean if a field has been set.

### GetChildMetadata

`func (o *AwsEntityMetadata) GetChildMetadata() AwsEntityChildMetadata`

GetChildMetadata returns the ChildMetadata field if non-nil, zero value otherwise.

### GetChildMetadataOk

`func (o *AwsEntityMetadata) GetChildMetadataOk() (*AwsEntityChildMetadata, bool)`

GetChildMetadataOk returns a tuple with the ChildMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildMetadata

`func (o *AwsEntityMetadata) SetChildMetadata(v AwsEntityChildMetadata)`

SetChildMetadata sets ChildMetadata field to given value.

### HasChildMetadata

`func (o *AwsEntityMetadata) HasChildMetadata() bool`

HasChildMetadata returns a boolean if a field has been set.

### GetPostgresParams

`func (o *AwsEntityMetadata) GetPostgresParams() AwsPostgresEntityMetadata`

GetPostgresParams returns the PostgresParams field if non-nil, zero value otherwise.

### GetPostgresParamsOk

`func (o *AwsEntityMetadata) GetPostgresParamsOk() (*AwsPostgresEntityMetadata, bool)`

GetPostgresParamsOk returns a tuple with the PostgresParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresParams

`func (o *AwsEntityMetadata) SetPostgresParams(v AwsPostgresEntityMetadata)`

SetPostgresParams sets PostgresParams field to given value.

### HasPostgresParams

`func (o *AwsEntityMetadata) HasPostgresParams() bool`

HasPostgresParams returns a boolean if a field has been set.

### GetRdsParams

`func (o *AwsEntityMetadata) GetRdsParams() AwsRdsEntityMetadata`

GetRdsParams returns the RdsParams field if non-nil, zero value otherwise.

### GetRdsParamsOk

`func (o *AwsEntityMetadata) GetRdsParamsOk() (*AwsRdsEntityMetadata, bool)`

GetRdsParamsOk returns a tuple with the RdsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsParams

`func (o *AwsEntityMetadata) SetRdsParams(v AwsRdsEntityMetadata)`

SetRdsParams sets RdsParams field to given value.

### HasRdsParams

`func (o *AwsEntityMetadata) HasRdsParams() bool`

HasRdsParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


