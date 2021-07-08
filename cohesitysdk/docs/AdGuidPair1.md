# AdGuidPair1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **NullableString** | Specifies the destination guid in production AD object corresponding to source. If empty, it assumed to be &#39;source&#39; guid. | [optional] 
**Source** | Pointer to **NullableString** | Specifies the source guid string of an AD object in mounted AD snapshot. This cannot be empty. | [optional] 

## Methods

### NewAdGuidPair1

`func NewAdGuidPair1() *AdGuidPair1`

NewAdGuidPair1 instantiates a new AdGuidPair1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdGuidPair1WithDefaults

`func NewAdGuidPair1WithDefaults() *AdGuidPair1`

NewAdGuidPair1WithDefaults instantiates a new AdGuidPair1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *AdGuidPair1) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AdGuidPair1) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AdGuidPair1) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *AdGuidPair1) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *AdGuidPair1) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *AdGuidPair1) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetSource

`func (o *AdGuidPair1) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AdGuidPair1) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AdGuidPair1) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AdGuidPair1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AdGuidPair1) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AdGuidPair1) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


