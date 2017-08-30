import unittest
import client
from faker import Factory

class TestReport(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.client = client.ClientAPI("Dandoh@gmail.com")

    def test_upper(self):
        self.assertEqual('foo'.upper(), 'FOO')


    def test_GetNewReport(self):
        ##Get today report
        get_today_report_query = """
        query{
            reportToday{
                reportId
                todoes{
                    todoId
                    state
                }
            }
        }
        """
        res = self.client.send(get_today_report_query)
        todoes = res['data']['reportToday']['todoes']
        states = [x['state'] for x in todoes]
        ##Number of task finished
        num = sum(states)
        #print(res)
        #self.assertTrue(num == 1)


        ##Create new report automatically: Change the time in main.go and run main.go again.
        print("respone for res", res)
        #self.assertTrue(num == 0 and len(todoes) == 2)






if __name__ == '__main__':
    unittest.main()