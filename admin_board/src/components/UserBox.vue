<template>
  <div>
    <router-link :to="{ name: 'user_plans', params: { id: user.id }}">
      <b-card>
        <p>名前: {{ user.name }}</p>
        <p>性別: {{ user.sex }}</p>
        <p>メールアドレス: {{ user.email }}</p>
        <p>誕生日: {{ user.birthday }}</p>
        <p>大学: {{ user.university }}</p>
      </b-card>
    </router-link>
  </div>
</template>

<script>
export default {
  props: [
    'user'
  ]
}
</script>
