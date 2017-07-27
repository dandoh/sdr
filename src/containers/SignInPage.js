import React from 'react';
import {withRouter} from "react-router";
import {LinkContainer} from 'react-router-bootstrap'

class SignInPage extends React.Component {
  constructor() {
    super();
    this.state = {
      username: "",
      password: "",
    };
    this.handleChangeUsername = this.handleChangeUsername.bind(this);
    this.handleChangePassword = this.handleChangePassword.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit} className="form-horizontal">
        <fieldset>
          <legend>Sign-in</legend>

          <div className="form-group">
            <label className="col-md-4 control-label" htmlFor="username">Username</label>
            <div className="col-md-4">
              <input id="username" name="username" type="text" placeholder="e.g dandoh"
                     className="form-control input-md" value={this.state.username}
                     onChange={this.handleChangeUsername}/>
            </div>
          </div>

          <div className="form-group">
            <label className="col-md-4 control-label" htmlFor="password">Password</label>
            <div className="col-md-4">
              <input id="password" name="password" type="password" placeholder="password"
                     className="form-control input-md" value={this.state.password}
                     onChange={this.handleChangePassword}/>

            </div>
          </div>

          <div className="form-group">
            <label className="col-md-4 control-label" htmlFor="sign-in"/>
            <div className="col-md-8">
              <button id="sign-in" name="sign-in" className="btn btn-primary">Sign in</button>
              <LinkContainer to={`/sign-up`}>
                <button id="sign-up" name="sign-up" className="btn btn-default">Sign up</button>
              </LinkContainer>
            </div>
          </div>

        </fieldset>
      </form>
    )
  }

  handleChangeUsername(event) {
    this.setState({username: event.target.value});
  }

  handleChangePassword(event) {
    this.setState({password: event.target.value});
  }

  handleSubmit(event) {
    event.preventDefault();
    fetch("http://localhost:8080/login", {
      headers: {
        'Accept': 'application/json, text/plain, */*',
        'Content-Type': 'application/json',
        // 'Access-Control-Allow-Credentials': 'true',
        // 'Access-Control-Allow-Origin':'*'
      },
      method: "POST",
      body: JSON.stringify({username: this.state.username, password: this.state.password})
    }).then(res => {
      if (res.status != 200) {
        alert("Can't sign in");
      } else {
        res.json().then(
          json => {
            localStorage.setItem('token', json.token);
            localStorage.setItem('userId', json.userId);
            this.props.router.replace("/");
          }
        )
      }
    })
  }
}

export default withRouter(SignInPage)