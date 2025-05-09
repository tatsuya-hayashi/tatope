# V1TableColumnDefinition

TableColumnDefinition contains information about a column returned in the Table.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | description is a human readable description of this column. | [default to '']
**format** | **str** | format is an optional OpenAPI type modifier for this column. A format modifies the type and imposes additional rules, like date or time formatting for a string. The &#39;name&#39; format is applied to the primary identifier column which has type &#39;string&#39; to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for more. | [default to '']
**name** | **str** | name is a human readable name for the column. | [default to '']
**priority** | **int** | priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a higher priority. | [default to 0]
**type** | **str** | type is an OpenAPI type definition for this column, such as number, integer, string, or array. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for more. | [default to '']

## Example

```python
from tatope.models.v1_table_column_definition import V1TableColumnDefinition

# TODO update the JSON string below
json = "{}"
# create an instance of V1TableColumnDefinition from a JSON string
v1_table_column_definition_instance = V1TableColumnDefinition.from_json(json)
# print the JSON string representation of the object
print(V1TableColumnDefinition.to_json())

# convert the object into a dict
v1_table_column_definition_dict = v1_table_column_definition_instance.to_dict()
# create an instance of V1TableColumnDefinition from a dict
v1_table_column_definition_from_dict = V1TableColumnDefinition.from_dict(v1_table_column_definition_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


