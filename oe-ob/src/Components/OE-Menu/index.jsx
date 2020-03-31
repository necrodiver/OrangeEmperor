import React, { Component, Fragment } from 'react';

export default class extends Component {
  constructor(refs) {
    super(refs);
    this.state = {
      name: 'OE-Menu'
    };
  }
  render() {
    return (
      <Fragment>
        <div className="oe-menu">
          <h1>{this.state.name}</h1>
        </div>
      </Fragment>
    );
  }
}
