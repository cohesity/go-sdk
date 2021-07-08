# RegisteredApplicationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationServer** | Pointer to [**NullableProtectionSourceNode**](ProtectionSourceNode.md) | Specifies the child subtree used to store additional application-level Objects. Different environments use the subtree to store application-level information. For example for SQL Server, this subtree stores the SQL Server instances running on a VM. | [optional] 
**RegisteredProtectionSource** | Pointer to [**NullableProtectionSource**](ProtectionSource.md) | Specifies the Protection Source like a VM or Physical Server that registered the Application Server. | [optional] 

## Methods

### NewRegisteredApplicationServer

`func NewRegisteredApplicationServer() *RegisteredApplicationServer`

NewRegisteredApplicationServer instantiates a new RegisteredApplicationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredApplicationServerWithDefaults

`func NewRegisteredApplicationServerWithDefaults() *RegisteredApplicationServer`

NewRegisteredApplicationServerWithDefaults instantiates a new RegisteredApplicationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationServer

`func (o *RegisteredApplicationServer) GetApplicationServer() ProtectionSourceNode`

GetApplicationServer returns the ApplicationServer field if non-nil, zero value otherwise.

### GetApplicationServerOk

`func (o *RegisteredApplicationServer) GetApplicationServerOk() (*ProtectionSourceNode, bool)`

GetApplicationServerOk returns a tuple with the ApplicationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationServer

`func (o *RegisteredApplicationServer) SetApplicationServer(v ProtectionSourceNode)`

SetApplicationServer sets ApplicationServer field to given value.

### HasApplicationServer

`func (o *RegisteredApplicationServer) HasApplicationServer() bool`

HasApplicationServer returns a boolean if a field has been set.

### SetApplicationServerNil

`func (o *RegisteredApplicationServer) SetApplicationServerNil(b bool)`

 SetApplicationServerNil sets the value for ApplicationServer to be an explicit nil

### UnsetApplicationServer
`func (o *RegisteredApplicationServer) UnsetApplicationServer()`

UnsetApplicationServer ensures that no value is present for ApplicationServer, not even an explicit nil
### GetRegisteredProtectionSource

`func (o *RegisteredApplicationServer) GetRegisteredProtectionSource() ProtectionSource`

GetRegisteredProtectionSource returns the RegisteredProtectionSource field if non-nil, zero value otherwise.

### GetRegisteredProtectionSourceOk

`func (o *RegisteredApplicationServer) GetRegisteredProtectionSourceOk() (*ProtectionSource, bool)`

GetRegisteredProtectionSourceOk returns a tuple with the RegisteredProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredProtectionSource

`func (o *RegisteredApplicationServer) SetRegisteredProtectionSource(v ProtectionSource)`

SetRegisteredProtectionSource sets RegisteredProtectionSource field to given value.

### HasRegisteredProtectionSource

`func (o *RegisteredApplicationServer) HasRegisteredProtectionSource() bool`

HasRegisteredProtectionSource returns a boolean if a field has been set.

### SetRegisteredProtectionSourceNil

`func (o *RegisteredApplicationServer) SetRegisteredProtectionSourceNil(b bool)`

 SetRegisteredProtectionSourceNil sets the value for RegisteredProtectionSource to be an explicit nil

### UnsetRegisteredProtectionSource
`func (o *RegisteredApplicationServer) UnsetRegisteredProtectionSource()`

UnsetRegisteredProtectionSource ensures that no value is present for RegisteredProtectionSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


