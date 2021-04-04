import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-layout';

export default () => (
  <DefaultFooter
    copyright="Ziyue Pan"
    links={[
      {
        title: 'Contact Ziyue',
        href: 'mailto:ziyuepan99@outlook.com',
        blankTarget: true,
      },
      {
        title: <GithubOutlined />,
        href: 'https://github.com/PAN-Ziyue',
        blankTarget: true,
      },
      {
        title: 'Source Code',
        href: 'https://github.com/PAN-Ziyue/IOTino',
        blankTarget: true,
      },
    ]}
  />
);
