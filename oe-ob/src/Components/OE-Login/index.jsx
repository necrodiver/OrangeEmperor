import React, { Component, Fragment } from 'react';

export default class extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: 'OE-Login'
    };
  }
  render() {
    return (
      <Fragment>
        <div className="oe-login">
          <h1>{this.state.name}</h1>
        </div>
      </Fragment>
    );
  }
}
