/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import {LinkContainer} from "react-router-bootstrap";
import {Button} from 'react-bootstrap';


export default function CommentItem({comment}) {
  return (
    <div className="container">
      <div className="row">
        <div className="col-sm-1">
          <div className="thumbnail">
            <img className="img-responsive user-photo" src="https://ssl.gstatic.com/accounts/ui/avatar_2x.png"/>
          </div>
        </div>

        <div className="col-sm-8">
          <div className="panel panel-default">
            <div className="panel-heading">
              <strong>{comment.user.name}</strong> <span className="text-muted">commented</span>
            </div>
            <div className="panel-body">
              {comment.content}
            </div>
          </div>
        </div>
      </div>

    </div>
  )
}



