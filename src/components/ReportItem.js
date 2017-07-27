/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import {LinkContainer} from "react-router-bootstrap";
import {Button} from 'react-bootstrap';

export default function ReportItem({report}) {
  return (
    <LinkContainer to={`/group/${report.group.groupId}/report/${report.reportId}`} className="list-group-item">
      <Button style={{width: "100%", overflow: "auto"}}>
        <div style={{float: "left", width: "30%"}}>{report.user.name}</div>
        <div style={{float: "left", width: "40%"}}>
          <ul>
            {report.todoes.map((todo, i) => {
              return (
                <li key={i}>{todo.content}</li>
              )
            })}
          </ul>
        </div>
        <div style={{float: "left"}}>{report.date}</div>
      </Button>
    </LinkContainer>
  )
}



