import React from 'react';
import {withRouter} from "react-router";
import {LinkContainer} from 'react-router-bootstrap'

class SignUpPage extends React.Component {

  render() {
    return (
      <form className="form-horizontal">
        <fieldset>
          <legend>Sign-up</legend>

          <div className="form-group">
            <label className="col-md-4 control-label" htmlFor="email">Email</label>
            <div className="col-md-4">
              <input id="email" name="email" type="email" placeholder="e.g dandoh@gmail.com"
                     className="form-control input-md"/>
            </div>
          </div>

          <div className="form-group">
            <label className="col-md-4 control-label" htmlFor="username">Username</label>
            <div className="col-md-4">
              <input id="username" name="username" type="text" placeholder="e.g dandoh"
                     className="form-control input-md"/>
            </div>
          </div>

          <div className="form-group">
            <label className="col-md-4 control-label" htmlFor="password">Password</label>
            <div className="col-md-4">
              <input id="password" name="password" type="password" placeholder="password"
                     className="form-control input-md"/>

            </div>
          </div>

          <div className="form-group">
            <label className="col-md-4 control-label" htmlFor="sign-in"/>
            <div className="col-md-8">
              <button id="sign-up" name="sign-up" className="btn btn-primary">Sign up</button>
              <LinkContainer to={`/sign-in`}>
                <button id="sign-in" name="sign-in" className="btn btn-default">Sign in</button>
              </LinkContainer>
            </div>
          </div>

        </fieldset>
      </form>
    )
  }
}

export default SignUpPage