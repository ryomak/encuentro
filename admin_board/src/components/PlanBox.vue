<template>
  <div>
    <b-card>
      <p>日付: {{ plan.date }}</p>
      <p>場所: {{ plan.place }}</p>
      <p>飲酒: {{ plan.drink }}</p>
      <p>コース: {{ plan.course }}</p>
      <p>ステータス: {{ plan.status }}</p>
    </b-card>
  </div>
</template>

<script>
export default {
  props: [
    'plan'
  ]
}
</script>
