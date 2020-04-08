import './index.less';
import React, { Component, Fragment } from 'react';
import { Menu } from 'antd';
import PropTypes from 'prop-types';
import Icon from '@ant-design/icons';
const { SubMenu } = Menu;

export default class extends Component {
  static defaultProps = {
    collapsed: true,
  };
  static propTypes = {
    collapsed: PropTypes.bool.isRequired,
  };
  state = {
    collapsed: true,
    menuKeys: ['Activity'],
    subMenuKeys: [],
    menuList: [
      {
        keyCode: 'Home',
        name: '首页',
        url: 'Home',
        icon: '&#xe605;',
        childList: [],
      },
      {
        keyCode: 'Activity',
        name: '活动管理',
        url: 'Activity',
        icon: '&#xe61d;',
        childList: [
          {
            keyCode: 'Activity-Index',
            name: '首页活动',
            url: 'Activity/Index',
          },
          {
            keyCode: 'Activity-TimeLimited',
            name: '限时活动',
            url: 'Activity/TimeLimited',
          },
        ],
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
            url: 'Specialty/priceSetting',
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
  constructor(props) {
    super(props);
    this.state.collapsed = props.collapsed;
  }
  // 菜单选择事件
  menuSelect = ({ item, key, keyPath, selectedKeys, domEvent }) => {
    this.setState({
      menuKeys: [key],
    });
  };
  // 菜单整体选择事件
  openChangeSubMenu = (openkeys) => {
    this.setState({
      subMenuKeys: openkeys,
    });
  };
  componentWillReceiveProps(nextProps) {
    this.setState({ collapsed: nextProps.collapsed });
  }
  render() {
    let menuItem;
    console.log('this.state.collapsed', this.state.collapsed);
    return (
      <Fragment>
        <div className="menu-logo">
          {!this.state.collapsed ? '大' : '小'}logo
        </div>
        {this.state.menuList && this.state.menuList.length > 0 && (
          <Menu
            selectedKeys={this.state.menuKeys}
            openkeys={this.state.subMenuKeys}
            mode="inline"
            theme="dark"
            onSelect={this.menuSelect}
            onOpenChange={this.openChangeSubMenu}
          >
            {
              (menuItem = (childItem, isSubMenu) => {
                return (
                  <Menu.Item key={childItem.keyCode}>
                    {childItem.icon && (
                      <i
                        className="iconfont"
                        dangerouslySetInnerHTML={{ __html: childItem.icon }}
                      ></i>
                    )}
                    {(!this.state.collapsed || isSubMenu) && (
                      <span className="pl-6">{childItem.name}</span>
                    )}
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
                        {!this.state.collapsed && (
                          <span className="pl-6">{item.name}</span>
                        )}
                      </span>
                    }
                  >
                    {item.childList.map((childItem) => {
                      return menuItem(childItem, true);
                    })}
                  </SubMenu>
                );
              }
              return menuItem(item, false);
            })}
          </Menu>
        )}
      </Fragment>
    );
  }
}
