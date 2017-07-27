import React from 'react';
import {withRouter} from "react-router";
import '../style/TodoStyle.css'
import {graphql} from 'react-apollo'
import gql from 'graphql-tag'
import Loading from '../components/Loading'
import Error from '../components/Error'
import TodoList from '../components/TodoList';
import CommentList from '../components/CommentList';

import '../style/TodoStyle.css'
import '../style/CommentStyle.css'

class ReportDetailPage extends React.Component {
  constructor() {
    super();
    this.state = {
      summerization: "",
      comment: ""
    };

    this.onTick = this.onTick.bind(this);
    this.onDelete = this.onDelete.bind(this);
    this.onEnterTodo = this.onEnterTodo.bind(this);
    this.onChangeSummerization = this.onChangeSummerization.bind(this);
    this.onUpdateClick = this.onUpdateClick.bind(this);
    this.onCommentChange = this.onCommentChange.bind(this);
    this.onSubmitComment = this.onSubmitComment.bind(this);
  }


  render() {
    return (
      <div>
        {this.reportSection()}
        {/*comment here*/}
        <br/>
        <br/>
        <div className="container">
          <div className="row">
            <h4>Comment</h4>
          </div>
          <div className="row">
            <div className="col-md-6">
              <div className="widget-area no-padding blank">
                <div className="status-upload">
                  <form onSubmit={this.onSubmitComment}>
                    <textarea placeholder="Leave a comment" onChange={this.onCommentChange} value={this.state.comment}>
                    </textarea>
                    <button type="submit"
                            className="btn btn-success green"><i className="fa fa-share"/>
                      Comment
                    </button>
                  </form>
                </div>
              </div>
            </div>

          </div>
        </div>

        {this.commentSection()}
      </div>

    )
  }

  commentSection() {
    let {loading, error, getReportById} = this.props.data;
    if (error) {
      return (<Error/>)
    } else if (loading) {
      return (<Loading/>)
    } else {
      return (
        <div className="container">
          <CommentList comments={getReportById.comments}/>
        </div>
      )
    }
  }

  reportSection() {
    let {loading, error, getReportById} = this.props.data;
    if (error) {
      return (<Error/>)
    } else if (loading) {
      return (<Loading/>)
    } else {
      const ownUserId = localStorage.getItem("userId");
      const report = getReportById;
      const isMine = (ownUserId == report.user.userId);


      const mineLayout = (
        <div className="row container">
          <div className="col-md-6">
            <div className="todolist not-done">
              <h1>Daily report:</h1>
              <h4>Todo list:</h4>
              <input type="text" className="form-control add-todo"
                     placeholder="Add more task" onKeyDown={this.onEnterTodo}/>
              <TodoList todoes={this.state.todoes} hasTick={true}
                        editable={true} onTick={this.onTick} onDelete={this.onDelete}/>
              <h4>Summerization:</h4>
              <textarea className="form-control animated" onChange={this.onChangeSummerization}>
              {this.state.summerization}
              </textarea>
              <button className="btn btn-info pull-right"
                      onClick={this.onUpdateClick}
                      style={{marginTop: "10px"}} type="button">Update
              </button>
            </div>
          </div>
        </div>
      ); // if this report belong to the current user

      const otherLayout = (

        <div className="row container">
          <div className="col-md-6">
            <div className="todolist not-done">
              <h1>{`${report.user.name}'s Daily report:`}</h1>
              <h4>Todo list:</h4>
              <TodoList todoes={this.state.todoes} hasTick={true}
                        editable={false}/>
              <h4>Summerization:</h4>
              <p>{this.state.summerization}</p>
            </div>
          </div>
        </div>
      );

      return isMine ? mineLayout : otherLayout
    }
  }

  onTick(todo) {
    if (todo.state == 0) {
      todo.state = 1;
    } else {
      todo.state = 0;
    }
    this.setState({
      todoes: [...this.state.todoes]
    })
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

  onChangeSummerization(e) {
    e.preventDefault();
    this.setState({summerization: e.target.value})
  }

  componentWillReceiveProps(newProps) {
    // note: set data to the state after receive from server
    let {getReportById} = newProps.data;
    let report = getReportById;
    if (report && !this.state.todoes) {
      this.setState({
        todoes: report.todoes || [],
        summerization: report.summerization,
      });
    }
  }

  onUpdateClick() {
    this.props.updateReport({
      variables: {
        reportId: parseInt(this.props.params.reportId),
        contentTodoes: this.state.todoes.map(todo => todo.content),
        states: this.state.todoes.map(todo => todo.state),
        summerization: this.state.summerization,
        status: "Not decided this field yet" // !!
      }
    }).then(res => {
      alert("Update daily report successfully");
    }, err => {
      alert("Can't update report");
    })
  }

  onCommentChange(e) {
    e.preventDefault();
    this.setState({comment: e.target.value});
  }

  onSubmitComment(e) {
    e.preventDefault();
    this.props.createComment({
      variables: {
        content: this.state.comment,
        reportId: this.props.params.reportId
      }
    }).then(res => {
      this.props.data.refetch();
      this.setState({
        comment: "",
      });
    }, err => {
      alert("Something wrong happened, dude!")
    })
  }


}

const getReportDetailQuery = gql`query 
  GetReportQuery($id: Int) {
    getReportById(id: $id) {
      reportId
      user {
        userId
        name
      }
      todoes {
        todoId
        content
        state
      }
      summerization
      comments {
        user {
          name
          userId
        }
        commentId
        content
      }
    }
  }`;

const withData = graphql(getReportDetailQuery, {
  options: (ownProps) => {
    return {
      variables: {
        id: parseInt(ownProps.params.reportId)
      },
      forceFetch: true,
    }
  }
});

const updateReport = gql`mutation 
  UpdateReport($reportId: Int, $contentTodoes: [String], $states: [Int], $summerization: String, $status: String) {
    updateReport(contentTodoes: $contentTodoes, states: $states, reportId: $reportId,
                 summerization: $summerization, status: $status)
  }
`;


const createComment = gql`mutation
  CreateComment($content: String, $reportId: Int) {
    createComment(content: $content, reportId: $reportId)
  }`;

const withUpdateReport = graphql(updateReport, {name: 'updateReport'});

const withCreateComment = graphql(createComment, {name: 'createComment'});


export default withCreateComment(
  withUpdateReport(
    withData(
      withRouter(
        ReportDetailPage
      )
    )
  )
);