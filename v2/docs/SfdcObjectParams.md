# SfdcObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsAdded** | Pointer to **NullableInt64** | Specifies the number of records added for the Object. | [optional] 
**RecordsModified** | Pointer to **NullableInt64** | Specifies the number of records updated for the Object. | [optional] 
**RecordsRemoved** | Pointer to **NullableInt64** | Specifies the number of records removed from the Object. | [optional] 

## Methods

### NewSfdcObjectParams

`func NewSfdcObjectParams() *SfdcObjectParams`

NewSfdcObjectParams instantiates a new SfdcObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcObjectParamsWithDefaults

`func NewSfdcObjectParamsWithDefaults() *SfdcObjectParams`

NewSfdcObjectParamsWithDefaults instantiates a new SfdcObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsAdded

`func (o *SfdcObjectParams) GetRecordsAdded() int64`

GetRecordsAdded returns the RecordsAdded field if non-nil, zero value otherwise.

### GetRecordsAddedOk

`func (o *SfdcObjectParams) GetRecordsAddedOk() (*int64, bool)`

GetRecordsAddedOk returns a tuple with the RecordsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsAdded

`func (o *SfdcObjectParams) SetRecordsAdded(v int64)`

SetRecordsAdded sets RecordsAdded field to given value.

### HasRecordsAdded

`func (o *SfdcObjectParams) HasRecordsAdded() bool`

HasRecordsAdded returns a boolean if a field has been set.

### SetRecordsAddedNil

`func (o *SfdcObjectParams) SetRecordsAddedNil(b bool)`

 SetRecordsAddedNil sets the value for RecordsAdded to be an explicit nil

### UnsetRecordsAdded
`func (o *SfdcObjectParams) UnsetRecordsAdded()`

UnsetRecordsAdded ensures that no value is present for RecordsAdded, not even an explicit nil
### GetRecordsModified

`func (o *SfdcObjectParams) GetRecordsModified() int64`

GetRecordsModified returns the RecordsModified field if non-nil, zero value otherwise.

### GetRecordsModifiedOk

`func (o *SfdcObjectParams) GetRecordsModifiedOk() (*int64, bool)`

GetRecordsModifiedOk returns a tuple with the RecordsModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsModified

`func (o *SfdcObjectParams) SetRecordsModified(v int64)`

SetRecordsModified sets RecordsModified field to given value.

### HasRecordsModified

`func (o *SfdcObjectParams) HasRecordsModified() bool`

HasRecordsModified returns a boolean if a field has been set.

### SetRecordsModifiedNil

`func (o *SfdcObjectParams) SetRecordsModifiedNil(b bool)`

 SetRecordsModifiedNil sets the value for RecordsModified to be an explicit nil

### UnsetRecordsModified
`func (o *SfdcObjectParams) UnsetRecordsModified()`

UnsetRecordsModified ensures that no value is present for RecordsModified, not even an explicit nil
### GetRecordsRemoved

`func (o *SfdcObjectParams) GetRecordsRemoved() int64`

GetRecordsRemoved returns the RecordsRemoved field if non-nil, zero value otherwise.

### GetRecordsRemovedOk

`func (o *SfdcObjectParams) GetRecordsRemovedOk() (*int64, bool)`

GetRecordsRemovedOk returns a tuple with the RecordsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsRemoved

`func (o *SfdcObjectParams) SetRecordsRemoved(v int64)`

SetRecordsRemoved sets RecordsRemoved field to given value.

### HasRecordsRemoved

`func (o *SfdcObjectParams) HasRecordsRemoved() bool`

HasRecordsRemoved returns a boolean if a field has been set.

### SetRecordsRemovedNil

`func (o *SfdcObjectParams) SetRecordsRemovedNil(b bool)`

 SetRecordsRemovedNil sets the value for RecordsRemoved to be an explicit nil

### UnsetRecordsRemoved
`func (o *SfdcObjectParams) UnsetRecordsRemoved()`

UnsetRecordsRemoved ensures that no value is present for RecordsRemoved, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


