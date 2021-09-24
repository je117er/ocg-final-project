<template>
  <section>
    <b-table
      @click="clinicDetails"
      :paginated=true
      :per-page=10
      :data="data">
      <b-table-column field="name" label="Name" centered v-slot="props">
        {{ props.row.name }}
      </b-table-column>
      <b-table-column field="address" label="Address" centered v-slot="props">
        {{ props.row.address }}
      </b-table-column>
      <b-table-column field="location" label="location" centered v-slot="props">
        {{ props.row.location }}
      </b-table-column>
      <b-table-column field="status" label="status" centered v-slot="props">
        <b-checkbox :value="props.row.status" disabled
                    type="is-info">
        </b-checkbox>
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
    clinicDetails(item) {
      const clinicID = item.id;
      this.$router.push({
        name: 'clinic-detail',
        params: {
          id: `${clinicID}`,
        },
      });
    },
  },
  created() {
    Vue.axios.get('http://localhost:8088/admin/clinics').then((response) => {
      this.data = response.data;
      console.log(this.data);
    }).catch((err) => {
      console.log(err);
    });
  },
};
</script>
