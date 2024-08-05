# SapHanaAgentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentType** | **string** | Specifies the type of agent. kJava agent type is supported only for kScript package type. | 
**PackageType** | **string** | Specifies the type of installer. | 

## Methods

### NewSapHanaAgentParams

`func NewSapHanaAgentParams(agentType string, packageType string, ) *SapHanaAgentParams`

NewSapHanaAgentParams instantiates a new SapHanaAgentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSapHanaAgentParamsWithDefaults

`func NewSapHanaAgentParamsWithDefaults() *SapHanaAgentParams`

NewSapHanaAgentParamsWithDefaults instantiates a new SapHanaAgentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentType

`func (o *SapHanaAgentParams) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *SapHanaAgentParams) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *SapHanaAgentParams) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.


### GetPackageType

`func (o *SapHanaAgentParams) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *SapHanaAgentParams) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *SapHanaAgentParams) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


