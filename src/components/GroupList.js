/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import GroupItem from "./GroupItem";
export default function GroupList({groups, activeGroupId}) {
  return (
    <div>
      {groups.map((group) => {
          return (
            <GroupItem key={group.groupId} group={group} activeGroupId={activeGroupId}/>
          )
        }
      )}
    </div>
  )
}
