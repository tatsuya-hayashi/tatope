# coding: utf-8

"""
    tatoperator

    Python SDK for Tat-Operator

    The version of the OpenAPI document: v1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from tatope.models.v1_table_row import V1TableRow

class TestV1TableRow(unittest.TestCase):
    """V1TableRow unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> V1TableRow:
        """Test V1TableRow
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `V1TableRow`
        """
        model = V1TableRow()
        if include_optional:
            return V1TableRow(
                cells = [
                    None
                    ],
                conditions = [
                    tatope.models.v1/table_row_condition.v1.TableRowCondition(
                        message = '', 
                        reason = '', 
                        status = '', 
                        type = '', )
                    ],
                object = None
            )
        else:
            return V1TableRow(
                cells = [
                    None
                    ],
        )
        """

    def testV1TableRow(self):
        """Test V1TableRow"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
