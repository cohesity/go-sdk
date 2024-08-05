# McmCohesionClaimRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceId** | **string** | Specifies the Id of the cohesion appliance with AWS. | 
**ApplianceName** | **string** | Specifies the name of the cohesion appliance. | 
**ClaimToken** | **string** | Claim token used for authentication. | 
**ClusterId** | **int64** | Specifies the cluster id. | 
**ClusterIncarnationId** | **int64** | Specifies the cluster incarnation id. | 

## Methods

### NewMcmCohesionClaimRequestParams

`func NewMcmCohesionClaimRequestParams(applianceId string, applianceName string, claimToken string, clusterId int64, clusterIncarnationId int64, ) *McmCohesionClaimRequestParams`

NewMcmCohesionClaimRequestParams instantiates a new McmCohesionClaimRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmCohesionClaimRequestParamsWithDefaults

`func NewMcmCohesionClaimRequestParamsWithDefaults() *McmCohesionClaimRequestParams`

NewMcmCohesionClaimRequestParamsWithDefaults instantiates a new McmCohesionClaimRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceId

`func (o *McmCohesionClaimRequestParams) GetApplianceId() string`

GetApplianceId returns the ApplianceId field if non-nil, zero value otherwise.

### GetApplianceIdOk

`func (o *McmCohesionClaimRequestParams) GetApplianceIdOk() (*string, bool)`

GetApplianceIdOk returns a tuple with the ApplianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceId

`func (o *McmCohesionClaimRequestParams) SetApplianceId(v string)`

SetApplianceId sets ApplianceId field to given value.


### GetApplianceName

`func (o *McmCohesionClaimRequestParams) GetApplianceName() string`

GetApplianceName returns the ApplianceName field if non-nil, zero value otherwise.

### GetApplianceNameOk

`func (o *McmCohesionClaimRequestParams) GetApplianceNameOk() (*string, bool)`

GetApplianceNameOk returns a tuple with the ApplianceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceName

`func (o *McmCohesionClaimRequestParams) SetApplianceName(v string)`

SetApplianceName sets ApplianceName field to given value.


### GetClaimToken

`func (o *McmCohesionClaimRequestParams) GetClaimToken() string`

GetClaimToken returns the ClaimToken field if non-nil, zero value otherwise.

### GetClaimTokenOk

`func (o *McmCohesionClaimRequestParams) GetClaimTokenOk() (*string, bool)`

GetClaimTokenOk returns a tuple with the ClaimToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimToken

`func (o *McmCohesionClaimRequestParams) SetClaimToken(v string)`

SetClaimToken sets ClaimToken field to given value.


### GetClusterId

`func (o *McmCohesionClaimRequestParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmCohesionClaimRequestParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmCohesionClaimRequestParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.


### GetClusterIncarnationId

`func (o *McmCohesionClaimRequestParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *McmCohesionClaimRequestParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *McmCohesionClaimRequestParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


