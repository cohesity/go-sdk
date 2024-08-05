# StorageDomainCloudDownWaterFallParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdPercentage** | Pointer to **NullableInt32** | Specifies the threshold percentage for cloud down water fall. This value indicates how full a storage domain is before cohesity cluster down water fall its data to cloud tier. This field is only appliciable if physicalQuota is set. | [optional] 
**ThresholdSecs** | Pointer to **NullableInt32** | Specifes a time in seconds when cloud down water fall starts. | [optional] 

## Methods

### NewStorageDomainCloudDownWaterFallParams

`func NewStorageDomainCloudDownWaterFallParams() *StorageDomainCloudDownWaterFallParams`

NewStorageDomainCloudDownWaterFallParams instantiates a new StorageDomainCloudDownWaterFallParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDomainCloudDownWaterFallParamsWithDefaults

`func NewStorageDomainCloudDownWaterFallParamsWithDefaults() *StorageDomainCloudDownWaterFallParams`

NewStorageDomainCloudDownWaterFallParamsWithDefaults instantiates a new StorageDomainCloudDownWaterFallParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdPercentage

`func (o *StorageDomainCloudDownWaterFallParams) GetThresholdPercentage() int32`

GetThresholdPercentage returns the ThresholdPercentage field if non-nil, zero value otherwise.

### GetThresholdPercentageOk

`func (o *StorageDomainCloudDownWaterFallParams) GetThresholdPercentageOk() (*int32, bool)`

GetThresholdPercentageOk returns a tuple with the ThresholdPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercentage

`func (o *StorageDomainCloudDownWaterFallParams) SetThresholdPercentage(v int32)`

SetThresholdPercentage sets ThresholdPercentage field to given value.

### HasThresholdPercentage

`func (o *StorageDomainCloudDownWaterFallParams) HasThresholdPercentage() bool`

HasThresholdPercentage returns a boolean if a field has been set.

### SetThresholdPercentageNil

`func (o *StorageDomainCloudDownWaterFallParams) SetThresholdPercentageNil(b bool)`

 SetThresholdPercentageNil sets the value for ThresholdPercentage to be an explicit nil

### UnsetThresholdPercentage
`func (o *StorageDomainCloudDownWaterFallParams) UnsetThresholdPercentage()`

UnsetThresholdPercentage ensures that no value is present for ThresholdPercentage, not even an explicit nil
### GetThresholdSecs

`func (o *StorageDomainCloudDownWaterFallParams) GetThresholdSecs() int32`

GetThresholdSecs returns the ThresholdSecs field if non-nil, zero value otherwise.

### GetThresholdSecsOk

`func (o *StorageDomainCloudDownWaterFallParams) GetThresholdSecsOk() (*int32, bool)`

GetThresholdSecsOk returns a tuple with the ThresholdSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdSecs

`func (o *StorageDomainCloudDownWaterFallParams) SetThresholdSecs(v int32)`

SetThresholdSecs sets ThresholdSecs field to given value.

### HasThresholdSecs

`func (o *StorageDomainCloudDownWaterFallParams) HasThresholdSecs() bool`

HasThresholdSecs returns a boolean if a field has been set.

### SetThresholdSecsNil

`func (o *StorageDomainCloudDownWaterFallParams) SetThresholdSecsNil(b bool)`

 SetThresholdSecsNil sets the value for ThresholdSecs to be an explicit nil

### UnsetThresholdSecs
`func (o *StorageDomainCloudDownWaterFallParams) UnsetThresholdSecs()`

UnsetThresholdSecs ensures that no value is present for ThresholdSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


