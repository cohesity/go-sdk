# MssqlFileObjectProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object being protected. If this is a non leaf level object, then the object will be auto-protected unless leaf objects are specified for exclusion. | 

## Methods

### NewMssqlFileObjectProtection

`func NewMssqlFileObjectProtection(id NullableInt64, ) *MssqlFileObjectProtection`

NewMssqlFileObjectProtection instantiates a new MssqlFileObjectProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlFileObjectProtectionWithDefaults

`func NewMssqlFileObjectProtectionWithDefaults() *MssqlFileObjectProtection`

NewMssqlFileObjectProtectionWithDefaults instantiates a new MssqlFileObjectProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MssqlFileObjectProtection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MssqlFileObjectProtection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MssqlFileObjectProtection) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *MssqlFileObjectProtection) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MssqlFileObjectProtection) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


