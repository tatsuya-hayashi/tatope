# V1Table

Table is a tabular representation of a set of API resources. The server transforms the object into a set of preferred columns for quickly reviewing the objects.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**column_definitions** | [**List[V1TableColumnDefinition]**](V1TableColumnDefinition.md) | columnDefinitions describes each column in the returned items array. The number of cells per row will always match the number of column definitions. | 
**kind** | **str** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**metadata** | [**V1ListMeta**](V1ListMeta.md) |  | [optional] 
**rows** | [**List[V1TableRow]**](V1TableRow.md) | rows is the list of items in the table. | 

## Example

```python
from tatope.models.v1_table import V1Table

# TODO update the JSON string below
json = "{}"
# create an instance of V1Table from a JSON string
v1_table_instance = V1Table.from_json(json)
# print the JSON string representation of the object
print(V1Table.to_json())

# convert the object into a dict
v1_table_dict = v1_table_instance.to_dict()
# create an instance of V1Table from a dict
v1_table_from_dict = V1Table.from_dict(v1_table_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


