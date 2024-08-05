# VMWareCDPFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsxiVersion** | **NullableString** | Specifies the version of the ESXi host where filter needs to be installed. | 

## Methods

### NewVMWareCDPFilterParams

`func NewVMWareCDPFilterParams(esxiVersion NullableString, ) *VMWareCDPFilterParams`

NewVMWareCDPFilterParams instantiates a new VMWareCDPFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMWareCDPFilterParamsWithDefaults

`func NewVMWareCDPFilterParamsWithDefaults() *VMWareCDPFilterParams`

NewVMWareCDPFilterParamsWithDefaults instantiates a new VMWareCDPFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsxiVersion

`func (o *VMWareCDPFilterParams) GetEsxiVersion() string`

GetEsxiVersion returns the EsxiVersion field if non-nil, zero value otherwise.

### GetEsxiVersionOk

`func (o *VMWareCDPFilterParams) GetEsxiVersionOk() (*string, bool)`

GetEsxiVersionOk returns a tuple with the EsxiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxiVersion

`func (o *VMWareCDPFilterParams) SetEsxiVersion(v string)`

SetEsxiVersion sets EsxiVersion field to given value.


### SetEsxiVersionNil

`func (o *VMWareCDPFilterParams) SetEsxiVersionNil(b bool)`

 SetEsxiVersionNil sets the value for EsxiVersion to be an explicit nil

### UnsetEsxiVersion
`func (o *VMWareCDPFilterParams) UnsetEsxiVersion()`

UnsetEsxiVersion ensures that no value is present for EsxiVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


