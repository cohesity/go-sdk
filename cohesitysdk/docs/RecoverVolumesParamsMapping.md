# RecoverVolumesParamsMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DstGuid** | Pointer to **NullableString** | The destination, pertains to the newly rebuilt system. | [optional] 
**SrcGuid** | Pointer to **NullableString** | The source, pertains to the original backup. | [optional] 

## Methods

### NewRecoverVolumesParamsMapping

`func NewRecoverVolumesParamsMapping() *RecoverVolumesParamsMapping`

NewRecoverVolumesParamsMapping instantiates a new RecoverVolumesParamsMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVolumesParamsMappingWithDefaults

`func NewRecoverVolumesParamsMappingWithDefaults() *RecoverVolumesParamsMapping`

NewRecoverVolumesParamsMappingWithDefaults instantiates a new RecoverVolumesParamsMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDstGuid

`func (o *RecoverVolumesParamsMapping) GetDstGuid() string`

GetDstGuid returns the DstGuid field if non-nil, zero value otherwise.

### GetDstGuidOk

`func (o *RecoverVolumesParamsMapping) GetDstGuidOk() (*string, bool)`

GetDstGuidOk returns a tuple with the DstGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstGuid

`func (o *RecoverVolumesParamsMapping) SetDstGuid(v string)`

SetDstGuid sets DstGuid field to given value.

### HasDstGuid

`func (o *RecoverVolumesParamsMapping) HasDstGuid() bool`

HasDstGuid returns a boolean if a field has been set.

### SetDstGuidNil

`func (o *RecoverVolumesParamsMapping) SetDstGuidNil(b bool)`

 SetDstGuidNil sets the value for DstGuid to be an explicit nil

### UnsetDstGuid
`func (o *RecoverVolumesParamsMapping) UnsetDstGuid()`

UnsetDstGuid ensures that no value is present for DstGuid, not even an explicit nil
### GetSrcGuid

`func (o *RecoverVolumesParamsMapping) GetSrcGuid() string`

GetSrcGuid returns the SrcGuid field if non-nil, zero value otherwise.

### GetSrcGuidOk

`func (o *RecoverVolumesParamsMapping) GetSrcGuidOk() (*string, bool)`

GetSrcGuidOk returns a tuple with the SrcGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcGuid

`func (o *RecoverVolumesParamsMapping) SetSrcGuid(v string)`

SetSrcGuid sets SrcGuid field to given value.

### HasSrcGuid

`func (o *RecoverVolumesParamsMapping) HasSrcGuid() bool`

HasSrcGuid returns a boolean if a field has been set.

### SetSrcGuidNil

`func (o *RecoverVolumesParamsMapping) SetSrcGuidNil(b bool)`

 SetSrcGuidNil sets the value for SrcGuid to be an explicit nil

### UnsetSrcGuid
`func (o *RecoverVolumesParamsMapping) UnsetSrcGuid()`

UnsetSrcGuid ensures that no value is present for SrcGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


