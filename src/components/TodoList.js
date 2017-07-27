/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import TodoItem from './TodoItem';

export default function TodoList({
  todoes, editable,
  hasTick, onTick, onDelete
}) {
  return (
    <ul id="sortable" className="list-unstyled">
      {todoes.map((todo, i) => (<TodoItem
        key={i}
        todo={todo}
        editable={editable}
        hasTick={hasTick}
        onTick={onTick}
        onDelete={onDelete}/>))}
    </ul>
  )
}



