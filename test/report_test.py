import unittest
import client
from faker import Factory

class TestReport(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.client = client.ClientAPI("Dandoh@gmail.com")

    def test_upper(self):
        self.assertEqual('foo'.upper(), 'FOO')


    def test_report_todo_comment_note(self):



        ####Test get report today##############
        reportId = 1
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
        data = res['data']['report']['comments']
        data = [x['content'] for x in data]
        print("comment: ", res)
        self.assertTrue(comment in data)





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
        leng = len(res['data']['report']['todoes'])
        todo = res['data']['report']['todoes'][leng-1]

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
        leng= len (res['data']['report']['todoes'])
        todo = res['data']['report']['todoes'][leng-1]

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
        todoIds = [x['todoId'] for x in todoes]
        self.assertFalse(todoId in todoIds)



    def test_reports(self):

        ##Test get all reports of user
        get_all_reports_of_user="""
        query{
            reports{
                reportId
                todoes{
                    estimateTime
                }
            }
        }
        """
        res = self.client.send(get_all_reports_of_user)
        print(res)
        self.assertIsNotNone(res['data'], msg = None)

        ##Test get all old reports in (date1, date2)
        fromDate = "2017-08-16"
        toDate = "2017-08-17"
        get_all_old_reports_in_period="""
        query{
            oldReports(fromDate: "%s", toDate: "%s"){
                reportId
            }
        }
        """%(fromDate, toDate)

        res = self.client.send(get_all_old_reports_in_period)

        self.assertTrue(len(res['data']['oldReports']) ==0 )


        ##Change the date
        toDate = "3000-08-17"

        get_all_old_reports_in_period="""
        query{
            oldReports(fromDate: "%s", toDate: "%s"){
                reportId
            }
        }
        """%(fromDate, toDate)


        res = self.client.send(get_all_old_reports_in_period)

        self.assertFalse(len(res['data']['oldReports']) ==0 )

        ####Test get reports Today of group
        groupId = 1
        get_reports_to_day_of_group="""
        query{
            reportsTodayOfGroup(groupId: %d){
                reportId
                user{
                    userId
                    name
                }
            }
        }
        """%(groupId)

        print("report today groups: ", res)
        res = self.client.send(get_reports_to_day_of_group)
        self.assertIsNotNone(res['data'], msg = None)
        #print(res)





if __name__ == '__main__':
    unittest.main()