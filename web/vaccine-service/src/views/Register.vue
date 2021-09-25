<template>
  <div class="container is-fullhd">
    <h1 class="title">Register for vaccination</h1>
    <section class="is-small">
        <form-wizard
          @on-complete="onComplete"
          title=""
          subtitle=""
          shape="tab"
          @on-validate="handleValidation"
          @on-error="handleErrorMessage"
          color="#556ee6">
          <template slot="step" slot-scope="props">
            <wizard-step :tab="props.tab"
                         :transition="props.transition"
                         :key="props.tab.title"
                         :index="props.index">
            </wizard-step>
          </template>
          <tab-content
            title="Email"
            icon="ti-user"
            :before-change="checkEmailExists"
          >
            <div class="tile is-ancestor">
              <div class="tile is-parent">
                <article class="tile is-child">
                </article>
              </div>
              <div class="tile is-parent">
                <article class="tile is-child">
                  <b-field label="Email" class="has-text-left" position="is-centered">
                    <b-input
                      type="email"
                      maxlength="50"
                      v-model="email"
                    ></b-input>
                  </b-field>
                </article>
              </div>
              <div class="tile is-parent">
                <article class="tile is-child">
                </article>
              </div>
            </div>
          </tab-content>
          <tab-content
            title="Personal details"
            :before-change="checkApproved"
            icon="ti-user">
            <b-table
              :data="conditions"
            >
              <b-table-column field="id" label="ID" centered v-slot="props">
                {{ props.row.id}}
              </b-table-column>
              <b-table-column
                field="description"
                label="Description"
                v-slot="props"
              >
                <div class="has-text-left">
                  {{ props.row.description}}
                </div>
              </b-table-column>
              <b-table-column field="selected" label="Yes" centered v-slot="props">
                <div class="has-text-right">
<!--                  <b-checkbox v-model="props.row.selected"></b-checkbox>-->
                  <input type="radio" id="one" value="1" v-model="props.row.selected">
                </div>
              </b-table-column>
              <b-table-column field="selected" label="No" centered v-slot="props">
                <div class="has-text-right">
                  <input type="radio" id="one" value="0" v-model="props.row.selected">
                </div>
              </b-table-column>
            </b-table>
          </tab-content>
          <tab-content
            title="Order Info"
            icon="ti-user"
            >
            <div class="tile is-ancestor">
              <div class="tile is-parent">
                <article class="tile is-child">
                </article>
              </div>
              <div class="tile is-parent">
                <article class="tile is-child">
                  <b-field label="Vaccine" class="has-text-left" position="is-centered">
                    <b-select
                      placeholder="Select a vaccine"
                      v-model="selectedProduct"
                      >
                      <option
                        v-for="product in productList"
                        :value="product.Name"
                        :key="product.ID"
                        >
                       {{ product.Name }}
                      </option>
                    </b-select>
                  </b-field>
                  <div v-if="clinicList">
                    <b-field
                      label="Clinic" class="has-text-left" position="is-centered">
                      <b-select
                        placeholder="Select a clinic"
                        v-model="selectedClinic"
                      >
                        <option
                          v-for="clinic in clinicList"
                          :value="clinic.clinic_name"
                          :key="clinic.clinic_name"
                        >
                          {{ clinic.clinic_name }}
                        </option>
                      </b-select>
                    </b-field>
                  </div>
                  <div v-else>
                    <p class="is-danger">
                      Currently {{selectedProduct}} vaccines are running out of stock at all clinics
                    </p>
                  </div>
                </article>
              </div>
              <div class="tile is-parent">
                <article class="tile is-child">
                  {{ selectedProduct}}
                </article>
              </div>
            </div>
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
<!--              </div>-->
            </section>
          </tab-content>
          <div v-if="errorMsg">
           <span class="error">{{ errorMsg }}</span>
          </div>
          <template slot="footer" slot-scope="props">
            <div class="wizard-footer-left">
            </div>
            <div class="wizard-footer-right">
              <wizard-button
                @click.native="props.nextTab()"
                class="wizard-footer-right finish-button"
                :style="props.fillButtonStyle">{{props.isLastStep ? 'Done' : 'Next'}}
              </wizard-button>
            </div>
          </template>
        </form-wizard>
        <form>
        </form>
      </section>

  </div>
</template>

<script>
import Vue from 'vue';
import {
  FormWizard, TabContent, WizardStep, WizardButton,
} from 'vue-form-wizard';
import 'vue-form-wizard/dist/vue-form-wizard.min.css';

export default {
  name: 'Register',
  component: {
    FormWizard,
    TabContent,
    WizardStep,
    WizardButton,
  },
  data() {
    return {
      constraints: [],
      conditions: [],
      errorMsg: '',
      email: '',
      productList: [],
      clinicList: [],
      selectedProduct: '',
      selectedClinic: '',
    };
  },
  watch: {
    selectedProduct() {
      Vue.axios.get(`http://localhost:8088/clinic?vaccineName=${this.selectedProduct}`).then((response) => {
        this.clinicList = response.data;
        console.log(response.data);
      })
        .catch((e) => {
          console.log(e);
        });
    },
  },
  methods: {
    onComplete() {
      alert('yay!');
    },
    handleErrorMessage(errorMsg) {
      this.errorMsg = String(errorMsg).substring(7);
    },
    handleValidation(isValid, tabIndex) {
      console.log(`Tab: ${tabIndex} valid: ${isValid}`);
    },
    checkEmailExists() {
      return new Promise((resolve, reject) => {
        if (this.email) {
          Vue.axios.get(`http://localhost:8088/customer?email=${this.email}`, this.form)
            .then((response) => {
              console.log(response.data);
              this.fetchProductList();
              resolve(true);
            })
            .catch((error) => {
              console.log(error);
              reject(new Error('Email already exists!'));
            });
        } else {
          reject(new Error('No email was supplied!'));
        }
      });
    },
    fetchProductList() {
      Vue.axios.get('http://localhost:8088/constraints').then((emailResponse) => {
        this.constraints = emailResponse.data;
        this.conditions = this.constraints.map((e) => ({
          id: e.ID,
          description: e.Description,
          selected: 1,
        }));
      })
        .catch((e) => {
          console.log(e);
        });
    },
    checkApproved() {
      return new Promise((resolve, reject) => {
        const unapproved = this.conditions
          .filter((e, i) => (Number(this.constraints[i].VaccineEligible) !== Number(e.selected)
            && this.constraints[i].VaccineEligible !== 2));
        if (unapproved.length >= 1) {
          reject(new Error("Based on the survey's results, you are not eligible for vaccination at the moment. "
              + 'Please consult your healthcare provider for more information.'));
        } else {
          Vue.axios.get('http://localhost:8088/productlist').then((response) => {
            this.productList = response.data;
            console.log(response.data);
          })
            .catch((e) => {
              console.log(e);
            });
          resolve(true);
        }
      });
    },
  },
  created() {
  },
};
</script>

<style scoped>
span.error{
  color:#e74c3c;
  font-size:20px;
  display:flex;
  justify-content:center;
}
</style>
