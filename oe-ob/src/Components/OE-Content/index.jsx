import React, { Component, Fragment } from 'react';
import { Switch, Route } from 'react-router-dom';

import Home from './Home/index';
import ActivityHome from './Activity/Home/index';
import ActivityTimeLimited from './Activity/TimeLimited/index';
import SpecialtyPriceSetting from './Specialty/PriceSetting/index';
import User from './User/index';
import SystemConfig from './SystemConfig/index';

export default class extends Component {
  render() {
    return (
      <Fragment>
        <Switch>
          <Route exact path="/" component={Home}></Route>
          <Route path="/Home" component={Home}></Route>
          <Route path="/Activity/Home" component={ActivityHome} />
          <Route path="/Activity/TimeLimited" component={ActivityTimeLimited} />
          <Route
            path="/Specialty/PriceSetting"
            component={SpecialtyPriceSetting}
          />
          <Route path="/User" component={User} />
          <Route path="/SystemConfig" component={SystemConfig} />
        </Switch>
      </Fragment>
    );
  }
}
