# AwsIAmRoleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **NullableString** | Specifies the Account Id of the external target. | 
**IAmRole** | **NullableString** | Specifies the I Am Role of the external target. | 

## Methods

### NewAwsIAmRoleParams

`func NewAwsIAmRoleParams(accountId NullableString, iAmRole NullableString, ) *AwsIAmRoleParams`

NewAwsIAmRoleParams instantiates a new AwsIAmRoleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsIAmRoleParamsWithDefaults

`func NewAwsIAmRoleParamsWithDefaults() *AwsIAmRoleParams`

NewAwsIAmRoleParamsWithDefaults instantiates a new AwsIAmRoleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AwsIAmRoleParams) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsIAmRoleParams) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AwsIAmRoleParams) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *AwsIAmRoleParams) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *AwsIAmRoleParams) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetIAmRole

`func (o *AwsIAmRoleParams) GetIAmRole() string`

GetIAmRole returns the IAmRole field if non-nil, zero value otherwise.

### GetIAmRoleOk

`func (o *AwsIAmRoleParams) GetIAmRoleOk() (*string, bool)`

GetIAmRoleOk returns a tuple with the IAmRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAmRole

`func (o *AwsIAmRoleParams) SetIAmRole(v string)`

SetIAmRole sets IAmRole field to given value.


### SetIAmRoleNil

`func (o *AwsIAmRoleParams) SetIAmRoleNil(b bool)`

 SetIAmRoleNil sets the value for IAmRole to be an explicit nil

### UnsetIAmRole
`func (o *AwsIAmRoleParams) UnsetIAmRole()`

UnsetIAmRole ensures that no value is present for IAmRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


