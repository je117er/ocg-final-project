<template>
  <section>
    <b-table
      @click="customerDetails"
      :paginated=true
      :per-page=10
      :data="data">
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
</template>

<script>
import Vue from 'vue';

export default {
  data() {
    return {
      data: [],
    };
  },
  methods: {
    customerDetails(item) {
      const customerID = item.id;
      this.$router.push({
        name: 'customer-detail',
        params: {
          id: `${customerID}`,
        },
      });
    },
  },
  created() {
    Vue.axios.get('http://localhost:8088/admin/customers').then((response) => {
      this.data = response.data;
    }).catch((err) => {
      console.log(err);
    });
  },
};
</script>
