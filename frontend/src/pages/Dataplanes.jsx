import React, { useState } from 'react';
import { List, Button } from 'antd';

export default function Dataplanes() {
  const [planes, setPlanes] = useState([]);

  const createPlane = () => {
    setPlanes([...planes, `dataplane-${planes.length + 1}`]);
  };

  return (
    <div>
      <Button type="primary" onClick={createPlane}>Create Dataplane</Button>
      <List dataSource={planes} renderItem={item => <List.Item>{item}</List.Item>} />
    </div>
  );
}
