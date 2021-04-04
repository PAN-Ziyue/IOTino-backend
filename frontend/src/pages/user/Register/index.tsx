import { Form, Button, Input, Popover, Progress, message } from 'antd';
import type { FC } from 'react';
import { useState, useEffect } from 'react';
import type { Dispatch } from 'umi';
import { Link, connect, history, FormattedMessage, formatMessage, SelectLang } from 'umi';
import Footer from '@/components/Footer';
import type { StateType } from './model';
import styles from './style.less';

const FormItem = Form.Item;

export type UserRegisterParams = {
  email: string;
  password: string;
  confirm: string;
};

const passwordStatusMap = {
  ok: (
    <div className={styles.success}>
      <FormattedMessage id="userandregister.strength.strong" />
    </div>
  ),
  pass: (
    <div className={styles.warning}>
      <FormattedMessage id="userandregister.strength.medium" />
    </div>
  ),
  poor: (
    <div className={styles.error}>
      <FormattedMessage id="userandregister.strength.short" />
    </div>
  ),
};

const passwordProgressMap: {
  ok: 'success';
  pass: 'normal';
  poor: 'exception';
} = {
  ok: 'success',
  pass: 'normal',
  poor: 'exception',
};

type RegisterProps = {
  dispatch: Dispatch;
  userAndRegister: StateType;
  submitting: boolean;
};

export const Register: FC<RegisterProps> = ({
  submitting,
  dispatch,
  userAndRegister,
}) => {
  const [visible, setvisible]: [boolean, any] = useState(false);
  const [popover, setpopover]: [boolean, any] = useState(false);
  const confirmDirty = false;
  let interval: number | undefined;
  const [form] = Form.useForm();
  useEffect(() => {
    if (!userAndRegister) {
      return;
    }
    const account = form.getFieldValue('email');
    if (userAndRegister.status === 'ok') {
      message.success('注册成功！');
      history.push({
        pathname: '/user/register-result',
        state: {
          account,
        },
      });
    }
  }, [userAndRegister]);
  useEffect(
    () => () => {
      clearInterval(interval);
    },
    [],
  );
  const getPasswordStatus = () => {
    const value = form.getFieldValue('password');
    if (value && value.length > 9) {
      return 'ok';
    }
    if (value && value.length > 5) {
      return 'pass';
    }
    return 'poor';
  };
  const onFinish = (values: Record<string, any>) => {
    dispatch({
      type: 'userAndRegister/submit',
      payload: {
        ...values,
      },
    });
  };
  const checkConfirm = (_: any, value: string) => {
    const promise = Promise;
    if (value && value !== form.getFieldValue('password')) {
      return promise.reject(formatMessage({ id: 'userandregister.password.twice' }));
    }
    return promise.resolve();
  };
  const checkPassword = (_: any, value: string) => {
    const promise = Promise;
    // 没有值的情况
    if (!value) {
      setvisible(!!value);
      return promise.reject(formatMessage({ id: 'userandregister.password.required' }));
    }
    // 有值的情况
    if (!visible) {
      setvisible(!!value);
    }
    setpopover(!popover);
    if (value.length < 6) {
      return promise.reject('');
    }
    if (value && confirmDirty) {
      form.validateFields(['confirm']);
    }
    return promise.resolve();
  };
  const renderPasswordProgress = () => {
    const value = form.getFieldValue('password');
    const passwordStatus = getPasswordStatus();
    return value && value.length ? (
      <div className={styles[`progress-${passwordStatus}`]}>
        <Progress
          status={passwordProgressMap[passwordStatus]}
          className={styles.progress}
          strokeWidth={6}
          percent={value.length * 10 > 100 ? 100 : value.length * 10}
          showInfo={false}
        />
      </div>
    ) : null;
  };

  return (
    <div className={styles.container}>
      <div className={styles.lang}>{SelectLang && <SelectLang />}</div>
      <div className={styles.content}>
        <div className={styles.top}>
          <div className={styles.header}>
            <Link to="/">
              <img alt="logo" className={styles.logo} src="/logo.svg" />
              <span className={styles.title}>IOTino</span>
            </Link>
          </div>
          <div className={styles.desc}>老和山最好的物联网托管平台</div>
        </div>

        <div className={styles.main}>
          <Form form={form} name="UserRegister" onFinish={onFinish}>
            <FormItem
              name="email"
              rules={[
                {
                  required: true,
                  message: formatMessage({ id: 'userandregister.email.required' }),
                },
                {
                  type: 'email',
                  message: formatMessage({ id: 'userandregister.email.wrong-format' }),
                },
              ]}
            >
              <Input size="large" placeholder={formatMessage({ id: 'userandregister.email.placeholder' })} />
            </FormItem>
            <Popover
              getPopupContainer={(node) => {
                if (node && node.parentNode) {
                  return node.parentNode as HTMLElement;
                }
                return node;
              }}
              content={
                visible && (
                  <div style={{ padding: '4px 0' }}>
                    {passwordStatusMap[getPasswordStatus()]}
                    {renderPasswordProgress()}
                    <div style={{ marginTop: 10 }}>
                      <FormattedMessage id="userandregister.strength.msg" />
                    </div>
                  </div>
                )
              }
              overlayStyle={{ width: 240 }}
              placement="right"
              visible={visible}
            >
              <FormItem
                name="password"
                className={
                  form.getFieldValue('password') &&
                  form.getFieldValue('password').length > 0 &&
                  styles.password
                }
                rules={[
                  {
                    validator: checkPassword,
                  },
                ]}
              >
                <Input
                  size="large"
                  type="password"
                  placeholder={formatMessage({ id: 'userandregister.password.placeholder' })}
                />
              </FormItem>
            </Popover>
            <FormItem
              name="confirm"
              rules={[
                {
                  required: true,
                  message: formatMessage({ id: 'userandregister.confirm-password.required' }),
                },
                {
                  validator: checkConfirm,
                },
              ]}
            >
              <Input
                size="large"
                type="password"
                placeholder={formatMessage({ id: 'userandregister.confirm-password.placeholder' })}
              />
            </FormItem>
            <FormItem>
              <Button
                size="large"
                loading={submitting}
                className={styles.submit}
                type="primary"
                htmlType="submit"
              >
                <FormattedMessage id="userandregister.register.register" />
              </Button>
              <Link className={styles.login} to="/user/login">
                <FormattedMessage id="userandregister.register.sign-in" />
              </Link>
            </FormItem>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  );
};

export default connect(
  ({
    userAndRegister,
    loading,
  }: {
    userAndRegister: StateType;
    loading: {
      effects: Record<string, boolean>;
    };
  }) => ({
    userAndRegister,
    submitting: loading.effects['userAndRegister/submit'],
  }),
)(Register);
