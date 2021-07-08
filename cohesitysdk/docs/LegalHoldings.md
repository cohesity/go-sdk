# LegalHoldings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoldForLegalPurpose** | Pointer to **NullableBool** | Specifies whether the source is put on legal hold or not. | [optional] 
**ProtectionSourceId** | Pointer to **NullableInt64** | Specifies an Protection Source Id in the snapshot. | [optional] 

## Methods

### NewLegalHoldings

`func NewLegalHoldings() *LegalHoldings`

NewLegalHoldings instantiates a new LegalHoldings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalHoldingsWithDefaults

`func NewLegalHoldingsWithDefaults() *LegalHoldings`

NewLegalHoldingsWithDefaults instantiates a new LegalHoldings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoldForLegalPurpose

`func (o *LegalHoldings) GetHoldForLegalPurpose() bool`

GetHoldForLegalPurpose returns the HoldForLegalPurpose field if non-nil, zero value otherwise.

### GetHoldForLegalPurposeOk

`func (o *LegalHoldings) GetHoldForLegalPurposeOk() (*bool, bool)`

GetHoldForLegalPurposeOk returns a tuple with the HoldForLegalPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldForLegalPurpose

`func (o *LegalHoldings) SetHoldForLegalPurpose(v bool)`

SetHoldForLegalPurpose sets HoldForLegalPurpose field to given value.

### HasHoldForLegalPurpose

`func (o *LegalHoldings) HasHoldForLegalPurpose() bool`

HasHoldForLegalPurpose returns a boolean if a field has been set.

### SetHoldForLegalPurposeNil

`func (o *LegalHoldings) SetHoldForLegalPurposeNil(b bool)`

 SetHoldForLegalPurposeNil sets the value for HoldForLegalPurpose to be an explicit nil

### UnsetHoldForLegalPurpose
`func (o *LegalHoldings) UnsetHoldForLegalPurpose()`

UnsetHoldForLegalPurpose ensures that no value is present for HoldForLegalPurpose, not even an explicit nil
### GetProtectionSourceId

`func (o *LegalHoldings) GetProtectionSourceId() int64`

GetProtectionSourceId returns the ProtectionSourceId field if non-nil, zero value otherwise.

### GetProtectionSourceIdOk

`func (o *LegalHoldings) GetProtectionSourceIdOk() (*int64, bool)`

GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceId

`func (o *LegalHoldings) SetProtectionSourceId(v int64)`

SetProtectionSourceId sets ProtectionSourceId field to given value.

### HasProtectionSourceId

`func (o *LegalHoldings) HasProtectionSourceId() bool`

HasProtectionSourceId returns a boolean if a field has been set.

### SetProtectionSourceIdNil

`func (o *LegalHoldings) SetProtectionSourceIdNil(b bool)`

 SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil

### UnsetProtectionSourceId
`func (o *LegalHoldings) UnsetProtectionSourceId()`

UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


