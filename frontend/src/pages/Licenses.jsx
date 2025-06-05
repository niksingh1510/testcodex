import React from 'react';
import { Form, Input, Button } from 'antd';

export default function Licenses() {
  const onFinish = values => {
    console.log('activate license', values);
  };

  return (
    <Form layout="inline" onFinish={onFinish}>
      <Form.Item name="key" rules={[{ required: true, message: 'Enter license key' }]}>
        <Input placeholder="License key" />
      </Form.Item>
      <Form.Item>
        <Button type="primary" htmlType="submit">Activate</Button>
      </Form.Item>
    </Form>
  );
}
