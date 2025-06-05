import React, { useState } from 'react';
import { Layout, Menu } from 'antd';
import Licenses from './pages/Licenses';
import Instances from './pages/Instances';
import Dataplanes from './pages/Dataplanes';

const { Header, Content, Sider } = Layout;

export default function App() {
  const [page, setPage] = useState('licenses');

  const renderContent = () => {
    switch (page) {
      case 'instances':
        return <Instances />;
      case 'dataplanes':
        return <Dataplanes />;
      case 'licenses':
      default:
        return <Licenses />;
    }
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider>
        <Menu theme="dark" mode="inline" defaultSelectedKeys={['licenses']} onClick={({ key }) => setPage(key)}>
          <Menu.Item key="licenses">Licenses</Menu.Item>
          <Menu.Item key="instances">Instances</Menu.Item>
          <Menu.Item key="dataplanes">Dataplanes</Menu.Item>
        </Menu>
      </Sider>
      <Layout>
        <Header style={{ color: '#fff' }}>MDC Console</Header>
        <Content style={{ margin: '16px' }}>
          {renderContent()}
        </Content>
      </Layout>
    </Layout>
  );
}
