# ApplicationsWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]MapReduceInfo**](MapReduceInfo.md) | Applications specifies the list of available map-reduce applications in analytics workbench. | [optional] 

## Methods

### NewApplicationsWrapper

`func NewApplicationsWrapper() *ApplicationsWrapper`

NewApplicationsWrapper instantiates a new ApplicationsWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsWrapperWithDefaults

`func NewApplicationsWrapperWithDefaults() *ApplicationsWrapper`

NewApplicationsWrapperWithDefaults instantiates a new ApplicationsWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ApplicationsWrapper) GetApplications() []MapReduceInfo`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationsWrapper) GetApplicationsOk() (*[]MapReduceInfo, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationsWrapper) SetApplications(v []MapReduceInfo)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ApplicationsWrapper) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *ApplicationsWrapper) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *ApplicationsWrapper) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


