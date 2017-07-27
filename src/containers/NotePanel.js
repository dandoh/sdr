import React from 'react';
import {withRouter} from "react-router";
import {graphql} from 'react-apollo'
import gql from 'graphql-tag'
import Loading from '../components/Loading'
import Error from '../components/Error'

class NotePanel extends React.Component {
  render() {
    return (
      <div>
        <h3>Personal Note:</h3>
        {this.content()}
      </div>
    )
  }

  content() {
    let {loading, error, getNote, updateNote} = this.props.data;
    if (error) {
      return (<Error/>)
    } else if (loading) {
      return (<Loading/>)
    } else {
      return <textarea rows="20" cols="40" onBlur={this.handleUpdateNote}>{getNote}</textarea>
    }
  }

  handleUpdateNote = (event) => {
    let newNote = event.target.value;
    this.props.updateNote({
      variables: {
        note: newNote,
      }
    })
  };
}

const getNote = gql`query
  GetNote{
    getNote
  }
`;

const updateNote = gql`mutation
  UpdateNote($note: String){
    updateNote(note: $note)
  }
`;

const withData = graphql(getNote, {
  options: {
    forceFetch: true,
  }
});

const withMutation = graphql(updateNote, {name: 'updateNote'});

export default withMutation(withData(
  withRouter(NotePanel)))