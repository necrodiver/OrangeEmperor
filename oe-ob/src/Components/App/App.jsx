import React, { Component, Fragment } from 'react';
import './App.css';
import './App.less';
import OEHeader from '../OE-Header/index';
import OEMenu from '../OE-Menu/index';
import OELogin from '../OE-Login/index';

export default class extends Component {
  constructor() {
    super();
    this.state = {
      isLogin: true
    };
  }
  render() {
    return (
      <div className="oe-app">
        {!this.state.isLogin ? (
          <OELogin />
        ) : (
          <Fragment>
            <header className="oe-header">
              <OEHeader />
            </header>
            <section  className="oe-section-body">
              <aside className="oe-menu">
                <OEMenu />
              </aside>
              <main className="oe-content">
                <h1>这里是内容区域</h1>
              </main>
            </section>
          </Fragment>
        )}
      </div>
    );
  }
}
