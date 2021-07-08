# FilteringPolicyProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowFilters** | Pointer to **[]string** | List of filters to allow matched objects for backup. | [optional] 
**DenyFilters** | Pointer to **[]string** | List of filters to deny matched objects for backup. | [optional] 

## Methods

### NewFilteringPolicyProto

`func NewFilteringPolicyProto() *FilteringPolicyProto`

NewFilteringPolicyProto instantiates a new FilteringPolicyProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilteringPolicyProtoWithDefaults

`func NewFilteringPolicyProtoWithDefaults() *FilteringPolicyProto`

NewFilteringPolicyProtoWithDefaults instantiates a new FilteringPolicyProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowFilters

`func (o *FilteringPolicyProto) GetAllowFilters() []string`

GetAllowFilters returns the AllowFilters field if non-nil, zero value otherwise.

### GetAllowFiltersOk

`func (o *FilteringPolicyProto) GetAllowFiltersOk() (*[]string, bool)`

GetAllowFiltersOk returns a tuple with the AllowFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFilters

`func (o *FilteringPolicyProto) SetAllowFilters(v []string)`

SetAllowFilters sets AllowFilters field to given value.

### HasAllowFilters

`func (o *FilteringPolicyProto) HasAllowFilters() bool`

HasAllowFilters returns a boolean if a field has been set.

### SetAllowFiltersNil

`func (o *FilteringPolicyProto) SetAllowFiltersNil(b bool)`

 SetAllowFiltersNil sets the value for AllowFilters to be an explicit nil

### UnsetAllowFilters
`func (o *FilteringPolicyProto) UnsetAllowFilters()`

UnsetAllowFilters ensures that no value is present for AllowFilters, not even an explicit nil
### GetDenyFilters

`func (o *FilteringPolicyProto) GetDenyFilters() []string`

GetDenyFilters returns the DenyFilters field if non-nil, zero value otherwise.

### GetDenyFiltersOk

`func (o *FilteringPolicyProto) GetDenyFiltersOk() (*[]string, bool)`

GetDenyFiltersOk returns a tuple with the DenyFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyFilters

`func (o *FilteringPolicyProto) SetDenyFilters(v []string)`

SetDenyFilters sets DenyFilters field to given value.

### HasDenyFilters

`func (o *FilteringPolicyProto) HasDenyFilters() bool`

HasDenyFilters returns a boolean if a field has been set.

### SetDenyFiltersNil

`func (o *FilteringPolicyProto) SetDenyFiltersNil(b bool)`

 SetDenyFiltersNil sets the value for DenyFilters to be an explicit nil

### UnsetDenyFilters
`func (o *FilteringPolicyProto) UnsetDenyFilters()`

UnsetDenyFilters ensures that no value is present for DenyFilters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


