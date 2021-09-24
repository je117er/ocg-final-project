<template>
  <section class="section is-medium">
    <div class="columns">
      <h5 class="title is-5">1. Personal information</h5>
    </div>
    <b-field grouped class="has-text-left">
      <b-field label="Name" type="name" expanded>
        <b-input></b-input>
      </b-field>
      <b-field label="Date of Birth" expanded>
        <b-datepicker
          placeholder="Type or select a date"
          icon="calendar-today"
          editable
        >
        </b-datepicker>
      </b-field>
      <b-field label="Gender">
        <b-select>
          <option>Male</option>
          <option>Female</option>
        </b-select>

      </b-field>
      <b-field label="Phone Number" expanded>
        <b-input></b-input>
      </b-field>
    </b-field>
    <b-field class="has-text-left" grouped>
      <b-field label="Email" type="email">
        <b-input></b-input>
      </b-field>
      <b-field label="Insurance Number">
        <b-input></b-input>
      </b-field>
      <b-field label="Address" expanded>
        <b-input></b-input>
      </b-field>
      <b-field label="Ethnicity" expanded>
        <b-input></b-input>
      </b-field>
    </b-field>
    <b-field class="has-text-left" grouped>
      <b-field label="Commune">
        <b-input></b-input>
      </b-field>
      <b-field label="District">
        <b-input></b-input>
      </b-field>
      <b-field label="City">
        <b-input></b-input>
      </b-field>
      <b-field label="Nationality" expanded>
        <b-input></b-input>
      </b-field>
    </b-field>
    <div class="buttons">
      <b-button tag="router-link"
                to="/"
                outlined
                type="is-link is-danger">
        Cancel
      </b-button>
      <b-button tag="router-link"
                to="/expo"
                type="is-info">
        Continue
      </b-button>
    </div>
  </section>
</template>

<script>
export default {
  name: 'PersonalInfo',
};
</script>

<style scoped>

</style>
