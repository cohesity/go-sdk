# AzureSqlPackageOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compression** | Pointer to **NullableString** | Specifies the compression option supported by SQL package export command during Azure SQL backup. | [optional] 

## Methods

### NewAzureSqlPackageOptions

`func NewAzureSqlPackageOptions() *AzureSqlPackageOptions`

NewAzureSqlPackageOptions instantiates a new AzureSqlPackageOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSqlPackageOptionsWithDefaults

`func NewAzureSqlPackageOptionsWithDefaults() *AzureSqlPackageOptions`

NewAzureSqlPackageOptionsWithDefaults instantiates a new AzureSqlPackageOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompression

`func (o *AzureSqlPackageOptions) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *AzureSqlPackageOptions) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *AzureSqlPackageOptions) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *AzureSqlPackageOptions) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### SetCompressionNil

`func (o *AzureSqlPackageOptions) SetCompressionNil(b bool)`

 SetCompressionNil sets the value for Compression to be an explicit nil

### UnsetCompression
`func (o *AzureSqlPackageOptions) UnsetCompression()`

UnsetCompression ensures that no value is present for Compression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


