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

    def test_change_group_Info(self):
        groupId = 1
        get_all_user_in_group="""
        query{
            usersOfGroup(groupId : %d){
                userId
            }
        }
        """%(groupId)

        res = self.client.send(get_all_user_in_group)
        self.assertTrue(len(res['data']['usersOfGroup']) == 3)

        change_group_info="""
        mutation{
            changeGroupInfo(groupId : %d, groupName : "%s", purpose : "%s", emails : ["%s", "%s"])
        }
        """%(1, "groupnewHaha", "Purposehahahaha", "De@gmail.com", "Shiki@gmail.com")

        res = self.client.send(change_group_info)

        print("change group info respone:", res)

        res = self.client.send(get_all_user_in_group)
        print(res)
        self.assertTrue(len(res['data']['usersOfGroup']) == 2)





if __name__ == '__main__':
    unittest.main()
