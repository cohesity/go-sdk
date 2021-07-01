# GuidPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestGuid** | Pointer to **NullableString** | Specifies the guid of an AD object in the Production AD. | [optional] 
**SourceGuid** | Pointer to **NullableString** | Specifies the guid of an AD object in the Snapshot AD. | [optional] 

## Methods

### NewGuidPair

`func NewGuidPair() *GuidPair`

NewGuidPair instantiates a new GuidPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuidPairWithDefaults

`func NewGuidPairWithDefaults() *GuidPair`

NewGuidPairWithDefaults instantiates a new GuidPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestGuid

`func (o *GuidPair) GetDestGuid() string`

GetDestGuid returns the DestGuid field if non-nil, zero value otherwise.

### GetDestGuidOk

`func (o *GuidPair) GetDestGuidOk() (*string, bool)`

GetDestGuidOk returns a tuple with the DestGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestGuid

`func (o *GuidPair) SetDestGuid(v string)`

SetDestGuid sets DestGuid field to given value.

### HasDestGuid

`func (o *GuidPair) HasDestGuid() bool`

HasDestGuid returns a boolean if a field has been set.

### SetDestGuidNil

`func (o *GuidPair) SetDestGuidNil(b bool)`

 SetDestGuidNil sets the value for DestGuid to be an explicit nil

### UnsetDestGuid
`func (o *GuidPair) UnsetDestGuid()`

UnsetDestGuid ensures that no value is present for DestGuid, not even an explicit nil
### GetSourceGuid

`func (o *GuidPair) GetSourceGuid() string`

GetSourceGuid returns the SourceGuid field if non-nil, zero value otherwise.

### GetSourceGuidOk

`func (o *GuidPair) GetSourceGuidOk() (*string, bool)`

GetSourceGuidOk returns a tuple with the SourceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuid

`func (o *GuidPair) SetSourceGuid(v string)`

SetSourceGuid sets SourceGuid field to given value.

### HasSourceGuid

`func (o *GuidPair) HasSourceGuid() bool`

HasSourceGuid returns a boolean if a field has been set.

### SetSourceGuidNil

`func (o *GuidPair) SetSourceGuidNil(b bool)`

 SetSourceGuidNil sets the value for SourceGuid to be an explicit nil

### UnsetSourceGuid
`func (o *GuidPair) UnsetSourceGuid()`

UnsetSourceGuid ensures that no value is present for SourceGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


