# HeliosTenantAssignmentsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | The tenant id. | [optional] 
**Assignments** | Pointer to [**NullableHeliosBaseTenantAssignmentsResult**](HeliosBaseTenantAssignmentsResult.md) |  | [optional] 

## Methods

### NewHeliosTenantAssignmentsResult

`func NewHeliosTenantAssignmentsResult() *HeliosTenantAssignmentsResult`

NewHeliosTenantAssignmentsResult instantiates a new HeliosTenantAssignmentsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTenantAssignmentsResultWithDefaults

`func NewHeliosTenantAssignmentsResultWithDefaults() *HeliosTenantAssignmentsResult`

NewHeliosTenantAssignmentsResultWithDefaults instantiates a new HeliosTenantAssignmentsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *HeliosTenantAssignmentsResult) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HeliosTenantAssignmentsResult) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HeliosTenantAssignmentsResult) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HeliosTenantAssignmentsResult) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HeliosTenantAssignmentsResult) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HeliosTenantAssignmentsResult) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAssignments

`func (o *HeliosTenantAssignmentsResult) GetAssignments() HeliosBaseTenantAssignmentsResult`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *HeliosTenantAssignmentsResult) GetAssignmentsOk() (*HeliosBaseTenantAssignmentsResult, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *HeliosTenantAssignmentsResult) SetAssignments(v HeliosBaseTenantAssignmentsResult)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *HeliosTenantAssignmentsResult) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignmentsNil

`func (o *HeliosTenantAssignmentsResult) SetAssignmentsNil(b bool)`

 SetAssignmentsNil sets the value for Assignments to be an explicit nil

### UnsetAssignments
`func (o *HeliosTenantAssignmentsResult) UnsetAssignments()`

UnsetAssignments ensures that no value is present for Assignments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


