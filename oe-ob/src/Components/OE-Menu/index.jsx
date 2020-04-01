import React, { Component, Fragment } from 'react';

import { Menu, Button } from 'antd';
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  AppstoreOutlined,
  PieChartOutlined,
  DesktopOutlined,
  ContainerOutlined,
  MailOutlined
} from '@ant-design/icons';

const { SubMenu } = Menu;

export default class extends Component {
  state = {
    collapsed: false,
    menuKeys: ['1'],
    menuOpenKeys: ['sub1']
  };

  constructor(props) {
    super(props);
    this.setState({
      collapsed: props.collapsed
    });
  }
  render() {
    return (
      <Fragment>
        <Menu
          defaultSelectedKeys={this.state.menuKeys}
          defaultOpenKeys={this.state.menuOpenKeys}
          mode="inline"
          theme="dark"
          inlineCollapsed={this.state.collapsed}
        >
          <Menu.Item key="1">
            <PieChartOutlined />
            <span>选项 1</span>
          </Menu.Item>
          <Menu.Item key="2">
            <DesktopOutlined />
            <span>选项 2</span>
          </Menu.Item>
          <Menu.Item key="3">
            <ContainerOutlined />
            <span>选项 3</span>
          </Menu.Item>
          <SubMenu
            key="sub1"
            title={
              <span>
                <MailOutlined />
                <span>菜单 1</span>
              </span>
            }
          >
            <Menu.Item key="5">选项 5</Menu.Item>
            <Menu.Item key="6">选项 6</Menu.Item>
            <Menu.Item key="7">选项 7</Menu.Item>
            <Menu.Item key="8">选项 8</Menu.Item>
          </SubMenu>
          <SubMenu
            key="sub2"
            title={
              <span>
                <AppstoreOutlined />
                <span>菜单 2</span>
              </span>
            }
          >
            <Menu.Item key="9">选项 9</Menu.Item>
            <Menu.Item key="10">选项 10</Menu.Item>
            <SubMenu key="sub3" title="Submenu">
              <Menu.Item key="11">选项 11</Menu.Item>
              <Menu.Item key="12">选项 12</Menu.Item>
            </SubMenu>
          </SubMenu>
        </Menu>
      </Fragment>
    );
  }
}
