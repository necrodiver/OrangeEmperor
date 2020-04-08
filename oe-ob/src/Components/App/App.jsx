import React, { Component, Fragment } from 'react';
import './App.less';
import OEHeader from '../OE-Header/index';
import OEMenu from '../OE-Menu/index';
import OELogin from '../OE-Login/index';

import { Layout } from 'antd';
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons';

const { Header, Sider, Content } = Layout;

export default class extends Component {
  state = {
    collapsed: false,
    isLogin: true,
  };

  toggle = () => {
    this.setState({
      collapsed: !this.state.collapsed,
    });
  };
  render() {
    return (
      <div className="oe-app">
        {!this.state.isLogin ? (
          <OELogin />
        ) : (
          <Fragment>
            <Layout style={{ height: '100vh' }}>
              <Sider
                trigger={null}
                collapsible
                collapsed={this.state.collapsed}
              >
                <OEMenu collapsed={this.state.collapsed} />
              </Sider>
              <Layout className="site-layout">
                <Header
                  className="site-layout-background oe-header-box"
                  style={{ padding: 0 }}
                >
                  {React.createElement(
                    this.state.collapsed
                      ? MenuUnfoldOutlined
                      : MenuFoldOutlined,
                    {
                      className: 'trigger oe-header-btnmenu',
                      onClick: this.toggle,
                    }
                  )}
                  <OEHeader></OEHeader>
                </Header>
                <Content className="site-layout-background oe-content-box">
                  Content
                </Content>
              </Layout>
            </Layout>
          </Fragment>
        )}
      </div>
    );
  }
}
