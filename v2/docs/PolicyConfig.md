# PolicyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]ObjectPolicy**](ObjectPolicy.md) | Specifies the list of protection policies for protecting an object with multiple policies. | [optional] 

## Methods

### NewPolicyConfig

`func NewPolicyConfig() *PolicyConfig`

NewPolicyConfig instantiates a new PolicyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyConfigWithDefaults

`func NewPolicyConfigWithDefaults() *PolicyConfig`

NewPolicyConfigWithDefaults instantiates a new PolicyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *PolicyConfig) GetPolicies() []ObjectPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *PolicyConfig) GetPoliciesOk() (*[]ObjectPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *PolicyConfig) SetPolicies(v []ObjectPolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *PolicyConfig) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *PolicyConfig) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *PolicyConfig) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


