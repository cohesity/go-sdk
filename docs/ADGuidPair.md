# ADGuidPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **NullableString** | Destination guid in production AD object corresponding to source. If empty, it assumed to be &#39;source&#39; guid. | [optional] 
**Source** | Pointer to **NullableString** | Source guid string of an AD object in mounted AD snapshot. This cannot be empty. | [optional] 

## Methods

### NewADGuidPair

`func NewADGuidPair() *ADGuidPair`

NewADGuidPair instantiates a new ADGuidPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADGuidPairWithDefaults

`func NewADGuidPairWithDefaults() *ADGuidPair`

NewADGuidPairWithDefaults instantiates a new ADGuidPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *ADGuidPair) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ADGuidPair) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ADGuidPair) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ADGuidPair) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *ADGuidPair) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *ADGuidPair) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetSource

`func (o *ADGuidPair) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ADGuidPair) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ADGuidPair) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ADGuidPair) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ADGuidPair) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ADGuidPair) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


