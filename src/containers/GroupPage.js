import React from 'react';
import {withRouter} from "react-router";
import {graphql} from 'react-apollo'
import gql from 'graphql-tag'
import ReportList from '../components/ReportList'
import Loading from '../components/Loading'
import Error from '../components/Error'
import {LinkContainer} from "react-router-bootstrap";
import {Button} from "react-bootstrap"

class GroupPage extends React.Component {

  render() {
    let {groupId} = this.props.params;
    let {loading, error, getReportsByGroupId} = this.props.data;
    if (error) {
      return (<Error/>)
    } else if (loading) {
      return (<Loading/>)
    } else {
      let reports = getReportsByGroupId;
      return (
        <div>
          <ReportList reports={reports}/>
          <LinkContainer to={`/group/${groupId}/create_report`}>
            <Button>Create new report</Button>
          </LinkContainer>
          <LinkContainer to={`/group/${groupId}/users`}>
            <Button>Users</Button>
          </LinkContainer>
        </div>
      )
    }
  }
}

const getReportsQuery = gql`query 
  GetReportsQuery($id: Int) {
    getReportsByGroupId(id: $id) {
      reportId
      summerization
      user {
        name
      }
      todoes {
        content
        state
      }
    }
  }`;

const withData = graphql(getReportsQuery, {
  options: (ownProps) => {
    return {
      variables: {
        id: parseInt(ownProps.params.groupId)
      },
      forceFetch: true,
    }
  }
})(withRouter(GroupPage));


export default withData;