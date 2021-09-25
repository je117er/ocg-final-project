<template>
  <b-tabs>
    <b-tab-item label="Personal Information" icon="">
      <template>
        <section class="section is-medium">
          <h2 v-if="isInfoUpdated" style="color: darkgreen; font-weight: bold;">Update success!</h2>
          <b-field grouped class="has-text-left">
            <b-field label="Name" type="name" expanded>
              <b-input v-model="customer.name"></b-input>
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
              <b-input v-model="customer.phoneNumber"></b-input>
            </b-field>
          </b-field>
          <b-field class="has-text-left" grouped>
            <b-field label="Email" type="email">
              <b-input v-model="customer.email"></b-input>
            </b-field>
            <b-field label="Insurance Number">
              <b-input v-model="customer.insuranceNumber"></b-input>
            </b-field>
            <b-field label="Address" expanded>
              <b-input v-model="customer.address"></b-input>
            </b-field>
            <b-field label="Ethnicity" expanded>
              <b-input v-model="customer.ethnicity"></b-input>
            </b-field>
          </b-field>
          <b-field class="has-text-left" grouped>
            <b-field label="Commune">
              <b-input v-model="customer.commune"></b-input>
            </b-field>
            <b-field label="District">
              <b-input v-model="customer.district"></b-input>
            </b-field>
            <b-field label="City">
              <b-input v-model="customer.city"></b-input>
            </b-field>
            <b-field label="Nationality" expanded>
              <b-input v-model="customer.nationality"></b-input>
            </b-field>
          </b-field>
          <div class="buttons is-centered">
            <b-button @click="updateCustomer">
              Save
            </b-button>
            <b-button tag="router-link"
                      to="/"
                      outlined
                      type="is-link is-danger">
              Cancel
            </b-button>
          </div>
        </section>
      </template>
    </b-tab-item>
    <b-tab-item label="Medical History" icon="">
      <section>
        <b-table
          :paginated=true
          :per-page=10
          :data="medical">
          <b-table-column field="ID" label="ID" width="40" numeric centered v-slot="props">
            {{ props.row.ID }}
          </b-table-column>
          <b-table-column field="Code" label="Code" centered v-slot="props">
            {{ props.row.Code }}
          </b-table-column>
          <b-table-column field="Description" label="Description" centered v-slot="props">
            {{ props.row.Description }}
          </b-table-column>
          <b-table-column field="Status" label="Status" centered v-slot="props">
            <b-checkbox :value="props.row.ConditionStatus" disabled
                        type="is-info">
            </b-checkbox>
          </b-table-column>
        </b-table>
      </section>
    </b-tab-item>
    <b-tab-item label="Vaccination" icon="">
      <section>
        <h2 v-if="isVaccinationUpdated" style="color: darkgreen; font-weight: bold;">Update
          success!</h2>
        <b-table
          :paginated=true
          :per-page=10
          :data="vaccination">
          <b-table-column field="dose" label="Does" width="40" numeric centered v-slot="props">
            {{ props.row.dose }}
          </b-table-column>
          <b-table-column field="vaccineName" label="vaccineName" centered v-slot="props">
            {{ props.row.vaccineName }}
          </b-table-column>
          <b-table-column field="clinic" label="clinic" centered v-slot="props">
            {{ props.row.clinic }}
          </b-table-column>
          <b-table-column field="completed" label="completed" centered v-slot="props">
            <b-checkbox :value="props.row.completed" v-model="props.row.completed"
                        type="is-info">
            </b-checkbox>
          </b-table-column>
          <b-table-column field="registrationTime" label="registrationTime" centered
                          v-slot="props">
            {{ props.row.registrationTime }}
          </b-table-column>
          <b-table-column field="vaccinationTime" label="vaccinationTime" centered
                          v-slot="props">
            {{ props.row.vaccinationTime }}
          </b-table-column>
        </b-table>
        <div class="buttons is-centered" @click="UpdateVaccination">
          <b-button>
            Save
          </b-button>
          <b-button tag="router-link"
                    to="/"
                    outlined
                    type="is-link is-danger">
            Cancel
          </b-button>
        </div>
      </section>
    </b-tab-item>
    <b-tab-item label="Order" icon="">
      <div class="container">
        <div class="notification">
        </div>
      </div>
    </b-tab-item>
  </b-tabs>
</template>

<script>
import Vue from 'vue';

export default {
  data() {
    return {
      customer: [],
      medical: [],
      vaccination: [],
      loaded: false,
      isVaccinationUpdated: false,
      isInfoUpdated: false,
    };
  },
  methods: {
    updateCustomer() {
      this.loaded = false;
      Vue.axios.put('http://localhost:8088/admin/customer', this.customer).then((response) => {
        this.customer = response.data;
        this.loaded = true;
        this.isInfoUpdated = true;
      }).catch((err) => {
        console.log(err);
      });
    },
    UpdateVaccination() {
      this.isVaccinationUpdated = false;
      this.loaded = false;
      const count = this.vaccination.filter((item) => item.completed).length;
      Vue.axios.put('http://localhost:8088/admin/booking', {
        bookingID: this.vaccination[0].bookingID,
        completed: count,
      }).then((response) => {
        this.vaccination = response.data;
        this.loaded = true;
        this.isVaccinationUpdated = true;
      }).catch((err) => {
        this.isVaccinationUpdated = false;
        console.log(err);
      });
    },
  },
  created() {
    Vue.axios.get(`http://localhost:8088/admin/customer/${this.$route.params.id}`).then((response) => {
      this.customer = response.data;
      this.loaded = true;
    }).catch((err) => {
      console.log(err);
    });

    Vue.axios.get(`http://localhost:8088/admin/medicalHistory?customerID=${this.$route.params.id}`).then((response) => {
      this.medical = response.data;
      this.loaded = true;
    }).catch((err) => {
      console.log(err);
    });

    Vue.axios.get(`http://localhost:8088/admin/booking?customerID=${this.$route.params.id}`).then((response) => {
      this.vaccination = response.data;
      this.loaded = true;
    }).catch((err) => {
      console.log(err);
    });
  },
};
</script>
