# RigelConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the id of the connection. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the connection. | [optional] 

## Methods

### NewRigelConnection

`func NewRigelConnection() *RigelConnection`

NewRigelConnection instantiates a new RigelConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRigelConnectionWithDefaults

`func NewRigelConnectionWithDefaults() *RigelConnection`

NewRigelConnectionWithDefaults instantiates a new RigelConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RigelConnection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RigelConnection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RigelConnection) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RigelConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RigelConnection) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RigelConnection) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *RigelConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RigelConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RigelConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RigelConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RigelConnection) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RigelConnection) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


