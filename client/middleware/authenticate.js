import store2 from 'store2';
let local_storage = store2.namespace('auth');

export default function ({ redirect }) {
    // ユーザーが認証されていないとき
    if (!local_storage.AccessToken) {
       return redirect('/login')
    }
  }