/**
 * Created by Dandoh on 7/25/17.
 */
import React from "react";
import {graphql} from "react-apollo";
import gql from "graphql-tag";
import {withRouter} from "react-router";
import {Button, Glyphicon} from "react-bootstrap";
import {LinkContainer} from "react-router-bootstrap";
import GroupList from "../components/GroupList";

class NavigationPanel extends React.Component {
  constructor(props) {
    super(props);
    this.state = {newGroupName: ''};
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.logOut = this.logOut.bind(this);
  }


  render() {
    const subMenuStyle = {maxWidth: 400, margin: '0 auto 10px'};

    let activeGroupId = this.getActiveGroupId();
    let {loading, getGroups} = this.props.data;
    let groups = getGroups;

    const loadingView = (
      <div>Loading...</div>
    );

    const groupsView = (
      <div style={subMenuStyle}>
        <GroupList groups={groups} activeGroupId={activeGroupId}/>
      </div>
    );

    return (
      <div>
        <h4>Groups:</h4>
        {loading ? loadingView : groupsView}
        <form onSubmit={this.handleSubmit} className="form-horizontal">
          <fieldset>

            <h4>Create new group</h4>

            <div className="form-group">
              <label className="col-md-4 control-label" htmlFor="name">Group Name</label>
              <div className="col-md-5">
                <input id="newGroupName" name="newGroupName" type="text"
                       placeholder="group name" className="form-control input-md"
                       value={this.state.newGroupName} onChange={this.handleChange}/>
              </div>
            </div>

            <div className="form-group">
              <label className="col-md-4 control-label" htmlFor="singlebutton"/>
              <div className="col-md-4">
                <button id="singlebutton" name="singlebutton" className="btn btn-success">Create</button>
              </div>
            </div>

          </fieldset>
        </form>
        <Button onClick={this.logOut} bsStyle="default" block>Sign out</Button>
      </div>
    )
  }

  logOut() {
    localStorage.clear();
    this.props.router.replace("/sign-in");
  }

  getActiveGroupId() {
    let {location} = this.props;
    let regex = /^\/group\/(\d+)/gi;
    let match = regex.exec(location.pathname);
    let activeGroupId = -1;
    if (match) {
      activeGroupId = parseInt(match[1]);
    }

    return activeGroupId;
  }

  handleChange(event) {
    this.setState({newGroupName: event.target.value});
  }

  handleSubmit(event) {
    event.preventDefault();
    let newGroupName = this.state.newGroupName;
    this.props.createNewGroup({
      variables: {
        userId: 1,
        name: newGroupName,
      }
    }).then(() => {
      this.props.data.refetch();
      this.setState({newGroupName: ""});
    })
  }
}


const getGroupsQuery = gql` query {
  getGroups {
    groupId
    name
  }
}`;

const withData = graphql(getGroupsQuery, {
  options: {
    forceFetch: true,
  }
});

const createNewGroup = gql`mutation createNewGroup($name: String!){
  addGroup(name: $name)
}`;

const withMutation = graphql(createNewGroup, {name: 'createNewGroup'});

export default withMutation(withData(withRouter(NavigationPanel)));