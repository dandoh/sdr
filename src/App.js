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
      <div className="row">
        <div className="col-sm-3" >
          <NavigationPanel location={location}/>
        </div>
        <div className="col-sm-5" key={key}>
          {children ? children : null}
        </div>
        <div className="col-sm-4">
          <NotePanel/>
        </div>
      </div>
    )
  }
}

export default App;
