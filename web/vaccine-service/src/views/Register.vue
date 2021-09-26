<template>
  <div class="container is-fullhd">
    <h1 class="title">Register for vaccination</h1>
    <section class="is-medium">
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
          <b-field label="Email" class="has-text-left" position="is-centered">
            <b-input
              type="email"
              maxlength="50"
              v-model="customerInfo.email"
            ></b-input>
          </b-field>
        </tab-content>
        <tab-content
          title="Personal details"
          :before-change="checkApproved"
          icon="ti-user">
          <b-table
            :data="conditions"
          >
            <b-table-column field="id" label="ID" centered v-slot="props">
              {{ props.row.code}}
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
                <input type="radio" id="one" value="1" v-model="props.row.condition_status">
              </div>
            </b-table-column>
            <b-table-column field="selected" label="No" centered v-slot="props">
              <div class="has-text-right">
                <input type="radio" id="one" value="0" v-model="props.row.condition_status">
              </div>
            </b-table-column>
          </b-table>
        </tab-content>
        <tab-content
          title="Order Info"
          icon="ti-user"
        >
          <div class="tile is-ancestor">
            <div class="tile is-vertical is-parent">
              <div class="tile is-child box">
                <p class="title">Vaccination</p>
                <div class="has-text-left" style="color: red">
                  <p>
                    <u>Note:</u>
                  </p>
                  <ul>
                    <li>
                      - Those who are under 18 years old are allowed to choose Pfizer only
                      as this is the only vaccine that's available to this age group at the moment.
                    </li>
                  </ul>
                </div>
                <br/>
                <b-field
                  label="Date of Birth"
                  class="has-text-left"
                  style="width: 143px"
                >
                  <b-datepicker
                    placeholder="Select a date"
                    :selectable-dates="selectableBirthDates"
                    v-model="customerInfo.selectedBirthDate"
                  >
                  </b-datepicker>
                </b-field>
                  <b-field label="Select a vaccine" class="has-text-left" position="is-centered">
                  <b-select
                    placeholder="Select a vaccine"
                    v-model="selectedProduct"
                  >
                    <option
                      v-for="product in fetchAllOrOnlyPfizer"
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
                  <div v-if="selectedClinic">
                    <b-field grouped group-multiline>
                      <b-field label="Select a date" class="has-text-left">
                        <b-datepicker
                          v-model="selectedDate"
                          :selectable-dates="selectableDates"
                        >
                        </b-datepicker>
                      </b-field>
                      <b-field label="Session" class="has-text-left" position="is-centered">
                        <b-select
                          placeholder="Select a session"
                          v-model="selectedSession"
                        >
                          <option
                            v-for="session in availableSessions"
                            :value="session.time_period"
                            :key="session.ID"
                          >
                            {{ sessionType[session.time_period]}}
                          </option>
                        </b-select>
                      </b-field>
                    </b-field>
                  </div>
                </div>
                <div v-else>
                  <p class="is-danger">
                    Currently {{selectedProduct}} vaccines are running out of stock at all clinics
                  </p>
                </div>
              </div>
              <div class="tile is-child box">
                <p class="title">Personal Information</p>
                <b-field grouped class="has-text-left">
                  <b-field label="Name" type="name">
                    <b-input v-model="customerInfo.customer_name"></b-input>
                  </b-field>
                  <b-field label="Gender">
                    <b-select v-model="customerInfo.gender">
                      <option>Male</option>
                      <option>Female</option>
                    </b-select>

                  </b-field>
                  <b-field label="Insurance Number">
                    <b-input v-model="customerInfo.insurance_number"></b-input>
                  </b-field>
                  <b-field label="Ethnicity" expanded>
                    <b-input v-model="customerInfo.ethnicity"></b-input>
                  </b-field>
                </b-field>
                <b-field class="has-text-left" grouped group-multiline>
                  <b-field label="Phone Number" >
                    <b-input v-model="customerInfo.phone_number"></b-input>
                  </b-field>
                  <b-field label="Email" type="email">
                    <b-input v-model="customerInfo.email"></b-input>
                  </b-field>
                  <b-field label="Address" expanded>
                    <b-input v-model="customerInfo.address"></b-input>
                  </b-field>
                </b-field>
                <b-field class="has-text-left" grouped group-multiline>
                  <b-field label="Commune">
                    <b-input v-model="customerInfo.commune"></b-input>
                  </b-field>
                  <b-field label="District">
                    <b-input v-model="customerInfo.district"></b-input>
                  </b-field>
                  <b-field label="City">
                    <b-input v-model="customerInfo.city"></b-input>
                  </b-field>
                  <b-field label="Nationality">
                    <b-input v-model="customerInfo.nationality"></b-input>
                  </b-field>
                </b-field>
              </div>
              <div v-if="selectedProduct" class="tile is-child box">
               <p class="title">Order Summary</p>
                <ul class="has-text-left is-size-5">
                  <li>
                    Customer Name: {{customerInfo.customer_name}}
                  </li>
                  <li>
                    <ul>
                      <li>
                        Vaccine Name: {{selectedProduct}}
                      </li>
                      <li>
                        Total Price: ${{totalPrice}}
                      </li>
                    </ul>
                  </li>
                  <li>
                    Clinic: {{selectedClinic}}
                  </li>
                  <li>
                    Appointment Date: {{dateStr}}. Session: {{sessionStr}}
                  </li>
                </ul>
              </div>
              </div>
            </div>
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
              :style="props.fillButtonStyle">{{props.isLastStep ? 'Check out' : 'Next'}}
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
    const sessionType = {
      1: 'Morning',
      2: 'Afternoon',
      3: 'Evening',
    };
    return {
      constraints: [],
      conditions: [],
      errorMsg: '',
      email: '',
      productList: [],
      clinicList: [],
      selectedProduct: '',
      selectedClinic: '',
      selectedDate: new Date(),
      selectableDates: [],
      availableSessions: [],
      dateList: [],
      selectedSession: 0,
      minimumAge: 0,
      sessionType,
      customerInfo: {
        customer_name: '',
        gender: '',
        insurance_number: '',
        ethnicity: '',
        phone_number: '',
        email: '',
        address: '',
        commune: '',
        district: '',
        city: '',
        nationality: '',
        selectedBirthDate: new Date(new Date().setFullYear(new Date().getFullYear() - 18)),
      },
    };
  },
  computed: {
    fetchAllOrOnlyPfizer() {
      // eslint-disable-next-line max-len
      const notAllowBirthDatePfizer = new Date(new Date().setFullYear(new Date().getFullYear() - 18));
      if (this.customerInfo.selectedBirthDate.getTime() > notAllowBirthDatePfizer.getTime()) {
        return this.productList.filter((e) => e.Name.includes('Pfizer'));
      }
      return this.productList;
    },
    totalPrice() {
      // eslint-disable-next-line max-len
      const productInfo = this.productList.filter((e) => e.Name === this.selectedProduct)[0];
      const totalBill = productInfo.Price * productInfo.ImmunizationSchedule;
      console.log(totalBill);
      // return `$${totalBill}`;
      return totalBill;
    },
    sessionStr() {
      return this.sessionType[this.selectedSession];
    },
    dateStr() {
      return this.selectedDate.toDateString();
    },
  },
  watch: {
    customerInfo: {
      handler(val) {
        console.log(JSON.stringify(val));
      },
      deep: true,
    },
    conditions: {
      handler(val) {
        console.log(JSON.stringify(val));
      },
      deep: true,
    },
    selectedProduct() {
      Vue.axios.get(`http://localhost:8088/clinic?vaccineName=${this.selectedProduct}`).then((response) => {
        this.clinicList = response.data;
        console.log(response.data);
      })
        .catch((e) => {
          console.log(e);
        });
    },
    selectedClinic() {
      Vue.axios.get(`http://localhost:8088/session?clinicName=${this.selectedClinic}`).then((response) => {
        this.dateList = response.data;
        this.selectableDates = this.dateList.map((e) => new Date(e.current_date));
        [this.selectedDate] = this.selectableDates;
        // eslint-disable-next-line max-len
        this.getDateList();
      })
        .catch((e) => {
          console.log(e);
        });
    },
    selectedDate() {
      this.availableSessions = [];
      this.getDateList();
    },
  },
  methods: {
    selectableBirthDates(day) {
      const twelveYearsBeforeNow = new Date();
      twelveYearsBeforeNow.setFullYear(twelveYearsBeforeNow.getFullYear() - this.minimumAge);
      return day.getTime() < twelveYearsBeforeNow.getTime();
    },
    getDateList() {
      // eslint-disable-next-line max-len
      // this.dateList.forEach((e) => console.log(new Date(e.current_date)));
      // eslint-disable-next-line max-len
      this.availableSessions = this.dateList.filter((e) => this.selectedDate.getTime() === new Date(e.current_date).getTime());
      // this.availableSessions.forEach((e) => console.log(new Date(e.current_date)));
    },
    onComplete() {
      const clinic = this.clinicList.filter((e) => e.clinic_name === this.selectedClinic)[0];
      // eslint-disable-next-line max-len
      const session = this.availableSessions.filter((e) => e.time_period === this.selectedSession)[0];
      const orderRequest = {};
      orderRequest.condition_order_request = this.conditions;
      orderRequest.customer_order_request = this.customerInfo;
      orderRequest.time_period = this.selectedSession;
      orderRequest.doses_completed = 0;
      orderRequest.vaccine_name = this.selectedProduct;
      orderRequest.authorized_interval = clinic.authorized_interval;
      orderRequest.lot_number = clinic.lot_number;
      orderRequest.clinic_name = this.selectedClinic;
      orderRequest.price = this.selectedClinic.price;
      orderRequest.sent_reminder_email = 0;
      orderRequest.stock_item_id = clinic.stock_item_id;
      orderRequest.total_bill = this.totalPrice;
      orderRequest.session_capacity_id = session.session_id;
      orderRequest.date_registered = new Date();
      orderRequest.date_booked = this.selectedDate;

      Vue.axios.put('http://localhost:8088/order', orderRequest).then((response) => {
        console.log(response.data);
      }).catch((e) => {
        console.log(e);
      });

      // orderRequest.date_registered = new Date();
      console.log(JSON.stringify(orderRequest));
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
        if (this.customerInfo.email) {
          Vue.axios.get(`http://localhost:8088/customer?email=${this.customerInfo.email}`, this.form)
            .then((response) => {
              console.log(response.data);
              this.fetchConstraintList();
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
    fetchConstraintList() {
      Vue.axios.get('http://localhost:8088/constraints').then((emailResponse) => {
        this.constraints = emailResponse.data;
        this.conditions = this.constraints.map((e) => ({
          code: e.ID,
          description: e.Description,
          condition_status: 1,
        }));
      })
        .catch((e) => {
          console.log(e);
        });
    },
    fetchProductList() {
      Vue.axios.get('http://localhost:8088/productlist').then((response) => {
        this.productList = response.data;
        console.log(response.data);
      })
        .catch((e) => {
          console.log(e);
        });
    },
    checkApproved() {
      return new Promise((resolve, reject) => {
        const unapproved = this.conditions
          // eslint-disable-next-line max-len
          .filter((e, i) => (Number(this.constraints[i].VaccineEligible) !== Number(e.condition_status)
            && this.constraints[i].VaccineEligible !== 2));
        if (unapproved.length >= 1) {
          reject(new Error("Based on the survey's results, you are not eligible for vaccination at the moment. "
            + 'Please consult your healthcare provider for more information.'));
        } else {
          this.fetchProductList();
          this.minimumAge = Math.min(...this.productList.map((e) => e.AuthorizedAges), -1);
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
