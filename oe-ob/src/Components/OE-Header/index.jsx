import React, { Fragment, Component } from 'react';
export default class extends Component {
  constructor(refs) {
    super(refs);
    this.state = {
      name: 'OE-Header'
    };
  }
  render() {
    return (
      <Fragment>
        <div className="oe-header">
          <h1>{this.state.name}</h1>
        </div>
      </Fragment>
    );
  }
}
