import React from 'react';
import {withRouter} from "react-router";
import {graphql} from "react-apollo";
import gql from "graphql-tag";
import './TodoStyle.css'
import {Button, Glyphicon} from "react-bootstrap";
import TodoList from '../components/TodoList';

class CreateReportPage extends React.Component {
  constructor() {
    super();
    this.state = {
      todoes: []
    };

    this.onTick = this.onTick.bind(this);
    this.onDelete = this.onDelete.bind(this);
    this.onEnterTodo = this.onEnterTodo.bind(this);
    this.submitReport = this.submitReport.bind(this);
  }

  onTick(todo) {
    // nothing
  }


  onDelete(deleletedTodo) {
    this.setState({
      todoes: this.state.todoes.filter(todo => todo != deleletedTodo)
    });
  }


  onEnterTodo(e) {
    if (e.which == 13) {
      let content = e.target.value;
      let newTodo = {
        content,
        state: 0
      };
      e.target.value = "";
      this.setState({
        todoes: this.state.todoes.concat(newTodo)
      });
    }
  }


  render() {

    return (
      <div>
        <div className="row container">
          <div className="col-md-6">
            <div className="todolist not-done">
              <h1>Daily report:</h1>
              <h2>Plan your day work:</h2>
              <input type="text" className="form-control add-todo"
                     placeholder="Add a task" onKeyDown={this.onEnterTodo}
              />
              <TodoList todoes={this.state.todoes} tickable={false}
                        editable={true} onTick={this.onTick} onDelete={this.onDelete}/>
            </div>
            <button className="btn btn-info pull-right" style={{marginTop: "10px"}}
                    onClick={this.submitReport}
                    type="button">Submit
            </button>
          </div>
        </div>

      </div>
    )

  }

  submitReport(e) {
    e.preventDefault();
    this.props.createNewReport({
      variables: {
        contentTodoes: this.state.todoes.map(todo => todo.content),
        states: this.state.todoes.map(todo => todo.state),
        summerization: "",
        groupId: parseInt(this.props.params.groupId)
      }
    }).then(res => {
      this.props.router.goBack();
    }, err => {
      alert("Can't create daily report");
    })
  }
}

const createNewReport = gql`mutation 
  CreateNewReport($contentTodoes: [String], $states: [Int], $groupId: Int) {
    createReport(contentTodoes: $contentTodoes, states: $states, groupId: $groupId)
  }
`;

const withMutation = graphql(createNewReport,
  {name: 'createNewReport'}
);

export default withMutation(withRouter(CreateReportPage))