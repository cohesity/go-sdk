# ObjectsProtectedByPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectsProtected** | Pointer to [**[]ObjectsByEnv**](ObjectsByEnv.md) | Protected Objects. | [optional] 
**PolicyId** | Pointer to **NullableString** | Id of the policy. | [optional] 
**PolicyName** | Pointer to **NullableString** | Name of the policy. | [optional] 

## Methods

### NewObjectsProtectedByPolicy

`func NewObjectsProtectedByPolicy() *ObjectsProtectedByPolicy`

NewObjectsProtectedByPolicy instantiates a new ObjectsProtectedByPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsProtectedByPolicyWithDefaults

`func NewObjectsProtectedByPolicyWithDefaults() *ObjectsProtectedByPolicy`

NewObjectsProtectedByPolicyWithDefaults instantiates a new ObjectsProtectedByPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectsProtected

`func (o *ObjectsProtectedByPolicy) GetObjectsProtected() []ObjectsByEnv`

GetObjectsProtected returns the ObjectsProtected field if non-nil, zero value otherwise.

### GetObjectsProtectedOk

`func (o *ObjectsProtectedByPolicy) GetObjectsProtectedOk() (*[]ObjectsByEnv, bool)`

GetObjectsProtectedOk returns a tuple with the ObjectsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsProtected

`func (o *ObjectsProtectedByPolicy) SetObjectsProtected(v []ObjectsByEnv)`

SetObjectsProtected sets ObjectsProtected field to given value.

### HasObjectsProtected

`func (o *ObjectsProtectedByPolicy) HasObjectsProtected() bool`

HasObjectsProtected returns a boolean if a field has been set.

### SetObjectsProtectedNil

`func (o *ObjectsProtectedByPolicy) SetObjectsProtectedNil(b bool)`

 SetObjectsProtectedNil sets the value for ObjectsProtected to be an explicit nil

### UnsetObjectsProtected
`func (o *ObjectsProtectedByPolicy) UnsetObjectsProtected()`

UnsetObjectsProtected ensures that no value is present for ObjectsProtected, not even an explicit nil
### GetPolicyId

`func (o *ObjectsProtectedByPolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ObjectsProtectedByPolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ObjectsProtectedByPolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ObjectsProtectedByPolicy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *ObjectsProtectedByPolicy) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *ObjectsProtectedByPolicy) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetPolicyName

`func (o *ObjectsProtectedByPolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ObjectsProtectedByPolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ObjectsProtectedByPolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ObjectsProtectedByPolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ObjectsProtectedByPolicy) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ObjectsProtectedByPolicy) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


