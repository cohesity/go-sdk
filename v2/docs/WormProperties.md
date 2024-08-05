# WormProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsArchiveWormCompliant** | Pointer to **NullableBool** | Specifies whether this archive run is WORM compliant | [optional] 
**WormExpiryTimeUsecs** | Pointer to **NullableInt64** | Specifies the time at which the WORM protection expires. | [optional] 
**WormNonComplianceReason** | Pointer to **NullableString** | Specifies reason of archive not being worm compliant. | [optional] 

## Methods

### NewWormProperties

`func NewWormProperties() *WormProperties`

NewWormProperties instantiates a new WormProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWormPropertiesWithDefaults

`func NewWormPropertiesWithDefaults() *WormProperties`

NewWormPropertiesWithDefaults instantiates a new WormProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsArchiveWormCompliant

`func (o *WormProperties) GetIsArchiveWormCompliant() bool`

GetIsArchiveWormCompliant returns the IsArchiveWormCompliant field if non-nil, zero value otherwise.

### GetIsArchiveWormCompliantOk

`func (o *WormProperties) GetIsArchiveWormCompliantOk() (*bool, bool)`

GetIsArchiveWormCompliantOk returns a tuple with the IsArchiveWormCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchiveWormCompliant

`func (o *WormProperties) SetIsArchiveWormCompliant(v bool)`

SetIsArchiveWormCompliant sets IsArchiveWormCompliant field to given value.

### HasIsArchiveWormCompliant

`func (o *WormProperties) HasIsArchiveWormCompliant() bool`

HasIsArchiveWormCompliant returns a boolean if a field has been set.

### SetIsArchiveWormCompliantNil

`func (o *WormProperties) SetIsArchiveWormCompliantNil(b bool)`

 SetIsArchiveWormCompliantNil sets the value for IsArchiveWormCompliant to be an explicit nil

### UnsetIsArchiveWormCompliant
`func (o *WormProperties) UnsetIsArchiveWormCompliant()`

UnsetIsArchiveWormCompliant ensures that no value is present for IsArchiveWormCompliant, not even an explicit nil
### GetWormExpiryTimeUsecs

`func (o *WormProperties) GetWormExpiryTimeUsecs() int64`

GetWormExpiryTimeUsecs returns the WormExpiryTimeUsecs field if non-nil, zero value otherwise.

### GetWormExpiryTimeUsecsOk

`func (o *WormProperties) GetWormExpiryTimeUsecsOk() (*int64, bool)`

GetWormExpiryTimeUsecsOk returns a tuple with the WormExpiryTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormExpiryTimeUsecs

`func (o *WormProperties) SetWormExpiryTimeUsecs(v int64)`

SetWormExpiryTimeUsecs sets WormExpiryTimeUsecs field to given value.

### HasWormExpiryTimeUsecs

`func (o *WormProperties) HasWormExpiryTimeUsecs() bool`

HasWormExpiryTimeUsecs returns a boolean if a field has been set.

### SetWormExpiryTimeUsecsNil

`func (o *WormProperties) SetWormExpiryTimeUsecsNil(b bool)`

 SetWormExpiryTimeUsecsNil sets the value for WormExpiryTimeUsecs to be an explicit nil

### UnsetWormExpiryTimeUsecs
`func (o *WormProperties) UnsetWormExpiryTimeUsecs()`

UnsetWormExpiryTimeUsecs ensures that no value is present for WormExpiryTimeUsecs, not even an explicit nil
### GetWormNonComplianceReason

`func (o *WormProperties) GetWormNonComplianceReason() string`

GetWormNonComplianceReason returns the WormNonComplianceReason field if non-nil, zero value otherwise.

### GetWormNonComplianceReasonOk

`func (o *WormProperties) GetWormNonComplianceReasonOk() (*string, bool)`

GetWormNonComplianceReasonOk returns a tuple with the WormNonComplianceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormNonComplianceReason

`func (o *WormProperties) SetWormNonComplianceReason(v string)`

SetWormNonComplianceReason sets WormNonComplianceReason field to given value.

### HasWormNonComplianceReason

`func (o *WormProperties) HasWormNonComplianceReason() bool`

HasWormNonComplianceReason returns a boolean if a field has been set.

### SetWormNonComplianceReasonNil

`func (o *WormProperties) SetWormNonComplianceReasonNil(b bool)`

 SetWormNonComplianceReasonNil sets the value for WormNonComplianceReason to be an explicit nil

### UnsetWormNonComplianceReason
`func (o *WormProperties) UnsetWormNonComplianceReason()`

UnsetWormNonComplianceReason ensures that no value is present for WormNonComplianceReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


