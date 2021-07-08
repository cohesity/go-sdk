# IsilonAccessZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the id of the access zone. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the access zone. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path of the access zone in ifs. This should include the leading \&quot;/ifs/\&quot;. | [optional] 

## Methods

### NewIsilonAccessZone

`func NewIsilonAccessZone() *IsilonAccessZone`

NewIsilonAccessZone instantiates a new IsilonAccessZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonAccessZoneWithDefaults

`func NewIsilonAccessZoneWithDefaults() *IsilonAccessZone`

NewIsilonAccessZoneWithDefaults instantiates a new IsilonAccessZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IsilonAccessZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IsilonAccessZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IsilonAccessZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IsilonAccessZone) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IsilonAccessZone) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IsilonAccessZone) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *IsilonAccessZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IsilonAccessZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IsilonAccessZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IsilonAccessZone) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IsilonAccessZone) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IsilonAccessZone) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *IsilonAccessZone) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IsilonAccessZone) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IsilonAccessZone) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IsilonAccessZone) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *IsilonAccessZone) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *IsilonAccessZone) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


