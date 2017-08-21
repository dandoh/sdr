import unittest
import client
from faker import Factory

class TestSubscribe(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.client = client.ClientAPI("Dandoh@gmail.com")
        cls.client2 = client.ClientAPI("De@gmail.com")
        cls.client3 = client.ClientAPI("Shiki@gmail.com")

    def test_upper(self):
        self.assertEqual('foo'.upper(), 'FOO')


    def testGetAllUser(self):
        get_all_users_query="""
        query{
            users{
                userId
                email
            }
        }
        """
        res = self.client.send(get_all_users_query)
        print(res)
        self.assertTrue(len(res['data']['users']) == 3)








if __name__ == '__main__':
    unittest.main()
