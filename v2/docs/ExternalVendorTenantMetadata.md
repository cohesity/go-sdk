# ExternalVendorTenantMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IbmTenantMetadataParams** | Pointer to [**IbmTenantMetadataParams**](IbmTenantMetadataParams.md) |  | [optional] 
**Type** | **NullableString** | Specifies the type of the external vendor. The type specific parameters must be specified the provided type. | 

## Methods

### NewExternalVendorTenantMetadata

`func NewExternalVendorTenantMetadata(type_ NullableString, ) *ExternalVendorTenantMetadata`

NewExternalVendorTenantMetadata instantiates a new ExternalVendorTenantMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalVendorTenantMetadataWithDefaults

`func NewExternalVendorTenantMetadataWithDefaults() *ExternalVendorTenantMetadata`

NewExternalVendorTenantMetadataWithDefaults instantiates a new ExternalVendorTenantMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbmTenantMetadataParams

`func (o *ExternalVendorTenantMetadata) GetIbmTenantMetadataParams() IbmTenantMetadataParams`

GetIbmTenantMetadataParams returns the IbmTenantMetadataParams field if non-nil, zero value otherwise.

### GetIbmTenantMetadataParamsOk

`func (o *ExternalVendorTenantMetadata) GetIbmTenantMetadataParamsOk() (*IbmTenantMetadataParams, bool)`

GetIbmTenantMetadataParamsOk returns a tuple with the IbmTenantMetadataParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbmTenantMetadataParams

`func (o *ExternalVendorTenantMetadata) SetIbmTenantMetadataParams(v IbmTenantMetadataParams)`

SetIbmTenantMetadataParams sets IbmTenantMetadataParams field to given value.

### HasIbmTenantMetadataParams

`func (o *ExternalVendorTenantMetadata) HasIbmTenantMetadataParams() bool`

HasIbmTenantMetadataParams returns a boolean if a field has been set.

### GetType

`func (o *ExternalVendorTenantMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalVendorTenantMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalVendorTenantMetadata) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalVendorTenantMetadata) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalVendorTenantMetadata) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


