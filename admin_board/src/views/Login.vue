<template>
  <div>
    <b-form>
      <b-form-group label="Email address:"
                    label-for="emailForm">
        <b-form-input id="emailForm"
                      type="email"
                      required
                      placeholder="Enter email"
                      :value="email"
                      @input="setEmail" />
      </b-form-group>
      <b-form-group label="Password:"
                    label-for="passwordForm">
        <b-form-input id="passwordForm"
                      type="password"
                      required
                      placeholder="Enter password"
                      :value="password"
                      @input="setPassword" />
      </b-form-group>
      <router-link to="/users">
        <b-button type="submit" variant="primary" @click="login">Submit</b-button>
      </router-link>
    </b-form>
  </div>
</template>
<script>
import { mapActions, mapState } from 'vuex'

export default {
  computed: {
    ...mapState({
      email: state => state.auth.email,
      password: state => state.auth.password
    })
  },
  methods: {
    ...mapActions([
      'setEmail',
      'setPassword',
      'login'
    ])
  }


}


</script>
