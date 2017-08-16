import json
import http.client

class ClientAPI(object):
    def __init__(self):
        self.client = http.client.HTTPConnection("localhost:8080")

        # send login credential
        headers = {'Content-type': 'application/json'}
        sign_in_json = json.dumps({'email': "Dandoh@gmail.com", 'password': 'haha'})
        self.client.request('POST', '/signin', sign_in_json, headers)
        response = json.loads(self.client.getresponse().
                read().decode('utf-8'))
        # store token for later uses
        self.token = response["token"]
        self.user_id = response["userId"]

        
    def send(self, query):
        #print('Sending query:')
        #print(query)
        headers = {'Authorization': 'Bearer ' + self.token}
        self.client.request('POST', '/graphql', json.dumps({'query': query}), headers)
        response = json.loads(self.client.getresponse().
                read().decode('utf-8'))
        return response

