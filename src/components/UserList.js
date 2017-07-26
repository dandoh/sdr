/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import UserItem from "./UserItem";
export default function UserList({users}) {
  return (
    <div>
      {users.map((user) => {
          return (
            <UserItem key={user.userId} user={user}/>
          )
        }
      )}
    </div>
  )
}