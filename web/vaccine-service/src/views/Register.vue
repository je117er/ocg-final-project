<template>
  <div class="container is-fullhd">
    <h1 class="title">Register for vaccination</h1>
      <section class="is-small">
        <form>
          <div class="has-text-left is-small"
               v-for="(condition, index) in conditions" :key="condition.code">
<!--            <field name="condition.code">-->
              <div>
              {{ condition.description }}
                <div class="has-text-right block">
                  <b-radio v-model="condition.selected"
                           :value="yes(index)">
                    Yes
                  </b-radio>
                  <b-radio v-model="condition.selected"
                           :value="no(index)">
                    No
                  </b-radio>
                </div>
              </div>
<!--            </field>-->
          </div>
        </form>
        <p class="content">
          <b>Selection:</b>
          {{ radio }}
        </p>
      </section>

  </div>
</template>

<script>
import Vue from 'vue';

export default {
  name: 'Register',
  data() {
    return {
      radio: 'jack',
      radio1: 'jo',
      constraints: [],
      conditions: [],
    };
  },
  methods: {
    yes(value) {
      return `${value}_yes`;
    },
    no(value) {
      return `${value}_no`;
    },
  },
  created() {
    Vue.axios.get('http://localhost:8088/constraints').then((response) => {
      this.constraints = response.data;
      this.conditions = this.constraints.map((e) => ({
        code: e.Code,
        description: e.Description,
        selected: false,
      }));

      console.log(this.conditions);
    })
      .catch((e) => {
        console.log(e);
      });
  },
};
</script>

<style scoped>

</style>
