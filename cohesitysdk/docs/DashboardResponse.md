# DashboardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboard** | Pointer to [**Dashboard**](Dashboard.md) |  | [optional] 
**Dashboards** | Pointer to [**[]Dashboard**](Dashboard.md) | Specifies a list of dashboards of all the clusters in the SPOG setup if the query parameter allClusters is set to true. Otherwise this field is not populated. When populated the dashboard field is also populated with aggregated dashboard values. | [optional] 

## Methods

### NewDashboardResponse

`func NewDashboardResponse() *DashboardResponse`

NewDashboardResponse instantiates a new DashboardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardResponseWithDefaults

`func NewDashboardResponseWithDefaults() *DashboardResponse`

NewDashboardResponseWithDefaults instantiates a new DashboardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboard

`func (o *DashboardResponse) GetDashboard() Dashboard`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *DashboardResponse) GetDashboardOk() (*Dashboard, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *DashboardResponse) SetDashboard(v Dashboard)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *DashboardResponse) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetDashboards

`func (o *DashboardResponse) GetDashboards() []Dashboard`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *DashboardResponse) GetDashboardsOk() (*[]Dashboard, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *DashboardResponse) SetDashboards(v []Dashboard)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *DashboardResponse) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### SetDashboardsNil

`func (o *DashboardResponse) SetDashboardsNil(b bool)`

 SetDashboardsNil sets the value for Dashboards to be an explicit nil

### UnsetDashboards
`func (o *DashboardResponse) UnsetDashboards()`

UnsetDashboards ensures that no value is present for Dashboards, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


