<template>
  <b-tabs>
    <b-tab-item label="Customers" icon="">
      <section>
        <b-table
          @click="customerDetails"
          :paginated=true
          :per-page=10
          :data="customers">
          <b-table-column field="id" label="ID" width="40" numeric centered v-slot="props">
            {{ props.row.id }}
          </b-table-column>
          <b-table-column field="name" label="Full Name" centered v-slot="props">
            {{ props.row.name }}
          </b-table-column>
          <b-table-column field="address" label="Address" centered v-slot="props">
            {{ props.row.address }}
          </b-table-column>
          <b-table-column field="dob" label="Date of birth" centered v-slot="props">
            {{ props.row.dob }}
          </b-table-column>
          <b-table-column field="gender" label="Gender" centered v-slot="props">
            {{ props.row.gender }}
          </b-table-column>
          <b-table-column field="nationality" label="Nationality" centered v-slot="props">
            {{ props.row.nationality }}
          </b-table-column>
        </b-table>
      </section>
    </b-tab-item>
    <b-tab-item label="Info" icon="">
    </b-tab-item>
    <b-tab-item label="Work Schedule" icon="">
      <div class="container">
        <h2 v-if="isUpdateSession" style="color: darkgreen; font-weight: bold;">Update Success!</h2>
        <div class="notification">
          <b-datepicker v-model="currentDateUI"
                        inline
                        size="medium">
          </b-datepicker>

          <div class="tile is-ancestor">
            <div class="tile is-parent">
              <article class="tile is-child box">
                <p class="title">Morning</p>
                <div>
                  Status <input type="checkbox" v-model="morningStatus">
                  Capacity <input type="number" v-model="morningCapacity">
                </div>
              </article>
            </div>
            <div class="tile is-parent">
              <article class="tile is-child box">
                <p class="title">Afternoon</p>
                <div>
                  Status <input type="checkbox" v-model="afternoonStatus">
                  Capacity <input type="number" v-model="afternoonCapacity">
                </div>
              </article>
            </div>
            <div class="tile is-parent">
              <article class="tile is-child box">
                <p class="title">Evening</p>
                <div>
                  Status <input type="checkbox" v-model="eveningStatus">
                  Capacity <input type="number" v-model="eveningCapacity">
                </div>
              </article>
            </div>
          </div>
        </div>
        <div class="buttons is-centered">
          <b-button @click="updateSession">
            Save
          </b-button>
          <b-button tag="router-link"
                    to="/"
                    outlined
                    type="is-link is-danger">
            Cancel
          </b-button>
        </div>
      </div>
    </b-tab-item>
    <b-tab-item label="Merchandise Control" icon="">
    </b-tab-item>
  </b-tabs>
</template>

<script>
import Vue from 'vue';

const TYPE = {
  MORNING: 0,
  AFTERNOON: 1,
  EVENING: 2,
};

export default {
  data() {
    return {
      customers: [],
      loaded: false,
      currentDateUI: new Date(),
      isFetchSession: false,
      session: [],
      morningStatus: true,
      morningCapacity: 100,
      afternoonStatus: true,
      afternoonCapacity: 100,
      eveningStatus: true,
      eveningCapacity: 100,
      isUpdateSession: false,
    };
  },
  watch: {
    currentDateUI(newVal, oldVal) {
      console.log(newVal, oldVal);
      const newValMonth = newVal.getMonth() + 1;
      const newValYear = newVal.getFullYear();

      const oldValMonth = oldVal.getMonth() + 1;
      const oldValYear = oldVal.getFullYear();
      if (newValMonth !== oldValMonth || newValYear !== oldValYear) {
        Vue.axios.get(`http://localhost:8088/admin/session?month=${this.currentDateUI.getMonth() + 1}`).then((response) => {
          this.session = response.data;
          this.loaded = true;
          this.isFetchSession = false;
        }).catch((err) => {
          console.log(err);
        });
      }

      this.updateDataCapacity(TYPE.MORNING);
      this.updateDataCapacity(TYPE.AFTERNOON);
      this.updateDataCapacity(TYPE.EVENING);

      this.updateDataStatus(TYPE.MORNING);
      this.updateDataStatus(TYPE.AFTERNOON);
      this.updateDataStatus(TYPE.EVENING);
    },
    isFetchSession(newVal) {
      console.log(newVal);
      if (newVal) {
        Vue.axios.get(`http://localhost:8088/admin/session?month=${this.currentDateUI.getMonth() + 1}`).then((response) => {
          this.session = response.data;
          this.loaded = true;
          this.isFetchSession = false;
        }).catch((err) => {
          console.log(err);
        });
      }
    },
  },
  methods: {
    updateDataCapacity(type) {
      // eslint-disable-next-line max-len
      const data = this.session.filter((item) => new Date(item.currentDate).getDate() === this.currentDateUI.getDate() && item.type === type);
      switch (type) {
        case TYPE.MORNING:
          this.morningCapacity = data.length !== 0 ? data[0].capacity : 100;
          break;
        case TYPE.AFTERNOON:
          this.afternoonCapacity = data.length !== 0 ? data[0].capacity : 100;
          break;
        case TYPE.EVENING:
          this.eveningCapacity = data.length !== 0 ? data[0].capacity : 100;
          break;
        default:
          break;
      }
    },

    updateDataStatus(type) {
      // eslint-disable-next-line max-len
      const data = this.session.filter((item) => new Date(item.currentDate).getDate() === this.currentDateUI.getDate() && item.type === type);
      switch (type) {
        case TYPE.MORNING:
          this.morningStatus = data.length !== 0 ? data[0].status : true;
          break;
        case TYPE.AFTERNOON:
          this.afternoonStatus = data.length !== 0 ? data[0].status : true;
          break;
        case TYPE.EVENING:
          this.eveningStatus = data.length !== 0 ? data[0].status : true;
          break;
        default:
          break;
      }
    },

    customerDetails(item) {
      const customerID = item.id;
      this.$router.push({
        name: 'customer-detail',
        params: {
          id: `${customerID}`,
        },
      });
    },

    updateSession() {
      // eslint-disable-next-line max-len
      const currentData = this.session.filter((item) => new Date(item.currentDate).getDate() === this.currentDateUI.getDate());
      currentData.forEach((element) => {
        // eslint-disable-next-line no-param-reassign
        element.currentDateStr = (new Date(element.currentDate)).toISOString().slice(0, 10);
        switch (element.type) {
          case TYPE.MORNING:
            // eslint-disable-next-line no-param-reassign
            element.capacity = Number(this.morningCapacity);
            // eslint-disable-next-line no-param-reassign
            element.status = this.morningStatus;
            break;
          case TYPE.AFTERNOON:
            // eslint-disable-next-line no-param-reassign
            element.capacity = Number(this.afternoonCapacity);
            // eslint-disable-next-line no-param-reassign
            element.status = this.afternoonStatus;
            break;
          case TYPE.EVENING:
            // eslint-disable-next-line no-param-reassign
            element.capacity = Number(this.eveningCapacity);
            // eslint-disable-next-line no-param-reassign
            element.status = this.eveningStatus;
            break;
          default:
            break;
        }
      });
      console.log(currentData);
      Vue.axios.post('http://localhost:8088/admin/session', currentData).then(() => {
        this.loaded = true;
        this.isFetchSession = true;
        this.isUpdateSession = true;
      }).catch((err) => {
        console.log(err);
        this.isFetchSession = true;
      });
    },
  },
  created() {
    Vue.axios.get(`http://localhost:8088/admin/customer?clinicId=${this.$route.params.id}`).then((response) => {
      this.customers = response.data;
      this.loaded = true;
    }).catch((err) => {
      console.log(err);
    });

    Vue.axios.get(`http://localhost:8088/admin/session?month=${this.currentDateUI.getMonth() + 1}`).then((response) => {
      this.session = response.data;
      this.loaded = true;
    }).catch((err) => {
      console.log(err);
    });
  },
};
</script>
