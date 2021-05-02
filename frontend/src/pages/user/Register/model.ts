import type { Effect, Reducer } from 'umi';
import { history } from 'umi';
import { message } from 'antd';
import { Register } from './service';

export type StateType = {
  status?: 'ok' | 'error';
  currentAuthority?: 'user' | 'guest' | 'admin';
};

export type ModelType = {
  namespace: string;
  state: StateType;
  effects: {
    submit: Effect;
  };
  reducers: {
    registerHandle: Reducer<StateType>;
  };
};

const Model: ModelType = {
  namespace: 'userAndRegister',

  state: {
    status: undefined,
  },

  effects: {
    *submit({ payload }, { call, put }) {
      const response = yield call(Register, payload);
      yield put({
        type: 'registerHandle',
        payload: response,
      });

      if (response === 'ok') {
        message.success('注册成功！');
        history.push({
          pathname: '/user/register-result',
          state: {
            // account
          },
        });
      } else {
        console.log('fuck')
      }
    },
  },

  reducers: {
    registerHandle(state, { payload }) {
      return {
        ...state,
        status: payload.status,
      };
    },
  },
};

export default Model;
