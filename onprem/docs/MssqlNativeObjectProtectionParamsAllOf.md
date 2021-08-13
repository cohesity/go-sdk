# MssqlNativeObjectProtectionParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]MssqlNativeObjectProtection**](MssqlNativeObjectProtection.md) | Specifies the list of objects to be protected. | 
**NumStreams** | Pointer to **NullableInt32** | Specifies the number of streams to be used. | [optional] 
**WithClause** | Pointer to **NullableString** | Specifies the WithClause to be used. | [optional] 

## Methods

### NewMssqlNativeObjectProtectionParamsAllOf

`func NewMssqlNativeObjectProtectionParamsAllOf(objects []MssqlNativeObjectProtection, ) *MssqlNativeObjectProtectionParamsAllOf`

NewMssqlNativeObjectProtectionParamsAllOf instantiates a new MssqlNativeObjectProtectionParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlNativeObjectProtectionParamsAllOfWithDefaults

`func NewMssqlNativeObjectProtectionParamsAllOfWithDefaults() *MssqlNativeObjectProtectionParamsAllOf`

NewMssqlNativeObjectProtectionParamsAllOfWithDefaults instantiates a new MssqlNativeObjectProtectionParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *MssqlNativeObjectProtectionParamsAllOf) GetObjects() []MssqlNativeObjectProtection`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MssqlNativeObjectProtectionParamsAllOf) GetObjectsOk() (*[]MssqlNativeObjectProtection, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MssqlNativeObjectProtectionParamsAllOf) SetObjects(v []MssqlNativeObjectProtection)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *MssqlNativeObjectProtectionParamsAllOf) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *MssqlNativeObjectProtectionParamsAllOf) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetNumStreams

`func (o *MssqlNativeObjectProtectionParamsAllOf) GetNumStreams() int32`

GetNumStreams returns the NumStreams field if non-nil, zero value otherwise.

### GetNumStreamsOk

`func (o *MssqlNativeObjectProtectionParamsAllOf) GetNumStreamsOk() (*int32, bool)`

GetNumStreamsOk returns a tuple with the NumStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStreams

`func (o *MssqlNativeObjectProtectionParamsAllOf) SetNumStreams(v int32)`

SetNumStreams sets NumStreams field to given value.

### HasNumStreams

`func (o *MssqlNativeObjectProtectionParamsAllOf) HasNumStreams() bool`

HasNumStreams returns a boolean if a field has been set.

### SetNumStreamsNil

`func (o *MssqlNativeObjectProtectionParamsAllOf) SetNumStreamsNil(b bool)`

 SetNumStreamsNil sets the value for NumStreams to be an explicit nil

### UnsetNumStreams
`func (o *MssqlNativeObjectProtectionParamsAllOf) UnsetNumStreams()`

UnsetNumStreams ensures that no value is present for NumStreams, not even an explicit nil
### GetWithClause

`func (o *MssqlNativeObjectProtectionParamsAllOf) GetWithClause() string`

GetWithClause returns the WithClause field if non-nil, zero value otherwise.

### GetWithClauseOk

`func (o *MssqlNativeObjectProtectionParamsAllOf) GetWithClauseOk() (*string, bool)`

GetWithClauseOk returns a tuple with the WithClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithClause

`func (o *MssqlNativeObjectProtectionParamsAllOf) SetWithClause(v string)`

SetWithClause sets WithClause field to given value.

### HasWithClause

`func (o *MssqlNativeObjectProtectionParamsAllOf) HasWithClause() bool`

HasWithClause returns a boolean if a field has been set.

### SetWithClauseNil

`func (o *MssqlNativeObjectProtectionParamsAllOf) SetWithClauseNil(b bool)`

 SetWithClauseNil sets the value for WithClause to be an explicit nil

### UnsetWithClause
`func (o *MssqlNativeObjectProtectionParamsAllOf) UnsetWithClause()`

UnsetWithClause ensures that no value is present for WithClause, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


