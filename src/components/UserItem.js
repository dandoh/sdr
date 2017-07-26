/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import {LinkContainer} from "react-router-bootstrap";
import {Button} from "react-bootstrap";

export default function UserItem({user}) {
  return (
    <h5>{user.name + "--------" + user.email}</h5>
  )
}



