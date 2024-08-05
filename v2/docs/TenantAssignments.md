# TenantAssignments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignments** | Pointer to [**NullableTenantAssignmentsResult**](TenantAssignmentsResult.md) |  | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id. | [optional] 

## Methods

### NewTenantAssignments

`func NewTenantAssignments() *TenantAssignments`

NewTenantAssignments instantiates a new TenantAssignments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAssignmentsWithDefaults

`func NewTenantAssignmentsWithDefaults() *TenantAssignments`

NewTenantAssignmentsWithDefaults instantiates a new TenantAssignments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignments

`func (o *TenantAssignments) GetAssignments() TenantAssignmentsResult`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *TenantAssignments) GetAssignmentsOk() (*TenantAssignmentsResult, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *TenantAssignments) SetAssignments(v TenantAssignmentsResult)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *TenantAssignments) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignmentsNil

`func (o *TenantAssignments) SetAssignmentsNil(b bool)`

 SetAssignmentsNil sets the value for Assignments to be an explicit nil

### UnsetAssignments
`func (o *TenantAssignments) UnsetAssignments()`

UnsetAssignments ensures that no value is present for Assignments, not even an explicit nil
### GetTenantId

`func (o *TenantAssignments) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantAssignments) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantAssignments) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantAssignments) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantAssignments) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantAssignments) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


