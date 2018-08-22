User.seed(:id,
  {
    name: '新垣結衣',
    sex:'female',
    email: 'yuiyui1@sample.com',
    password: 'yuiyui',
    password_confirmation: 'yuiyui',
    university: 'けんたチュキチュキクラブ',
    birthday: '1988/06/11',
  },
  {
    name: '吉岡里帆',
    sex: 'female',
    email: 'rihoriho@sample.com',
    password: 'rihoriho',
    password_confirmation: 'rihoriho',
    university: 'けんたチュキチュキクラブ',
    birthday: '1993/01/15',
  }
)

Plan.seed(:id,
  {
    user: User.find_by({ name: '新垣結衣' }),
    date: '2018/07/16',
    place: '大阪',
    drink: 'あり',
    course: '高学歴コース',
    status: 0
  },
  {
    user: User.find_by({ name: '新垣結衣' }),
    date: '2018/07/18',
    place: '兵庫',
    drink: 'あり',
    course: '低学歴コース',
    status: 0
  },
  {
    user: User.find_by({ name: '吉岡里帆' }),
    date: '2018/07/18',
    place: '大阪',
    drink: 'なし',
    course: '高学歴コース',
    status: 0
  },
  {
    user: User.find_by({ name: '吉岡里帆' }),
    date: '2018/08/16',
    place: '京都',
    drink: 'あり',
    course: '低学歴コース',
    status: 0
  }
)
