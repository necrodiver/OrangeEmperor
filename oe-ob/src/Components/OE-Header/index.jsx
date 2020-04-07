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
          <span>{this.state.name}</span>
        </div>
      </Fragment>
    );
  }
}
