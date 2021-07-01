# SwiftParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeystoneId** | Pointer to **NullableInt64** | Specifies the associated Keystone configuration id. | [optional] 
**OperatorRoles** | Pointer to **[]string** | Specifies a list of operator roles. | [optional] 

## Methods

### NewSwiftParams

`func NewSwiftParams() *SwiftParams`

NewSwiftParams instantiates a new SwiftParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwiftParamsWithDefaults

`func NewSwiftParamsWithDefaults() *SwiftParams`

NewSwiftParamsWithDefaults instantiates a new SwiftParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeystoneId

`func (o *SwiftParams) GetKeystoneId() int64`

GetKeystoneId returns the KeystoneId field if non-nil, zero value otherwise.

### GetKeystoneIdOk

`func (o *SwiftParams) GetKeystoneIdOk() (*int64, bool)`

GetKeystoneIdOk returns a tuple with the KeystoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneId

`func (o *SwiftParams) SetKeystoneId(v int64)`

SetKeystoneId sets KeystoneId field to given value.

### HasKeystoneId

`func (o *SwiftParams) HasKeystoneId() bool`

HasKeystoneId returns a boolean if a field has been set.

### SetKeystoneIdNil

`func (o *SwiftParams) SetKeystoneIdNil(b bool)`

 SetKeystoneIdNil sets the value for KeystoneId to be an explicit nil

### UnsetKeystoneId
`func (o *SwiftParams) UnsetKeystoneId()`

UnsetKeystoneId ensures that no value is present for KeystoneId, not even an explicit nil
### GetOperatorRoles

`func (o *SwiftParams) GetOperatorRoles() []string`

GetOperatorRoles returns the OperatorRoles field if non-nil, zero value otherwise.

### GetOperatorRolesOk

`func (o *SwiftParams) GetOperatorRolesOk() (*[]string, bool)`

GetOperatorRolesOk returns a tuple with the OperatorRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorRoles

`func (o *SwiftParams) SetOperatorRoles(v []string)`

SetOperatorRoles sets OperatorRoles field to given value.

### HasOperatorRoles

`func (o *SwiftParams) HasOperatorRoles() bool`

HasOperatorRoles returns a boolean if a field has been set.

### SetOperatorRolesNil

`func (o *SwiftParams) SetOperatorRolesNil(b bool)`

 SetOperatorRolesNil sets the value for OperatorRoles to be an explicit nil

### UnsetOperatorRoles
`func (o *SwiftParams) UnsetOperatorRoles()`

UnsetOperatorRoles ensures that no value is present for OperatorRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


