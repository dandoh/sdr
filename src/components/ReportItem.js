/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
export default function ReportItem({report}) {
  return (
    <div>
      <div style={{width: "100%", overflow: "auto"}}>
        <div style={{float: "left", width: "30%"}}>{report.user.name}</div>
        <div style={{float: "left", width: "30%"}}>{JSON.stringify(report.todoes)}</div>
        <div style={{float: "left"}}>{report.date}</div>
      </div>
    </div>
  )
}



