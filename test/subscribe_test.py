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


    def testSubscribe(self):
        #############Get number comments##########################33
        reportId = 1
        reportId2 = 2
        commentContent = "Hi! You do a good job today"
        #While don't have any subsribes
        get_subscribes_query ="""
        query{
            subscribes{
                report{
                    reportId
                }
                userCommentLast{
                    userId
                }
                numberCommentsNotSeen
                lastUpdatedAt
            }
        }
        """

        add_comment_mutation="""
        mutation{
            addComment(content: "%s", reportId: %d)
        }
        """%(commentContent, reportId )
        res = self.client.send(add_comment_mutation)
        add_comment_mutation="""
        mutation{
            addComment(content: "%s", reportId: %d)
        }
        """%(commentContent, reportId2 )

        res = self.client.send(add_comment_mutation)
        res = self.client.send(get_subscribes_query)
        print("ressssssssss", res)
        self.assertTrue(len(res['data']['subscribes']) == 0)


        #While there is some body else comment
        res = self.client2.send(add_comment_mutation)
        res = self.client2.send(add_comment_mutation)


        ##Get the reports that are in interaction
        res = self.client.send(get_subscribes_query)
        self.assertTrue(res['data']['subscribes'][0]['numberCommentsNotSeen'] == 2)


        ##Other person comment
        res = self.client3.send(add_comment_mutation)
        res = self.client.send(get_subscribes_query)
        self.assertTrue(res['data']['subscribes'][0]['numberCommentsNotSeen'] == 3)
        print("hahahahah", res)
        self.assertTrue(res['data']['subscribes'][0]['userCommentLast']['userId'] == 3)

        ##When myself comment to the report
        res = self.client.send(add_comment_mutation)
        res = self.client.send(get_subscribes_query)
        self.assertTrue(len(res['data']['subscribes']) == 0)


        ##others comment to other report
        add_comment_mutation="""
        mutation{
            addComment(content: "%s", reportId: %d)
        }
        """%(commentContent, reportId2 )
        res = self.client2.send(add_comment_mutation)

        add_comment_mutation="""
        mutation{
            addComment(content: "%s", reportId: %d)
        }
        """%(commentContent, reportId )
        res = self.client3.send(add_comment_mutation)
        res = self.client.send(get_subscribes_query)
        self.assertTrue(len(res['data']['subscribes']) == 2)

        ###################Test saveSubscribes##################3
        save_query_mutation = """
        mutation{
            saveSubscribe(reportId : %d)
        }
        """%(reportId)
        res = self.client.send(save_query_mutation)
        res = self.client.send(get_subscribes_query)

        self.assertTrue(len(res['data']['subscribes']) == 1)

        save_query_mutation = """
        mutation{
            saveSubscribe(reportId : %d)
        }
        """%(reportId2)
        res = self.client.send(save_query_mutation)
        res = self.client.send(get_subscribes_query)
        self.assertTrue(len(res['data']['subscribes']) == 0)








if __name__ == '__main__':
    unittest.main()
