import React, { useState } from 'react';
import { List, Button } from 'antd';

export default function Instances() {
  const [instances, setInstances] = useState([]);

  const createInstance = () => {
    setInstances([...instances, `instance-${instances.length + 1}`]);
  };

  return (
    <div>
      <Button type="primary" onClick={createInstance}>Create Instance</Button>
      <List dataSource={instances} renderItem={item => <List.Item>{item}</List.Item>} />
    </div>
  );
}
