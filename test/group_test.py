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
        ##########Generate data for valid group###########
        faker = Factory.create()
        nameGroup = faker.bs()
        purposeGroup =  faker.catch_phrase()


        ############create new Valid group###############
        valid_group = {'name': nameGroup, 'purpose': purposeGroup}
        add_group_mutation = """
        mutation {
            addGroup(name: "%s", purpose: "%s")
        }
        """ % (valid_group['name'], valid_group['purpose'])
        res = self.client.send(add_group_mutation)


        #######Check if  a user in a created group is  valid########
        groupId = res['data']['addGroup']
        get_users_of_group_query = """
        query{
            usersOfGroup(groupId : %d){
                    userId    
            }
        }
        """ %(groupId)
        res = self.client.send(get_users_of_group_query)

        self.assertTrue(res['data']['usersOfGroup'][0]['userId'] == 1)


        #############Check: Can't create the new group having same name with the old one###############
        existed_group = {'name': nameGroup, 'purpose' : purposeGroup}
        add_group_mutation = """
        mutation {
            addGroup(name: "%s", purpose: "%s")
        }
        """ % (existed_group['name'], existed_group['purpose'])
        res = self.client.send(add_group_mutation)

        self.assertFalse(res['data']['addGroup'])



    def test_add_user_to_group(self):
        #####Generate random email######
        faker = Factory.create()
        email = faker.bs()
        groupId = 2

        ###### Check: Can not add unvalid(random) mail to group #######
        add_user_to_group_mutation = """
        mutation{
         addUserToGroup(email: "%s", groupId: %d)
        }
        """ %(email, groupId)
        res = self.client.send(add_user_to_group_mutation)

        self.assertFalse(res['data']['addUserToGroup'])



        ####Add valid user in group ####
        email = "De@gmail.com"

        add_user_to_group_mutation = """
        mutation{
         addUserToGroup(email: "%s", groupId: %d)
        }
        """ %(email, groupId)

        res = self.client.send(add_user_to_group_mutation)


        ###Check if added user has been in group yet
        get_users_of_group_query = """
        query{
            usersOfGroup(groupId : %d){
                    email
            }
        }
        """ %(groupId)
        res = self.client.send(get_users_of_group_query)
        emails = res['data']['usersOfGroup']
        emails = [ x['email'] for x in emails]
        self.assertTrue(email in emails)


        ### Can not add new existed user####
        email = "De@gmail.com"

        add_user_to_group_mutation = """
        mutation{
         addUserToGroup(email: "%s", groupId: %d)
        }
        """ %(email, groupId)

        res = self.client.send(add_user_to_group_mutation)

        self.assertFalse(res['data']['addUserToGroup'])

    def test_add_users_to_group(self):

        #Insert 2 NEW members to group 2 by valid emails
        email1 = "Shiki@gmail.com"
        email2 = "De@gmail.com"
        groupId = 3

        add_users_to_group_mutation = """
        mutation{
         addUsersToGroup(emails: ["%s","%s"], groupId: %d)
        }
        """ %(email1, email2, groupId)

        res = self.client.send(add_users_to_group_mutation)
        print("this is res: ", res)

        ###Check if added user has been in group yet
        get_users_of_group_query = """
        query{
            usersOfGroup(groupId : %d){
                    email
            }
        }
        """ %(groupId)
        res = self.client.send(get_users_of_group_query)
        emails = res['data']['usersOfGroup']
        emails = [ x['email'] for x in emails]
        self.assertTrue(email1 in emails and email2 in emails)


        ### Check :  can not add existed member to this group
        add_users_to_group_mutation = """
        mutation{
         addUsersToGroup(emails: ["%s","%s"], groupId: %d)
        }
        """ %(email1, email2, groupId)

        res = self.client.send(add_users_to_group_mutation)

        self.assertFalse(res['data']['addUsersToGroup'])


    def test_delete_user_in_group(self):

        email1 = "De@gmail.com"

        #Delete this user in group 1
        delete_user_in_group = """
        mutation{
            deleteUserInGroup(email: "%s", groupId: %d)
        }
        """%(email1, 1)
        res = self.client.send(delete_user_in_group)

        self.assertTrue(res['data']['deleteUserInGroup'])

        #Delete this user again
        delete_user_in_group = """
        mutation{
            deleteUserInGroup(email: "%s", groupId: %d)
        }
        """%(email1, 1)
        res = self.client.send(delete_user_in_group)

        self.assertFalse(res['data']['deleteUserInGroup'])















if __name__ == '__main__':
    unittest.main()
