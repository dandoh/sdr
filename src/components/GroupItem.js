/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import {LinkContainer} from "react-router-bootstrap";
import {Button} from "react-bootstrap";

export default function GroupItem({group, activeGroupId}) {
  if (group.groupId == activeGroupId) {
    return (
      <LinkContainer active to={`/group/${group.groupId}`}>
        <Button bsStyle="primary" block>{group.name}</Button>
      </LinkContainer>
    )
  } else {
    return (
      <LinkContainer to={`/group/${group.groupId}`}>
        <Button block>{group.name}</Button>
      </LinkContainer>
    )
  }
}



