# ObjectPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | Specifies the id of protection policy. | 
**ProtectionType** | **NullableString** | Specifies the protection type. | 

## Methods

### NewObjectPolicy

`func NewObjectPolicy(id NullableString, protectionType NullableString, ) *ObjectPolicy`

NewObjectPolicy instantiates a new ObjectPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPolicyWithDefaults

`func NewObjectPolicyWithDefaults() *ObjectPolicy`

NewObjectPolicyWithDefaults instantiates a new ObjectPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectPolicy) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ObjectPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ObjectPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProtectionType

`func (o *ObjectPolicy) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *ObjectPolicy) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *ObjectPolicy) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.


### SetProtectionTypeNil

`func (o *ObjectPolicy) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *ObjectPolicy) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


