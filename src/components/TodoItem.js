/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import {LinkContainer} from "react-router-bootstrap";
import {Button, Glyphicon} from "react-bootstrap";

export default function TodoItem({todo, editable, hasTick, onTick, onDelete}) {
  const tickBox = (
    <input onClick={(e) => {
      onTick(todo)
    }} type="checkbox" checked={todo.state == 1} disabled={!editable}/>
  );
  const deleteButton = (
    <button onClick={(e) => {
      e.preventDefault();
      onDelete(todo)
    }} className="remove-item btn btn-default btn-xs pull-right">
      <Glyphicon glyph="remove"/>
    </button>
  );

  const textStyle = {textDecoration: todo.state == 1 ? "line-through" : "none"};
  return (
    <li className="ui-state-default">
      <div className="checkbox">
        {hasTick ? tickBox : null}
        <label style={textStyle}>  {todo.content}</label>
        {editable ? deleteButton : null}
      </div>
    </li>
  )
}



