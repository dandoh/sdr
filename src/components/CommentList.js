/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import CommentItem from "./CommentItem";
export default function CommentList({comments}) {
  return (
    <div>
      {comments.map(comment => {
        return (
          <CommentItem comment={comment} key={comment.commentId}/>
        )
      })}
    </div>
  )
}