# Run server before run test script
# Dependency: pip install Faker
import unittest
import client
from faker import Factory

class TestGroup(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.client = client.ClientAPI()

    def test_upper(self):
        self.assertEqual('foo'.upper(), 'FOO')

    def test_add_new_group(self):
        faker = Factory.create()
        valid_group = {'name': faker.bs(), 'purpose': faker.catch_phrase()}
        add_group_mutation = """
        mutation {
            addGroup(name: "%s", purpose: "%s")
        }
        """ % (valid_group['name'], valid_group['purpose'])
        res = self.client.send(add_group_mutation)
        print(res)

if __name__ == '__main__':
    unittest.main()
