import unittest
import client
from faker import Factory

class TestReport(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.client = client.ClientAPI()

    def test_upper(self):
        self.assertEqual('foo'.upper(), 'FOO')


    def test_report_todo_comment_note(self):

        ####Create new report and Get new reportID
        create_report_mutation = """
        mutation {
            createReport
        }
        """
        res = self.client.send(create_report_mutation)
        reportId = res['data']['createReport']
        print("report cua tao: ", reportId)

        ####Test get report today##############
        get_reportToday_query = """
        query{
            reportToday{
                reportId
            }
        }
        """
        res = self.client.send(get_reportToday_query)
        self.assertTrue(res['data']['reportToday']['reportId'] == reportId)






        ###############Test update note###############

        note = "Note: Remember to finish classify tasks"
        update_note_mutation = """
        mutation{
            updateNote(note: "%s", reportId: %d)
        }
        """%(note, reportId)

        self.client.send(update_note_mutation)

        get_todo_query = """
        query{
            report(reportId: %d){
                note
                todoes{
                    todoId
                    state
                    content
                    estimateTime
                    spentTime
                }
                comments{
                    content
                }
            }
        }
        """%(reportId)

        res = self.client.send(get_todo_query)
        self.assertTrue(res['data']['report']['note'] == note)


        #################Test CreateComment##############
        comment = "This is comment: Try harder!"
        create_comment_mutation = """
        mutation{
            addComment(content: "%s", reportId: %d)
        }
        """%(comment, reportId)
        res = self.client.send(create_comment_mutation)

        res = self.client.send(get_todo_query)

        self.assertTrue(res['data']['report']['comments'][0]['content'] == comment)





        ###############Test add todo  into the created report##################
        content = "Learn statistic"
        state = 1
        estimateTime = 120
        spentTime = 120
        add_todo_mutation = """
        mutation{
            addTodo(content: "%s", state: %d, estimateTime: %d, spentTime: %d, reportId: %d)
        }
        """%(content, state, estimateTime, spentTime, reportId)
        res = self.client.send(add_todo_mutation)
        todoId = res['data']['addTodo']



        res = self.client.send(get_todo_query)
        todo = res['data']['report']['todoes'][0]

        self.assertTrue(todo['todoId']== todoId and todo['state']== state and todo['content'] == content
                    and todo['estimateTime'] == estimateTime and todo['spentTime'] == spentTime)



       ################### Test Update created todo #######################
        newContent = "Learn statistic and Machine learning"
        newState = 0
        newEstimateTime = 200
        newSpentTime = 100
        update_todo_mutation = """
        mutation{
            updateTodo(todoId: %d, content: "%s", state: %d, estimateTime: %d, spentTime: %d)
        }
        """%(todoId, newContent, newState, newEstimateTime, newSpentTime)

        res = self.client.send(update_todo_mutation)

        res = self.client.send(get_todo_query)
        todo = res['data']['report']['todoes'][0]

        self.assertTrue(todo['todoId']== todoId and todo['state']== newState and todo['content'] == newContent
                        and todo['estimateTime'] == newEstimateTime and todo['spentTime'] == newSpentTime)



        ############# Test delete Todo#################
        delete_todo_mutation="""
        mutation{
            deleteTodo(todoId: %d)
        }
        """%(todoId)
        res = self.client.send(delete_todo_mutation)

        res = self.client.send(get_todo_query)

        print("deleteTodo :   ", res)
        todoes = res['data']['report']['todoes']
        self.assertTrue(len(todoes) == 0)







if __name__ == '__main__':
    unittest.main()