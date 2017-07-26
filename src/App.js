import React from 'react';
import NavigationPanel from './containers/NavigationPanel'
import NotePanel from './containers/NotePanel'

class App extends React.Component {
  render() {

    let {children, location} = this.props;
    let key = -1;
    let regex = /^\/group\/(\d+)$/gi;
    let match = regex.exec(location.pathname);
    if (match) key = match[1];
    return (
      <div style={{width: "100%", overflow: "auto"}}>
        <div style={{float: "left", width: "20%"}}>
          <NavigationPanel location={location}/>
        </div>
        <div style={{float: "left", width: "45%"}} key={key}>
          {children ? children : null}
        </div>
        <div style={{float: "left"}}>
          <NotePanel/>
        </div>
      </div>
    )
  }
}

export default App;
