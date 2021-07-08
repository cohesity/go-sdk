# ApplicationSpecialParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationEntityIds** | Pointer to **[]int64** | Array of Ids of Application Entities like SQL/Oracle instances, and databases that should be protected in a Protection Source.  Specifies the subset of application entities like SQL/Oracle instances, and databases to protect in a Protection Source of type &#39;kSQL&#39;/&#39;kOracle&#39;. If not specified, all application entities on the Protection Source. | [optional] 

## Methods

### NewApplicationSpecialParameters

`func NewApplicationSpecialParameters() *ApplicationSpecialParameters`

NewApplicationSpecialParameters instantiates a new ApplicationSpecialParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSpecialParametersWithDefaults

`func NewApplicationSpecialParametersWithDefaults() *ApplicationSpecialParameters`

NewApplicationSpecialParametersWithDefaults instantiates a new ApplicationSpecialParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationEntityIds

`func (o *ApplicationSpecialParameters) GetApplicationEntityIds() []int64`

GetApplicationEntityIds returns the ApplicationEntityIds field if non-nil, zero value otherwise.

### GetApplicationEntityIdsOk

`func (o *ApplicationSpecialParameters) GetApplicationEntityIdsOk() (*[]int64, bool)`

GetApplicationEntityIdsOk returns a tuple with the ApplicationEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEntityIds

`func (o *ApplicationSpecialParameters) SetApplicationEntityIds(v []int64)`

SetApplicationEntityIds sets ApplicationEntityIds field to given value.

### HasApplicationEntityIds

`func (o *ApplicationSpecialParameters) HasApplicationEntityIds() bool`

HasApplicationEntityIds returns a boolean if a field has been set.

### SetApplicationEntityIdsNil

`func (o *ApplicationSpecialParameters) SetApplicationEntityIdsNil(b bool)`

 SetApplicationEntityIdsNil sets the value for ApplicationEntityIds to be an explicit nil

### UnsetApplicationEntityIds
`func (o *ApplicationSpecialParameters) UnsetApplicationEntityIds()`

UnsetApplicationEntityIds ensures that no value is present for ApplicationEntityIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


