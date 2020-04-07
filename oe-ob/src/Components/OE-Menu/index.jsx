import './index.less';
import React, { Component, Fragment } from 'react';
import { Menu } from 'antd';
import PropTypes from 'prop-types';
const { SubMenu } = Menu;

export default class extends Component {
  // static defaultProps = {
  //   collapsed: true
  // };
  // static propTypes = {
  //   collapsed: PropTypes.bool.isRequired
  // };
  state = {
    collapsed: true,
    menuKeys: ['Activity'],
    menuOpenKeys: [],
    menuList: [
      {
        keyCode: 'Activity',
        name: '活动管理',
        url: 'Activity',
        icon: '&#xe61d;',
        childList: [],
      },
      {
        keyCode: 'Specialty',
        name: '产品管理',
        url: 'Specialty',
        icon: '&#xe60b;',
        childList: [
          {
            keyCode: 'Specialty-PriceSetting',
            name: '价格配置',
            url: 'priceSetting',
            childList: [],
          },
        ],
      },
      {
        keyCode: 'User',
        name: '用户管理',
        url: 'User',
        icon: '&#xe650;',
        childList: [],
      },
      {
        keyCode: 'SystemConfig',
        name: '系统配置',
        url: 'SystemConfig',
        icon: '&#xe62a;',
        childList: [],
      },
    ],
  };
  menuSelect = ({ item, key, keyPath, selectedKeys, domEvent }) => {
    console.log('选项key', key);
  };
  render() {
    let menuItem;
    return (
      <Fragment>
        <div className="menu-logo">这里放logo</div>
        {this.state.menuList && this.state.menuList.length > 0 && (
          <Menu
            defaultSelectedKeys={this.state.menuKeys}
            defaultOpenKeys={this.state.menuOpenKeys}
            mode="inline"
            theme="dark"
            onSelect={this.menuSelect}
          >
            {
              (menuItem = (childItem) => {
                return (
                  <Menu.Item key={childItem.keyCode}>
                    {childItem.icon && (
                      <i
                        className="iconfont"
                        dangerouslySetInnerHTML={{ __html: childItem.icon }}
                      ></i>
                    )}
                    <span className="pl-6">{childItem.name}</span>
                  </Menu.Item>
                );
              })
            }
            {this.state.menuList.map((item) => {
              if (item.childList && item.childList.length > 0) {
                return (
                  <SubMenu
                    key={item.keyCode}
                    title={
                      <span>
                        <i
                          className="iconfont"
                          dangerouslySetInnerHTML={{ __html: item.icon }}
                        ></i>
                        <span className="pl-6">{item.name}</span>
                      </span>
                    }
                  >
                    {item.childList.map((childItem) => {
                      return menuItem(childItem);
                    })}
                  </SubMenu>
                );
              }
              return menuItem(item);
            })}
          </Menu>
        )}
      </Fragment>
    );
  }
}
