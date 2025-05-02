# V1Table

Table is a tabular representation of a set of API resources. The server transforms the object into a set of preferred columns for quickly reviewing the objects.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**column_definitions** | [**[V1TableColumnDefinition]**](V1TableColumnDefinition.md) | columnDefinitions describes each column in the returned items array. The number of cells per row will always match the number of column definitions. | 
**rows** | [**[V1TableRow]**](V1TableRow.md) | rows is the list of items in the table. | 
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**metadata** | [**V1ListMeta**](V1ListMeta.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


