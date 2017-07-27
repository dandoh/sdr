/**
 * Created by Dandoh on 7/21/17.
 */
import React from "react";
import ReportItem from "./ReportItem";
export default function ReportList({reports}) {
  return (
    <div>
      <ul className="list-group">
        {reports.map((report) => {
            return (
              <ReportItem key={report.reportId}
                          report={report}/>
            )
          }
        )}
      </ul>
    </div>
  )
}