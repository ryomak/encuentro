<template>
  <div>
    <b-form>
      <b-form-group label="Email address:"
                    label-for="emailForm">
        <b-form-input id="emailForm"
                      type="email"
                      required
                      placeholder="Enter email" />
      </b-form-group>
      <b-form-group label="Password:"
                    label-for="passwordForm">
        <b-form-input id="passwordForm"
                      type="password"
                      required
                      placeholder="Enter password" />
      </b-form-group>
      <b-button type="submit" variant="primary">Submit</b-button>
    </b-form>
  </div>
</template>
