import React from 'react';
import {withRouter} from "react-router";
import {graphql} from "react-apollo";
import gql from "graphql-tag";

import UserList from '../components/UserList'
import Loading from '../components/Loading'
import Error from '../components/Error'

class UsersPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {userEmail: ''};
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  render() {
    let {loading, error, getUsersByGroupId} = this.props.data;
    if (error) {
      return (<Error/>)
    } else if (loading) {
      return (<Loading/>)
    } else {
      let users = getUsersByGroupId;
      return (
        <div>
          <UserList users={users}/>
          <form onSubmit={this.handleSubmit} className="form-horizontal">
            <fieldset>

              <h4>Add new user</h4>

              <div className="form-group">
                <label className="col-md-4 control-label" htmlFor="name">User Email</label>
                <div className="col-md-5">
                  <input id="userEmail" name="userEmail" type="text"
                         placeholder="e.g dainhan605@gmail.com" className="form-control input-md"
                         value={this.state.userEmail} onChange={this.handleChange}/>
                </div>
              </div>

              <div className="form-group">
                <label className="col-md-4 control-label" htmlFor="singlebutton"/>
                <div className="col-md-4">
                  <button id="singlebutton" name="singlebutton" className="btn btn-success">Add</button>
                </div>
              </div>

            </fieldset>
          </form>
        </div>
      )
    }
  }

  handleChange(event) {
    this.setState({userEmail: event.target.value});
  }

  handleSubmit(event) {
    event.preventDefault();
    let userEmail = this.state.userEmail;
    this.props.addUserByEmail({
      variables: {
        groupId: this.props.params.groupId,
        email: userEmail,
      }
    }).then((response) => {
      this.setState({userEmail: ""});
      this.props.data.refetch();
    }, (err) => {
      alert("Invalid user email");
    });
  }


}

const getGroupsQuery = gql` query GetUsersQuery($groupId: Int){
  getUsersByGroupId(id: $groupId) {
    userId
    name
    email
    note
  }
}`;

const withData = graphql(getGroupsQuery, {
  options: (ownProps) => {
    return {
      variables: {
        groupId: parseInt(ownProps.params.groupId)
      },
      forceFetch: true,
    }
  }
});

const addUserByEmail = gql`mutation AddUser($email: String, $groupId: Int){
  addUserByEmail(groupId: $groupId, email: $email)
}`;

const withMutation = graphql(addUserByEmail, {name: 'addUserByEmail'});

export default withMutation(withData(withRouter(UsersPage)));